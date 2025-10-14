# Monitoring

Metrics, Traces and Logs are critical for determining the health and performance of your applications.

## Logs

In all of the examples in this repo, we've used a technique called structured logging. This is where a log entry is more like an event with key/value pairs attached, for example (formatted for readability, in practice this would be a single line):

```json
{
  "level": "info",
  "ts": "2024-06-20T12:00:00Z",
  "msg": "User logged in",
  "userID": "12345",
  "ipAddress": "127.0.0.1"
}
```

Go has built-in structured logging in the `slog` package.

`slog` supports different handlers, which determine how logs are formatted and where the output is sent.

Usually, logs are sent to stderr, and then collected by a log aggregation system. In the cloud, container and compute services like Fargate, Lambda, and Cloud Run automatically collect logs sent to stdout/stderr and send them to the cloud's logging service.

In a self-managed environment, you might use a log aggregation tool like Logstash to push data to a self-hosted Loki system.

The two main handlers are `TextHandler` and `JSONHandler`. `TextHandler` is designed to be more human-readable, while `JSONHandler` is designed to be machine-readable and easily parsed by log aggregation systems.

For server applications, I always default to `JSONHandler`.

```go
import (
    "log/slog"
)

func main() {
    log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
    log.Debug("This is a debug message")
    log.Info("This is an info message", slog.String("key", "value"))
    err := errors.New("something went wrong")
    log.Error("This is an error message", slog.Any("error", err))
}
```

Go provides a global logger via `slog.Default()`, which you can configure at the start of your application, but I usually avoid the global logger and instead create a logger instance and pass it around as needed to ensure that no other code can change the logger configuration unexpectedly.

## Metrics

Metrics are numerical values that represent some aspect of your application's performance or behavior over time. They are typically collected at regular intervals and can be used to monitor the health of your application, identify trends, and trigger alerts when certain thresholds are crossed.

It used to be common to use a platform or framework specific library such as the AWS CloudWatch SDK, the DataDog Agent, the AWS X-Ray SDK, or the Prometheus client libraries to instrument your applications.

However, this made it difficult to switch monitoring systems or to support multiple systems simultaneously.

The OpenTelemetry project provides a vendor-neutral way to instrument your applications for metrics, traces, and logs, and then export that data to a variety of backends.

If you use the OpenTelemetry SDKs to instrument your application, you can switch between different monitoring systems by changing the exporter configuration, without having to change your application code.

### OpenTelemetry and Lambda

However, you should be aware that if you use OpenTelemetry with AWS Lambda, you may incur additional cold start latency of around 200ms while the Lambda Layer used to export OpenTelemetry data to CloudWatch is initialized, see https://github.com/open-telemetry/opentelemetry-lambda/issues/263

If the Lambda workload is latency-sensitive, I would recommend using Embedded Metrics Format (EMF) to send metrics directly to CloudWatch, see https://aws.amazon.com/blogs/mt/introducing-embedded-metrics-format-for-structured-logging-in-amazon-cloudwatch/ - I've used https://github.com/prozz/aws-embedded-metrics-golang in production before.

### OpenTelemetry Metrics

The `./metrics` package shows how to create counters and histograms.

```go
type Metrics struct {
	HTTPRequestsTotal      metric.Int64Counter
	RequestDurationSeconds metric.Float64Histogram
	UptimeTotalSeconds     metric.Int64Counter
}

func New(ctx context.Context) (metrics Metrics, err error) {
	// Setup Prometheus to export metrics.
	exp, err := prometheus.New()
	if err != nil {
		return metrics, fmt.Errorf("failed to create prometheus exporter: %w", err)
	}
	// A MeterProvider is required to create Meters like Counters, Gauges, etc.
	// We've configured Prometheus to read from this MeterProvider, so any
	// instruments created from this MeterProvider will be exported to Prometheus.
	meter := metricsdk.
		NewMeterProvider(metricsdk.WithReader(exp)).
		Meter("github.com/a-h/go-workshop/200/observability")

	// Create counters.
	metrics.HTTPRequestsTotal, err = meter.Int64Counter("http_requests_total")
	if err != nil {
		return metrics, fmt.Errorf("failed to create http_requests_total counter: %w", err)
	}
	metrics.RequestDurationSeconds, err = meter.Float64Histogram("http_request_duration")
	if err != nil {
		return metrics, fmt.Errorf("failed to create http_request_duration_seconds histogram: %w", err)
	}
	metrics.UptimeTotalSeconds, err = meter.Int64Counter("uptime_total")
	if err != nil {
		return metrics, fmt.Errorf("failed to create uptime_total counter: %w", err)
	}

	return metrics, nil
}
```

The Prometheus exporter is configured to read from the `MeterProvider`, so any instruments created from this `MeterProvider` will have their values made available to the Prometheus exporter.

But what does that mean? Not much really - it just means that the Prometheus exporter has the values stored in RAM.

Prometheus uses a pull model, where it scrapes the metrics from the application at regular intervals (default is every 15 seconds). So, the application needs to be running and accessible for Prometheus to scrape the metrics.

To scrape the application, it needs to be running a HTTP server that exposes the metrics in a format that Prometheus understands.

Rather than serve metrics on the same endpoint as our customer traffic, it's best to create a separate server for the scrape endpoint, and run it on a different port.

```go
func ListenAndServe(addr string) (err error) {
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.Handler())
	s := &http.Server{
		Addr:              addr,
		Handler:           m,
		ReadHeaderTimeout: 5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		IdleTimeout:       5 * time.Minute,
	}
	return s.ListenAndServe()
}
```

From our `run` function, we can create our counters and histograms, then start the scrape endpoint in a separate goroutine so that it runs concurrently with our main HTTP server.

```go
// Configure Prometheus metrics.
m, err := metrics.New(ctx)
if err != nil {
    return fmt.Errorf("failed to create metrics: %w", err)
}
// Use a non-standard scrape port.
go metrics.ListenAndServe(":9191")
```

### Middleware

To record metrics for each HTTP request, we can create middleware that wraps our HTTP handlers.

It simply records the start time, increments the request counter, calls the next handler, then records the duration to the histogram.

```go
func NewMetricsMiddleware(httpRequestsTotal metric.Int64Counter, requestDurationSeconds metric.Float64Histogram, next http.Handler) http.Handler {
	return MetricsMiddleware{
		httpRequestsTotal:      httpRequestsTotal,
		requestDurationSeconds: requestDurationSeconds,
		next:                   next,
	}
}

type MetricsMiddleware struct {
	httpRequestsTotal      metric.Int64Counter
	requestDurationSeconds metric.Float64Histogram
	next                   http.Handler
}

func (c MetricsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	c.httpRequestsTotal.Add(r.Context(), 1)
	c.next.ServeHTTP(w, r)
	c.requestDurationSeconds.Record(r.Context(), time.Since(start).Seconds())
}
```

## Traces

Traces are a way to track the flow of a request through a distributed system. They provide visibility into how requests are processed, where time is spent, and where errors occur.

OpenTelemetry provides a way to instrument your application for tracing, and then export that data to a variety of backends.

```go
// Configure tracing to output to stdout (not useful, actually).
// If you're using AWS, you'd use X-Ray here instead.
traceExporter, err := stdouttrace.New()
if err != nil {
    return fmt.Errorf("failed to create trace exporter: %w", err)
}
tp := tracesdk.NewTracerProvider(tracesdk.WithBatcher(traceExporter))
otel.SetTracerProvider(tp)
tracer := tp.Tracer("github.com/a-h/go-workshop/200/observability")
````

This example outputs trace spans to stdout, which isn't very useful compared to a real tracing backend like AWS X-Ray, Jaeger, or Zipkin which all provide a way to visualise traces that cross application boundaries, and highlight where time is spent and errors occur.

Again, with Lambda, you might want to stick with X-Ray directly rather than using OpenTelemetry, to avoid the additional cold start latency.

To create spans, you can use the `tracer` instance created above.

```go
func HelloHandler(tracer trace.Tracer) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx, span := tracer.Start(r.Context(), "HelloHandler")
        defer span.End()

        // Simulate some work.
        time.Sleep(100 * time.Millisecond)

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Hello, world!\n"))

        // Add attributes to the span.
        span.SetAttributes(attribute.String("handler", "hello"))
    })
}
```

However, you don't to do this yourself, the OpenTelemetry project provides HTTP middleware that does this for you.

```go
m := http.NewServeMux()
// Register routes...
handler := otelhttp.NewHandler(m, "server")
// Start the server.
http.ListenAndServe(":8080", handler)
```

## Connecting to Prometheus and Grafana

This repo contains a `compose.yaml` file that starts Prometheus and Grafana in Docker containers.

To start them, run:

```bash
docker compose up
```

Prometheus will be available on port 9090, and Grafana on port 3001. The `./prometheus/prometheus.yml` file configures Prometheus to scrape the metrics of the example app endpoint on port 9191.

The `compose.yaml` also contains instructions to build and run the Go application as a Docker container. It makes careful use Docker layers, and the the multi-stage build feature of Docker to create a small final image that only contains the compiled Go binary and a few necessary files such as the TLS certificate authority bundle and timezone data.

You can access your application on http://localhost:8080 and see the scrape endpoint of your application on http://localhost:9191/metrics

Use these to access endpoints:

- http://localhost:8080/hello - This is your basic 200 route
- http://localhost:9191/metrics - Where Prometheus metrics are displayed
- http://localhost:8080/not-found - An endpoint that produces a 404 error
- http://localhost:8080/internal-err - An endpoint that produces a 500 error
- http://localhost:9090/query - Prometheus
- http://localhost:3001/ - Grafana dashboard

### Checking that Prometheus can see your app

Go to http://localhost:9090/targets in your browser, you should see your app listed as a target, and it should be `UP`.

### Checking that Grafana can see Prometheus

Go to http://localhost:3001/ in your browser and logging with username `admin` and password `grafana`.

Visit http://localhost:3001/connections/datasources - you should see Prometheus listed as a data source. If you don't, you can add one manually.

## Viewing data

Create a new dashboard in Grafana, and add a new metric to it. If you limit the metrics to `otel_scope_name="github.com/a-h/go-workshop/200/observability"` you should see the metrics created in the `./metrics` package.

The other metrics are created by the OpenTelemetry SDK and the Prometheus exporter.

## Dashboard guide

To see the Grafana dashboard for uptime use the following:
- Navigate to http://localhost:3001/ - the Grafana dashboard.
- You'll be presented with a login screen.
    - Username: `admin` 
    - Password: `grafana`
- You should be on the home screen now. Click dashboards on the left or navigate to http://localhost:3001/dashboards
- Click the blue `New` button in the top right, and click `New Dashboard`.
- `Start your new dashboard by adding a visualization` should be on your screen now. Click the big blue `+ Add visualization`.
- Select `Prometheus` in the data source pop up.
- Your dashboard should now be created, make sure you're on the query tab at the bottom, click add query and select `uptime_total` from the metric dropdown (If you're adding your own metric, then change this).
- Finally, click `Run Queries` button, you should then see a graph of the uptime.

## Tasks

### docker-up

Start the grafana and prometheus docker containers. Grafana port 3001 and Prometheus port 9090

```bash
docker compose up
```

### run-app

Start the go app. It will serve on port 8080 

```bash
go run .
```

### get-hello

Call to the `/hello` endpoint.

```bash
curl --verbose "http://localhost:8080/hello"
```

### get-metrics

Call to the `metrics` endpoint.

```bash
curl --verbose "http://localhost:8080/metrics"
```

### notfound-err

Call to the `/err` endpoint.

```bash
curl --verbose "http://localhost:8080/err"
```

### internal-err

Call to the `/internal-err` endpoint.

```bash
curl --verbose "http://localhost:8080/internal-err"
```

## Next steps

- Add another metric.
- Query them in Prometheus.
- Create a dashboard in grafana.
