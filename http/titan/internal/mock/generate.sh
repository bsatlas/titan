#!/bin/bash

mockgen \
	-destination http.go \
	-package mock \
	-mock_names "Handler=Handler,ResponseWriter=ResponseWriter" \
	net/http Handler,ResponseWriter
