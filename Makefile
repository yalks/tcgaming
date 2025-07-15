.PHONY: test build clean fmt lint coverage

# Test the package
test:
	go test -v ./...

# Build the example
build:
	go build -o bin/example ./examples/main.go

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Format code
fmt:
	go fmt ./...
	gofmt -s -w .

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Run tests with coverage
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run the example
run-example: build
	./bin/example