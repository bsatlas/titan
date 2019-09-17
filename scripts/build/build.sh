#!/bin/bash

DIRTY=$(git status --porcelain --untracked-files=no)
COMMITHASH=$(git rev-parse HEAD 2>/dev/null || true)
if [ "$DIRTY" ]; then
	COMMIT="$COMMITHASH-dirty"
else
	COMMIT="$COMMITHASH"
fi
PACKAGE="github.com/atlaskerr/titan/cmd/titan/version"
VERSION="v0.0.0"
LDFLAGS="-X $PACKAGE.Version=$VERSION -X $PACKAGE.Commit=$COMMIT"

go build -ldflags "$LDFLAGS" -o bin/titan ./cmd/titan
