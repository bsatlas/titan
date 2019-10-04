package http

import (
	"github.com/prometheus/client_golang/prometheus"
)

// totalRequestsCounter records total requests handled by HTTP endpoints.
func totalRequestsCounter() *prometheus.CounterVec {
	opts := prometheus.CounterOpts{
		Name: "titan_http_requests_total",
		Help: "Total number of requests served.",
	}
	labels := []string{"code", "endpoint"}
	return prometheus.NewCounterVec(opts, labels)
}

// requestsInFlightGauge records total requests currently being handled by
// HTTP endpoints.
func requestsInFlightGauge() *prometheus.GaugeVec {
	opts := prometheus.GaugeOpts{
		Name: "titan_http_requests_in_flight",
		Help: "Current number of requests being served.",
	}
	labels := []string{"endpoint"}
	return prometheus.NewGaugeVec(opts, labels)
}

// requestDurationHistogram records the duration of requests handled by HTTP
// endpoints.
func requestDurationHistogram() *prometheus.HistogramVec {
	opts := prometheus.HistogramOpts{
		Name:    "titan_http_request_duration_seconds",
		Help:    "Duration of requests served.",
		Buckets: []float64{.1, 0.25, .5, 1, 2.5, 5, 10},
	}
	labels := []string{"endpoint", "code"}
	return prometheus.NewHistogramVec(opts, labels)
}

// requestSizeHistogram records the approximate size of a request.
func requestSizeHistogram() *prometheus.HistogramVec {
	opts := prometheus.HistogramOpts{
		Name: "titan_http_request_size_bytes",
		Help: "Approximate request size.",
		Buckets: []float64{
			200, 500, 700, 1000, 1200, 1500, 2000, 5000, 10000, 15000,
		},
	}
	labels := []string{"endpoint"}
	return prometheus.NewHistogramVec(opts, labels)
}

// responseSizeHistogram records the approximate size of a response.
func responseSizeHistogram() *prometheus.HistogramVec {
	opts := prometheus.HistogramOpts{
		Name: "titan_http_response_size_bytes",
		Help: "Approximate response size.",
		Buckets: []float64{
			200, 500, 700, 1000, 1200, 1500, 2000, 5000, 10000, 15000,
		},
	}
	labels := []string{"endpoint"}
	return prometheus.NewHistogramVec(opts, labels)
}
