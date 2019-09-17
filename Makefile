.PHONY: build
build:
	scripts/build/build.sh

.PHONY: test
test:
	scripts/testing/static-analysis.sh
	scripts/testing/unit-tests.sh
