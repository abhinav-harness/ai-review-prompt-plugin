# Changelog

All notable changes to the Drone AI Review Plugin will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2026-01-12

### Added
- Initial release of Drone AI Review Plugin
- Generate AI code review prompts from pull request diffs
- Configurable review types:
  - Bug detection (`enable_bugs`)
  - Performance analysis (`enable_performance`)
  - Scalability concerns (`enable_scalability`)
  - Code smell detection (`enable_code_smell`)
- Support for custom review rules via `.harness/rules/review.md`
- Configurable maximum comment count
- Customizable output directory
- Docker-based plugin for easy Drone CI integration
- Multi-stage Docker build for optimized image size
- **Multi-architecture support**: `linux/amd64`, `linux/arm64`, `linux/arm/v7`
- Multi-arch build script (`build-multiarch.sh`)
- Comprehensive documentation and usage examples
- Test script for local development
- Git diff command with line number extraction
- Structured JSON output format specification
- Code suggestion markdown support

### Features
- Automatic detection of Drone environment variables
- Fallback to default values for all configuration options
- Template-based prompt generation with Go templates
- Error handling and directory creation
- Detailed logging of configuration and execution

### Documentation
- README.md with complete setup and usage instructions
- USAGE.md with detailed examples and integration guides
- plugin.yml with Drone plugin metadata
- .drone.yml example pipeline configuration
- LICENSE file (MIT)

[1.0.0]: https://github.com/abhinav-harness/ai-review-prompt-plugin/releases/tag/v1.0.0

