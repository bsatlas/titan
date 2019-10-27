package upload

import (
	"net/http"
	//	"github.com/atlaskerr/titan/metrics"
)

// ChunkServer implements the blob upload cancellation endpoint.
type ChunkServer struct {
	handlers chunkServerHandlers
	//	metrics  *metrics.Collector
}

type chunkServerHandlers struct {
	undefined http.Handler
}

func (s *ChunkServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.handlers.undefined.ServeHTTP(w, req)
}
