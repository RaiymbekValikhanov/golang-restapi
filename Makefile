.PHONY: build
build:
	go build -o ./bin/ -v ./cmd/apiserver

.DEFAULT_GOAL := build
