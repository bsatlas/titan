# build titan
.PHONY: build
build:
	scripts/build/build.sh

# clean built binaries
.PHONY: clean
clean:
	scripts/build/clean.sh

# run tests for go code
.PHONY: test
test: lint unit-test

# run unit tests for go code
.PHONY: unit-test
unit-test:
	scripts/testing/unit-tests.sh

# run unit tests for go code
.PHONY: lint 
lint:
	scripts/testing/static-analysis.sh

# generate mocks
.PHONY: mocks
mocks:
	scripts/testing/mock/generate.sh

# remove all generated mocks
.PHONY: clean-mocks
clean-mocks:
	scripts/testing/mock/clean.sh
