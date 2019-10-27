package upload

import (
	"net/http"
	//	"github.com/atlaskerr/titan/metrics"
)

// CancelServer implements the blob upload cancellation endpoint.
type CancelServer struct {
	handlers cancelServerHandlers
	//	metrics  *metrics.Collector
}

type cancelServerHandlers struct {
	undefined http.Handler
}

func (s *CancelServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.handlers.undefined.ServeHTTP(w, req)
}
