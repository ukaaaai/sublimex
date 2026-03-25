.PHONY: proto-all proto-go proto-python

# パス設定
PROTO_DIR = proto
PROTO_FILE = embedding.proto
GO_PB_DIR = mcp-server/pb
PY_DIR = embedding-server
PY_PB_DIR = $(PY_DIR)/pb

# 全生成
proto-all: proto-go proto-python

# Go側の生成
proto-go:
	@mkdir -p $(GO_PB_DIR)
	protoc --proto_path=$(PROTO_DIR) \
		--go_out=$(GO_PB_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(GO_PB_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/$(PROTO_FILE)
	@echo "✅ Go code generated in $(GO_PB_DIR)"

# Python側の生成
proto-python:
	@mkdir -p $(PY_PB_DIR)
	@touch $(PY_PB_DIR)/__init__.py
	uv run --project $(PY_DIR) python -m grpc_tools.protoc \
		-I$(PROTO_DIR) \
		--python_out=$(PY_PB_DIR) \
		--pyi_out=$(PY_PB_DIR) \
		--grpc_python_out=$(PY_PB_DIR) \
		$(PROTO_DIR)/$(PROTO_FILE)
	@echo "✅ Python code & type stubs (.pyi) generated in $(PY_PB_DIR)"