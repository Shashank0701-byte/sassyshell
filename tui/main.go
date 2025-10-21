package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7C3AED")).
		MarginBottom(1)

	subtitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6B7280")).
		MarginBottom(2)
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sassysh-tui <command>")
		fmt.Println("Commands:")
		fmt.Println("  setup    Run interactive setup wizard")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "setup":
		runSetupWizard()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func runSetupWizard() {
	fmt.Print(titleStyle.Render("🚀 SassyShell TUI Setup"))
	fmt.Print(subtitleStyle.Render("Let's configure your LLM provider with style!"))

	var (
		provider string
		model    string
		apiKey   string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your LLM provider").
				Description("Select the AI service you'd like to use").
				Options(
					huh.NewOption("Google Gemini", "google_genai").Selected(true),
					huh.NewOption("OpenAI GPT", "openai"),
					huh.NewOption("Ollama (Local)", "ollama"),
					huh.NewOption("Anthropic Claude", "anthropic"),
					huh.NewOption("Groq", "groq"),
				).
				Value(&provider),
		),

		huh.NewGroup(
			huh.NewInput().
				Title("Model Name").
				Description("Enter the specific model you want to use").
				Placeholder("gemini-2.5-flash").
				Value(&model).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("model name is required")
					}
					return nil
				}),
		),

		huh.NewGroup(
			huh.NewInput().
				Title("API Key").
				Description("Enter your API key (input will be hidden)").
				EchoMode(huh.EchoModePassword).
				Value(&apiKey).
				Validate(func(s string) error {
					if s == "" {
						return fmt.Errorf("API key is required")
					}
					if len(s) < 10 {
						return fmt.Errorf("API key seems too short")
					}
					return nil
				}),
		),
	).WithTheme(huh.ThemeCharm())

	err := form.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Set default model based on provider if not specified
	if model == "" {
		switch provider {
		case "google_genai":
			model = "gemini-2.5-flash"
		case "openai":
			model = "gpt-4"
		case "ollama":
			model = "llama2"
		case "anthropic":
			model = "claude-3-sonnet-20240229"
		case "groq":
			model = "mixtral-8x7b-32768"
		}
	}

	// Save configuration
	err = saveConfiguration(provider, model, apiKey)
	if err != nil {
		fmt.Printf("❌ Error saving configuration: %v\n", err)
		os.Exit(1)
	}

	// Success message
	successStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#10B981")).
		MarginTop(1)

	fmt.Print(successStyle.Render("✅ Setup complete!"))
	fmt.Println("\n📁 Configuration saved to ~/.config/sassyshell/.env")
	fmt.Println("🎉 You can now use sassysh ask to query your AI assistant!")
}
