PROTO_SRC = ../../grpc/src/main/proto
PROTO_VERSION = 3.5.1

BUILD_DIR = build
BIN_DIR = $(BUILD_DIR)/bin
SERVICES = \
    $(PROTO_SRC)/service/search.pb.go \
    $(PROTO_SRC)/service/event_service.pb.go \
    $(PROTO_SRC)/service/metadata_service.pb.go \
    $(PROTO_SRC)/service/task_service.pb.go \
    $(PROTO_SRC)/service/workflow_service.pb.go

PROTOC=protoc

proto: models $(SERVICES)

build:
	go fmt ./...
	go build ./...

test:
	go test ./...

# Helpers
$(SERVICES): %.pb.go: %.proto
	$(PROTOC) -I $(PROTO_SRC) $< --go_out=plugins=grpc,paths=source_relative,import_path=conductor:conductor

models:
	$(PROTOC) -I $(PROTO_SRC) $(PROTO_SRC)/model/*.proto --go_out=paths=source_relative,import_path=conductor:conductor