package http

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Collector is a custom prometheus collector for titan.
type Collector struct {
	TotalRequests    *prometheus.CounterVec
	RequestsInFlight *prometheus.GaugeVec
	RequestDuration  *prometheus.HistogramVec
	RequestSize      *prometheus.HistogramVec
	ResponseSize     *prometheus.HistogramVec
}

// NewCollector initializes a Collector.
func NewCollector() *Collector {
	c := &Collector{
		TotalRequests:    totalRequestsCounter(),
		RequestsInFlight: requestsInFlightGauge(),
		RequestDuration:  requestDurationHistogram(),
		RequestSize:      requestSizeHistogram(),
		ResponseSize:     responseSizeHistogram(),
	}
	return c
}

// Describe implements prometheus.Collector.
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	c.TotalRequests.Describe(ch)
	c.RequestsInFlight.Describe(ch)
	c.RequestDuration.Describe(ch)
	c.RequestSize.Describe(ch)
	c.ResponseSize.Describe(ch)
}

// Collect implements prometheus.Collector.
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.TotalRequests.Collect(ch)
	c.RequestsInFlight.Collect(ch)
	c.RequestDuration.Collect(ch)
	c.RequestSize.Collect(ch)
	c.ResponseSize.Collect(ch)
}
