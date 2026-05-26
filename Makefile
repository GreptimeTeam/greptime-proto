.PHONY: all rust go go-deps java cpp build-builder-image

PROTOC_CONTAINER=namely/protoc-all:1.51_2
BUILDER_CONTAINER=greptime/protoc-all-local:latest
BUILDER_DOCKERFILE=./docker/protoc-all/Dockerfile
BUILDER_CONTEXT=./docker/protoc-all

all: rust go java cpp

build-builder-image:
	docker build -f ${BUILDER_DOCKERFILE} -t ${BUILDER_CONTAINER} ${BUILDER_CONTEXT}

rust: build-builder-image
	docker run --rm -t -w /greptime-proto \
		--user $(shell id -u):$(shell id -g) \
		--entrypoint ./scripts/generate-rust.sh \
		-v ${PWD}:/greptime-proto ${BUILDER_CONTAINER}

go: go-deps
	docker run --rm -t -w /greptime-proto \
		--entrypoint ./scripts/generate-go.sh \
		-v ${PWD}:/greptime-proto ${PROTOC_CONTAINER}

go-deps:
	go mod download

java:
	docker run --rm -t -w /greptime-proto \
		--entrypoint ./scripts/generate-java.sh \
		-v ${PWD}:/greptime-proto ${PROTOC_CONTAINER}

cpp:
	docker run --rm -t -w /greptime-proto \
		--entrypoint ./scripts/generate-cpp.sh \
		-v ${PWD}:/greptime-proto ${PROTOC_CONTAINER}
