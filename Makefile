PROJECT = $(shell basename $(CURDIR))

default: test

test:
	go test -v -race ./src/...

dep:
	go get -u github.com/golang/dep/cmd/dep

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(PROJECT) ./cmd/$(PROJECT);

local-build:
	@go build -o build/$(PROJECT) ./cmd/$(PROJECT);

run: local-build
	./build/$(PROJECT) -arn=$(AWS_CLOUD_DIRECTORY_ARN) -schemaArn=$(AWS_CLOUD_DIRECTORY_SCHEMA_ARN)

.PHONY: default test dep build local-build run
