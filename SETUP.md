# Setup and Deployment Guide

## Project Information

**Repository**: [https://github.com/abhinav-harness/ai-review-prompt-plugin](https://github.com/abhinav-harness/ai-review-prompt-plugin)

**Container Registry**: Build and push to your own registry

## Multi-Architecture Support

This plugin supports the following architectures:
- **linux/amd64** - Intel/AMD 64-bit (x86_64)
- **linux/arm64** - ARM 64-bit (ARM v8)
- **linux/arm/v7** - ARM 32-bit (ARM v7)

## Quick Start

### 1. Pull the Pre-built Image

```bash
docker pull your-registry/drone-ai-review:latest
```

The correct architecture will be automatically selected based on your platform.

### 2. Use in Drone CI

Add to your `.drone.yml`:

```yaml
kind: pipeline
type: docker
name: ai-code-review

steps:
  - name: generate-review-prompt
    image: your-registry/drone-ai-review:latest
    settings:
      enable_bugs: true
      enable_performance: true
      enable_scalability: true
      enable_code_smell: true
      comment_count: 15

trigger:
  event:
    - pull_request
```

## Building from Source

### Standard Build (Current Architecture)

```bash
# Clone the repository
git clone https://github.com/abhinav-harness/ai-review-prompt-plugin.git
cd ai-review-prompt-plugin

# Build the binary
go build -o drone-ai-review .

# Build Docker image
docker build -t drone-ai-review:latest .
```

### Multi-Architecture Build

#### Using the Automated Script

```bash
# Dry run (no push)
./build-multiarch.sh

# Build and load to local Docker (amd64 only)
LOAD=true ./build-multiarch.sh

# Build and push to registry
PUSH=true REGISTRY=your-registry IMAGE_NAME=drone-ai-review ./build-multiarch.sh

# Custom platforms
PLATFORMS=linux/amd64,linux/arm64 PUSH=true ./build-multiarch.sh
```

#### Manual Multi-Arch Build

```bash
# Create buildx builder (one-time setup)
docker buildx create --name multiarch-builder --use
docker buildx inspect --bootstrap

# Build for all platforms
docker buildx build \
  --platform linux/amd64,linux/arm64,linux/arm/v7 \
  -t your-registry/drone-ai-review:latest \
  --push .
```

## Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PLUGIN_REPO_NAME` | Auto-detected | Repository name |
| `PLUGIN_SOURCE_BRANCH` | Auto-detected | Source branch |
| `PLUGIN_TARGET_BRANCH` | Auto-detected | Target branch |
| `PLUGIN_MERGE_BASE_SHA` | Auto-detected | Merge base SHA |
| `PLUGIN_SOURCE_SHA` | Auto-detected | Source SHA |
| `PLUGIN_ENABLE_BUGS` | `true` | Enable bug detection |
| `PLUGIN_ENABLE_PERFORMANCE` | `true` | Enable performance reviews |
| `PLUGIN_ENABLE_SCALABILITY` | `true` | Enable scalability reviews |
| `PLUGIN_ENABLE_CODE_SMELL` | `true` | Enable code smell detection |
| `PLUGIN_COMMENT_COUNT` | `10` | Max comments per PR |
| `PLUGIN_OUTPUT_DIR` | `../output` | Output directory |
| `PLUGIN_CUSTOM_RULES_PATH` | `.harness/rules/review.md` | Custom rules file |

### Drone Settings

All environment variables can be configured via Drone settings:

```yaml
settings:
  enable_bugs: true
  enable_performance: true
  enable_scalability: false
  enable_code_smell: true
  comment_count: 20
  output_dir: /drone/src/output
```

## Testing

### Running Unit Tests

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./plugin/

# Run with coverage report
go test -cover ./plugin/

# Generate coverage HTML report
go test -coverprofile=coverage.out ./plugin/
go tool cover -html=coverage.out
```

### Local Testing

```bash
# Build the binary
go build -o drone-ai-review .

# Set up environment
export PLUGIN_REPO_NAME="my-repo"
export PLUGIN_SOURCE_BRANCH="feature-branch"
export PLUGIN_TARGET_BRANCH="main"
export PLUGIN_MERGE_BASE_SHA="abc123"
export PLUGIN_SOURCE_SHA="def456"
export PLUGIN_COMMENT_COUNT=15

# Run the plugin
./drone-ai-review
```

### Docker Testing

```bash
docker run --rm \
  -e PLUGIN_REPO_NAME="test-repo" \
  -e PLUGIN_SOURCE_BRANCH="feature" \
  -e PLUGIN_TARGET_BRANCH="main" \
  -e PLUGIN_MERGE_BASE_SHA="abc123" \
  -e PLUGIN_SOURCE_SHA="def456" \
  -e PLUGIN_COMMENT_COUNT=15 \
  -v $(pwd)/output:/workspace/../output \
  your-registry/drone-ai-review:latest
```

## Deployment Checklist

- [ ] Fork/clone the repository
- [ ] Update go.mod with your module path (if customizing)
- [ ] Run tests: `go test -v ./plugin/`
- [ ] Test locally with environment variables
- [ ] Build Docker image: `docker build -t drone-ai-review .`
- [ ] Test Docker image locally
- [ ] Build multi-arch images: `./build-multiarch.sh`
- [ ] Push to your registry
- [ ] Update Drone pipeline to use your image
- [ ] Test in Drone CI environment

## Troubleshooting

### Docker Buildx Not Available

```bash
# Install Docker Desktop (includes buildx) or enable buildx
docker buildx version
```

### Multi-Arch Build Fails

```bash
# Set up QEMU for emulation
docker run --privileged --rm tonistiigi/binfmt --install all

# Recreate builder
docker buildx rm multiarch-builder
docker buildx create --name multiarch-builder --use
```

### Image Pull Failed in Drone

- Ensure image name is correct and matches your registry
- For private registries, configure image pull secrets in Drone
- Verify the package is public in GitHub Container Registry settings

## Support

- **Issues**: https://github.com/abhinav-harness/ai-review-prompt-plugin/issues
- **Documentation**: See [README.md](README.md) and [USAGE.md](USAGE.md)
- **Contributing**: See [CONTRIBUTING.md](CONTRIBUTING.md)

