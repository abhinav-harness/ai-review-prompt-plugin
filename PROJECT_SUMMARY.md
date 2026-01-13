# Drone AI Review Plugin - Project Summary

## Repository Information

- **GitHub Repository**: [https://github.com/abhinav-harness/ai-review-prompt-plugin](https://github.com/abhinav-harness/ai-review-prompt-plugin)
- **Container Registry**: Build and push to your own registry
- **License**: MIT
- **Language**: Go 1.21+

## Overview

A Drone CI plugin that generates AI-powered code review prompts from pull request diffs with configurable review types and multi-architecture Docker support.

## Key Features

### ✅ Multi-Architecture Support
- **linux/amd64** (Intel/AMD 64-bit)
- **linux/arm64** (ARM 64-bit)
- **linux/arm/v7** (ARM 32-bit)

### ✅ Configurable Review Types (All with Toggle Knobs)
- **Bug Detection** (`enable_bugs`) - Default: `true`
- **Performance Analysis** (`enable_performance`) - Default: `true`
- **Scalability Concerns** (`enable_scalability`) - Default: `true`
- **Code Smell Detection** (`enable_code_smell`) - Default: `true`

### ✅ Flexible Configuration
- Maximum comment count (default: 10)
- Custom output directory (default: `../output`)
- Custom rules file path (default: `.harness/rules/review.md`)
- Auto-detection of Drone environment variables

## Project Structure

```
ai-review-plugin-prompt/
├── plugin/
│   ├── settings.go                 # Configuration parser
│   ├── settings_test.go            # Settings unit tests
│   ├── template.go                 # Prompt template
│   ├── template_test.go            # Template unit tests
│   ├── writer.go                   # File output logic
│   └── writer_test.go              # Writer unit tests
├── main.go                         # Plugin entry point
├── Dockerfile                      # Multi-arch Docker build
├── plugin.yml                      # Drone plugin metadata
├── go.mod                          # Go module
├── build-multiarch.sh             # Multi-arch build script
├── README.md                       # Main documentation
├── USAGE.md                        # Usage examples
├── SETUP.md                        # Setup & deployment guide
├── CONTRIBUTING.md                 # Contribution guidelines
├── CHANGELOG.md                    # Version history
├── LICENSE                         # MIT License
├── .gitignore                      # Git ignore patterns
└── .dockerignore                   # Docker ignore patterns
```

## Multi-Architecture Implementation

### Dockerfile Changes
- Uses `--platform=$BUILDPLATFORM` for cross-compilation
- Supports `TARGETOS`, `TARGETARCH`, `TARGETVARIANT` build arguments
- Optimized with `-ldflags="-w -s"` for smaller binaries
- Multi-stage build for minimal final image size

### Build Script (`build-multiarch.sh`)
- Automated multi-arch builds with Docker Buildx
- Support for push to registry or load locally
- Configurable platforms, registry, and image name
- Environment variable configuration

### Multi-Arch Build Script
- Build for multiple architectures using `./build-multiarch.sh`
- Supports linux/amd64, linux/arm64, linux/arm/v7
- Push to your own registry with `PUSH=true`
- Automatic semantic versioning from git tags

## Usage Examples

### Basic Drone Pipeline

```yaml
kind: pipeline
type: docker
name: code-review

steps:
  - name: generate-review-prompt
    image: your-registry/drone-ai-review:latest
    settings:
      enable_bugs: true
      enable_performance: true
      comment_count: 15

trigger:
  event:
    - pull_request
```

### Custom Configuration

```yaml
steps:
  - name: review
    image: your-registry/drone-ai-review:latest
    settings:
      enable_bugs: true
      enable_performance: true
      enable_scalability: false
      enable_code_smell: true
      comment_count: 20
      output_dir: /drone/src/review-output
      custom_rules_path: .config/review-rules.md
```

## Building & Deployment

### Local Build
```bash
go build -o drone-ai-review .
```

### Docker Build (Single Arch)
```bash
docker build -t drone-ai-review:latest .
```

### Multi-Arch Build
```bash
# Using the script
./build-multiarch.sh

# Build and push
PUSH=true REGISTRY=your-registry ./build-multiarch.sh

# Manual
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 \
  -t your-registry/drone-ai-review:latest --push .
```

## Output

### Generated Files

1. **`task.txt`** - AI review prompt containing:
   - Git diff command with line number parsing
   - Review guidelines based on enabled types
   - Custom rules integration
   - JSON output format specification
   - Code suggestion markdown templates

2. **`review.json`** - Expected output from AI (created by AI model):
   ```json
   {
     "reviews": [
       {
         "file_path": "path/to/file",
         "line_number_start": 123,
         "line_number_end": 125,
         "type": "bug|performance|scalability|code_smell",
         "review": "Review comment with suggestions"
       }
     ]
   }
   ```

## Configuration Parameters

| Parameter | Environment Variable | Default | Description |
|-----------|---------------------|---------|-------------|
| `repo_name` | `PLUGIN_REPO_NAME` | auto | Repository name |
| `source_branch` | `PLUGIN_SOURCE_BRANCH` | auto | Source branch |
| `target_branch` | `PLUGIN_TARGET_BRANCH` | auto | Target branch |
| `merge_base_sha` | `PLUGIN_MERGE_BASE_SHA` | auto | Merge base SHA |
| `source_sha` | `PLUGIN_SOURCE_SHA` | auto | Source SHA |
| `enable_bugs` | `PLUGIN_ENABLE_BUGS` | `true` | Bug detection |
| `enable_performance` | `PLUGIN_ENABLE_PERFORMANCE` | `true` | Performance reviews |
| `enable_scalability` | `PLUGIN_ENABLE_SCALABILITY` | `true` | Scalability reviews |
| `enable_code_smell` | `PLUGIN_ENABLE_CODE_SMELL` | `true` | Code smell detection |
| `comment_count` | `PLUGIN_COMMENT_COUNT` | `10` | Max comments |
| `output_dir` | `PLUGIN_OUTPUT_DIR` | `../output` | Output directory |
| `custom_rules_path` | `PLUGIN_CUSTOM_RULES_PATH` | `.harness/rules/review.md` | Rules file |

## Testing

### Unit Tests
```bash
# Run all tests
go test ./...

# Verbose output
go test -v ./plugin/

# With coverage
go test -cover ./plugin/
```

### Local Testing
```bash
# Build and run locally
go build -o drone-ai-review .
export PLUGIN_REPO_NAME="test-repo"
export PLUGIN_COMMENT_COUNT=15
./drone-ai-review
```

### Docker Testing
```bash
docker run --rm \
  -e PLUGIN_REPO_NAME="test-repo" \
  -e PLUGIN_COMMENT_COUNT=15 \
  -v $(pwd)/output:/workspace/../output \
  your-registry/drone-ai-review:latest
```

## Documentation

- **[README.md](README.md)** - Main documentation with features and setup
- **[USAGE.md](USAGE.md)** - Detailed usage examples and integrations
- **[SETUP.md](SETUP.md)** - Setup and deployment guide
- **[CONTRIBUTING.md](CONTRIBUTING.md)** - Contribution guidelines
- **[CHANGELOG.md](CHANGELOG.md)** - Version history

## Next Steps

1. **Push to GitHub**:
   ```bash
   git init
   git add .
   git commit -m "Initial commit: Drone AI Review Plugin with multi-arch support"
   git remote add origin https://github.com/abhinav-harness/ai-review-prompt-plugin.git
   git push -u origin main
   ```

2. **Create First Release**:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

3. **Build Multi-Arch Images**: Run `./build-multiarch.sh` to build for all platforms

4. **Push to Registry**: Use `PUSH=true ./build-multiarch.sh` to push images

5. **Test in Drone**: Use the published image in your Drone pipeline

5. **Configure Custom Rules**: Add `.harness/rules/review.md` to your repositories

## Support & Contributing

- **Issues**: https://github.com/abhinav-harness/ai-review-prompt-plugin/issues
- **Pull Requests**: Welcome! See [CONTRIBUTING.md](CONTRIBUTING.md)
- **License**: MIT - See [LICENSE](LICENSE)

---

**Built with ❤️ for automated code reviews**

