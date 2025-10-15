#!/bin/bash
set -e

echo "🚀 Building rules-lint for all platforms..."

# Create dist directory
mkdir -p dist

PLATFORMS=("darwin" "linux" "windows")
ARCHS=("amd64" "arm64")
VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")

for platform in "${PLATFORMS[@]}"; do
  for arch in "${ARCHS[@]}"; do
    echo "📦 Building for $platform/$arch..."
    
    output_name="dist/rules-lint-$platform-$arch"
    if [ "$platform" = "windows" ]; then
      output_name="$output_name.exe"
    fi
    
    # Build the binary
    GOOS=$platform GOARCH=$arch go build -ldflags="-X 'main.Tag=$VERSION'" -o $output_name ./cmd/lint
    
    # Make executable (except Windows)
    if [ "$platform" != "windows" ]; then
      chmod +x $output_name
    fi
    
    echo "✅ Built $output_name"
  done
done

echo "🎉 All binaries built successfully!"