package ready_test

import (
	"testing"

	"github.com/atlaskerr/titan/http/health/ready"
	"github.com/atlaskerr/titan/http/health/ready/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
	opentracing "github.com/opentracing/opentracing-go"
)

func TestNewServerNoCore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []ready.ServerOption{
		ready.OptionMetricsCollector(metrics.NewCollector()),
		ready.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		ready.OptionTracer(opentracing.NoopTracer{}),
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
		ready.OptionTracer(opentracing.NoopTracer{}),
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
		ready.OptionTracer(opentracing.NoopTracer{}),
	}
	_, err := ready.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoTracer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []ready.ServerOption{
		ready.OptionCore(mock.NewReadiness(ctrl)),
		ready.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		ready.OptionMetricsCollector(metrics.NewCollector()),
	}
	_, err := ready.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
