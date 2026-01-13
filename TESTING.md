# Testing Guide

## Overview

The Drone AI Review Plugin includes comprehensive Go unit tests with **87.9% code coverage**.

## Running Tests

### Using Make (Recommended)

```bash
# Run all tests
make test

# Run tests with coverage report (generates coverage.html)
make test-coverage

# Format code
make fmt

# Run all checks (clean, format, tidy, test, build)
make all
```

### Using Go Directly

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./plugin/

# Run with coverage
go test -cover ./plugin/

# Generate HTML coverage report
go test -coverprofile=coverage.out ./plugin/
go tool cover -html=coverage.out
```

## Test Files

### `plugin/settings_test.go`
Tests configuration parsing and environment variable handling:
- **TestNewSettings** - Tests settings creation from environment variables
  - Default values
  - Custom values via PLUGIN_* environment variables
  - Drone CI environment variables (DRONE_*)
- **TestGetBoolEnv** - Tests boolean environment variable parsing
- **TestGetIntEnv** - Tests integer environment variable parsing

**Coverage**: All environment variable parsing logic and default values

### `plugin/template_test.go`
Tests prompt template generation:
- **TestPromptTemplate** - Tests template parsing and execution
- **TestPromptTemplateConditionals** - Tests conditional sections
  - All review types enabled
  - Selective review types
  - No review types enabled
  - Mix of enabled/disabled types

**Coverage**: Template parsing, variable interpolation, conditional rendering

### `plugin/writer_test.go`
Tests file output functionality:
- **TestWritePromptFile** - Tests prompt file generation
  - Basic prompt generation with all features
  - Selective review types
  - All review types disabled
  - Content verification
- **TestWritePromptFileCreatesDirectory** - Tests directory creation
  - Nested directory creation
  - File output verification

**Coverage**: File I/O, directory creation, template execution, error handling

## Test Coverage

Current coverage: **87.9%**

```bash
$ go test -cover ./plugin/
ok      github.com/abhinav-harness/ai-review-prompt-plugin/plugin       0.574s  coverage: 87.9% of statements
```

### Coverage by File
- `settings.go` - Environment variable parsing and configuration
- `template.go` - Prompt template with conditionals
- `writer.go` - File output and directory creation

## Local Testing

### Build and Run Locally

```bash
# Build the binary
make build

# Set up test environment
export PLUGIN_REPO_NAME="test-repo"
export PLUGIN_SOURCE_BRANCH="feature-branch"
export PLUGIN_TARGET_BRANCH="main"
export PLUGIN_MERGE_BASE_SHA="abc123def456"
export PLUGIN_SOURCE_SHA="789ghi012jkl"
export PLUGIN_COMMENT_COUNT=15
export PLUGIN_ENABLE_BUGS=true
export PLUGIN_ENABLE_PERFORMANCE=true
export PLUGIN_ENABLE_SCALABILITY=false
export PLUGIN_ENABLE_CODE_SMELL=true
export PLUGIN_OUTPUT_DIR="./output"

# Run the plugin
./drone-ai-review

# Check the output
cat ./output/task.txt
```

### Docker Testing

```bash
# Build Docker image
make docker-build

# Run in Docker
docker run --rm \
  -e PLUGIN_REPO_NAME="test-repo" \
  -e PLUGIN_SOURCE_BRANCH="feature" \
  -e PLUGIN_TARGET_BRANCH="main" \
  -e PLUGIN_MERGE_BASE_SHA="abc123" \
  -e PLUGIN_SOURCE_SHA="def456" \
  -e PLUGIN_COMMENT_COUNT=15 \
  -v $(pwd)/output:/workspace/../output \
  drone-ai-review:latest
```

## Test Best Practices

### Adding New Tests

1. **Create test file** with `_test.go` suffix
2. **Use table-driven tests** for multiple scenarios
3. **Clean up** test artifacts (temporary files/directories)
4. **Test edge cases** and error conditions
5. **Verify coverage** remains high

Example test structure:
```go
func TestNewFeature(t *testing.T) {
    tests := []struct {
        name     string
        input    Input
        expected Output
        wantErr  bool
    }{
        {
            name:     "valid input",
            input:    validInput,
            expected: expectedOutput,
            wantErr:  false,
        },
        // more test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Running Specific Tests

```bash
# Run specific test
go test -v ./plugin/ -run TestNewSettings

# Run specific subtest
go test -v ./plugin/ -run TestNewSettings/default_values

# Run with race detection
go test -race ./plugin/

# Run with timeout
go test -timeout 30s ./plugin/
```

## Debugging Tests

### Verbose Output
```bash
go test -v ./plugin/
```

### Print Statements
Use `t.Log()` or `t.Logf()` instead of `fmt.Print()`:
```go
t.Logf("Debug info: %v", value)
```

### Run Single Test
```bash
go test -v ./plugin/ -run TestWritePromptFile
```

### Enable Test Caching
Go caches successful test runs. To clear cache:
```bash
go clean -testcache
```

## Coverage Goals

- **Minimum target**: 80% coverage
- **Current coverage**: 87.9%
- **Ideal target**: 90%+

Areas with full coverage:
- ✅ Environment variable parsing
- ✅ Template rendering
- ✅ File I/O operations
- ✅ Configuration validation

## Contributing Tests

When contributing:
1. Add tests for all new features
2. Maintain or improve coverage percentage
3. Ensure all tests pass: `make test`
4. Run coverage check: `make test-coverage`
5. Follow existing test patterns

## Troubleshooting

### Tests Fail Locally
```bash
# Clean and rebuild
make clean
make build
make test
```

### Permission Errors
```bash
# Ensure test directories are writable
chmod -R 755 .
```

### Stale Cache
```bash
go clean -testcache
go test ./...
```

For more information, see:
- [Testing in Go](https://golang.org/pkg/testing/)
- [Go Test Coverage](https://go.dev/blog/cover)

