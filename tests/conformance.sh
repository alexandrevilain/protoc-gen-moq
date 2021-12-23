#!/bin/bash

echo "Running conformance tests"

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
WORK_DIR="conformance"
TEST_DIR="$DIR/$WORK_DIR"

mkdir $TEST_DIR && cd $TEST_DIR

function cleanup {      
  rm -rf "$TEST_DIR"
  echo "Deleted temp working directory $TEST_DIR"
}
trap cleanup EXIT

curl https://raw.githubusercontent.com/grpc/grpc-go/master/test/grpc_testing/test.proto -o conformance.proto

echo "Running protoc"
protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
    --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	--moq_out=. \
	--moq_opt=paths=source_relative \
    conformance.proto

echo "Running go vet"
go vet .