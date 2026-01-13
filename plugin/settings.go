package plugin

import (
	"os"
	"strconv"
)

// Settings defines the plugin input parameters
type Settings struct {
	// Git repository information
	RepoName     string
	SourceBranch string
	TargetBranch string
	MergeBaseSha string
	SourceSha    string

	// Review type flags (all enabled by default)
	EnableBugs        bool
	EnablePerformance bool
	EnableScalability bool
	EnableCodeSmell   bool

	// Review configuration
	CommentCount     int
	OutputFile       string
	ReviewOutputFile string
	CustomRulesPath  string
}

// NewSettings creates a new Settings instance from environment variables
func NewSettings() Settings {
	return Settings{
		RepoName:     getEnv("PLUGIN_REPO_NAME", getEnv("DRONE_REPO_NAME", "")),
		SourceBranch: getEnv("PLUGIN_SOURCE_BRANCH", getEnv("DRONE_SOURCE_BRANCH", "")),
		TargetBranch: getEnv("PLUGIN_TARGET_BRANCH", getEnv("DRONE_TARGET_BRANCH", "")),
		MergeBaseSha: getEnv("PLUGIN_MERGE_BASE_SHA", getEnv("DRONE_COMMIT_BEFORE", "")),
		SourceSha:    getEnv("PLUGIN_SOURCE_SHA", getEnv("DRONE_COMMIT_SHA", "")),

		EnableBugs:        getBoolEnv("PLUGIN_ENABLE_BUGS", true),
		EnablePerformance: getBoolEnv("PLUGIN_ENABLE_PERFORMANCE", true),
		EnableScalability: getBoolEnv("PLUGIN_ENABLE_SCALABILITY", true),
		EnableCodeSmell:   getBoolEnv("PLUGIN_ENABLE_CODE_SMELL", true),

		CommentCount:     getIntEnv("PLUGIN_COMMENT_COUNT", 10),
		OutputFile:       getEnv("PLUGIN_OUTPUT_FILE", "../output/task.txt"),
		ReviewOutputFile: getEnv("PLUGIN_REVIEW_OUTPUT_FILE", "../output/review.json"),
		CustomRulesPath:  getEnv("PLUGIN_CUSTOM_RULES_PATH", ".harness/rules/review.md"),
	}
}

// getEnv retrieves an environment variable with a default fallback
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getBoolEnv retrieves a boolean environment variable with a default fallback
func getBoolEnv(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}

// getIntEnv retrieves an integer environment variable with a default fallback
func getIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

