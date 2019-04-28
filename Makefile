.DEFAULT_GOAL := default

NAME := gh

build:
	go build -o $(NAME)

install:
	go install

check:
	go vet . ./internal/...
	golint ./internal/...
	golint main.go

clean:
	go clean
	rm -f $(NAME)

default:
	make check
	make build

.PHONY: build install check clean
