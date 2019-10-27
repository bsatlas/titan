package upload

import (
	"net/http"
	//	"github.com/atlaskerr/titan/metrics"
)

// StatusServer implements the blob upload completion endpoint.
type StatusServer struct {
	handlers statusServerHandlers
	//	metrics  *metrics.Collector
}

type statusServerHandlers struct {
	undefined http.Handler
}

func (s *StatusServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.handlers.undefined.ServeHTTP(w, req)
}
