.PHONY: build
build:
	go build -o protoc-gen-moq .

.PHONY: install
install:
	go install .

.PHONY: unit-test
unit-test:
	go test ./...

.PHONY: conformance-test
conformance-test:
	./tests/conformance.sh

.PHONY: test
test: unit-test conformance-test

