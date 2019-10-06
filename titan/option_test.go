package titan_test

import (
	"testing"

	"github.com/atlaskerr/titan/titan"
)

func TestNewCoreNoMetricsCollector(t *testing.T) {
	opts := []titan.CoreOption{}
	_, err := titan.NewCore(opts...)
	if err == nil {
		t.Fail()
	}
}
