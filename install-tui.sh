#!/bin/bash

# SassyShell TUI Installation Script

set -e

echo "🚀 Installing SassyShell TUI..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher first."
    echo "   Visit: https://golang.org/doc/install"
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | cut -d' ' -f3 | sed 's/go//')
REQUIRED_VERSION="1.21"

if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]; then
    echo "❌ Go version $GO_VERSION is too old. Please upgrade to Go $REQUIRED_VERSION or higher."
    exit 1
fi

echo "✅ Go version $GO_VERSION detected"

# Navigate to TUI directory
cd "$(dirname "$0")/tui"

# Install dependencies
echo "📦 Installing dependencies..."
go mod tidy

# Build the binary
echo "🔨 Building sassysh-tui..."
go build -o build/sassysh-tui .

# Create installation directory
INSTALL_DIR="$HOME/.local/bin"
mkdir -p "$INSTALL_DIR"

# Install binary
echo "📥 Installing to $INSTALL_DIR..."
cp build/sassysh-tui "$INSTALL_DIR/"

# Make executable
chmod +x "$INSTALL_DIR/sassysh-tui"

echo ""
echo "✅ Installation complete!"
echo ""
echo "📋 Next steps:"
echo "   1. Add $INSTALL_DIR to your PATH if not already added:"
echo "      export PATH=\"\$HOME/.local/bin:\$PATH\""
echo ""
echo "   2. Run the setup wizard:"
echo "      sassysh-tui setup"
echo ""
echo "   3. Use the regular sassysh command for queries:"
echo "      sassysh ask \"your question\""
echo ""
echo "🎉 Happy coding with SassyShell TUI!"
