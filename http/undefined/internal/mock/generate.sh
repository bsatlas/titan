#!/bin/bash

mockgen \
	-destination http.go \
	-package mock \
	-mock_names "ResponseWriter=ResponseWriter" \
	net/http Handler,ResponseWriter
