package http

import (
	"net/http"
)

// ComputeRequestSize returns the size of an http.Request. Used for recording
// Collector.RequestSize.
func ComputeRequestSize(req *http.Request) float64 {
	size := 0
	if req.URL != nil {
		size += len(req.URL.String())
	}
	size += len(req.Method)
	size += len(req.Proto)
	size += len(req.Host)
	for headerKey, headerValues := range req.Header {
		size += len(headerKey)
		for _, headerValue := range headerValues {
			size += len(headerValue)
		}
	}
	if req.ContentLength != -1 {
		size += int(req.ContentLength)
	}
	return float64(size)
}
