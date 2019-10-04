package titan_test

import (
	"testing"

	"github.com/atlaskerr/titan/http/titan"
	"github.com/atlaskerr/titan/http/titan/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

func TestNewServerNoUndefinedHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []titan.ServerOption{
		titan.OptionOCIHandler(mock.NewHandler(ctrl)),
		titan.OptionMetricsHandler(mock.NewHandler(ctrl)),
		titan.OptionLiveHandler(mock.NewHandler(ctrl)),
		titan.OptionReadyHandler(mock.NewHandler(ctrl)),
		titan.OptionMetricsCollector(metrics.NewCollector()),
	}
	_, err := titan.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoOCIHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []titan.ServerOption{
		titan.OptionMetricsHandler(mock.NewHandler(ctrl)),
		titan.OptionLiveHandler(mock.NewHandler(ctrl)),
		titan.OptionReadyHandler(mock.NewHandler(ctrl)),
		titan.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		titan.OptionMetricsCollector(metrics.NewCollector()),
	}
	_, err := titan.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoMetricsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []titan.ServerOption{
		titan.OptionOCIHandler(mock.NewHandler(ctrl)),
		titan.OptionLiveHandler(mock.NewHandler(ctrl)),
		titan.OptionReadyHandler(mock.NewHandler(ctrl)),
		titan.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		titan.OptionMetricsCollector(metrics.NewCollector()),
	}
	_, err := titan.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoReadyHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []titan.ServerOption{
		titan.OptionOCIHandler(mock.NewHandler(ctrl)),
		titan.OptionLiveHandler(mock.NewHandler(ctrl)),
		titan.OptionMetricsHandler(mock.NewHandler(ctrl)),
		titan.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		titan.OptionMetricsCollector(metrics.NewCollector()),
	}
	_, err := titan.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoLiveHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []titan.ServerOption{
		titan.OptionOCIHandler(mock.NewHandler(ctrl)),
		titan.OptionMetricsHandler(mock.NewHandler(ctrl)),
		titan.OptionReadyHandler(mock.NewHandler(ctrl)),
		titan.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		titan.OptionMetricsCollector(metrics.NewCollector()),
	}
	_, err := titan.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoMetricsCollector(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []titan.ServerOption{
		titan.OptionOCIHandler(mock.NewHandler(ctrl)),
		titan.OptionLiveHandler(mock.NewHandler(ctrl)),
		titan.OptionReadyHandler(mock.NewHandler(ctrl)),
		titan.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		titan.OptionMetricsHandler(mock.NewHandler(ctrl)),
	}
	_, err := titan.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
