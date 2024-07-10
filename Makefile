build:
	@go build -o bin/gokafka

run: build
	@./bin/gokafka

test:
	@go test ./... -v