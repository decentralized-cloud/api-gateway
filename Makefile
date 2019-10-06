SHELL = /bin/bash
OS = $(shell uname -s | tr '[:upper:]' '[:lower:]')

# Build variables
BINARY_NAME = api-gateway
BUILD_DIR ?= bin
PACKAGE_DIR = cmd/api-gateway
VERSION ?= dev
COMMIT ?= $(shell git rev-parse HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)
PREFIX = github.com/micro-business/go-core/pkg/util
LDFLAGS += -X $(PREFIX).version=$(VERSION) -X $(PREFIX).commit=$(COMMIT) -X $(PREFIX).date=$(BUILD_DATE) -X $(PREFIX).platform=$(GOOS)/$(GOARCH)
REPORTS_DIR ?= reports
CI_SERVICE ?=
COVERALLS_TOKEN ?=

# Go variables
export CGO_ENABLED ?= 0
export GOOS ?= $(OS)
export GOARCH ?= amd64
GOFILES = $(shell find . -type f -name '*.go' -not -path "*/mock/*.go")

.PHONY: all
all: dep build install ## Get deps, build, and install binary

.PHONY: clean
clean: ## Clean the working area and the project
	rm -rf $(BUILD_DIR)/
	rm -rf $(REPORTS_DIR)

.PHONY: dep
dep: ## Install dependencies
	@go get golang.org/x/tools/cmd/cover
	@go get github.com/mattn/goveralls
	@go mod tidy
	@go get -v -t ./...

.PHONY: build
build: GOARGS += -tags "$(GOTAGS)" -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)
build: ## Build the binary
	@go build -v $(GOARGS) $(PACKAGE_DIR)/main.go

.PHONY: install
install: ## Install the api-gateway binary to /usr/local/bin
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin

.PHONY: format
format: ## Format the source
	@goimports -w $(GOFILES)

.PHONY: test
test: ## Run unit tests
	mkdir -p $(REPORTS_DIR)
	rm -f $(REPORTS_DIR)/*
	@go test -ldflags "$(LDFLAGS)" -v -covermode=count -coverprofile="$(REPORTS_DIR)/coverage.out" ./...

.PHONY: publish-test-results
publish-test-results: ## Publish test results
	@goveralls -coverprofile="$(REPORTS_DIR)/coverage.out" -service=$(CI_SERVICE) -repotoken $(COVERALLS_TOKEN)

.PHONY: list
list: ## List all make targets
	@$(MAKE) -pRrn : -f $(MAKEFILE_LIST) 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | sort

.PHONY: help
.DEFAULT_GOAL := help
help: ## Get help output
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)

