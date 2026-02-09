.PHONY: build run test clean

# Build the application
build:
	go build -o bin/youtube-tui ./cmd/youtube-tui

# Run the application
run:
	go run ./cmd/youtube-tui

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod download
