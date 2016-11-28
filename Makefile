SHELL := /bin/bash

directory := $(shell echo ./releases/`date +'%Y.%m.%d'`)
release:
	mkdir -p $(directory)

	GOOS=darwin GOARCH=amd64 go build -o $(directory)/darwin-amd64 -v ./cmd/datadog
	GOOS=linux GOARCH=amd64 go build -o $(directory)/linux-amd64 -v ./cmd/datadog
