#!/bin/sh

find . -path '**/mock/**' -name generate.go -execdir go generate {} \;
