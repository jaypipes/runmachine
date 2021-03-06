PROTO := proto
VENDOR := vendor
VERSION := $(shell git describe --tags --always --dirty)
BUILD_DIR := $(shell pwd)/build
PROTO_DIR := $(shell pwd)/proto
PROTO_DEFS_DIR := $(shell pwd)/proto/defs
GO_BIN_DIR := $(GOPATH)/bin
GO_PROTOC_BIN := $(GO_BIN_DIR)/protoc-gen-go
PKGS := $(shell go list ./... | grep -v /$(VENDOR)/ | grep -v /$(PROTO)/)
SRC = $(shell find . -type f -name '*.go' -not -path "*/$(VENDOR)/*" -not -path "*/$(PROTO_DIR)/*")

.PHONY: test
test: generated fmtcheck vet
	@echo "Running all go tests ... "
	@go test $(PKGS)

$(GO_PROTOC_BIN):
	@go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: generated
# Generates protobuffer code
generated: $(GO_PROTOC_BIN)
	@echo -n "Generating protobuffer code from proto definitions ... "
	@protoc -I $(PROTO_DEFS_DIR) \
	       $(PROTO_DEFS_DIR)/*.proto \
	       --go_out=plugins=grpc:$(PROTO_DIR) && echo "ok."

$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	$(GOMETALINTER) --install

.PHONY: lint
lint: $(GOMETALINTER)
	$(GOMETALINTER) ./... --vendor

.PHONY: fmt
fmt:
	@echo "Running gofmt on all sources ..."
	@gofmt -s -l -w $(SRC)

.PHONY: fmtcheck
fmtcheck:
	@bash -c "diff -u <(echo -n) <(gofmt -d $(SRC))"

.PHONY: vet
vet:
	@go vet $(PKGS)

.PHONY: cover
cover:
	$(shell [ -e coverage.out ] && rm coverage.out)
	@echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PKGS),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out -o=coverage-all.html

build-base:
	@echo "building base Docker image ..."
	docker build -q --label built-by=runmachine.io -t runmachine.io/runmachine/base . -f cmd/Dockerfile

build-metadata: build-base
	@echo "building runm-metadata Docker image ..."
	docker build -q --label built-by=runmachine.io -t runmachine.io/runmachine/metadata:$(VERSION) . -f cmd/runm-metadata/Dockerfile

build-resource: build-base
	@echo "building runm-resource Docker image ..."
	docker build -q --label built-by=runmachine.io -t runmachine.io/runmachine/resource:$(VERSION) . -f cmd/runm-resource/Dockerfile

build-api: build-base
	@echo "building runm-api Docker image ..."
	docker build -q --label built-by=runmachine.io -t runmachine.io/runmachine/api:$(VERSION) . -f cmd/runm-api/Dockerfile

build-cli:
	@echo "building runm CLI Docker image to $(BUILD_BIN_DIR)/runm ..."
	bash $(BUILD_DIR)/build_runm.sh

build: build-base build-metadata build-resource build-api build-cli

.PHONY: clean
clean:
	@echo "Cleaning up all built Docker images ..."
	@for i in $( docker image list | grep runm | awk '{print $3}' ); do \
		docker image rm $i --force; \
	done
	@docker image prune --force
