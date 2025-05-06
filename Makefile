.PHONY: build test

build:
	go build -o bin/gin-helloworld

test:
	go test -v ./...

run:
	go run main.go