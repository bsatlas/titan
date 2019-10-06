package manifest

import (
	"net/http"

	"github.com/atlaskerr/titan/metrics"
)

// Server implements the OCI manifest endpoint.
// See: https://github.com/opencontainers/distribution-spec/blob/master/spec.md#manifest
type Server struct {
	handlers handlers
	metrics  *metrics.Collector
}

type handlers struct {
	undefined http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		s.handlers.undefined.ServeHTTP(w, req)
		return
	}
}
