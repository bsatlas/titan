package metrics

import (
	"errors"
	"net/http"

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

// OptionUndefinedHandler sets the http.Handler to use for the unknown endpoint.
func OptionUndefinedHandler(handler http.Handler) ServerOption {
	var fn ServerOption = func(s *Server) {
		s.handlers.undefined = handler
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
	if srv.handlers.undefined == nil {
		return nil, errors.New("no unknown handler defined")
	}
	return srv, nil
}
