GOPATH?=$(shell go env GOPATH)
SERVICE_STYLE=zero-swagger

GO ?= go
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
LDFLAGS := -s -w

PLATFORM := $(shell go env GOOS)
#CGO_ENABLED := $(shell go env CGO_ENABLED)
GOARCH := $(shell go env GOARCH)

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: lint
lint:
	golangci-lint run -D staticcheck

.PHONY: build
build:
	go env -u CGO_ENABLED GOOS GOARCH
	go build -ldflags "$(LDFLAGS)" -o $(SERVICE_STYLE) $(SERVICE_STYLE).go
	@echo "Build project $(PLATFORM)-$(GOARCH) successfully"
	mv $(SERVICE_STYLE) $(GOPATH)/bin/$(SERVICE_STYLE)
	@echo "move project $(SERVICE_STYLE) successfully"

.PHONY: build-win
build-win:
	env CGO_ENABLED=0 GOOS=windows go build -ldflags "$(LDFLAGS)" -o $(SERVICE_STYLE).exe $(SERVICE_STYLE).go
	go env -u CGO_ENABLED GOOS GOARCH
	@echo "Build project for windows successfully"

.PHONY: build-mac
build-mac:
	env CGO_ENABLED=0 GOOS=darwin go build -ldflags "$(LDFLAGS)" -o $(SERVICE_STYLE) $(SERVICE_STYLE).go
	go env -u CGO_ENABLED GOOS GOARCH
	@echo "Build project for MacOS successfully"

.PHONY: build-linux
build-linux:
	env CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" -o $(SERVICE_STYLE) $(SERVICE_STYLE).go
	go env -u CGO_ENABLED GOOS GOARCH
	@echo "Build project for linux successfully"