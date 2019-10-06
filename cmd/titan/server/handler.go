package server

import (
	"fmt"

	"github.com/atlaskerr/titan/http/blob"
	"github.com/atlaskerr/titan/http/live"
	"github.com/atlaskerr/titan/http/manifest"
	"github.com/atlaskerr/titan/http/metrics"
	"github.com/atlaskerr/titan/http/oci"
	"github.com/atlaskerr/titan/http/ready"
	"github.com/atlaskerr/titan/http/tag"
	"github.com/atlaskerr/titan/http/titan"
	"github.com/atlaskerr/titan/http/undefined"

	"github.com/prometheus/client_golang/prometheus"
)

func cmpTitanHandler(s *service) error {
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

func cmpLiveHandler(s *service) error {
	opts := []live.ServerOption{
		live.OptionCore(s.core),
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

func cmpReadyHandler(s *service) error {
	opts := []ready.ServerOption{
		ready.OptionCore(s.core),
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

func cmpMetricsHandler(s *service) error {
	gatherer := prometheus.NewRegistry()
	err := gatherer.Register(s.collector)
	if err != nil {
		return fmt.Errorf("failed to register collector: %s", err)
	}
	opts := []metrics.ServerOption{
		metrics.OptionCore(gatherer),
		metrics.OptionMetricsCollector(s.collector),
		metrics.OptionUndefinedHandler(s.handlers.undefined),
	}
	h, err := metrics.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize metrics endpoint: %s", err)
	}
	s.handlers.metrics = h
	return nil
}

func cmpManifestHandler(s *service) error {
	opts := []manifest.ServerOption{
		manifest.OptionMetricsCollector(s.collector),
		manifest.OptionUndefinedHandler(s.handlers.undefined),
	}
	h, err := manifest.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize oci endpoint: %s", err)
	}
	s.handlers.manifest = h
	return nil
}

func cmpBlobHandler(s *service) error {
	opts := []blob.ServerOption{
		blob.OptionMetricsCollector(s.collector),
		blob.OptionUndefinedHandler(s.handlers.undefined),
	}
	h, err := blob.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize oci endpoint: %s", err)
	}
	s.handlers.blob = h
	return nil
}

func cmpTagHandler(s *service) error {
	opts := []tag.ServerOption{
		tag.OptionMetricsCollector(s.collector),
		tag.OptionUndefinedHandler(s.handlers.undefined),
	}
	h, err := tag.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize oci endpoint: %s", err)
	}
	s.handlers.tag = h
	return nil
}

func cmpOCIHandler(s *service) error {
	opts := []oci.ServerOption{
		oci.OptionMetricsCollector(s.collector),
		oci.OptionManifestHandler(s.handlers.manifest),
		oci.OptionBlobHandler(s.handlers.blob),
		oci.OptionTagHandler(s.handlers.tag),
		oci.OptionUndefinedHandler(s.handlers.undefined),
	}
	h, err := oci.NewServer(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize oci endpoint: %s", err)
	}
	s.handlers.oci = h
	return nil
}

func cmpUndefinedHandler(s *service) error {
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
