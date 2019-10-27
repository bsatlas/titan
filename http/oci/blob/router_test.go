package blob_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/atlaskerr/titan/http/oci/blob"
	"github.com/atlaskerr/titan/http/oci/blob/internal/mock"

	"github.com/golang/mock/gomock"
)

type routerTestComponents struct {
	router           *blob.Router
	undefinedHandler *mock.Handler
	downloadHandler  *mock.Handler
	existHandler     *mock.Handler
	removeHandler    *mock.Handler
	uploadHandler    *mock.Handler
	responseWriter   *mock.ResponseWriter
	types            routerTestMatchers
}

type routerTestMatchers struct {
	request        gomock.Matcher
	responseWriter gomock.Matcher
}

func setupRouterTestComponents(t *testing.T,
	ctrl *gomock.Controller) routerTestComponents {
	t.Helper()
	cmp := routerTestComponents{
		undefinedHandler: mock.NewHandler(ctrl),
		downloadHandler:  mock.NewHandler(ctrl),
		existHandler:     mock.NewHandler(ctrl),
		removeHandler:    mock.NewHandler(ctrl),
		uploadHandler:    mock.NewHandler(ctrl),
		responseWriter:   mock.NewResponseWriter(ctrl),
		types: routerTestMatchers{
			request:        gomock.AssignableToTypeOf(new(http.Request)),
			responseWriter: gomock.Any(),
		},
	}
	opts := []blob.RouterOption{
		blob.OptionUndefinedHandler(cmp.undefinedHandler),
		blob.OptionDownloadHandler(cmp.downloadHandler),
		blob.OptionExistHandler(cmp.existHandler),
		blob.OptionRemoveHandler(cmp.removeHandler),
		blob.OptionUploadHandler(cmp.uploadHandler),
	}
	router, err := blob.NewRouter(opts...)
	if err != nil {
		t.Fatal(err)
	}
	cmp.router = router
	return cmp
}

func TestRouterUpload(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/sha246:41af286dc0b172ed2f1ca934fd2278de4a1192302ffa07087cea2682e7d372e3/c1f6a2a0-24ce-4b34-9237-d4e2dea1ee27",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupRouterTestComponents(t, ctrl)
	cmp.uploadHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.router.ServeHTTP(cmp.responseWriter, request)
}

func TestRouterMethodGet(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/sha246:41af286dc0b172ed2f1ca934fd2278de4a1192302ffa07087cea2682e7d372e3",
		},
		Method: "GET",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupRouterTestComponents(t, ctrl)
	cmp.downloadHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.router.ServeHTTP(cmp.responseWriter, request)
}

func TestRouterMethodHead(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/sha246:41af286dc0b172ed2f1ca934fd2278de4a1192302ffa07087cea2682e7d372e3",
		},
		Method: "HEAD",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupRouterTestComponents(t, ctrl)
	cmp.existHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.router.ServeHTTP(cmp.responseWriter, request)
}

func TestRouterMethodDelete(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/sha246:41af286dc0b172ed2f1ca934fd2278de4a1192302ffa07087cea2682e7d372e3",
		},
		Method: "DELETE",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupRouterTestComponents(t, ctrl)
	cmp.removeHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.router.ServeHTTP(cmp.responseWriter, request)
}

func TestRouterMethodUndefined(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/sha246:41af286dc0b172ed2f1ca934fd2278de4a1192302ffa07087cea2682e7d372e3",
		},
		Method: "PUT",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupRouterTestComponents(t, ctrl)
	cmp.undefinedHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.router.ServeHTTP(cmp.responseWriter, request)
}
