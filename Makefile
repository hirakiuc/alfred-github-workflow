.DEFAULT_GOAL := default
ALFRED_WORKFLOW_DIR := ${HOME}/Library/Application\ Support/Alfred\ 3/Alfred.alfredpreferences/workflows
DIR_NAME := `basename ${PWD}`

NAME := gh

dev-deps:
	go get -u golang.org/x/lint/golint

deps:
	go mod download

build:
	go build -o $(NAME)

install:
	go install

check:
	golangci-lint run --enable-all -D dupl ./...

clean:
	go clean
	rm -f $(NAME)

default:
	make check
	make build

link:
	ln -sv $(pwd) ${ALFRED_WORKFLOW_DIR}/

unlink:
	rm -i ${ALFRED_WORKFLOW_DIR}/${DIR_NAME}

.PHONY: build install check clean
