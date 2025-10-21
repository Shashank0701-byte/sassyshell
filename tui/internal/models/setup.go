package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SetupModel represents the state of the setup wizard
type SetupModel struct {
	step     int
	provider string
	model    string
	apiKey   string
	err      error
	quitting bool
}

// SetupStep represents different steps in the setup process
type SetupStep int

const (
	StepProvider SetupStep = iota
	StepModel
	StepAPIKey
	StepConfirm
	StepComplete
)

// NewSetupModel creates a new setup model
func NewSetupModel() SetupModel {
	return SetupModel{
		step: int(StepProvider),
	}
}

// Init initializes the model
func (m SetupModel) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model
func (m SetupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.step < int(StepComplete) {
				m.step++
			}
		case "esc":
			if m.step > 0 {
				m.step--
			}
		}
	}
	return m, nil
}

// View renders the current state of the model
func (m SetupModel) View() string {
	if m.quitting {
		return "👋 Setup cancelled. Run again anytime!\n"
	}

	var s string

	// Header
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7C3AED")).
		MarginBottom(1)

	s += headerStyle.Render("🚀 SassyShell Interactive Setup")
	s += "\n\n"

	// Progress indicator
	progressStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6B7280"))

	s += progressStyle.Render(fmt.Sprintf("Step %d of 4", m.step+1))
	s += "\n\n"

	// Step content
	switch SetupStep(m.step) {
	case StepProvider:
		s += m.renderProviderStep()
	case StepModel:
		s += m.renderModelStep()
	case StepAPIKey:
		s += m.renderAPIKeyStep()
	case StepConfirm:
		s += m.renderConfirmStep()
	case StepComplete:
		s += m.renderCompleteStep()
	}

	// Footer
	s += "\n\n"
	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#9CA3AF")).
		Italic(true)

	if m.step < int(StepComplete) {
		s += helpStyle.Render("Press Enter to continue • Esc to go back • Ctrl+C to quit")
	}

	return s
}

func (m SetupModel) renderProviderStep() string {
	return "Select your LLM provider:\n\n" +
		"1. Google Gemini (recommended)\n" +
		"2. OpenAI GPT\n" +
		"3. Ollama (local)\n" +
		"4. Anthropic Claude\n" +
		"5. Groq\n"
}

func (m SetupModel) renderModelStep() string {
	return fmt.Sprintf("Provider: %s\n\nEnter model name:", m.provider)
}

func (m SetupModel) renderAPIKeyStep() string {
	return fmt.Sprintf("Provider: %s\nModel: %s\n\nEnter your API key:", m.provider, m.model)
}

func (m SetupModel) renderConfirmStep() string {
	return fmt.Sprintf("Please confirm your settings:\n\n"+
		"Provider: %s\n"+
		"Model: %s\n"+
		"API Key: %s\n\n"+
		"Save configuration?", 
		m.provider, m.model, "***"+m.apiKey[len(m.apiKey)-4:])
}

func (m SetupModel) renderCompleteStep() string {
	successStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#10B981"))

	return successStyle.Render("✅ Setup complete!") + "\n\n" +
		"Configuration saved to ~/.config/sassyshell/.env\n" +
		"You can now use 'sassysh ask' to query your AI assistant!"
}
