package titan

import (
	"github.com/atlaskerr/titan/metrics"
)

// Core defines high-level methods for titan's external operations.
type Core struct {
	metrics *metrics.Collector
}
