package oci

import (
	"net/http"
	"path"
	"strings"

	titancontext "github.com/atlaskerr/titan/context"
	"github.com/atlaskerr/titan/metrics"
)

// Server is titan's OCI endpoint.
type Server struct {
	handlers handlers
	metrics  *metrics.Collector
}

type handlers struct {
	undefined http.Handler
	manifest  http.Handler
	blob      http.Handler
	tag       http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		s.handlers.undefined.ServeHTTP(w, req)
		return
	}
	ctx := req.Context()
	var namespace string
	namespace, req.URL.Path = shiftPath(req.URL.Path)
	if req.URL.Path == "/" {
		s.handlers.undefined.ServeHTTP(w, req)
		return
	}
	ctx = titancontext.WithNamespace(ctx, namespace)
	var project string
	project, req.URL.Path = shiftPath(req.URL.Path)
	if req.URL.Path == "/" {
		s.handlers.undefined.ServeHTTP(w, req)
		return
	}
	ctx = titancontext.WithProject(ctx, project)
	var repo string
	repo, req.URL.Path = shiftPath(req.URL.Path)
	if req.URL.Path == "/" {
		s.handlers.undefined.ServeHTTP(w, req)
		return
	}
	ctx = titancontext.WithRepo(ctx, repo)
	req = req.WithContext(ctx)
	var handler http.Handler
	var resource string
	resource, req.URL.Path = shiftPath(req.URL.Path)
	switch resource {
	case "blobs":
		handler = s.handlers.blob
		handler.ServeHTTP(w, req)
		return
	case "manifests":
		handler = s.handlers.manifest
		handler.ServeHTTP(w, req)
		return
	case "tags":
		handler = s.handlers.tag
		handler.ServeHTTP(w, req)
		return
	default:
		handler = s.handlers.undefined
		handler.ServeHTTP(w, req)
		return
	}
}

// shiftPath is a helper function for routing HTTP requests.
func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
