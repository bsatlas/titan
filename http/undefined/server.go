package undefined

import (
	"net/http"
	"strconv"

	"github.com/atlaskerr/titan/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

// Server is titan's endpoint for undefined requests.
type Server struct {
	metrics *metrics.Collector
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var statusCode int
	switch req.Method {
	case "GET", "HEAD":
		statusCode = 404
	default:
		statusCode = 501
	}
	var requestLabels prometheus.Labels = map[string]string{
		"endpoint": "undefined",
	}
	w.WriteHeader(statusCode)
	requestLabels["code"] = strconv.Itoa(statusCode)
	s.metrics.HTTP.TotalRequests.With(requestLabels).Inc()
}
