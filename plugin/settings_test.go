package plugin

import (
	"os"
	"testing"
)

func TestNewSettings(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected Settings
	}{
		{
			name: "default values",
			envVars: map[string]string{
				"PLUGIN_REPO_NAME": "test-repo",
			},
			expected: Settings{
				RepoName:          "test-repo",
				EnableBugs:        true,
				EnablePerformance: true,
				EnableScalability: true,
				EnableCodeSmell:   true,
				CommentCount:      10,
				OutputDir:         "../output",
				CustomRulesPath:   ".harness/rules/review.md",
			},
		},
		{
			name: "custom values",
			envVars: map[string]string{
				"PLUGIN_REPO_NAME":          "custom-repo",
				"PLUGIN_SOURCE_BRANCH":      "feature-branch",
				"PLUGIN_TARGET_BRANCH":      "main",
				"PLUGIN_MERGE_BASE_SHA":     "abc123",
				"PLUGIN_SOURCE_SHA":         "def456",
				"PLUGIN_ENABLE_BUGS":        "false",
				"PLUGIN_ENABLE_PERFORMANCE": "true",
				"PLUGIN_ENABLE_SCALABILITY": "false",
				"PLUGIN_ENABLE_CODE_SMELL":  "true",
				"PLUGIN_COMMENT_COUNT":      "25",
				"PLUGIN_OUTPUT_DIR":         "./custom-output",
				"PLUGIN_CUSTOM_RULES_PATH":  ".config/rules.md",
			},
			expected: Settings{
				RepoName:          "custom-repo",
				SourceBranch:      "feature-branch",
				TargetBranch:      "main",
				MergeBaseSha:      "abc123",
				SourceSha:         "def456",
				EnableBugs:        false,
				EnablePerformance: true,
				EnableScalability: false,
				EnableCodeSmell:   true,
				CommentCount:      25,
				OutputDir:         "./custom-output",
				CustomRulesPath:   ".config/rules.md",
			},
		},
		{
			name: "drone environment variables",
			envVars: map[string]string{
				"DRONE_REPO_NAME":     "drone-repo",
				"DRONE_SOURCE_BRANCH": "feature",
				"DRONE_TARGET_BRANCH": "develop",
				"DRONE_COMMIT_BEFORE": "before123",
				"DRONE_COMMIT_SHA":    "after456",
			},
			expected: Settings{
				RepoName:          "drone-repo",
				SourceBranch:      "feature",
				TargetBranch:      "develop",
				MergeBaseSha:      "before123",
				SourceSha:         "after456",
				EnableBugs:        true,
				EnablePerformance: true,
				EnableScalability: true,
				EnableCodeSmell:   true,
				CommentCount:      10,
				OutputDir:         "../output",
				CustomRulesPath:   ".harness/rules/review.md",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear environment
			os.Clearenv()

			// Set test environment variables
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			// Create settings
			settings := NewSettings()

			// Verify settings
			if settings.RepoName != tt.expected.RepoName {
				t.Errorf("RepoName = %v, want %v", settings.RepoName, tt.expected.RepoName)
			}
			if settings.SourceBranch != tt.expected.SourceBranch {
				t.Errorf("SourceBranch = %v, want %v", settings.SourceBranch, tt.expected.SourceBranch)
			}
			if settings.TargetBranch != tt.expected.TargetBranch {
				t.Errorf("TargetBranch = %v, want %v", settings.TargetBranch, tt.expected.TargetBranch)
			}
			if settings.MergeBaseSha != tt.expected.MergeBaseSha {
				t.Errorf("MergeBaseSha = %v, want %v", settings.MergeBaseSha, tt.expected.MergeBaseSha)
			}
			if settings.SourceSha != tt.expected.SourceSha {
				t.Errorf("SourceSha = %v, want %v", settings.SourceSha, tt.expected.SourceSha)
			}
			if settings.EnableBugs != tt.expected.EnableBugs {
				t.Errorf("EnableBugs = %v, want %v", settings.EnableBugs, tt.expected.EnableBugs)
			}
			if settings.EnablePerformance != tt.expected.EnablePerformance {
				t.Errorf("EnablePerformance = %v, want %v", settings.EnablePerformance, tt.expected.EnablePerformance)
			}
			if settings.EnableScalability != tt.expected.EnableScalability {
				t.Errorf("EnableScalability = %v, want %v", settings.EnableScalability, tt.expected.EnableScalability)
			}
			if settings.EnableCodeSmell != tt.expected.EnableCodeSmell {
				t.Errorf("EnableCodeSmell = %v, want %v", settings.EnableCodeSmell, tt.expected.EnableCodeSmell)
			}
			if settings.CommentCount != tt.expected.CommentCount {
				t.Errorf("CommentCount = %v, want %v", settings.CommentCount, tt.expected.CommentCount)
			}
			if settings.OutputDir != tt.expected.OutputDir {
				t.Errorf("OutputDir = %v, want %v", settings.OutputDir, tt.expected.OutputDir)
			}
			if settings.CustomRulesPath != tt.expected.CustomRulesPath {
				t.Errorf("CustomRulesPath = %v, want %v", settings.CustomRulesPath, tt.expected.CustomRulesPath)
			}
		})
	}
}

func TestGetBoolEnv(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		value        string
		defaultValue bool
		expected     bool
	}{
		{"true string", "TEST_BOOL", "true", false, true},
		{"false string", "TEST_BOOL", "false", true, false},
		{"1 as true", "TEST_BOOL", "1", false, true},
		{"0 as false", "TEST_BOOL", "0", true, false},
		{"empty uses default true", "TEST_BOOL", "", true, true},
		{"empty uses default false", "TEST_BOOL", "", false, false},
		{"invalid uses default", "TEST_BOOL", "invalid", true, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			if tt.value != "" {
				os.Setenv(tt.key, tt.value)
			}
			result := getBoolEnv(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("getBoolEnv(%s, %v) = %v, want %v", tt.key, tt.defaultValue, result, tt.expected)
			}
		})
	}
}

func TestGetIntEnv(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		value        string
		defaultValue int
		expected     int
	}{
		{"valid number", "TEST_INT", "42", 10, 42},
		{"zero", "TEST_INT", "0", 10, 0},
		{"negative", "TEST_INT", "-5", 10, -5},
		{"empty uses default", "TEST_INT", "", 10, 10},
		{"invalid uses default", "TEST_INT", "abc", 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			if tt.value != "" {
				os.Setenv(tt.key, tt.value)
			}
			result := getIntEnv(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("getIntEnv(%s, %v) = %v, want %v", tt.key, tt.defaultValue, result, tt.expected)
			}
		})
	}
}
