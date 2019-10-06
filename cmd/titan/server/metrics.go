package server

import (
	titanMetrics "github.com/atlaskerr/titan/metrics"
)

func cmpCollector(s *service) error {
	c := titanMetrics.NewCollector()
	s.collector = c
	return nil
}
