package http_test

import (
	"net/http"
	"net/url"
	"testing"

	httpmetrics "github.com/atlaskerr/titan/metrics/http"
)

func TestRecordRequestSize(t *testing.T) {
	expected := 0
	var header http.Header = map[string][]string{
		"foo": []string{"bar", "baz"},
	}
	for headerKey, headerValues := range header {
		expected += len(headerKey)
		for _, headerValue := range headerValues {
			expected += len(headerValue)
		}
	}

	rURL := &url.URL{
		Scheme: "http",
		Host:   "localhost",
		Path:   "/foo",
	}
	expected += len(rURL.String())
	method := "GET"
	expected += len(method)
	proto := "HTTP/1.1"
	expected += len(proto)
	host := "localhost"
	expected += len(host)
	contentLength := 100
	expected += contentLength
	req := &http.Request{
		Method:        method,
		URL:           rURL,
		Proto:         proto,
		Host:          host,
		Header:        header,
		ContentLength: int64(contentLength),
	}
	got := httpmetrics.ComputeRequestSize(req)
	if got != float64(expected) {
		t.Fail()
	}
}
