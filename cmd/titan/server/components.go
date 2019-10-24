package server

import (
	"fmt"

	"github.com/atlaskerr/titan/http/blob"
	"github.com/atlaskerr/titan/http/live"
	"github.com/atlaskerr/titan/http/manifest"
	"github.com/atlaskerr/titan/http/metrics"
	"github.com/atlaskerr/titan/http/oci"
	"github.com/atlaskerr/titan/http/ready"
	"github.com/atlaskerr/titan/http/tag"
	"github.com/atlaskerr/titan/http/titan"
	"github.com/atlaskerr/titan/http/undefined"
	titanMetrics "github.com/atlaskerr/titan/metrics"
	core "github.com/atlaskerr/titan/titan"

	opentracing "github.com/opentracing/opentracing-go"
)

type component func(*service) error

type service struct {
	collector *titanMetrics.Collector
	tracer    opentracing.Tracer
	core      *core.Core
	handlers  handlers
}

type handlers struct {
	titan     *titan.Server
	manifest  *manifest.Server
	blob      *blob.Server
	tag       *tag.Server
	oci       *oci.Server
	metrics   *metrics.Server
	live      *live.Server
	ready     *ready.Server
	undefined *undefined.Server
}

var baseComponents = []component{
	cmpCollector,
	cmpTracer,
	cmpCore,
}

var handlerComponents = []component{
	cmpUndefinedHandler,
	cmpMetricsHandler,
	cmpLiveHandler,
	cmpReadyHandler,
	cmpTagHandler,
	cmpManifestHandler,
	cmpBlobHandler,
	cmpOCIHandler,
	cmpTitanHandler,
}

func newService() (*service, error) {
	s := &service{}
	// TODO(atlaskerr): Find a cleaner way to do this. Eventually there will be
	// need to start and stop components running in goroutines and whatnot.
	var components []component
	components = append(components, baseComponents...)
	components = append(components, handlerComponents...)
	for _, component := range components {
		err := component(s)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize component: %s", err)
		}
	}
	return s, nil
}
