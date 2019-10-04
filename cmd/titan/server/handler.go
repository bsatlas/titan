package server

import (
	"fmt"

	"github.com/atlaskerr/titan/http/live"
	"github.com/atlaskerr/titan/http/oci"
	"github.com/atlaskerr/titan/http/ready"
	"github.com/atlaskerr/titan/http/titan"
	"github.com/atlaskerr/titan/http/undefined"
)

func titanHandler(s *service) error {
	opts := []titan.ServerOption{
		titan.OptionMetricsCollector(s.collector),
		titan.OptionOCIHandler(s.handlers.oci),
		titan.OptionLiveHandler(s.handlers.live),
		titan.OptionReadyHandler(s.handlers.ready),
		titan.OptionMetricsHandler(s.handlers.metrics),
		titan.OptionUndefinedHandler(s.handlers.undefined),
	}
	h, err := titan.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize titan router: %s", err)
	}
	s.handlers.titan = h
	return nil
}

func liveHandler(s *service) error {
	opts := []live.ServerOption{
		live.OptionMetricsCollector(s.collector),
		live.OptionUndefinedHandler(s.handlers.undefined),
	}
	h, err := live.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize live endpoint: %s", err)
	}
	s.handlers.live = h
	return nil
}

func readyHandler(s *service) error {
	opts := []ready.ServerOption{
		ready.OptionMetricsCollector(s.collector),
		ready.OptionUndefinedHandler(s.handlers.undefined),
	}
	h, err := ready.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize ready endpoint: %s", err)
	}
	s.handlers.ready = h
	return nil
}

func ociHandler(s *service) error {
	opts := []oci.ServerOption{
		oci.OptionMetricsCollector(s.collector),
		oci.OptionUndefinedHandler(s.handlers.undefined),
	}
	h, err := oci.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize oci endpoint: %s", err)
	}
	s.handlers.oci = h
	return nil
}

func undefinedHandler(s *service) error {
	opts := []undefined.ServerOption{
		undefined.OptionMetricsCollector(s.collector),
	}
	h, err := undefined.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize unknown endpoint: %s", err)
	}
	s.handlers.undefined = h
	return nil
}
