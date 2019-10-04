package metrics

import (
	"compress/gzip"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/atlaskerr/titan/metrics"
	httpmetrics "github.com/atlaskerr/titan/metrics/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
)

// Server is titan's metrics endpoint.
type Server struct {
	gatherer prometheus.Gatherer
	handlers handlers
	metrics  *metrics.Collector
}

type handlers struct {
	undefined http.Handler
}

var endpointLabel prometheus.Labels = map[string]string{
	"endpoint": "metrics",
}

var gzipPool = sync.Pool{
	New: func() interface{} {
		return gzip.NewWriter(nil)
	},
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	requestDurationStart := time.Now()
	s.metrics.HTTP.RequestsInFlight.With(endpointLabel).Inc()
	defer func() {
		s.metrics.HTTP.RequestsInFlight.With(endpointLabel).Dec()
	}()
	requestSize := httpmetrics.ComputeRequestSize(req)
	s.metrics.HTTP.RequestSize.With(endpointLabel).Observe(requestSize)
	var requestLabels prometheus.Labels = map[string]string{
		"endpoint": "metrics",
	}
	var statusCode int
	if req.URL.Path != "/" {
		s.handlers.undefined.ServeHTTP(w, req)
		return
	}
	if req.Method != "GET" {
		s.handlers.undefined.ServeHTTP(w, req)
		return
	}
	mfs, err := s.gatherer.Gather()
	if err != nil {
		// Ignore the error if the registry returns something
		if len(mfs) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	header := w.Header()
	if gzipAccepted(req.Header) {
		tmpw := w.(io.Writer)
		header.Set("Content-Encoding", "gzip")
		gz := gzipPool.Get().(*gzip.Writer)
		defer gzipPool.Put(gz)

		gz.Reset(tmpw)
		defer gz.Close()

		tmpw = gz
	}
	contentType := expfmt.Negotiate(req.Header)
	enc := expfmt.NewEncoder(w, contentType)
	for _, mf := range mfs {
		if err := enc.Encode(mf); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	header.Set("Content-Type", string(contentType))
	statusCode = 200
	w.WriteHeader(statusCode)
	requestLabels["code"] = strconv.Itoa(statusCode)
	s.metrics.HTTP.TotalRequests.With(requestLabels).Inc()
	requestDuration := time.Since(requestDurationStart).Seconds()
	s.metrics.HTTP.RequestDuration.With(requestLabels).Observe(requestDuration)
}

// gzipAccepted returns whether the client will accept gzip-encoded content.
func gzipAccepted(header http.Header) bool {
	a := header.Get("Content-Encoding")
	parts := strings.Split(a, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "gzip" || strings.HasPrefix(part, "gzip;") {
			return true
		}
	}
	return false
}
