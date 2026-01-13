#!/bin/bash

# Multi-architecture Docker build script for Drone AI Review Plugin
# Supports: linux/amd64, linux/arm64, linux/arm/v7

set -e

# Configuration
IMAGE_NAME="${IMAGE_NAME:-drone-ai-review}"
IMAGE_TAG="${IMAGE_TAG:-latest}"
REGISTRY="${REGISTRY:-}"
PLATFORMS="${PLATFORMS:-linux/amd64,linux/arm64,linux/arm/v7}"

# Full image name
if [ -n "$REGISTRY" ]; then
    FULL_IMAGE_NAME="${REGISTRY}/${IMAGE_NAME}:${IMAGE_TAG}"
else
    FULL_IMAGE_NAME="${IMAGE_NAME}:${IMAGE_TAG}"
fi

echo "==========================================="
echo "Multi-Architecture Docker Build"
echo "==========================================="
echo "Image: ${FULL_IMAGE_NAME}"
echo "Platforms: ${PLATFORMS}"
echo "==========================================="
echo ""

# Check if docker buildx is available
if ! docker buildx version &> /dev/null; then
    echo "Error: docker buildx is not available"
    echo "Please install Docker Desktop or enable buildx"
    exit 1
fi

# Create or use existing buildx builder
BUILDER_NAME="multiarch-builder"
if ! docker buildx inspect ${BUILDER_NAME} &> /dev/null; then
    echo "Creating new buildx builder: ${BUILDER_NAME}"
    docker buildx create --name ${BUILDER_NAME} --use
else
    echo "Using existing buildx builder: ${BUILDER_NAME}"
    docker buildx use ${BUILDER_NAME}
fi

# Bootstrap the builder
docker buildx inspect --bootstrap

# Build options
BUILD_ARGS=""
if [ "$PUSH" = "true" ]; then
    echo "Building and pushing to registry..."
    BUILD_ARGS="--push"
elif [ "$LOAD" = "true" ]; then
    echo "Building and loading to local Docker..."
    # Note: --load only works with single platform
    BUILD_ARGS="--load"
    PLATFORMS="linux/amd64"
else
    echo "Building (dry run)..."
    BUILD_ARGS=""
fi

# Build the image
docker buildx build \
    --platform ${PLATFORMS} \
    --tag ${FULL_IMAGE_NAME} \
    ${BUILD_ARGS} \
    --progress=plain \
    .

echo ""
echo "==========================================="
echo "Build completed successfully!"
echo "==========================================="
echo "Image: ${FULL_IMAGE_NAME}"
echo "Platforms: ${PLATFORMS}"

if [ "$PUSH" = "true" ]; then
    echo "Status: Pushed to registry"
elif [ "$LOAD" = "true" ]; then
    echo "Status: Loaded to local Docker"
else
    echo "Status: Built (not pushed)"
    echo ""
    echo "To push to registry, run:"
    echo "  PUSH=true ./build-multiarch.sh"
    echo ""
    echo "To load to local Docker (single arch), run:"
    echo "  LOAD=true ./build-multiarch.sh"
fi

echo "==========================================="

