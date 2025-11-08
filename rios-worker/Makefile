.PHONY: build clean test install run-register run

# Binary name
BINARY_NAME=rios-worker

# Version info
VERSION ?= v0.1.0
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build flags
LDFLAGS=-ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT)"

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@CGO_ENABLED=0 go build $(LDFLAGS) -o $(BINARY_NAME) main.go
	@chmod +x $(BINARY_NAME)
	@echo "✅ Build complete: $(BINARY_NAME)"

# Build for all platforms
build-all:
	@echo "Building for multiple platforms..."
	@mkdir -p dist
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-amd64 main.go
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-linux-arm64 main.go
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-amd64 main.go
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-darwin-arm64 main.go
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o dist/$(BINARY_NAME)-windows-amd64.exe main.go
	@echo "✅ Multi-platform builds complete!"
	@ls -lh dist/

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -rf dist/
	@echo "✅ Clean complete"

# Install binary to $GOPATH/bin
install: build
	@echo "Installing to $(GOPATH)/bin/$(BINARY_NAME)..."
	@cp $(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME)
	@echo "✅ Installed successfully"

# Run register command (development)
run-register: build
	@./$(BINARY_NAME) register --api http://localhost:3000

# Run worker (development)
run: build
	@./$(BINARY_NAME) run --api http://localhost:3000

# Show help
help:
	@echo "RiOS Worker Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make build       - Build the binary"
	@echo "  make build-all   - Build for all platforms"
	@echo "  make deps        - Install dependencies"
	@echo "  make test        - Run tests"
	@echo "  make clean       - Remove build artifacts"
	@echo "  make install     - Install to \$$GOPATH/bin"
	@echo "  make run-register - Run register command (dev)"
	@echo "  make run         - Run worker (dev)"
	@echo "  make help        - Show this help"

