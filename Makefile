.PHONY: all rust go go-deps

BUILDER_CONTAINER=namely/protoc-all:1.51_1

all: rust go

rust:
	cargo build

go: go-deps
	docker run -t --rm -w /greptime-proto \
    	--entrypoint ./scripts/generate-go.sh \
    	-v ${PWD}:/greptime-proto ${BUILDER_CONTAINER}

go-deps:
	go mod download
