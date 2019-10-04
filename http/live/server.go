package live

import (
	"net/http"
	"strconv"
	"time"

	"github.com/atlaskerr/titan/metrics"
	httpmetrics "github.com/atlaskerr/titan/metrics/http"

	"github.com/prometheus/client_golang/prometheus"
)

// Server is titan's liveness endpoint.
type Server struct {
	core     Liveness
	handlers handlers
	metrics  *metrics.Collector
}

type handlers struct {
	undefined http.Handler
}

// Liveness defines the method for checking server liveness.
type Liveness interface {
	Live() bool
}

var endpointLabel prometheus.Labels = map[string]string{
	"endpoint": "live",
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
		"endpoint": "live",
	}
	var statusCode int
	if req.URL.Path != "/" {
		s.handlers.undefined.ServeHTTP(w, req)
		return
	}
	ok := s.core.Live()
	if ok {
		statusCode = 200
	} else {
		statusCode = 404
	}
	w.WriteHeader(statusCode)
	requestLabels["code"] = strconv.Itoa(statusCode)
	s.metrics.HTTP.TotalRequests.With(requestLabels).Inc()
	requestDuration := time.Since(requestDurationStart).Seconds()
	s.metrics.HTTP.RequestDuration.With(requestLabels).Observe(requestDuration)
}
