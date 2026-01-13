# Build stage
FROM --platform=$BUILDPLATFORM golang:1.21-alpine AS builder

# Build arguments for multi-arch support
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary for target architecture
RUN CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    GOARM=${TARGETVARIANT#v} \
    go build -a -installsuffix cgo -ldflags="-w -s" -o drone-ai-review .

# Final stage
FROM alpine:latest

# Install git (required for git diff commands)
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /workspace

# Copy binary from builder
COPY --from=builder /build/drone-ai-review /bin/drone-ai-review

# Set the entrypoint
ENTRYPOINT ["/bin/drone-ai-review"]

