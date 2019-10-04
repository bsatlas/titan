package live_test

import (
	"testing"

	"github.com/atlaskerr/titan/http/live"
	"github.com/atlaskerr/titan/http/live/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

func TestNewServerNoCore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []live.ServerOption{
		live.OptionMetricsCollector(metrics.NewCollector()),
		live.OptionUndefinedHandler(mock.NewHandler(ctrl)),
	}
	_, err := live.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoUndefinedHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []live.ServerOption{
		live.OptionMetricsCollector(metrics.NewCollector()),
		live.OptionCore(mock.NewLiveness(ctrl)),
	}
	_, err := live.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoMetricsCollector(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []live.ServerOption{
		live.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		live.OptionCore(mock.NewLiveness(ctrl)),
	}
	_, err := live.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
