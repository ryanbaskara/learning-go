# Direktori proto dan output
PROTO_SRC = proto
PROTO_OUT = go
PROTO_THIRD_PARTY = third_party

# Cari semua file .proto di dalam PROTO_DIR
PROTOS := $(shell find $(PROTO_SRC) -name "*.proto")

# Command utama
proto-go:
	@echo "Generating Go gRPC code..."
	@mkdir -p $(PROTO_OUT)
	@protoc \
		--proto_path=$(PROTO_SRC) \
		--proto_path=$(PROTO_THIRD_PARTY) \
		--go_out=$(PROTO_OUT) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT) --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=$(PROTO_OUT) --grpc-gateway_opt=paths=source_relative \
		$(PROTOS)
	@echo "✅ gRPC code generated in $(PROTO_OUT)"

# Bersihkan hasil generate
clean:
	@echo "Cleaning generated files..."
	@rm -rf $(PROTO_OUT)
	@echo "✅ Clean done"
