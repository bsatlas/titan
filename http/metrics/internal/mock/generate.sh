#!/bin/bash

mockgen \
	-destination http.go \
	-package mock \
	-mock_names "Handler=Handler,ResponseWriter=ResponseWriter" \
	net/http Handler,ResponseWriter

mockgen \
	-destination prometheus.go \
	-package mock \
	-mock_names "Gatherer=Gatherer" \
	github.com/prometheus/client_golang/prometheus Gatherer


