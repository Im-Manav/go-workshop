---
layout: center
---

# Monitoring

---
layout: center
---

## What's Monitoring? 
Involves collecting and reporting on specific metrics to identify issues in a system. Monitoring can provide alerts and data on a system's performance, but it doesn't necessarily explain what's causing problems.

---
layout: center
---
Why should we care?
---

- Application's performance
- Delight our users
- Troubleshooting
- Real-time data
- Security
---

## How are we going to do that?

Prometheus is an open-source monitoring and alerting tool designed for reliability and scalability. It collects and stores metrics as time-series data, provides a query language (PromQL).

Grafana is an open-source visualization and analytics tool that transforms monitoring data into interactive dashboards. It supports multiple data sources, including Prometheus, and helps teams monitor, analyze, and alert on metrics in real-time.

---


---
# Prometheus in Go

Prometheus has an offical go package
```go
var(
    uptimeTotal = prometheus.NewCounter(prometheus.CounterOpts{
            Name: "uptime_total",
            Help: "uptime of app in seconds",
        })
)
func main(){
    r := prometheus.NewRegistry()
    r.MustRegister(uptimeTotal)

    go func() {
        for {
            uptimeTotal.Inc()
            time.Sleep(time.Second)
        }
    }()
}
```
---

# Prometheus with API
```go
var (
    httpRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "http_requests_total",
        Help: "Count of all HTTP requests",
    })
)
func main() {
    r := prometheus.NewRegistry()
    r.MustRegister(httpRequestsTotal)
    r.MustRegister(httpErrorsTotal)


    foundHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        httpRequestsTotal.Inc()
        err := doWork()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(err.Error()))
            httpErrorsTotal.Inc()
            return
        }
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Hello from example application."))
    })

    mux := http.NewServeMux()
    mux.Handle("/hello", foundHandler)

     // Host it
    srv := &http.Server{
        Addr: ":8080",
        Handler: h2c.NewHandler(
            mux,
            &http2.Server{},
        )}

    srv.ListenAndServe()
}
```

---

# Red Method
- Rate (the number of requests per second)
- Errors (the number of those requests that are failing)
- Duration (the amount of time those requests take)

---

# Why Red?
- Good for microservices, and records things that directly affect users
- Good proxy to how happy your customers will be
- Generally, you want to focus on business metrics rather than technical ones (cpu used etc)

