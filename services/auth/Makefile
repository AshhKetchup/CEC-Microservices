# services/product/Makefile

PROTO_DIR := proto
PROTO_FILE := $(PROTO_DIR)/auth.proto
GEN_DIR := $(PROTO_DIR)/gen
GO_MODULE := cec/services/product

.PHONY: gen clean check-tools

check-tools:
	@which protoc > /dev/null || (echo "protoc missing. Install from https://grpc.io/docs/protoc-installation/" && exit 1)
	@which protoc-gen-go > /dev/null || (echo "Run: go install google.golang.org/protobuf/cmd/protoc-gen-go@latest" && exit 1)
	@which protoc-gen-go-grpc > /dev/null || (echo "Run: go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest" && exit 1)

gen: check-tools
	@mkdir -p $(GEN_DIR)
	protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILE)
	@echo "Generated files in: $(GEN_DIR)"

clean:
	rm -rf $(GEN_DIR)

help:
	@echo "Product Service Makefile"
	@echo "Usage:"
	@echo "  make gen      - Generate gRPC code"
	@echo "  make clean    - Remove generated files"
	@echo "  make check-tools - Verify required tools are installed"