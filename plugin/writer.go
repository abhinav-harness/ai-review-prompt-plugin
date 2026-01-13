package plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// WritePromptFile generates and writes the prompt file to the specified output directory
func WritePromptFile(settings Settings) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(settings.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Parse the template
	tmpl, err := template.New("prompt").Parse(PromptTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create the output file
	outputPath := filepath.Join(settings.OutputDir, "task.txt")
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	// Execute the template with settings
	if err := tmpl.Execute(file, settings); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	fmt.Printf("Successfully generated prompt file at: %s\n", outputPath)
	return nil
}
