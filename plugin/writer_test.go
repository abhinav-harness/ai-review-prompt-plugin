package plugin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWritePromptFile(t *testing.T) {
	tests := []struct {
		name        string
		settings    Settings
		wantError   bool
		checkOutput func(t *testing.T, outputDir string)
	}{
		{
			name: "basic prompt generation",
			settings: Settings{
				RepoName:          "test-repo",
				SourceBranch:      "feature",
				TargetBranch:      "main",
				MergeBaseSha:      "abc123",
				SourceSha:         "def456",
				EnableBugs:        true,
				EnablePerformance: true,
				EnableScalability: true,
				EnableCodeSmell:   true,
				CommentCount:      10,
				OutputDir:         filepath.Join(os.TempDir(), "test-output-1"),
				CustomRulesPath:   ".harness/rules/review.md",
			},
			wantError: false,
			checkOutput: func(t *testing.T, outputDir string) {
				content, err := os.ReadFile(filepath.Join(outputDir, "task.txt"))
				if err != nil {
					t.Fatalf("Failed to read output file: %v", err)
				}

				output := string(content)

				// Check for required content
				if !strings.Contains(output, "test-repo") {
					t.Error("Output should contain repo name")
				}
				if !strings.Contains(output, "abc123...def456") {
					t.Error("Output should contain git diff SHAs")
				}
				if !strings.Contains(output, "Look for critical bugs") {
					t.Error("Output should contain bug detection guideline")
				}
				if !strings.Contains(output, "Look for performance issues") {
					t.Error("Output should contain performance guideline")
				}
				if !strings.Contains(output, "Look for scalability issues") {
					t.Error("Output should contain scalability guideline")
				}
				if !strings.Contains(output, "Look for code smells") {
					t.Error("Output should contain code smell guideline")
				}
				if !strings.Contains(output, "10 comments") {
					t.Error("Output should contain comment count")
				}
				// Check that output directory is used in the review.json path
				expectedPath := filepath.Join(outputDir, "review.json")
				if !strings.Contains(output, expectedPath) {
					t.Errorf("Output should contain dynamic output path: %s", expectedPath)
				}
			},
		},
		{
			name: "selective review types",
			settings: Settings{
				RepoName:          "test-repo",
				SourceBranch:      "feature",
				TargetBranch:      "main",
				MergeBaseSha:      "abc123",
				SourceSha:         "def456",
				EnableBugs:        true,
				EnablePerformance: false,
				EnableScalability: false,
				EnableCodeSmell:   true,
				CommentCount:      15,
				OutputDir:         filepath.Join(os.TempDir(), "test-output-2"),
				CustomRulesPath:   ".harness/rules/review.md",
			},
			wantError: false,
			checkOutput: func(t *testing.T, outputDir string) {
				content, err := os.ReadFile(filepath.Join(outputDir, "task.txt"))
				if err != nil {
					t.Fatalf("Failed to read output file: %v", err)
				}

				output := string(content)

				// Check enabled guidelines
				if !strings.Contains(output, "Look for critical bugs") {
					t.Error("Output should contain bug detection guideline")
				}
				if !strings.Contains(output, "Look for code smells") {
					t.Error("Output should contain code smell guideline")
				}

				// Check disabled guidelines are not present
				if strings.Contains(output, "Look for performance issues") {
					t.Error("Output should not contain performance guideline when disabled")
				}
				if strings.Contains(output, "Look for scalability issues") {
					t.Error("Output should not contain scalability guideline when disabled")
				}

				if !strings.Contains(output, "15 comments") {
					t.Error("Output should contain custom comment count")
				}
			},
		},
		{
			name: "all review types disabled",
			settings: Settings{
				RepoName:          "test-repo",
				SourceBranch:      "feature",
				TargetBranch:      "main",
				MergeBaseSha:      "abc123",
				SourceSha:         "def456",
				EnableBugs:        false,
				EnablePerformance: false,
				EnableScalability: false,
				EnableCodeSmell:   false,
				CommentCount:      5,
				OutputDir:         filepath.Join(os.TempDir(), "test-output-3"),
				CustomRulesPath:   ".harness/rules/review.md",
			},
			wantError: false,
			checkOutput: func(t *testing.T, outputDir string) {
				content, err := os.ReadFile(filepath.Join(outputDir, "task.txt"))
				if err != nil {
					t.Fatalf("Failed to read output file: %v", err)
				}

				output := string(content)

				// Verify no specific review type guidelines
				if strings.Contains(output, "Look for critical bugs") {
					t.Error("Output should not contain bug guideline when disabled")
				}
				if strings.Contains(output, "Look for performance issues") {
					t.Error("Output should not contain performance guideline when disabled")
				}
				if strings.Contains(output, "Look for scalability issues") {
					t.Error("Output should not contain scalability guideline when disabled")
				}
				if strings.Contains(output, "Look for code smells") {
					t.Error("Output should not contain code smell guideline when disabled")
				}

				// Should still have general structure
				if !strings.Contains(output, "test-repo") {
					t.Error("Output should contain repo name")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clean up before test
			os.RemoveAll(tt.settings.OutputDir)

			// Run test
			err := WritePromptFile(tt.settings)

			// Check error expectation
			if (err != nil) != tt.wantError {
				t.Errorf("WritePromptFile() error = %v, wantError %v", err, tt.wantError)
				return
			}

			// Run output checks
			if !tt.wantError && tt.checkOutput != nil {
				tt.checkOutput(t, tt.settings.OutputDir)
			}

			// Clean up after test
			os.RemoveAll(tt.settings.OutputDir)
		})
	}
}

func TestWritePromptFileCreatesDirectory(t *testing.T) {
	tempDir := filepath.Join(os.TempDir(), "test-nested", "output", "dir")
	defer os.RemoveAll(filepath.Join(os.TempDir(), "test-nested"))

	settings := Settings{
		RepoName:          "test-repo",
		MergeBaseSha:      "abc123",
		SourceSha:         "def456",
		EnableBugs:        true,
		EnablePerformance: true,
		EnableScalability: true,
		EnableCodeSmell:   true,
		CommentCount:      10,
		OutputDir:         tempDir,
		CustomRulesPath:   ".harness/rules/review.md",
	}

	err := WritePromptFile(settings)
	if err != nil {
		t.Fatalf("WritePromptFile() failed: %v", err)
	}

	// Verify directory was created
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		t.Error("Output directory was not created")
	}

	// Verify file exists
	outputFile := filepath.Join(tempDir, "task.txt")
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Error("Output file was not created")
	}
}
