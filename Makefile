.PHONY: all rust go go-deps java

BUILDER_CONTAINER=namely/protoc-all:1.51_2

all: rust go java cpp

rust:
	cargo build

go: go-deps
	docker run --rm -t -w /greptime-proto \
		--entrypoint ./scripts/generate-go.sh \
		-v ${PWD}:/greptime-proto ${BUILDER_CONTAINER}

go-deps:
	go mod download

java:
	docker run --rm -t -w /greptime-proto \
		--entrypoint ./scripts/generate-java.sh \
		-v ${PWD}:/greptime-proto ${BUILDER_CONTAINER}

cpp:
	docker run --rm -t -w /greptime-proto \
		--entrypoint ./scripts/generate-cpp.sh \
		-v ${PWD}:/greptime-proto ${BUILDER_CONTAINER}
