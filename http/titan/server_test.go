package titan_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/atlaskerr/titan/http/titan"
	"github.com/atlaskerr/titan/http/titan/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

type serverTestComponents struct {
	server         *titan.Server
	ociHandler     *mock.Handler
	metricsHandler *mock.Handler
	readyHandler   *mock.Handler
	liveHandler    *mock.Handler
	unknownHandler *mock.Handler
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
		ociHandler:     mock.NewHandler(ctrl),
		metricsHandler: mock.NewHandler(ctrl),
		readyHandler:   mock.NewHandler(ctrl),
		liveHandler:    mock.NewHandler(ctrl),
		unknownHandler: mock.NewHandler(ctrl),
		responseWriter: mock.NewResponseWriter(ctrl),
		types: serverTestMatchers{
			request:        gomock.AssignableToTypeOf(new(http.Request)),
			responseWriter: gomock.Any(),
		},
	}
	opts := []titan.ServerOption{
		titan.OptionMetricsCollector(metrics.NewCollector()),
		titan.OptionOCIHandler(cmp.ociHandler),
		titan.OptionLiveHandler(cmp.liveHandler),
		titan.OptionReadyHandler(cmp.readyHandler),
		titan.OptionMetricsHandler(cmp.metricsHandler),
		titan.OptionUndefinedHandler(cmp.unknownHandler),
	}
	server, err := titan.NewServer(opts...)
	if err != nil {
		t.Fatal(err)
	}
	cmp.server = server
	return cmp
}

func TestServerUnknownPath(t *testing.T) {
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
	cmp.unknownHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerOCIPath(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/v2",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.ociHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerMetricsPath(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/metrics",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.metricsHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerReadyPath(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/ready",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.readyHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerLivePath(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/live",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.liveHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}
