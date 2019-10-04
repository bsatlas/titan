package undefined_test

import (
	"testing"

	"github.com/atlaskerr/titan/http/undefined"
)

func TestNewServerNoMetricsCollector(t *testing.T) {
	opts := []undefined.ServerOption{}
	_, err := undefined.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
