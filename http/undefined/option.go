package undefined

import (
	"errors"

	"github.com/atlaskerr/titan/metrics"
)

// ServerOption applies a parameter to a Server.
type ServerOption func(*Server)

// OptionMetricsCollector sets the http.Handler to use for the metrics endpoint.
func OptionMetricsCollector(collector *metrics.Collector) ServerOption {
	var fn ServerOption = func(s *Server) {
		s.metrics = collector
	}
	return fn
}

// NewServer returns a fully initialized Server.
func NewServer(options ...ServerOption) (*Server, error) {
	srv := &Server{}
	for _, addOption := range options {
		addOption(srv)
	}
	if srv.metrics == nil {
		return nil, errors.New("no metrics collector defined")
	}
	return srv, nil
}
