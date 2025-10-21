# SassyShell TUI Makefile

BINARY_NAME=sassysh-tui
BUILD_DIR=build
INSTALL_DIR=$(HOME)/.local/bin

.PHONY: all build clean install uninstall test deps help

# Default target
all: build

# Build the binary
build:
	@echo "🔨 Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "✅ Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Install dependencies
deps:
	@echo "📦 Installing dependencies..."
	go mod tidy
	go mod download
	@echo "✅ Dependencies installed"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	go clean
	@echo "✅ Clean complete"

# Install binary to user's local bin
install: build
	@echo "📥 Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	@mkdir -p $(INSTALL_DIR)
	cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/
	@echo "✅ Installation complete"
	@echo "💡 Add $(INSTALL_DIR) to your PATH if not already added"
	@echo "🚀 Run: $(BINARY_NAME) setup"

# Uninstall binary
uninstall:
	@echo "🗑️  Uninstalling $(BINARY_NAME)..."
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "✅ Uninstall complete"

# Run tests
test:
	@echo "🧪 Running tests..."
	go test ./...
	@echo "✅ Tests complete"

# Development build (with race detection)
dev:
	@echo "🔧 Building development version..."
	go build -race -o $(BUILD_DIR)/$(BINARY_NAME)-dev .
	@echo "✅ Development build complete: $(BUILD_DIR)/$(BINARY_NAME)-dev"

# Cross-compile for different platforms
build-all:
	@echo "🌍 Cross-compiling for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	
	# Linux AMD64
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 .
	
	# macOS AMD64
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 .
	
	# macOS ARM64 (Apple Silicon)
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 .
	
	# Windows AMD64
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe .
	
	@echo "✅ Cross-compilation complete"
	@ls -la $(BUILD_DIR)/

# Show help
help:
	@echo "SassyShell TUI Build System"
	@echo ""
	@echo "Available targets:"
	@echo "  build      - Build the binary"
	@echo "  deps       - Install Go dependencies"
	@echo "  clean      - Clean build artifacts"
	@echo "  install    - Install binary to ~/.local/bin"
	@echo "  uninstall  - Remove installed binary"
	@echo "  test       - Run tests"
	@echo "  dev        - Build development version with race detection"
	@echo "  build-all  - Cross-compile for multiple platforms"
	@echo "  help       - Show this help message"
	@echo ""
	@echo "Quick start:"
	@echo "  make deps && make build"
	@echo "  ./build/$(BINARY_NAME) setup"
