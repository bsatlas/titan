package blob

import (
	"net/http"
	//	"github.com/atlaskerr/titan/metrics"
)

// ExistHandler implements the blob existence check endpoint.
type ExistHandler struct {
	undefined http.Handler
	//	metrics  *metrics.Collector
}

func (h *ExistHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.undefined.ServeHTTP(w, req)
}
