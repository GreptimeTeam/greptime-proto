.PHONY: all rust go go-deps java cpp build-builder-image

BUILDER_CONTAINER=greptime/protoc-all-local:latest
BUILDER_DOCKERFILE=./docker/protoc-all/Dockerfile

all: rust go java cpp

build-builder-image:
	docker build -f ${BUILDER_DOCKERFILE} -t ${BUILDER_CONTAINER} .

rust: build-builder-image
	docker run --rm -t -w /greptime-proto \
		--entrypoint ./scripts/generate-rust.sh \
		-v ${PWD}:/greptime-proto ${BUILDER_CONTAINER}

go: build-builder-image go-deps
	docker run --rm -t -w /greptime-proto \
		--entrypoint ./scripts/generate-go.sh \
		-v ${PWD}:/greptime-proto ${BUILDER_CONTAINER}

go-deps:
	go mod download

java: build-builder-image
	docker run --rm -t -w /greptime-proto \
		--entrypoint ./scripts/generate-java.sh \
		-v ${PWD}:/greptime-proto ${BUILDER_CONTAINER}

cpp: build-builder-image
	docker run --rm -t -w /greptime-proto \
		--entrypoint ./scripts/generate-cpp.sh \
		-v ${PWD}:/greptime-proto ${BUILDER_CONTAINER}
