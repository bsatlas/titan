package oci

import (
	"github.com/atlaskerr/titan/metrics"
	"net/http"
)

// Server is titan's OCI endpoint.
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
