# ROOT Options:
ROOT_PACKAGE=github.com/envoyproxy/gateway

# Set Root Directory Path
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell  pwd -P))
endif

# Set Output Directory Path
ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(ROOT_DIR)/bin
$(shell mkdir -p $(OUTPUT_DIR))
endif

# Set the version number. you should not need to do this
# for the majority of scenarios.
ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --abbrev=0 --dirty --always --tags | sed 's/-/./g')
endif

# Supported Platforms for building multiarch binaries.
PLATFORMS ?= darwin_amd64 darwin_arm64 linux_amd64 linux_arm64 

# Set a specific PLATFORM
ifeq ($(origin PLATFORM), undefined)
	ifeq ($(origin GOOS), undefined)
		GOOS := $(shell go env GOOS)
	endif
	ifeq ($(origin GOARCH), undefined)
		GOARCH := $(shell go env GOARCH)
	endif
	PLATFORM := $(GOOS)_$(GOARCH)
	# Use linux as the default OS when building images
	IMAGE_PLAT := linux_$(GOARCH)
else
	GOOS := $(word 1, $(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2, $(subst _, ,$(PLATFORM)))
	IMAGE_PLAT := $(PLATFORM)
endif

# List commands in cmd directory for building targets
COMMANDS ?= $(wildcard ${ROOT_DIR}/cmd/*)
BINS ?= $(foreach cmd,${COMMANDS},$(notdir ${cmd}))

ifeq (${COMMANDS},)
  $(error Could not determine COMMANDS, set ROOT_DIR or run in source dir)
endif
ifeq (${BINS},)
  $(error Could not determine BINS, set ROOT_DIR or run in source dir)
endif

# Docker variables
# REGISTRY is the image registry to use for build and push image targets.
REGISTRY ?= docker.io/envoyproxy
# IMAGE is the image URL for build and push image targets.
IMAGE ?= ${REGISTRY}/gateway-dev
# REV is the short git sha of latest commit.
REV=$(shell git rev-parse --short HEAD)
# Tag is the tag to use for build and push image targets.
TAG ?= $(REV)

# Build the envoy-gateway binaries in target platforms
.PHONY: build.%
build.%:
	$(eval COMMAND := $(word 2,$(subst ., ,$*)))
	$(eval PLATFORM := $(word 1,$(subst ., ,$*)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "===========> Building binary $(COMMAND) $(VERSION) for $(OS) $(ARCH)"
	@mkdir -p $(OUTPUT_DIR)/$(OS)/$(ARCH)
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -o $(OUTPUT_DIR)/$(OS)/$(ARCH)/$(COMMAND) -ldflags "$(GO_LDFLAGS)" $(ROOT_PACKAGE)/cmd/$(COMMAND)

# Build the envoy-gateway binaries in hosted platforms
.PHONY: build
build:  $(addprefix build., $(addprefix $(PLATFORM)., $(BINS)))

# Build the envoy-gateway binaries in multi platforms
.PHONY: build.multiarch
build.multiarch:  $(foreach p,$(PLATFORMS),$(addprefix build., $(addprefix $(p)., $(BINS))))

.PHONY: clean
clean: ## Clean the building output files
	@echo "===========> Cleaning all build output"
	@rm -rf $(OUTPUT_DIR)

.PHONY: test
test: ## Run go unit tests
	@go test ./...

.PHONY: docker-build
docker-build: build ## Build the envoy-gateway docker image.
	$(eval IMAGE_NAME := $(COMMAND))
	$(eval IMAGE_PLAT := $(subst _,/,$(PLATFORM)))
	@echo "===========> Building docker image $(IMAGE_NAME) $(VERSION) for $(IMAGE_PLAT)"
	@DOCKER_BUILDKIT=1 docker build -t $(IMAGE):$(TAG) -f Dockerfile bin

.PHONY: docker-push
docker-push: ## Push the docker image for envoy-gateway.
	@docker push $(IMAGE):$(TAG)

.PHONY: help
help: ## Display this help
	@echo Envoy Gateway is an open source project for managing Envoy Proxy as a standalone or Kubernetes-based application gateway
	@echo
	@echo Targets:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9._-]+:.*?## / {printf "  %-25s %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.PHONY: lint
lint: ## Run lint checks
lint: lint-golint lint-yamllint lint-codespell

.PHONY: lint-golint
lint-golint:
	@echo Running Go linter ...
	@golangci-lint run --build-tags=e2e --config=tools/golangci-lint/.golangci.yml

.PHONY: lint-yamllint
lint-yamllint:
	@echo Running YAML linter ...
	## TODO(lianghao208): add other directories later
	@yamllint --config-file=tools/yamllint/.yamllint changelogs/

.PHONY: lint-codespell
lint-codespell: CODESPELL_SKIP := $(shell cat tools/codespell/.codespell.skip | tr \\n ',')
lint-codespell:
	@codespell --skip $(CODESPELL_SKIP) --ignore-words tools/codespell/.codespell.ignorewords --check-filenames --check-hidden -q2
