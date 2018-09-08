CIRCLE_BUILD_NUM ?= 0
VERSION?=0.0.$(CIRCLE_BUILD_NUM)
REF?=$(shell git rev-parse --short HEAD)

PREFIX?=$(shell pwd)
NAME := jira
PKG := github.com/endocrimes/jira-cli

BUILDDIR := ${PREFIX}/cross

# Populate version variables
# Add to compile time flags
CTIMEVAR=-X $(PKG)/internal/version.VERSION=$(VERSION) -X $(PKG)/cmd.GITCOMMIT=$(REF)
GO_LDFLAGS=-ldflags "-w $(CTIMEVAR)"

GOOSARCHES = darwin/amd64 darwin/386 freebsd/amd64 freebsd/386 linux/arm linux/arm64 linux/amd64 linux/386 windows/amd64

.PHONY: build
build: dist/$(NAME)

.PHONY: dist/$(NAME)
dist/$(NAME):
	@echo "+ $@"
	CGO_ENABLED=0 go build -tags "$(BUILDTAGS)" ${GO_LDFLAGS} -o $@ .

.PHONY: all
all: clean build lint test install

.PHONY: fmt
fmt:
	@echo "+ $@"
	@gometalinter --disable-all --enable gofmt ./...

.PHONY: test
test:
	@echo "+ $@"
	@gotestsum


.PHONY: lint
lint:
	@echo "+ $@"
	@gometalinter ./...

.PHONY: cover
cover:
	@echo "+ $@"
	@go test -coverprofile=coverage.txt ./...

.PHONY: install
install:
	@echo "+ $@"
	go install -a -tags "$(BUILDTAGS)" ${GO_LDFLAGS} .

.PHONY: clean
clean:
	@echo "+ $@"
	$(RM) -r dist $(BUILDDIR)

define buildpretty
mkdir -p $(BUILDDIR)/$(1)/$(2);
GOOS=$(1) GOARCH=$(2) go build \
	 -o $(BUILDDIR)/$(1)/$(2)/$(NAME) \
	 -a -tags "$(BUILDTAGS)" .;
endef

.PHONY: cross
cross:
	@echo "+ $@"
	$(foreach GOOSARCH,$(GOOSARCHES), $(call buildpretty,$(subst /,,$(dir $(GOOSARCH))),$(notdir $(GOOSARCH))))

