package http_test

import (
	"testing"

	"github.com/atlaskerr/titan/metrics/http"

	"github.com/prometheus/client_golang/prometheus"
)

func TestCollector(t *testing.T) {
	c := http.NewCollector()
	reg := prometheus.NewPedanticRegistry()
	err := reg.Register(c)
	if err != nil {
		t.Fatal(err)
	}
	_, err = reg.Gather()
	if err != nil {
		t.Fatal(err)
	}
}
