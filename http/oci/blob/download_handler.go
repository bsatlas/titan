package blob

import (
	"net/http"
	//	"github.com/atlaskerr/titan/metrics"
)

// DownloadHandler implements the blob download endpoint.
type DownloadHandler struct {
	undefined http.Handler
}

func (h *DownloadHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.undefined.ServeHTTP(w, req)
}
