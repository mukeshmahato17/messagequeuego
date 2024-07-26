build:
	@go build -o bin/messagequeuego

run: build
	@./bin/messagequeuego

test:
	@go test ./... -v