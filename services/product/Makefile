# Makefile for generating gRPC code from proto files

PROTO_DIR := proto
OUT_DIR := proto/gen
PROTO_FILE := $(PROTO_DIR)/product.proto

# Commands
PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)

# Check if plugins are installed
check-tools:
	@if [ -z "$(PROTOC_GEN_GO)" ]; then \
		echo "protoc-gen-go is not installed. Run 'go install google.golang.org/protobuf/cmd/protoc-gen-go@latest'"; \
		exit 1; \
	fi
	@if [ -z "$(PROTOC_GEN_GO_GRPC)" ]; then \
		echo "protoc-gen-go-grpc is not installed. Run 'go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest'"; \
		exit 1; \
	fi

# Generate protobuf code
gen: check-tools
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_FILE)

# Clean generated files
clean:
	rm -rf $(OUT_DIR)/*.pb.go

.PHONY: generate clean check-tools
