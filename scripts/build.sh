#!/usr/bin/env bash

set -e

build() {
  GOOS=$1 GOARCH=$2 OUT_DIR=$3 OUT_BIN=$4
  echo "→ Building $OUT_BIN ($OUT_DIR)"
  GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 \
    go build -o dist/$OUT_DIR/$OUT_BIN ./$BIN_SRC
}

build_all() {
  BIN=$1
  build linux amd64 linux-amd64 $BIN
  build linux arm64 linux-arm64 $BIN
  build darwin amd64 darwin-amd64 $BIN
  build darwin arm64 darwin-arm64 $BIN
  build windows amd64 windows-amd64 $BIN.exe
}

copy_config() {
  OUT_DIR=$1
  echo "→ Copying config.toml"
  cp ./configs/config.toml dist/$OUT_DIR/config.toml
}

BIN_SRC=cmd/dictionary
build_all dictionary

BIN_SRC=cmd/setup
build_all setup

copy_config linux-amd64
copy_config linux-arm64
copy_config darwin-amd64
copy_config darwin-arm64
copy_config windows-amd64
