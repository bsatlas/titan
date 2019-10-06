package tag_test

import (
	"testing"

	"github.com/atlaskerr/titan/http/tag"
	"github.com/atlaskerr/titan/http/tag/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

func TestNewServerNoUndefinedHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []tag.ServerOption{
		tag.OptionMetricsCollector(metrics.NewCollector()),
	}
	_, err := tag.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoMetricsCollector(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []tag.ServerOption{
		tag.OptionUndefinedHandler(mock.NewHandler(ctrl)),
	}
	_, err := tag.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
