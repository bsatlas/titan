package ready

import (
	"net/http"

	"github.com/atlaskerr/titan/metrics"
)

// Server is titan's ready endpoint.
type Server struct {
	handlers handlers
	metrics  *metrics.Collector
}

type handlers struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {}
