# Makefile for generating protobuf files

# Base directories
PROTO_DIR = api/proto
OUTPUT_DIR = api/pb  # Keep the output directory as api/pb

# Protoc command options
PROTOC = protoc
PROTOC_FLAGS = --proto_path=$(PROTO_DIR) \
                --go_out=$(OUTPUT_DIR) --go_opt=paths=source_relative \
                --go-grpc_out=$(OUTPUT_DIR) --go-grpc_opt=paths=source_relative \
                --grpc-gateway_out=$(OUTPUT_DIR) --grpc-gateway_opt=paths=source_relative \
                --openapiv2_out=docs/api --openapiv2_opt=allow_delete_body=true \
                -I . \
                -I /usr/local/include \
                -I ./googleapis \
                -I $(shell go list -f '{{.Dir}}' -m github.com/grpc-ecosystem/grpc-gateway/v2) \
                -I $(shell go list -f '{{.Dir}}' -m google.golang.org/protobuf)

# Targets
.PHONY: generate

generate:
	@if [ -z "$$PROTO_FILE" ]; then \
		echo "Usage: make generate PROTO_FILE=path/to/your.proto"; \
		exit 1; \
	fi; \
	mkdir -p $(OUTPUT_DIR); \
	$(PROTOC) $(PROTOC_FLAGS) "$$PROTO_FILE"

run-grpc:
	go run cmd/main.go

run-gateway:
	go run cmd/gateway/main.go

docker-up:
	docker-compose up --build
