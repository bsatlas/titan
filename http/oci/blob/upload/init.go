package upload

import (
	"net/http"
	//	"github.com/atlaskerr/titan/metrics"
)

// InitServer implements the blob upload completion endpoint.
type InitServer struct {
	handlers initServerHandlers
	//	metrics  *metrics.Collector
}

type initServerHandlers struct {
	undefined http.Handler
}

func (s *InitServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.handlers.undefined.ServeHTTP(w, req)
}
