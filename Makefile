.PHONY: build test

build:
	go build -o bin/gin-helloworld

test:
	go install gotest.tools/gotestsum@latest
	gotestsum --junitfile test-result/junit.xml ./...
	# go test -v ./...

run:
	go run main.go