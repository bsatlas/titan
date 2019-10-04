package ready

import (
	"net/http"
	"strconv"
	"time"

	"github.com/atlaskerr/titan/metrics"
	httpmetrics "github.com/atlaskerr/titan/metrics/http"

	"github.com/prometheus/client_golang/prometheus"
)

// Server is titan's ready endpoint.
type Server struct {
	core     Readiness
	handlers handlers
	metrics  *metrics.Collector
}

type handlers struct {
	undefined http.Handler
}

// Readiness defines the method for checking server readiness.
type Readiness interface {
	Ready() bool
}

var endpointLabel prometheus.Labels = map[string]string{
	"endpoint": "ready",
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
		"endpoint": "ready",
	}
	if req.URL.Path != "/" {
		s.handlers.undefined.ServeHTTP(w, req)
		return
	}
	ready := s.core.Ready()
	var statusCode int
	if ready {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusNotFound
	}
	w.WriteHeader(statusCode)
	requestLabels["code"] = strconv.Itoa(statusCode)
	s.metrics.HTTP.TotalRequests.With(requestLabels).Inc()
	requestDuration := time.Since(requestDurationStart).Seconds()
	s.metrics.HTTP.RequestDuration.With(requestLabels).Observe(requestDuration)
}
