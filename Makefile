PKG_NAME := main

# Commands 
GO := go
GOFMT := gofmt
GOLINT := golint
GOVET := go vet
GOTEST := go test
GOCOVER := go tool cover
TEMP_DIR := ./tmp
COVERAGE_PROFILE := cover.out
BINARY := $(TEMP_DIR)/$(PKG_NAME)

# Flags
GOFLAGS := -v
GOFMTFLAGS := -s
GOLINTFLAGS := -set_exit_status
GOVETFLAGS := -all
GOTESTFLAGS := -v
GOCOVERFLAGS := -html

.PHONY: all
all: build test

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFMTFLAGS) -w .

.PHONY: lint
lint:
	$(GOLINT) $(GOLINTFLAGS) ./...

.PHONY: vet
vet:
	$(GOVET) $(GOVETFLAGS) ./...

.PHONY: clean
clean:
	rm -rf $(TEMP_DIR)

.PHONY: build
build: vet
	$(GO) build $(GOFLAGS) -o $(BINARY)

.PHONY: test
test:
	$(GOTEST) $(GOTESTFLAGS) ./...

.PHONY: cover
cover:
	mkdir -p $(TEMP_DIR)
	$(GOTEST) $(GOTESTFLAGS) -coverprofile=$(COVERAGE_PROFILE) -outputdir=$(TEMP_DIR) `go list ./... | grep -v ./internal | grep -v ./examples`
	$(GOCOVER) $(GOCOVERFLAGS) $(TEMP_DIR)/$(COVERAGE_PROFILE)


.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  all     - build and test"
	@echo "  fmt     - format source code"
	@echo "  lint    - lint source code"
	@echo "  vet     - vet source code"
	@echo "  clean   - clean up"
	@echo "  build   - build binary"
	@echo "  test    - run tests"
	@echo "  cover   - run tests with coverage"
	@echo "  help    - show this help message"