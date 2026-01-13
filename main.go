package main

import (
	"fmt"
	"os"

	"github.com/abhinav-harness/ai-review-prompt-plugin/plugin"
)

func main() {
	// Parse settings from environment variables
	settings := plugin.NewSettings()

	// Display configuration
	fmt.Println("Drone AI Review Plugin")
	fmt.Println("======================")
	fmt.Printf("Repository: %s\n", settings.RepoName)
	fmt.Printf("Source Branch: %s\n", settings.SourceBranch)
	fmt.Printf("Target Branch: %s\n", settings.TargetBranch)
	fmt.Printf("Merge Base SHA: %s\n", settings.MergeBaseSha)
	fmt.Printf("Source SHA: %s\n", settings.SourceSha)
	fmt.Printf("Output Directory: %s\n", settings.OutputDir)
	fmt.Printf("Comment Count: %d\n", settings.CommentCount)
	fmt.Printf("Enable Bugs: %v\n", settings.EnableBugs)
	fmt.Printf("Enable Performance: %v\n", settings.EnablePerformance)
	fmt.Printf("Enable Scalability: %v\n", settings.EnableScalability)
	fmt.Printf("Enable Code Smell: %v\n", settings.EnableCodeSmell)
	fmt.Println("======================")

	// Generate and write the prompt file
	if err := plugin.WritePromptFile(settings); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Plugin execution completed successfully!")
}
