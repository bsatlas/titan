package upload

import (
	"net/http"
	//	"github.com/atlaskerr/titan/metrics"
)

// Server implements the OCI blobs endpoint.
// See: https://github.com/opencontainers/distribution-spec/blob/master/spec.md#blob
type Server struct {
	handlers handlers
	//	metrics  *metrics.Collector
}

type handlers struct {
	begin     http.Handler
	status    http.Handler
	chunk     http.Handler
	cancel    http.Handler
	complete  http.Handler
	undefined http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		s.handlers.undefined.ServeHTTP(w, req)
		return
	}
	var handler http.Handler
	switch req.Method {
	case "GET":
		handler = s.handlers.status
	case "PATCH":
		handler = s.handlers.chunk
	case "DELETE":
		handler = s.handlers.cancel
	case "PUT":
		handler = s.handlers.complete
	case "POST":
		handler = s.handlers.begin
	default:
		handler = s.handlers.undefined
	}
	handler.ServeHTTP(w, req)
}

// shift is a helper function for routing HTTP requests.
//func shift(p string) (head, tail string) {
//	p = path.Clean("/" + p)
//	i := strings.Index(p[1:], "/") + 1
//	if i <= 0 {
//		return p[1:], "/"
//	}
//	return p[1:i], p[i:]
//}
