#!/bin/bash

EXCLUDE='(mock|cmd)'
PACKAGES=$(go list ./... | grep -Ev $EXCLUDE)
go test -cover -count=1 $PACKAGES
