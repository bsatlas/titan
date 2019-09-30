package unknown

import (
	"net/http"

	"github.com/atlaskerr/titan/metrics"
)

// Server is titan's endpoint for unknown requests.
type Server struct {
	handlers handlers
	metrics  *metrics.Collector
}

type handlers struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {}
