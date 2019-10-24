package server

import (
	titanMetrics "github.com/atlaskerr/titan/metrics"

	opentracing "github.com/opentracing/opentracing-go"
)

func cmpCollector(s *service) error {
	c := titanMetrics.NewCollector()
	s.collector = c
	return nil
}

func cmpTracer(s *service) error {
	s.tracer = opentracing.NoopTracer{}
	return nil
}
