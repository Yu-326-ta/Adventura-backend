BINARY := adventura_backend
MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
PATH := $(PATH):${MAKEFILE_DIR}bin
SHELL := env PATH="$(PATH)" /bin/bash

GOARCH = amd64

build:
	CGO_ENABLED=0 go build -o build/${BINARY} ./cmd

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o build/${BINARY}-linux-${GOARCH} ./cmd

mod:
	go mod download

.PHONY:	build build-linux mod
