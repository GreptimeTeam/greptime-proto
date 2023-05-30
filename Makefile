.PHONY: all rust go go-deps java

BUILDER_CONTAINER=namely/protoc-all:1.51_1

all: rust go java

rust:
	cargo build

go: go-deps
	docker run -t -w /greptime-proto \
		--entrypoint ./scripts/generate-go.sh \
		-v ${PWD}:/greptime-proto ${BUILDER_CONTAINER}

go-deps:
	go mod download

java:
	docker run -t -w /greptime-proto \
		--entrypoint ./scripts/generate-java.sh \
		-v ${PWD}:/greptime-proto ${BUILDER_CONTAINER}
