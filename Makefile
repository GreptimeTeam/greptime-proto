.PHONY: all rust go go-deps java cpp check-protoc

BUILDER_CONTAINER=namely/protoc-all:1.51_2
PROTOC_VERSION=3.21.12

all: rust go java cpp

check-protoc:
	@if ! command -v protoc >/dev/null 2>&1; then \
		echo "Error: protoc is not installed. Please install protoc $(PROTOC_VERSION)."; \
		exit 1; \
	fi
	@CURRENT_PROTOC_VERSION=$$(protoc --version | awk '{print $$2}'); \
	if [ "$$CURRENT_PROTOC_VERSION" != "$(PROTOC_VERSION)" ]; then \
		echo "Error: Required protoc version is $(PROTOC_VERSION), but found $$CURRENT_PROTOC_VERSION."; \
		echo "Please install protoc $(PROTOC_VERSION) to ensure generated code is consistent."; \
		exit 1; \
	fi

rust: check-protoc
	cargo run --manifest-path xtask/Cargo.toml -- generate-rust

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
