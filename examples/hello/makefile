.PHONY: generate
generate:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
    --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	--moq_out=. \
	--moq_opt=paths=source_relative \
    hello.proto