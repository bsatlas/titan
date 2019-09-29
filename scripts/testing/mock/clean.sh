#!/bin/sh

find . -type f -path '**/mock/**' \
	! -name 'generate.go' \
	! -name 'generate.sh' \
	-exec rm -f {} \;
