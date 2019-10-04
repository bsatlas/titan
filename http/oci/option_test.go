package oci_test

import (
	"testing"

	"github.com/atlaskerr/titan/http/oci"
	"github.com/atlaskerr/titan/http/oci/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

func TestNewServerNoUndefinedHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []oci.ServerOption{
		oci.OptionMetricsCollector(metrics.NewCollector()),
	}
	_, err := oci.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}

func TestNewServerNoMetricsCollector(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []oci.ServerOption{
		oci.OptionUndefinedHandler(mock.NewHandler(ctrl)),
	}
	_, err := oci.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
