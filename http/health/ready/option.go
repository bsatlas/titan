package ready

import (
	"errors"
	"net/http"

	"github.com/atlaskerr/titan/metrics"

	opentracing "github.com/opentracing/opentracing-go"
)

// ServerOption applies a parameter to a Server.
type ServerOption func(*Server)

// OptionCore sets the core to use for Readiness checks.
func OptionCore(core Readiness) ServerOption {
	var fn ServerOption = func(s *Server) {
		s.core = core
	}
	return fn
}

// OptionTracer sets the opentracing.Tracer to use for tracing.
func OptionTracer(tracer opentracing.Tracer) ServerOption {
	var fn ServerOption = func(s *Server) {
		s.tracer = tracer
	}
	return fn
}

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
	if srv.core == nil {
		return nil, errors.New("no core defined")
	}
	if srv.tracer == nil {
		return nil, errors.New("no tracer defined")
	}
	if srv.metrics == nil {
		return nil, errors.New("no metrics collector defined")
	}
	if srv.handlers.undefined == nil {
		return nil, errors.New("no unknown handler defined")
	}
	return srv, nil
}
