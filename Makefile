.PHONY: build run test docker


BINARY=go-blockchain


build:
go build -o $(BINARY) ./cmd/node


run:
go run ./cmd/node


test:
go test ./... -v


docker:
docker build -t $(BINARY) .
