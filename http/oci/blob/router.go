package blob

import (
	"net/http"
	"path"
	"strings"

	titancontext "github.com/atlaskerr/titan/context"
)

// Router implements the OCI blobs endpoint.
// See: https://github.com/opencontainers/distribution-spec/blob/master/spec.md#blob
type Router struct {
	download  http.Handler
	exist     http.Handler
	remove    http.Handler
	upload    http.Handler
	undefined http.Handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	var digest string
	digest, req.URL.Path = shift(req.URL.Path)
	ctx = titancontext.WithBlobDigest(ctx, digest)
	req = req.WithContext(ctx)
	var handler http.Handler
	if req.URL.Path != "/" {
		var uploadID string
		uploadID, req.URL.Path = shift(req.URL.Path)
		ctx = titancontext.WithBlobUploadID(ctx, uploadID)
		req = req.WithContext(ctx)
		r.upload.ServeHTTP(w, req)
		return
	}
	switch req.Method {
	case "GET":
		handler = r.download
	case "HEAD":
		handler = r.exist
	case "DELETE":
		handler = r.remove
	default:
		handler = r.undefined
	}
	handler.ServeHTTP(w, req)
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
