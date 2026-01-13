package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// WritePromptFile generates and writes the prompt file to the specified output file
func WritePromptFile(settings Settings) error {
	// Get the directory from the output file path
	outputDir := filepath.Dir(settings.OutputFile)

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Parse the template
	tmpl, err := template.New("prompt").Parse(PromptTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create the output file
	file, err := os.Create(settings.OutputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	// Execute the template with settings
	if err := tmpl.Execute(file, settings); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	fmt.Printf("Successfully generated prompt file at: %s\n", settings.OutputFile)
	return nil
}

