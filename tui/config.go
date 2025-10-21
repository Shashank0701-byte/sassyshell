package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// saveConfiguration saves the LLM configuration to the same location as the Python version
func saveConfiguration(provider, model, apiKey string) error {
	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	// Create config directory path (same as Python version)
	configDir := filepath.Join(homeDir, ".config", "sassyshell")
	
	// Create directory if it doesn't exist
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Create .env file path
	envFile := filepath.Join(configDir, ".env")

	// Prepare configuration content (same format as Python version)
	content := fmt.Sprintf("llm_model_provider=%s\nllm_model_name=%s\nllm_api_key=%s\n", 
		provider, model, apiKey)

	// Write configuration file
	err = os.WriteFile(envFile, []byte(content), 0600) // 0600 for security (owner read/write only)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
