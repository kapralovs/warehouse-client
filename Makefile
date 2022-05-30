.PHONY: build
build:
	go build -v ./cmd/app

.PHONY: run
run:
	go run ./cmd/app admin admin
	
.DEFAULT_GOAL := build