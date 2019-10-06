package server

import (
	"fmt"

	"github.com/atlaskerr/titan/titan"
)

func cmpCore(s *service) error {
	opts := []titan.CoreOption{
		titan.OptionMetricsCollector(s.collector),
	}
	c, err := titan.NewCore(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize core: %s", err)
	}
	s.core = c
	return nil
}
