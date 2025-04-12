# ------------------------------------------------------- ENVS ------------------------------------------------------- #

PROJECT := ebpfstruct

COMMIT_SHA := $(shell git rev-parse --short HEAD)
TIMESTAMP  := $(shell date --utc --iso-8601=seconds)
VERSION    ?= $(shell git describe --tags --always --dirty)
ROOT       := $(shell git rev-parse --show-toplevel)

# ------------------------------------------------------- GENERATE --------------------------------------------------- #

MOCKERY_VERSION := v3
MOCKERY := go run github.com/vektra/mockery/v3@$(MOCKERY_VERSION)

.PHONY: generate
generate:
	rm -rf pkg/mockebpfstruct
	$(MOCKERY)

# ------------------------------------------------------- FMT -------------------------------------------------------- #

GOFUMPT_VERSION := v0.6.0
GOFUMPT := go run mvdan.cc/gofumpt@$(GOFUMPT_VERSION)

.PHONY: fmt
fmt:
	$(GOFUMPT) -w .

# ------------------------------------------------------- LINT ------------------------------------------------------- #

GOLANGCI_LINT_VERSION := v2.0.2
GOLANGCI_LINT := go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
LINT_LICENSES := ./hack/find-files-without-licenses.sh

.PHONY: lint
lint:
	$(LINT_LICENSES)
	$(GOLANGCI_LINT) run --fix

# ------------------------------------------------------- PRE-PUSH --------------------------------------------------- #

.PHONY: pre-push
pre-push: generate fmt lint
	git status --porcelain
