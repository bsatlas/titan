package oci_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/atlaskerr/titan/http/oci"
	"github.com/atlaskerr/titan/http/oci/internal/mock"
	"github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
)

type serverTestComponents struct {
	server           *oci.Server
	undefinedHandler *mock.Handler
	manifestHandler  *mock.Handler
	blobHandler      *mock.Handler
	tagHandler       *mock.Handler
	responseWriter   *mock.ResponseWriter
	types            serverTestMatchers
}

type serverTestMatchers struct {
	request        gomock.Matcher
	responseWriter gomock.Matcher
}

func setupServerTestComponents(t *testing.T,
	ctrl *gomock.Controller) serverTestComponents {
	t.Helper()
	cmp := serverTestComponents{
		undefinedHandler: mock.NewHandler(ctrl),
		manifestHandler:  mock.NewHandler(ctrl),
		blobHandler:      mock.NewHandler(ctrl),
		tagHandler:       mock.NewHandler(ctrl),
		responseWriter:   mock.NewResponseWriter(ctrl),
		types: serverTestMatchers{
			request:        gomock.AssignableToTypeOf(new(http.Request)),
			responseWriter: gomock.Any(),
		},
	}
	opts := []oci.ServerOption{
		oci.OptionMetricsCollector(metrics.NewCollector()),
		oci.OptionUndefinedHandler(cmp.undefinedHandler),
		oci.OptionManifestHandler(cmp.manifestHandler),
		oci.OptionBlobHandler(cmp.blobHandler),
		oci.OptionTagHandler(cmp.tagHandler),
	}
	server, err := oci.NewServer(opts...)
	if err != nil {
		t.Fatal(err)
	}
	cmp.server = server
	return cmp
}

func TestServerUndefinedPath(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/test-namespace/test-project/test-repo/undefined",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.undefinedHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerManifestPath(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/test-namespace/test-project/test-repo/manifests",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.manifestHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerBlobPath(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/test-namespace/test-project/test-repo/blobs",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.blobHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerTagPath(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/test-namespace/test-project/test-repo/tags",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.tagHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerNoNamespace(t *testing.T) {
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
	cmp.undefinedHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerNoProject(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/test-namespace",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.undefinedHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerNoRepo(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/test-namespace/test-project",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.undefinedHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerNoResource(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/test-namespace/test-project/test-repo",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.undefinedHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}
