package plugin

import (
	"strings"
	"testing"
	"text/template"
)

func TestPromptTemplate(t *testing.T) {
	// Test that the template is valid and can be parsed
	tmpl, err := template.New("prompt").Parse(PromptTemplate)
	if err != nil {
		t.Fatalf("Failed to parse template: %v", err)
	}

	// Test with sample settings
	settings := Settings{
		RepoName:          "test-repo",
		SourceBranch:      "feature-branch",
		TargetBranch:      "main",
		MergeBaseSha:      "abc123",
		SourceSha:         "def456",
		EnableBugs:        true,
		EnablePerformance: true,
		EnableScalability: true,
		EnableCodeSmell:   true,
		CommentCount:      10,
		OutputDir:         "../output",
		CustomRulesPath:   ".harness/rules/review.md",
	}

	var result strings.Builder
	err = tmpl.Execute(&result, settings)
	if err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	output := result.String()

	// Verify template output contains expected content
	expectedStrings := []string{
		"test-repo",
		"abc123...def456",
		"git diff",
		"JSON response format",
		"code suggestion markdown",
		"../output/review.json", // Should contain the output directory path
		".harness/rules/review.md",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Template output missing expected string: %s", expected)
		}
	}
}

func TestPromptTemplateConditionals(t *testing.T) {
	tests := []struct {
		name             string
		settings         Settings
		shouldContain    []string
		shouldNotContain []string
	}{
		{
			name: "all review types enabled",
			settings: Settings{
				RepoName:          "test-repo",
				MergeBaseSha:      "abc",
				SourceSha:         "def",
				EnableBugs:        true,
				EnablePerformance: true,
				EnableScalability: true,
				EnableCodeSmell:   true,
				CommentCount:      10,
				CustomRulesPath:   ".harness/rules/review.md",
			},
			shouldContain: []string{
				"Look for critical bugs",
				"Look for performance issues",
				"Look for scalability issues",
				"Look for code smells",
			},
			shouldNotContain: []string{},
		},
		{
			name: "only bugs enabled",
			settings: Settings{
				RepoName:          "test-repo",
				MergeBaseSha:      "abc",
				SourceSha:         "def",
				EnableBugs:        true,
				EnablePerformance: false,
				EnableScalability: false,
				EnableCodeSmell:   false,
				CommentCount:      10,
				CustomRulesPath:   ".harness/rules/review.md",
			},
			shouldContain: []string{
				"Look for critical bugs",
			},
			shouldNotContain: []string{
				"Look for performance issues",
				"Look for scalability issues",
				"Look for code smells",
			},
		},
		{
			name: "performance and scalability only",
			settings: Settings{
				RepoName:          "test-repo",
				MergeBaseSha:      "abc",
				SourceSha:         "def",
				EnableBugs:        false,
				EnablePerformance: true,
				EnableScalability: true,
				EnableCodeSmell:   false,
				CommentCount:      15,
				CustomRulesPath:   ".harness/rules/review.md",
			},
			shouldContain: []string{
				"Look for performance issues",
				"Look for scalability issues",
			},
			shouldNotContain: []string{
				"Look for critical bugs",
				"Look for code smells",
			},
		},
		{
			name: "none enabled",
			settings: Settings{
				RepoName:          "test-repo",
				MergeBaseSha:      "abc",
				SourceSha:         "def",
				EnableBugs:        false,
				EnablePerformance: false,
				EnableScalability: false,
				EnableCodeSmell:   false,
				CommentCount:      5,
				CustomRulesPath:   ".harness/rules/review.md",
			},
			shouldContain: []string{
				"test-repo",
			},
			shouldNotContain: []string{
				"Look for critical bugs",
				"Look for performance issues",
				"Look for scalability issues",
				"Look for code smells",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl, err := template.New("prompt").Parse(PromptTemplate)
			if err != nil {
				t.Fatalf("Failed to parse template: %v", err)
			}

			var result strings.Builder
			err = tmpl.Execute(&result, tt.settings)
			if err != nil {
				t.Fatalf("Failed to execute template: %v", err)
			}

			output := result.String()

			// Check strings that should be present
			for _, expected := range tt.shouldContain {
				if !strings.Contains(output, expected) {
					t.Errorf("Output should contain: %s", expected)
				}
			}

			// Check strings that should not be present
			for _, unexpected := range tt.shouldNotContain {
				if strings.Contains(output, unexpected) {
					t.Errorf("Output should not contain: %s", unexpected)
				}
			}
		})
	}
}
