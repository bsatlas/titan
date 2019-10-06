package titan

import (
	"errors"

	"github.com/atlaskerr/titan/metrics"
)

// CoreOption applies a parameter to a Core.
type CoreOption func(*Core)

// OptionMetricsCollector sets the http.Handler to use for the metrics endpoint.
func OptionMetricsCollector(collector *metrics.Collector) CoreOption {
	var fn CoreOption = func(c *Core) {
		c.metrics = collector
	}
	return fn
}

// NewCore returns a fully initialized Core.
func NewCore(options ...CoreOption) (*Core, error) {
	core := &Core{}
	for _, addOption := range options {
		addOption(core)
	}
	if core.metrics == nil {
		return nil, errors.New("no metrics collector defined")
	}
	return core, nil
}
