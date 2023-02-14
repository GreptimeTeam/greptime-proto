.PHONY: all rust go go-deps install-protoc clean

all: rust go

rust:
	cargo build

go: install-protoc install-protobuf-go-plugins go-deps
	./scripts/generate-go.sh

go-deps:
	go mod download

install-protoc:
	./scripts/install-protoc.sh

install-protobuf-go-plugins:
	./scripts/install-protobuf-go-plugins.sh

clean:
	rm -rf bin
