lint:
	@staticcheck ./...

build:
	@go build -o bin/v3

run: build
	@./bin/v3

test:
	@go test -v ./...

node:
	@node tests/test.js