PROTO_FILES := proto/**/
GO_OUT := plugins=grpc:go/
PYTHON_OUT := lamia_shared/

.PHONY: generate-python
generate-python:
	@echo "--- Generate Python client code ... ---"
	python3 -m grpc_tools.protoc -I=. --python_out=$(PYTHON_OUT) --grpc_python_out=$(PYTHON_OUT) $(PROTO_FILES)*.proto
	@echo "--- Python client generated ---"
	@echo "--- Reformat Python generated codes ... ---"
	@find $(PYTHON_OUT)proto/ -type f -name "*_pb2_grpc.py" -exec sed -i'' -e 's/from proto.\([^ ]*\) import \(.*\)/from . import \2/g' {} +
	@echo "--- Reformat Python generated codes finished ---"

.PHONY: clean-python
clean-python:
	@echo "--- Cleaning Python generated codes ... ---"
	rm -rf lamia_shared/proto/**/*_pb*.py
	@echo "--- Cleaning Python generated codes finished ---"

.PHONY: generate-go
generate-go:
	@echo "--- Generate Go client code ... ---"
	@find $(PROTO_FILES) -type f -name "*.proto" -exec protoc --go_out=$(GO_OUT) {} \;
	@echo "--- Go client generated ---"

.PHONY: clean-go
clean-go:
	@echo "--- Cleaning go generated codes ... ---"
	rm -rf go/proto/**/*.pb.go
	@echo "--- Cleaning go generated codes finished ---"

.PHONY: generate-all
generate-all: generate-python generate-go

.PHONY: regenerate-all
regenerate-all: clean-all generate-all

.PHONY: clean-all
clean-all: clean-python clean-go