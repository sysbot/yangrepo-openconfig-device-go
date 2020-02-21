# Example usage:
# make

BIN_DIR := $(GOPATH)/bin
PUBLIC := /work/public

GOBUILD:=go build
GOARCH ?= amd64
GOCGO:=0
GOCLEAN:=go clean
GOGET:=go get
GOINSTALL:=go install
GOTEST:=go test
GOVET:=go vet
GOTOOL:=go tool
NAME ?= $(shell basename $(PWD) | tr '[:upper:]' '[:lower:]')

# containerize
PWD := $(shell pwd)
DOCKERFILE ?= Dockerfile
CONTAINER ?= $(NAME)
DOCKER_RUN_FLAGS :=--rm -v $$(pwd	):/work -w /work $(CONTAINER)

define GEN_OPENCONFIG
generator \
	-path=$(PUBLIC),yang \
	-output_dir=./ \
	-package_name=openconfig \
	-generate_fakeroot \
	-fakeroot_name=device \
	-compress_paths=true \
	-exclude_modules=ietf-interfaces \
	-generate_rename \
	-generate_append \
	-generate_getters \
	-generate_delete \
	-annotations \
	$(PUBLIC)/release/models/network-instance/openconfig-network-instance.yang \
	$(PUBLIC)/release/models/mpls/openconfig-mpls.yang \
	$(PUBLIC)/release/models/optical-transport/openconfig-optical-amplifier.yang \
	$(PUBLIC)/release/models/optical-transport/openconfig-terminal-device.yang \
	$(PUBLIC)/release/models/optical-transport/openconfig-transport-line-protection.yang \
	$(PUBLIC)/release/models/platform/openconfig-platform.yang \
	$(PUBLIC)/release/models/policy/openconfig-routing-policy.yang \
	$(PUBLIC)/release/models/lacp/openconfig-lacp.yang \
	$(PUBLIC)/release/models/system/openconfig-system.yang \
	$(PUBLIC)/release/models/lldp/openconfig-lldp.yang \
	$(PUBLIC)/release/models/stp/openconfig-spanning-tree.yang \
	$(PUBLIC)/release/models/interfaces/openconfig-interfaces.yang \
	$(PUBLIC)/release/models/interfaces/openconfig-if-ip.yang \
	$(PUBLIC)/release/models/interfaces/openconfig-if-aggregate.yang \
	$(PUBLIC)/release/models/interfaces/openconfig-if-ethernet.yang \
	$(PUBLIC)/release/models/interfaces/openconfig-if-ip-ext.yang \
	$(PUBLIC)/release/models/local-routing/openconfig-local-routing.yang \
	$(PUBLIC)/release/models/vlan/openconfig-vlan.yang \
	$(PUBLIC)/release/models/vlan/openconfig-vlan-types.yang
endef

## build inside container and install
all: build

## run golang code fmt except for generated code and vendor
lint: ; @goimports -w ./

## build a container from the Dockerfile
container:
	DOCKER_BUILDKIT=1 docker build -t $(CONTAINER) -f $(DOCKERFILE) .

test: container
	docker run $(DOCKER_RUN_FLAGS)

console: container
	docker run -it $(DOCKER_RUN_FLAGS)

container_build: container
	docker run $(DOCKER_RUN_FLAGS) make build

## generate openconfig need to run inside container
generate:
	$(GEN_OPENCONFIG)

build: generate lint

GREEN	 := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE	 := $(shell tput -Txterm setaf 7)
RESET	 := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=20
## Show this help
help:
	@echo ''
	@echo 'Usage:'
	@echo '	 ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "	${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.PHONY: help build generate lint console container container_build
