package metrics

import (
	"github.com/atlaskerr/titan/metrics/http"

	"github.com/prometheus/client_golang/prometheus"
)

// Collector is a custom prometheus collector for titan metrics.
type Collector struct {
	HTTP *http.Collector
}

// NewCollector initialiizes a Collector.
func NewCollector() *Collector {
	c := &Collector{
		HTTP: http.NewCollector(),
	}
	return c
}

// Describe implements prometheus.Collector.
func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	c.HTTP.Describe(ch)
}

// Collect implements prometheus.Collector.
func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	c.HTTP.Collect(ch)
}
