package metrics_test

import (
	"bytes"
	"compress/gzip"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/atlaskerr/titan/http/metrics"
	"github.com/atlaskerr/titan/http/metrics/internal/mock"
	titanMetrics "github.com/atlaskerr/titan/metrics"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

type serverTestComponents struct {
	server           *metrics.Server
	core             *mock.Gatherer
	undefinedHandler *mock.Handler
	responseWriter   *mock.ResponseWriter
	responseHeader   map[string][]string
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
		core:             mock.NewGatherer(ctrl),
		undefinedHandler: mock.NewHandler(ctrl),
		responseWriter:   mock.NewResponseWriter(ctrl),
		responseHeader:   map[string][]string{},
		types: serverTestMatchers{
			request:        gomock.AssignableToTypeOf(new(http.Request)),
			responseWriter: gomock.Any(),
		},
	}
	opts := []metrics.ServerOption{
		metrics.OptionCore(cmp.core),
		metrics.OptionMetricsCollector(titanMetrics.NewCollector()),
		metrics.OptionUndefinedHandler(cmp.undefinedHandler),
	}
	server, err := metrics.NewServer(opts...)
	if err != nil {
		t.Fatal(err)
	}
	cmp.server = server
	return cmp
}

var defaultMetricFamilies = []*dto.MetricFamily{
	{
		Name: proto.String("test"),
		Help: proto.String("test description"),
		Type: dto.MetricType_COUNTER.Enum(),
		Metric: []*dto.Metric{
			{
				Label: []*dto.LabelPair{
					{
						Name:  proto.String("label_one"),
						Value: proto.String("label_value_one"),
					},
				},
				Counter: &dto.Counter{
					Value: proto.Float64(1),
				},
			},
		},
	},
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
	cmp.undefinedHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerMethodNotGet(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/",
		},
		Method: "PUT",
		Body:   nil,
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	cmp.undefinedHandler.EXPECT().ServeHTTP(cmp.types.responseWriter, cmp.types.request)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerGatherError(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/",
		},
		Method: "GET",
		Body:   nil,
	}
	err := errors.New("error")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	gomock.InOrder(
		cmp.core.EXPECT().Gather().Return(nil, err),
		cmp.responseWriter.EXPECT().WriteHeader(500),
	)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerHappy(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/",
		},
		Header: map[string][]string{
			"Accept": []string{"text/plain;version=0.0.4"},
		},
		Method: "GET",
		Body:   nil,
	}
	mfs := defaultMetricFamilies
	buf := &bytes.Buffer{}
	enc := expfmt.NewEncoder(buf, expfmt.FmtText)
	for _, mf := range mfs {
		err := enc.Encode(mf)
		if err != nil {
			t.Fatal(err)
		}
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	gomock.InOrder(
		cmp.core.EXPECT().Gather().Return(mfs, nil),
		cmp.responseWriter.EXPECT().Header().Return(cmp.responseHeader),
		cmp.responseWriter.EXPECT().Write(buf.Bytes()).Return(buf.Len(), nil),
		cmp.responseWriter.EXPECT().WriteHeader(200),
	)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerGzipEncoding(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/",
		},
		Header: map[string][]string{
			"Accept":           []string{"text/plain;version=0.0.4"},
			"Content-Encoding": []string{"gzip"},
		},
		Method: "GET",
		Body:   nil,
	}
	mfs := defaultMetricFamilies
	buf := &bytes.Buffer{}
	gz := gzip.NewWriter(buf)
	enc := expfmt.NewEncoder(gz, expfmt.FmtText)
	for _, mf := range mfs {
		err := enc.Encode(mf)
		if err != nil {
			t.Fatal(err)
		}
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	gomock.InOrder(
		cmp.core.EXPECT().Gather().Return(mfs, nil),
		cmp.responseWriter.EXPECT().Header().Return(cmp.responseHeader),
		cmp.responseWriter.EXPECT().Write(buf.Bytes()).Return(buf.Len(), nil),
		cmp.responseWriter.EXPECT().WriteHeader(200),
	)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}

func TestServerWriteError(t *testing.T) {
	request := &http.Request{
		URL: &url.URL{
			Path: "/",
		},
		Header: map[string][]string{
			"Accept": []string{"text/plain;version=0.0.4"},
		},
		Method: "GET",
		Body:   nil,
	}
	encodeErr := errors.New("encode error")
	mfs := defaultMetricFamilies
	buf := &bytes.Buffer{}
	enc := expfmt.NewEncoder(buf, expfmt.FmtText)
	for _, mf := range mfs {
		err := enc.Encode(mf)
		if err != nil {
			t.Fatal(err)
		}
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmp := setupServerTestComponents(t, ctrl)
	gomock.InOrder(
		cmp.core.EXPECT().Gather().Return(mfs, nil),
		cmp.responseWriter.EXPECT().Header().Return(cmp.responseHeader),
		cmp.responseWriter.EXPECT().Write(buf.Bytes()).Return(0, encodeErr),
		cmp.responseWriter.EXPECT().WriteHeader(500),
	)
	cmp.server.ServeHTTP(cmp.responseWriter, request)
}
