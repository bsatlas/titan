#!/bin/bash

PACKAGE="github.com/atlaskerr/titan/http/ready"

mockgen \
	-destination http.go \
	-package mock \
	-mock_names "Handler=Handler,ResponseWriter=ResponseWriter" \
	net/http Handler,ResponseWriter

mockgen \
	-destination ready.go \
	-package mock \
	-mock_names "Readiness=Readiness" \
	$PACKAGE Readiness
