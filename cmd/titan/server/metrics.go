package server

import (
	titanMetrics "github.com/atlaskerr/titan/metrics"
)

func collector(s *service) error {
	c := titanMetrics.NewCollector()
	s.collector = c
	return nil
}
