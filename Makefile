build:
	@go build -o bin/messagequeuego

run: build
	@./bin/gokafka

test:
	@go test ./... -v