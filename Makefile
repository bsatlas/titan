# build titan
.PHONY: build
build:
	scripts/build/build.sh

# run unit tests for go code
.PHONY: test
test:
	scripts/testing/static-analysis.sh
	scripts/testing/unit-tests.sh

# generate mocks
.PHONY: mocks
mocks:
	scripts/testing/mock/generate.sh

# remove all generated mocks
.PHONY: clean-mocks
clean-mocks:
	scripts/testing/mock/clean.sh
