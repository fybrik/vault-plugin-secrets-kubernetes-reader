include Makefile.env

GOARCH = amd64
OS = linux

.DEFAULT_GOAL := all

DOCKER_HOSTNAME ?= ghcr.io
DOCKER_NAMESPACE ?= fybrik
DOCKER_TAG ?= 0.0.0
DOCKER_NAME ?= vault-plugin-secrets-kubernetes-reader

IMG := ${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/${DOCKER_NAME}:${DOCKER_TAG}

all: source-build

.PHONY: source-build
source-build:
	CGO_ENABLED=0 GOOS="$(OS)" GOARCH="$(GOARCH)" go build -o vault/plugins/vault-plugin-secrets-kubernetes-reader cmd/vault-plugin-secrets-kubernetes-reader/main.go

.PHONY: docker-build
docker-build: source-build
	docker build -f Dockerfile . -t ${IMG}

.PHONY: docker-push
docker-push:
	docker push ${IMG}

.PHONY: enable
enable:
	vault secrets enable -path=kubernetes-secrets-reader vault-plugin-secrets-kubernetes-reader 

.PHONY: clean
clean:
	rm -f ./vault/plugins/vault-plugin-secrets-kubernetes-reader

.PHONY: test
test:
	$(MAKE) kind
	cd testdata && ./setup.sh
	go test -v ./...

include hack/make-rules/verify.mk
include hack/make-rules/tools.mk
include hack/make-rules/cluster.mk

