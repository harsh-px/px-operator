#
# Based on http://chrismckenzie.io/post/deploying-with-golang/
#          https://raw.githubusercontent.com/lpabon/quartermaster/dev/Makefile
#

.PHONY: version all operator run dist clean

APP_NAME := operator
SHA := $(shell git rev-parse --short HEAD)
BRANCH := $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))
VER := $(shell git rev-parse --short HEAD)
ARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)
+GLIDEPATH := $(shell command -v glide 2> /dev/null)
DIR=.

ifndef TAGS
TAGS := daemon
endif

ifdef APP_SUFFIX
  VERSION = $(VER)-$(subst /,-,$(APP_SUFFIX))
else
ifeq (master,$(BRANCH))
  VERSION = $(VER)
else
  VERSION = $(VER)-$(BRANCH)
endif
endif

# Go setup
GO=go

# Sources and Targets
EXECUTABLES :=$(APP_NAME)
# Build Binaries setting main.version and main.build vars
LDFLAGS :=-ldflags "-X main.PX_OPERATOR_VERSION=$(VERSION) -extldflags '-z relro -z now'"
# Package target
PACKAGE :=$(DIR)/dist/$(APP_NAME)-$(VERSION).$(GOOS).$(ARCH).tar.gz
PKGS=$(shell go list ./... | grep -v vendor)
GOVET_PKGS=$(shell  go list ./... | grep -v vendor | grep -v pkg/client/informers/externalversions | grep -v versioned)

BASE_DIR := $(shell git rev-parse --show-toplevel)

BIN :=$(BASE_DIR)/bin
GOFMT := gofmt

.DEFAULT: all

all: pretest operator

vendor: glide.lock
ifndef GLIDEPATH
	echo "Installing Glide"
	curl https://glide.sh/get | sh
endif
	echo "Installing vendor directory"
	glide install -v

	echo "Building dependencies to make builds faster"
	go install github.com/harsh-px/px-operator/cmd/operator

glide.lock: glide.yaml
	echo "Glide.yaml has changed, updating glide.lock"
	glide update -v

# print the version
version:
	@echo $(VERSION)

# print the name of the app
name:
	@echo $(APP_NAME)

# print the package path
package:
	@echo $(PACKAGE)

operator: glide.lock vendor codegen
	mkdir -p $(BIN)
	go build $(LDFLAGS) -o $(BIN)/$(APP_NAME) cmd/operator/main.go

run: operator
	$(BIN)/$(APP_NAME)

test:
	go test -tags "$(TAGS)" $(TESTFLAGS) $(PKGS)

clean:
	@echo Cleaning Workspace...
	-sudo rm -rf $(BIN)
	rm -rf dist

fmt:
	@echo "Performing gofmt on following: $(PKGS)"
	@./hack/do-gofmt.sh $(PKGS)

checkfmt:
	@echo "Checking gofmt on following: $(PKGS)"
	@./hack/check-gofmt.sh $(PKGS)

lint:
	@echo "golint"
	go get -v github.com/golang/lint/golint
	for file in $$(find . -name '*.go' | grep -v vendor | \
																			grep -v '\.pb\.go' | \
																			grep -v '\.pb\.gw\.go' | \
																			grep -v 'externalversions' | \
																			grep -v 'versioned' | \
																			grep -v 'generated'); do \
		golint $${file}; \
		if [ -n "$$(golint $${file})" ]; then \
			exit 1; \
		fi; \
	done

vet:
	@echo "govet"
	go vet $(GOVET_PKGS)

errcheck:
	@echo "errcheck"
	go get -v github.com/kisielk/errcheck
	errcheck -tags "$(TAGS)" $(GOVET_PKGS)

codegen:
	@echo "Generating files"
	@./hack/update-codegen.sh

verifycodegen:
	@echo "Verifying generated files"
	@./hack/verify-codegen.sh

pretest: checkfmt vet lint errcheck verifycodegen

$(PACKAGE): all
	@echo Packaging Binaries...
	@mkdir -p tmp/$(APP_NAME)
	@cp $(BIN)/$(APP_NAME) tmp/$(APP_NAME)/
	@mkdir -p $(DIR)/dist/
	tar -czf $@ -C tmp $(APP_NAME);
	@rm -rf tmp
	@echo
	@echo Package $@ saved in dist directory

dist: $(PACKAGE) $(CLIENT_PACKAGE)

.PHONY: test clean name run version
