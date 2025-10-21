# SassyShell TUI Setup Guide

This document describes the new Terminal User Interface (TUI) implementation for SassyShell.

## Overview

The TUI provides a beautiful, interactive setup experience built with Go and the Bubbletea framework. It complements the existing Python CLI without replacing it.

## Features

- 🎨 **Interactive Setup Wizard**: Step-by-step configuration with rich UI
- 🔒 **Secure Input**: Masked API key entry
- ✨ **Modern Interface**: Colors, styling, and smooth navigation
- 🔄 **Full Compatibility**: Uses same config format as Python version
- 🚀 **Fast Performance**: Built in Go for optimal speed

## Installation

### Prerequisites

- Go 1.21 or higher
- Existing SassyShell Python installation (`pipx install sassyshell`)

### Quick Install (Windows)

```cmd
# Run the installation script
install-tui.bat
```

### Quick Install (Linux/macOS)

```bash
# Run the installation script
chmod +x install-tui.sh
./install-tui.sh
```

### Manual Installation

```bash
# Navigate to TUI directory
cd tui/

# Install dependencies
go mod tidy

# Build binary
go build -o build/sassysh-tui .

# Install to local bin (optional)
cp build/sassysh-tui ~/.local/bin/
```

## Usage

### Interactive Setup

```bash
# Run the beautiful TUI setup wizard
sassysh-tui setup
```

This will guide you through:
1. **Provider Selection**: Choose from Google Gemini, OpenAI, Ollama, etc.
2. **Model Configuration**: Specify the exact model name
3. **API Key Entry**: Secure, masked input
4. **Confirmation**: Review settings before saving

### Regular Usage

After setup, use the regular Python commands:

```bash
# Fast, non-TUI queries (unchanged)
sassysh ask "how to find files modified today"
```

## Architecture

```
SassyShell Ecosystem:
├── Python CLI (existing)     # Fast, scriptable commands
│   ├── sassysh ask          # Non-interactive queries
│   └── sassysh setup        # Simple prompts
└── Go TUI (new)             # Rich, interactive experience
    └── sassysh-tui setup    # Beautiful setup wizard
```

## Configuration Compatibility

Both implementations use the same configuration:

**Location**: `~/.config/sassyshell/.env`

**Format**:
```env
llm_model_provider=google_genai
llm_model_name=gemini-2.5-flash
llm_api_key=your_api_key_here
```

## Development

### Project Structure

```
tui/
├── main.go                  # Entry point with huh forms
├── config.go                # Configuration file handling
├── internal/
│   └── models/
│       └── setup.go         # Advanced Bubbletea models
├── go.mod                   # Go dependencies
├── Makefile                 # Build automation
└── README.md               # TUI-specific documentation
```

### Building

```bash
cd tui/
make build          # Build binary
make install        # Install to ~/.local/bin
make clean          # Clean build artifacts
make build-all      # Cross-compile for all platforms
```

## Future Enhancements

The TUI framework is designed to be extensible:

- [ ] **LLM Model Browser**: Interactive model selection with descriptions
- [ ] **Configuration Validator**: Test API connections during setup
- [ ] **History Browser**: Navigate past commands with TUI
- [ ] **Settings Manager**: Visual configuration editor
- [ ] **Plugin System**: Extensions for new providers

## Troubleshooting

### Go Not Found
```bash
# Install Go from https://golang.org/doc/install
# Ensure Go is in your PATH
go version
```

### Build Errors
```bash
# Clean and rebuild
cd tui/
make clean
make deps
make build
```

### Permission Issues
```bash
# Make sure installation directory exists
mkdir -p ~/.local/bin
# Add to PATH if needed
export PATH="$HOME/.local/bin:$PATH"
```

## Contributing

The TUI follows the same contribution guidelines as the main project:

1. **Zero Breaking Changes**: Existing Python CLI remains unchanged
2. **Additive Only**: New files only, no modifications to existing code
3. **Configuration Compatible**: Same format and location
4. **Extensible Design**: Framework for future TUI features

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed guidelines.
