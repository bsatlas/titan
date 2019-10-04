package ready_test

import (
	"testing"

	"github.com/atlaskerr/titan/http/ready"
	"github.com/atlaskerr/titan/http/ready/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

func TestNewServerNoCore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []ready.ServerOption{
		ready.OptionMetricsCollector(metrics.NewCollector()),
		ready.OptionUndefinedHandler(mock.NewHandler(ctrl)),
	}
	_, err := ready.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoUndefinedHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []ready.ServerOption{
		ready.OptionCore(mock.NewReadiness(ctrl)),
		ready.OptionMetricsCollector(metrics.NewCollector()),
	}
	_, err := ready.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoMetricsCollector(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []ready.ServerOption{
		ready.OptionCore(mock.NewReadiness(ctrl)),
		ready.OptionUndefinedHandler(mock.NewHandler(ctrl)),
	}
	_, err := ready.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
