package server

import (
	"fmt"

	"github.com/atlaskerr/titan/http/live"
	"github.com/atlaskerr/titan/http/oci"
	"github.com/atlaskerr/titan/http/ready"
	"github.com/atlaskerr/titan/http/titan"
	"github.com/atlaskerr/titan/http/unknown"
)

func titanHandler(s *service) error {
	opts := []titan.ServerOption{
		titan.OptionMetricsCollector(s.collector),
		titan.OptionOCIHandler(s.handlers.oci),
		titan.OptionLiveHandler(s.handlers.live),
		titan.OptionReadyHandler(s.handlers.ready),
		titan.OptionMetricsHandler(s.handlers.metrics),
		titan.OptionUnknownHandler(s.handlers.unknown),
	}
	h, err := titan.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize titan router: %s", err)
	}
	s.handlers.titan = h
	return nil
}

func liveHandler(s *service) error {
	h := &live.Server{}
	s.handlers.live = h
	return nil
}

func readyHandler(s *service) error {
	h := &ready.Server{}
	s.handlers.ready = h
	return nil
}

func ociHandler(s *service) error {
	h := &oci.Server{}
	s.handlers.oci = h
	return nil
}

func unknownHandler(s *service) error {
	h := &unknown.Server{}
	s.handlers.unknown = h
	return nil
}
