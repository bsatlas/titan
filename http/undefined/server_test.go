package undefined_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/atlaskerr/titan/http/undefined"
	"github.com/atlaskerr/titan/http/undefined/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

type serverTestComponents struct {
	server         *undefined.Server
	responseWriter *mock.ResponseWriter
	types          serverTestMatchers
}

type serverTestMatchers struct {
	request        gomock.Matcher
	responseWriter gomock.Matcher
}

func setupServerTestComponents(t *testing.T,
	ctrl *gomock.Controller) serverTestComponents {
	t.Helper()
	cmp := serverTestComponents{
		responseWriter: mock.NewResponseWriter(ctrl),
		types: serverTestMatchers{
			request:        gomock.AssignableToTypeOf(new(http.Request)),
			responseWriter: gomock.Any(),
		},
	}
	opts := []undefined.ServerOption{
		undefined.OptionMetricsCollector(metrics.NewCollector()),
	}
	server, err := undefined.NewServer(opts...)
	if err != nil {
		t.Fatal(err)
	}
	cmp.server = server
	return cmp
}

func TestServerUndefinedPath(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/unknown/path",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}
