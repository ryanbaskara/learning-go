# Direktori proto dan output
PROTO_DIR = proto
GEN_DIR = go

# Cari semua file .proto di dalam PROTO_DIR
PROTOS := $(shell find $(PROTO_DIR) -name "*.proto")

# Command utama
proto-go:
	@echo "Generating Go gRPC code..."
	@mkdir -p $(GEN_DIR)
	@protoc \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(GEN_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTOS)
	@echo "✅ gRPC code generated in $(GEN_DIR)"

# Bersihkan hasil generate
clean:
	@echo "Cleaning generated files..."
	@rm -rf $(GEN_DIR)
	@echo "✅ Clean done"
