package titan_test

import (
	"testing"

	"github.com/atlaskerr/titan/metrics"
	"github.com/atlaskerr/titan/titan"

	"github.com/golang/mock/gomock"
)

type coreTestComponents struct {
	core *titan.Core
}

func setupCoreTestComponents(t *testing.T,
	ctrl *gomock.Controller) coreTestComponents {
	t.Helper()
	cmp := coreTestComponents{}
	opts := []titan.CoreOption{
		titan.OptionMetricsCollector(metrics.NewCollector()),
	}
	core, err := titan.NewCore(opts...)
	if err != nil {
		t.Fatal(err)
	}
	cmp.core = core
	return cmp
}
