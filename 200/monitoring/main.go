package main

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	// Create counters
	httpRequestsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests",
	})
	httpErrorsTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_errors_total",
		Help: "Count of all HTTP errors",
	})
	uptimeTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "uptime_total",
		Help: "uptime of app in seconds",
	})
)

func main() {

	// Register Prometheus collectors, collect metrics
	r := prometheus.NewRegistry()
	r.MustRegister(httpRequestsTotal)
	r.MustRegister(httpErrorsTotal)
	r.MustRegister(uptimeTotal)

	// Go routine
	go func() {
		for {
			uptimeTotal.Inc()
			time.Sleep(time.Second)
		}
	}()

	// Create handlers (3 handlers that have different headers)
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
	notfoundHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.Inc()
		httpErrorsTotal.Inc()
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("NotFound"))
	})
	internalErrorHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpRequestsTotal.Inc()
		httpErrorsTotal.Inc()
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("InternalServerError"))
	})

	// Create the mux server and add the handlers (Wrapped by counter handler)
	mux := http.NewServeMux()
	mux.Handle("/hello", foundHandler)
	mux.Handle("/err", notfoundHandler)
	mux.Handle("/internal-err", internalErrorHandler)
	mux.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{Registry: r}))

	// Host it
	srv := &http.Server{
		Addr: ":8080",
		Handler: h2c.NewHandler(
			mux,
			&http2.Server{},
		)}

  log.Printf("server starting on port %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func doWork() error {
	// Simulate a wait time
	waitTime := time.Duration(rand.Intn(600)) * time.Millisecond
	time.Sleep(waitTime)

	// Randomize fail rate
	if rand.Float64() < 0.2 {
		return errors.New("doWork error")
	}

	return nil
}
