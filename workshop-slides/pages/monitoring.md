---
layout: center
---

# Monitoring

---
layout: center
---

## What's Monitoring? 
Collecting and reporting on specific metrics to identify issues in a system. Monitoring can provide alerts and data on a system's performance, but it doesn't necessarily explain what's causing problems.

---
layout: center
---

## Why should we care?

---

- Application's performance
- Delight our users
- Troubleshooting
- Real-time data
- Security

---

### How are we going to do that?

Prometheus is an open-source monitoring and alerting tool designed for reliability and scalability. It collects and stores metrics as time-series data, provides a query language (PromQL).

Grafana is an open-source visualization and analytics tool that transforms monitoring data into interactive dashboards. It supports multiple data sources, including Prometheus, and helps teams monitor, analyze, and alert on metrics in real-time.

---

## Prometheus in Go

Prometheus has an offical go package
```go{all|1-6|7-17|8-9|11-16}
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

## Prometheus with API
```go{all|1-6|7-38|8-10|13-37}
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
}
```
---
layout: center
---

## Let's get monitoring!

---

In the `monitoring` folder you find a `README.md` that will guide you through setting up the application.

```bash
monitoring
├── README.md
├── compose.yaml
├── grafana
│   └── datasource.yml
├── main.go
└── prometheus
    └── prometheus.yml
```
We will spend around 25 minutes.

---

## Tips

- Middlware to capture metrics (especially for APIs)
- RED method - dashboards, metrics
- OpenTelemetry
- Data driven iteration!
