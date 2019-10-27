package manifest_test

import (
	"testing"

	"github.com/atlaskerr/titan/http/oci/manifest"
	"github.com/atlaskerr/titan/http/oci/manifest/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

func TestNewServerNoUndefinedHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []manifest.ServerOption{
		manifest.OptionMetricsCollector(metrics.NewCollector()),
	}
	_, err := manifest.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoMetricsCollector(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []manifest.ServerOption{
		manifest.OptionUndefinedHandler(mock.NewHandler(ctrl)),
	}
	_, err := manifest.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
