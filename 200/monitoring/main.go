package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/a-h/go-workshop/200/monitoring/metrics"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/metric"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	ctx := context.Background()

	err := run(ctx, log)
	if err != nil {
		log.Error("failed to run example", slog.Any("error", err))
	}
}

func run(ctx context.Context, log *slog.Logger) (err error) {
	log.Info("starting example")

	// Configure Prometheus metrics.
	m, err := metrics.New(ctx)
	if err != nil {
		return fmt.Errorf("failed to create metrics: %w", err)
	}
	// Use a non-standard scrape port.
	go metrics.ListenAndServe(":9191")

	// Configure tracing to output to stdout (not useful, actually).
	// If you're using AWS, you'd use X-Ray here instead.
	traceExporter, err := stdouttrace.New()
	if err != nil {
		return fmt.Errorf("failed to create trace exporter: %w", err)
	}
	tp := tracesdk.NewTracerProvider(tracesdk.WithBatcher(traceExporter))
	otel.SetTracerProvider(tp)
	tracer := tp.Tracer("github.com/a-h/go-workshop/200/observability")

	// Count uptime in seconds. This isn't perfectly accurate, because
	// time.Sleep sleeps for _at least_ the given duration, but it's
	// close enough for this use case.
	go func() {
		// Manually define a trace for this background operation.
		ctx, span := tracer.Start(context.Background(), "background-operation")
		defer span.End()
		for {
			m.UptimeTotalSeconds.Add(ctx, 1)
			log.Debug("inremented uptime counter")
			time.Sleep(time.Second)
		}
	}()

	// Stand up the application.
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello!"))
	})
	mux.Handle("/not-found", http.NotFoundHandler())
	mux.HandleFunc("/internal-err", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	})

	// Wrap the mux with our counter middleware.
	withCounter := NewMetricsMiddleware(m.HTTPRequestsTotal, m.RequestDurationSeconds, mux)

	// Automatically trace all incoming requests.
	// This doesn't do anything useful unless you have a trace provider configured, e.g. AWS X-Ray.
	withOtelHTTP := otelhttp.NewHandler(withCounter, "http")

	s := &http.Server{
		Addr:              ":8080",
		Handler:           withOtelHTTP,
		ReadHeaderTimeout: 5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		IdleTimeout:       5 * time.Minute,
	}
	return s.ListenAndServe()
}

// The metrics middleware could be compressed to an anonymous function similar to this:
//
// totalMiddleware := func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			m.HTTPRequestsTotal.Add(r.Context(), 1)
//			next.ServeHTTP(w, r)
//		})
//	}
//
// But for clarity, it's been expanded out here.

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
