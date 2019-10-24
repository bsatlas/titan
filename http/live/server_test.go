package live_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/atlaskerr/titan/http/live"
	"github.com/atlaskerr/titan/http/live/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
	opentracing "github.com/opentracing/opentracing-go"
)

type serverTestComponents struct {
	server           *live.Server
	core             *mock.Liveness
	tracer           opentracing.Tracer
	undefinedHandler *mock.Handler
	responseWriter   *mock.ResponseWriter
	types            serverTestMatchers
}

type serverTestMatchers struct {
	request        gomock.Matcher
	responseWriter gomock.Matcher
	context        gomock.Matcher
}

func setupServerTestComponents(t *testing.T,
	ctrl *gomock.Controller) serverTestComponents {
	t.Helper()
	cmp := serverTestComponents{
		core:             mock.NewLiveness(ctrl),
		tracer:           opentracing.NoopTracer{},
		undefinedHandler: mock.NewHandler(ctrl),
		responseWriter:   mock.NewResponseWriter(ctrl),
		types: serverTestMatchers{
			request:        gomock.AssignableToTypeOf(new(http.Request)),
			responseWriter: gomock.Any(),
			context:        gomock.Any(),
		},
	}
	opts := []live.ServerOption{
		live.OptionMetricsCollector(metrics.NewCollector()),
		live.OptionUndefinedHandler(cmp.undefinedHandler),
		live.OptionCore(cmp.core),
		live.OptionTracer(cmp.tracer),
	}
	server, err := live.NewServer(opts...)
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
	cmp.undefinedHandler.EXPECT().ServeHTTP(
		cmp.types.responseWriter, cmp.types.request,
	)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerLiveOK(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	gomock.InOrder(
		cmp.core.EXPECT().Live(cmp.types.context).Return(true),
		cmp.responseWriter.EXPECT().WriteHeader(200),
	)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerLiveNotOK(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	gomock.InOrder(
		cmp.core.EXPECT().Live(cmp.types.context).Return(false),
		cmp.responseWriter.EXPECT().WriteHeader(404),
	)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}
