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
	$(MOCKERY)

# ------------------------------------------------------- FMT -------------------------------------------------------- #

.PHONY: fmt
fmt:
	$(GOFUMPT) -w .

# ------------------------------------------------------- LINT ------------------------------------------------------- #

.PHONY: lint
lint:
	$(LINT_LICENSES)
	$(GOLANGCI_LINT) run --fix

# ------------------------------------------------------- PRE-PUSH --------------------------------------------------- #

.PHONY: pre-push
pre-push: generate fmt lint test
	git status --porcelain
