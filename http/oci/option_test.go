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
		oci.OptionManifestHandler(mock.NewHandler(ctrl)),
		oci.OptionTagHandler(mock.NewHandler(ctrl)),
		oci.OptionBlobHandler(mock.NewHandler(ctrl)),
	}
	_, err := oci.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
func TestNewServerNoManifestHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []oci.ServerOption{
		oci.OptionMetricsCollector(metrics.NewCollector()),
		oci.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		oci.OptionTagHandler(mock.NewHandler(ctrl)),
		oci.OptionBlobHandler(mock.NewHandler(ctrl)),
	}
	_, err := oci.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
func TestNewServerNoBlobHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []oci.ServerOption{
		oci.OptionMetricsCollector(metrics.NewCollector()),
		oci.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		oci.OptionManifestHandler(mock.NewHandler(ctrl)),
		oci.OptionTagHandler(mock.NewHandler(ctrl)),
	}
	_, err := oci.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
func TestNewServerNoTagHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []oci.ServerOption{
		oci.OptionMetricsCollector(metrics.NewCollector()),
		oci.OptionUndefinedHandler(mock.NewHandler(ctrl)),
		oci.OptionManifestHandler(mock.NewHandler(ctrl)),
		oci.OptionBlobHandler(mock.NewHandler(ctrl)),
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
		oci.OptionManifestHandler(mock.NewHandler(ctrl)),
		oci.OptionTagHandler(mock.NewHandler(ctrl)),
		oci.OptionBlobHandler(mock.NewHandler(ctrl)),
	}
	_, err := oci.NewServer(opts...)
	if err == nil {
		t.Fail()
	}
}
