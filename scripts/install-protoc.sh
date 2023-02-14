#! /usr/bin/env bash

set -e

# Install protoc
LOCAL_DIR=bin
PROTOC_VERSION=21.7
DOWNLOAD_URL=https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}

# Create a directory to store the downloaded files.
if [ ! -d ${LOCAL_DIR} ]; then
  mkdir ${LOCAL_DIR}
fi

if [ -f "${LOCAL_DIR}/protoc/bin/protoc" ]; then
  echo "The protoc is already installed, skip the installation!"
  exit 0
fi

# Install protoc in macOS
if [ "$(uname)" == "Darwin" ]; then
    if [ "$(arch)" == "arm64" ]; then
      echo "Downloading macOS protoc ${PROTOC_VERSION} for arm64"
      curl ${DOWNLOAD_URL}/protoc-${PROTOC_VERSION}-osx-aarch_64.zip -o ${LOCAL_DIR}/protoc.zip -L
    else
      echo "Downloading macOS protoc ${PROTOC_VERSION} for x86_64"
      curl ${DOWNLOAD_URL}/protoc-${PROTOC_VERSION}-osx-x86_64.zip -o ${LOCAL_DIR}/protoc.zip -L
    fi
elif [ "$(uname -s)" == "Linux" ]; then
    if [ "$(uname -m)" == "aarch64" ]; then
      echo "Downloading Linux protoc ${PROTOC_VERSION} for arm64"
      curl ${DOWNLOAD_URL}/protoc-${PROTOC_VERSION}-linux-aarch_64.zip -o ${LOCAL_DIR}/protoc.zip -L
    else
      echo "Downloading Linux protoc ${PROTOC_VERSION} for x86_64"
      curl  ${DOWNLOAD_URL}/protoc-${PROTOC_VERSION}-linux-x86_64.zip -o ${LOCAL_DIR}/protoc.zip -L
    fi
else
    echo "Unsupported platform"
    exit 1
fi

echo "Unzipping protoc..."
unzip ${LOCAL_DIR}/protoc.zip -d ${LOCAL_DIR}/protoc
