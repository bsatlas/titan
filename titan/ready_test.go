package titan_test

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCoreReadyOK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupCoreTestComponents(t, ctrl)
	ok := cmp.core.Ready()
	if !ok {
		t.Fail()
	}
}
