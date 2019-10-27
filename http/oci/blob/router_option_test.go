package blob_test

import (
	"testing"

	"github.com/atlaskerr/titan/http/oci/blob"
	//	"github.com/atlaskerr/titan/http/oci/blob/internal/mock"

	"github.com/golang/mock/gomock"
)

func TestNewServerNoUndefinedHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := []blob.RouterOption{}
	_, err := blob.NewRouter(opts...)
	if err == nil {
		t.Fail()
	}
}

//func TestNewServerNoMetricsCollector(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	opts := []blob.RouterOption{
//		blob.OptionUndefinedHandler(mock.NewHandler(ctrl)),
//	}
//	_, err := blob.NewRouter(opts...)
//	if err == nil {
//		t.Fail()
//	}
//}
