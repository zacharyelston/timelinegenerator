BINARY_NAME=timelinegenerator
VERSION=$(shell git describe --tags --always --dirty)
BUILD_FLAGS=-ldflags "-X main.version=${VERSION}"

.PHONY: all build clean test lint

all: lint test build

build:
	go build ${BUILD_FLAGS} -o ${BINARY_NAME}

test:
	go test -v ./...

lint:
	golangci-lint run

clean:
	go clean
	rm -f ${BINARY_NAME}

install:
	go install ${BUILD_FLAGS}

run:
	go run main.go