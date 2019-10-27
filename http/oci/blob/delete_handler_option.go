package blob

import (
	"errors"

	"github.com/atlaskerr/titan/metrics"
)

// DeleteHandlerOption applies a parameter to a DeleteHandler.
type DeleteHandlerOption func(*DeleteHandler)

// DeleteOptionMetrics sets the http.Handler to use for the metrics
// endpoint.
func DeleteOptionMetrics(collector *metrics.Collector) DeleteHandlerOption {
	var fn DeleteHandlerOption = func(s *DeleteHandler) {
		s.metrics = collector
	}
	return fn
}

func NewDeleteHandler(options ...DeleteHandlerOption) (*DeleteHandler, error) {
	srv := &DeleteHandler{}
	for _, addOption := range options {
		addOption(srv)
	}
	if srv.metrics == nil {
		return nil, errors.New("no metrics collector defined")
	}
	return srv, nil
}
