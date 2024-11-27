---
layout: center
---

# Monitoring

---
layout: center
---

# What's Monitoring? 

Collecting and reporting on specific metrics to identify issues in a system. Monitoring can provide alerts and data on a system's performance, but it doesn't necessarily explain what's causing problems.

---
layout: center
---

## Why should we care?

---

- Application performance
- Business performance
- Delight our users
- Troubleshooting
- Real-time data
- Security

---

## Prometheus and Grafana architecture

<img src="/prom.webp" style="height: 90%">

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
```go{all|2-4|6-17|6|7|8|9-14|15-16}
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

## Exposing our Metrics
```go{all|4-6|6|8-14|15}    
main(){
    ...

    mux := http.NewServeMux()
    mux.Handle("/hello", foundHandler)
    mux.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{Registry: r}))

    srv := &http.Server{
        Addr: ":8080",
        Handler: h2c.NewHandler(
            mux,
            &http2.Server{},
        )}

    log.Fatal(srv.ListenAndServe())
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
- Data driven iteration!
