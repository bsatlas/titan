package metrics_test

import (
	"testing"

	"github.com/atlaskerr/titan/http/metrics"
	"github.com/atlaskerr/titan/http/metrics/internal/mock"
	titanMetrics "github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

func TestNewServerNoCore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []metrics.ServerOption{
		metrics.OptionMetricsCollector(titanMetrics.NewCollector()),
		metrics.OptionUndefinedHandler(mock.NewHandler(ctrl)),
	}
	_, err := metrics.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoUndefinedHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []metrics.ServerOption{
		metrics.OptionMetricsCollector(titanMetrics.NewCollector()),
		metrics.OptionCore(mock.NewGatherer(ctrl)),
	}
	_, err := metrics.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoMetricsCollector(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []metrics.ServerOption{
		metrics.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		metrics.OptionCore(mock.NewGatherer(ctrl)),
	}
	_, err := metrics.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
