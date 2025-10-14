package metrics

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	metricsdk "go.opentelemetry.io/otel/sdk/metric"
)

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
