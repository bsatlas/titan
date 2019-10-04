#!/bin/bash

PACKAGE="github.com/atlaskerr/titan/http/live"

mockgen \
	-destination http.go \
	-package mock \
	-mock_names "Handler=Handler,ResponseWriter=ResponseWriter" \
	net/http Handler,ResponseWriter

mockgen \
	-destination live.go \
	-package mock \
	-mock_names "Liveness=Liveness" \
	$PACKAGE Liveness
