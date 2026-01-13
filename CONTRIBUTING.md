# Contributing to Drone AI Review Plugin

Thank you for your interest in contributing! This document provides guidelines and instructions for contributing to the project.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for all contributors.

## How to Contribute

### Reporting Bugs

If you find a bug, please create an issue with:
- Clear description of the problem
- Steps to reproduce
- Expected vs actual behavior
- Environment details (Go version, Drone version, OS)
- Relevant logs or error messages

### Suggesting Features

Feature requests are welcome! Please provide:
- Clear description of the feature
- Use cases and benefits
- Potential implementation approach (if you have ideas)

### Pull Requests

1. **Fork the repository**
   ```bash
   git clone https://github.com/abhinav-harness/ai-review-prompt-plugin.git
   cd ai-review-prompt-plugin
   ```

2. **Create a feature branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Make your changes**
   - Follow the existing code style
   - Add tests if applicable
   - Update documentation

4. **Test your changes**
   ```bash
   # Run tests
   go test ./...
   
   # Run tests with verbose output
   go test -v ./plugin/
   
   # Run tests with coverage
   go test -cover ./plugin/
   
   # Build the binary
   go build -o drone-ai-review .
   
   # Test locally with environment variables
   export PLUGIN_REPO_NAME="test-repo"
   export PLUGIN_COMMENT_COUNT=15
   ./drone-ai-review
   
   # Build Docker image
   docker build -t drone-ai-review:test .
   ```

5. **Commit your changes**
   ```bash
   git add .
   git commit -m "feat: add new feature description"
   ```
   
   Use conventional commit messages:
   - `feat:` for new features
   - `fix:` for bug fixes
   - `docs:` for documentation changes
   - `refactor:` for code refactoring
   - `test:` for test additions/changes
   - `chore:` for maintenance tasks

6. **Push to your fork**
   ```bash
   git push origin feature/your-feature-name
   ```

7. **Create a Pull Request**
   - Provide a clear title and description
   - Reference any related issues
   - Ensure all checks pass

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Docker (for building images)
- Git

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/abhinav-harness/ai-review-prompt-plugin.git
   cd ai-review-prompt-plugin
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Build the project**
   ```bash
   go build -o drone-ai-review .
   ```

4. **Run tests**
   ```bash
   go test ./...
   ```

5. **Test locally**
   ```bash
   ./test-plugin.sh
   ```

## Project Structure

```
.
├── main.go              # Entry point
├── plugin/              # Plugin package
│   ├── settings.go      # Configuration and settings
│   ├── template.go      # Prompt template
│   └── writer.go        # File output logic
├── Dockerfile           # Docker build configuration
├── plugin.yml           # Drone plugin metadata
├── go.mod               # Go module definition
├── README.md            # Main documentation
├── USAGE.md             # Usage guide
└── build-multiarch.sh   # Multi-arch build script
```

## Coding Standards

### Go Code Style

- Follow standard Go formatting (`gofmt`, `goimports`)
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions focused and concise
- Handle errors appropriately

### Testing

- Add tests for new features
- Ensure existing tests pass
- Aim for good test coverage
- Use table-driven tests where appropriate

Example test structure:
```go
func TestSettings(t *testing.T) {
    tests := []struct {
        name     string
        envVars  map[string]string
        expected Settings
    }{
        // test cases
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Documentation

- Update README.md for major changes
- Add examples to USAGE.md
- Update CHANGELOG.md
- Document configuration options
- Include code comments for complex logic

## Adding New Features

### Adding a New Configuration Option

1. Add field to `Settings` struct in `plugin/settings.go`
2. Add getter function (e.g., `getBoolEnv`, `getEnv`)
3. Update `NewSettings()` function
4. Add to `plugin.yml` settings section
5. Update documentation in README.md
6. Add example in USAGE.md

### Modifying the Prompt Template

1. Edit `PromptTemplate` in `plugin/template.go`
2. Use Go template syntax for variables: `{{.VariableName}}`
3. Add conditional sections: `{{if .Flag}}...{{end}}`
4. Test with `test-plugin.sh`
5. Update documentation

### Adding Review Types

1. Add boolean flag to `Settings` struct
2. Update template with conditional section
3. Add to plugin.yml
4. Document in README.md
5. Add example usage in USAGE.md

## Release Process

1. Update version in CHANGELOG.md
2. Create a git tag: `git tag -a v1.x.x -m "Release v1.x.x"`
3. Push tag: `git push origin v1.x.x`
4. Build and push Docker image
5. Create GitHub release with changelog

## Questions?

If you have questions:
- Open an issue for discussion
- Check existing issues and documentation
- Reach out to maintainers

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

Thank you for contributing! 🎉

