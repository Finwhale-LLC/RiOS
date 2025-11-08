#!/bin/bash

# RiOS Worker Build Script

set -e

echo "🔨 Building RiOS Worker..."
echo ""

# Get version from git tag or use default
VERSION=${VERSION:-"v0.1.0"}
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Build flags
LDFLAGS="-s -w -X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.GitCommit=${GIT_COMMIT}"

# Detect OS
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64)
        ARCH="amd64"
        ;;
    aarch64|arm64)
        ARCH="arm64"
        ;;
esac

echo "Platform: ${OS}/${ARCH}"
echo "Version: ${VERSION}"
echo ""

# Build for current platform
echo "Building for ${OS}/${ARCH}..."
OUTPUT="rios-worker"
if [ "$OS" = "windows" ]; then
    OUTPUT="rios-worker.exe"
fi

CGO_ENABLED=0 go build -ldflags="${LDFLAGS}" -o "${OUTPUT}" main.go

echo "✅ Build complete: ${OUTPUT}"
echo ""

# Make executable
chmod +x "${OUTPUT}"

# Show size
SIZE=$(du -h "${OUTPUT}" | cut -f1)
echo "📦 Binary size: ${SIZE}"
echo ""

# Build for other platforms (optional)
read -p "Build for other platforms? (y/N) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo ""
    echo "Building for multiple platforms..."
    
    PLATFORMS=(
        "linux/amd64"
        "linux/arm64"
        "darwin/amd64"
        "darwin/arm64"
        "windows/amd64"
    )
    
    mkdir -p dist
    
    for PLATFORM in "${PLATFORMS[@]}"; do
        GOOS=${PLATFORM%/*}
        GOARCH=${PLATFORM#*/}
        OUTPUT_NAME="dist/rios-worker-${GOOS}-${GOARCH}"
        
        if [ "$GOOS" = "windows" ]; then
            OUTPUT_NAME="${OUTPUT_NAME}.exe"
        fi
        
        echo "Building ${GOOS}/${GOARCH}..."
        CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="${LDFLAGS}" -o "${OUTPUT_NAME}" main.go
    done
    
    echo ""
    echo "✅ Multi-platform builds complete!"
    echo "📁 Binaries are in ./dist/"
    ls -lh dist/
fi

echo ""
echo "🚀 Done! Run './rios-worker --help' to get started."

