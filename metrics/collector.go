package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Collector is a custom prometheus collector for titan metrics.
type Collector struct{}

// NewCollector initialiizes a Collector.
func NewCollector() *Collector {
	c := &Collector{}
	return c
}

// Describe implements prometheus.Collector.
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {}

// Collect implements prometheus.Collector.
func (c *Collector) Collect(ch chan<- prometheus.Metric) {}
