package titan

import (
	"net/http"
	"path"
	"strings"

	"github.com/atlaskerr/titan/metrics"
)

// Server is the main titan HTTP server.
type Server struct {
	handlers handlers
	metrics  *metrics.Collector
}

type handlers struct {
	metrics   http.Handler
	live      http.Handler
	ready     http.Handler
	oci       http.Handler
	undefined http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var resource string
	var handler http.Handler
	resource, req.URL.Path = shift(req.URL.Path)
	switch resource {
	case "v2":
		handler = s.handlers.oci
		handler.ServeHTTP(w, req)
		return
	case "metrics":
		handler = s.handlers.metrics
		handler.ServeHTTP(w, req)
		return
	case "live":
		handler = s.handlers.live
		handler.ServeHTTP(w, req)
		return
	case "ready":
		handler = s.handlers.ready
		handler.ServeHTTP(w, req)
		return
	default:
		handler = s.handlers.undefined
		handler.ServeHTTP(w, req)
		return
	}
}

// shift is a helper function for routing HTTP requests.
func shift(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
