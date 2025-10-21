# SassyShell TUI

A beautiful Terminal User Interface (TUI) for SassyShell built with [Bubbletea](https://github.com/charmbracelet/bubbletea).

## Features

- 🎨 **Interactive Setup Wizard**: Beautiful, step-by-step configuration
- 🔒 **Secure Input**: Masked API key entry
- ✨ **Rich UI**: Modern terminal interface with colors and styling
- 🔄 **Compatible**: Uses same configuration format as Python version
- 🚀 **Fast**: Built in Go for optimal performance

## Installation

### Prerequisites

- Go 1.21 or higher
- Existing SassyShell Python installation

### Build from Source

```bash
cd tui/
go mod tidy
go build -o sassysh-tui
```

### Usage

```bash
# Run interactive setup wizard
./sassysh-tui setup

# The configuration will be saved to ~/.config/sassyshell/.env
# and can be used by the main sassysh command
```

## Commands

- `setup` - Interactive setup wizard with TUI

## Architecture

The TUI is designed to complement, not replace, the existing Python implementation:

- **TUI Commands**: Interactive, rich terminal experience
- **CLI Commands**: Fast, scriptable, non-interactive (unchanged)

## Configuration Compatibility

The TUI saves configuration in the exact same format and location as the Python version:

```
~/.config/sassyshell/.env
```

This ensures seamless integration with existing `sassysh ask` commands.

## Development

### Project Structure

```
tui/
├── main.go              # Entry point and main setup logic
├── config.go            # Configuration file handling
├── internal/
│   └── models/
│       └── setup.go     # Bubbletea models for setup wizard
├── go.mod               # Go module definition
└── README.md           # This file
```

### Dependencies

- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Styling and layout
- `github.com/charmbracelet/huh` - Form components

### Building

```bash
go mod tidy
go build -o sassysh-tui
```

### Testing

```bash
go test ./...
```

## Future Enhancements

- [ ] LLM model selector with descriptions
- [ ] Configuration validation with test API calls
- [ ] History browser TUI
- [ ] Settings management interface
- [ ] Plugin architecture for extensions

## Contributing

This TUI implementation follows the same contribution guidelines as the main SassyShell project. See the main [CONTRIBUTING.md](../CONTRIBUTING.md) for details.
