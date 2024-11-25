---
layout: center
---

# Monitoring

---

# What's Monitoring? 

Collecting live data about our program.

---

# Promethus and Graphana

TODO: introduce these
Promethus is an open source monitoring system
Graphana is an observability platform and allows for visualisations

# Promethus in Go

Promethus has an offical go package
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

# Promethus in Go
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



// Should this be included somwhere else? Efectivly the conclusion

Red Method:
- Rate (the number of requests per second)
- Errors (the number of those requests that are failing)
- Duration (the amount of time those requests take)

Good for microservices, and records things that directly affect users
Good proxy to how happy your customers will be
Generally, you want to focus on business metrics rather than technical ones (cpu used etc)
Of course these can be useful in certain circumstances but shouldn't be your goal



