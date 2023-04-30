build:
	@go build -o bin/2pchess

run: build
	@./bin/2pchess

test:
	@go test -v ./...

proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    rpc/rpc.proto
	@echo "compiled rpc proto buffers"

.PHONY: proto