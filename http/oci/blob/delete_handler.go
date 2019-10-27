package blob

import (
	"github.com/atlaskerr/titan/metrics"
	"net/http"
)

// DeleteHandler implements the blob deletion endpoint.
type DeleteHandler struct {
	undefined http.Handler
	metrics   *metrics.Collector
}

func (h *DeleteHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.undefined.ServeHTTP(w, req)
}
