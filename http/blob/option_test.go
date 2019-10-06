package blob_test

import (
	"testing"

	"github.com/atlaskerr/titan/http/blob"
	"github.com/atlaskerr/titan/http/blob/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

func TestNewServerNoUndefinedHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []blob.ServerOption{
		blob.OptionMetricsCollector(metrics.NewCollector()),
	}
	_, err := blob.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoMetricsCollector(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []blob.ServerOption{
		blob.OptionUndefinedHandler(mock.NewHandler(ctrl)),
	}
	_, err := blob.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
