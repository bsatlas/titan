package titan

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

// OptionOCIHandler sets the http.Handler to use for the oci endpoint.
func OptionOCIHandler(handler http.Handler) ServerOption {
	var fn ServerOption = func(s *Server) {
		s.handlers.oci = handler
	}
	return fn
}

// OptionMetricsHandler sets the http.Handler to use for the metrics endpoint.
func OptionMetricsHandler(handler http.Handler) ServerOption {
	var fn ServerOption = func(s *Server) {
		s.handlers.metrics = handler
	}
	return fn
}

// OptionLiveHandler sets the http.Handler to use for the live endpoint.
func OptionLiveHandler(handler http.Handler) ServerOption {
	var fn ServerOption = func(s *Server) {
		s.handlers.live = handler
	}
	return fn
}

// OptionReadyHandler sets the http.Handler to use for the ready endpoint.
func OptionReadyHandler(handler http.Handler) ServerOption {
	var fn ServerOption = func(s *Server) {
		s.handlers.ready = handler
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
	if srv.handlers.oci == nil {
		return nil, errors.New("no oci handler defined")
	}
	if srv.handlers.metrics == nil {
		return nil, errors.New("no metrics handler defined")
	}
	if srv.handlers.live == nil {
		return nil, errors.New("no live handler defined")
	}
	if srv.handlers.ready == nil {
		return nil, errors.New("no ready handler defined")
	}
	if srv.handlers.undefined == nil {
		return nil, errors.New("no undefined handler defined")
	}
	return srv, nil
}
