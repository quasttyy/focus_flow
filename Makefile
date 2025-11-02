PROTO_DIR = task-service/proto
PROTO_FILES = $(PROTO_DIR)/messages.proto $(PROTO_DIR)/service.proto

generate:
	protoc -I $(PROTO_DIR) \
	  --go_out $(PROTO_DIR) --go_opt paths=source_relative \
	  --go-grpc_out $(PROTO_DIR) --go-grpc_opt paths=source_relative \
	  $(PROTO_FILES)
