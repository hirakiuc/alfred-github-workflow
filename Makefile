.DEFAULT_GOAL := default

build:
	go build

install:
	go install

check:
	go vet . ./internal/...
	golint ./internal/...
	golint main.go

clean:
	go clean

default:
	make check
	make build

.PHONY: build install check clean
