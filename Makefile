.PHONY: build
build:
	go build ./cmd/...

.PHONY: install
install:
	go install ./cmd/...

.PHONY: unit-test
unit-test:
	go test ./...

.PHONY: conformance-test
conformance-test:
	./tests/conformance.sh

.PHONY: test
test: unit-test conformance-test

