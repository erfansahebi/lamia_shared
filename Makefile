#proto:
#	protoc services/**/*.proto --go_out=plugins=grpc:.

PROTO_FILES := services/**/
GO_OUT := plugins=grpc:.
PYTHON_OUT := .

.PHONY: generate-python
generate-python:
	@echo "--- Generate Python client code ... ---"
	find $(PROTO_FILES) -name '*.proto' -exec python3 -m grpc_tools.protoc -I=$(PYTHON_OUT) --python_out=$(PYTHON_OUT) --grpc_python_out=$(PYTHON_OUT) {} +
	@echo "--- Python client generated ---"

.PHONY: generate-go
generate-go:
	@echo "--- Generate Go client code ... ---"
	protoc $(PROTO_FILES)*.proto --go_out=$(GO_OUT)
	@echo "--- Go client generated ---"

.PHONY: generate-all
generate-all: generate-python generate-go

.PHONY: regenerate-all
regenerate-all: clean generate-all

.PHONY: clean
clean:
	@echo "--- Cleaning generated codes ... ---"
	rm -rf services/**/*.go
	rm -rf services/**/*.py
	@echo "--- Cleaning generated codes finished ---"