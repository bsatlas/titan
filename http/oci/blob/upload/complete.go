package upload

import (
	"net/http"
	//	"github.com/atlaskerr/titan/metrics"
)

// CompleteServer implements the blob upload completion endpoint.
type CompleteServer struct {
	handlers completeServerHandlers
	//	metrics  *metrics.Collector
}

type completeServerHandlers struct {
	undefined http.Handler
}

func (s *CompleteServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.handlers.undefined.ServeHTTP(w, req)
}
