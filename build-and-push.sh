#!/bin/bash

# تنظیم متغیرها
REGISTRY="registry.hamdocker.ir/hesamksdev"
TAG="v1"

# لاگین به رجیستری
echo "Logging in to registry..."
docker login registry.hamdocker.ir/hesamksdev

echo "================================"
echo "Setting up buildx for multi-platform..."
echo "================================"

# Setup buildx for multi-platform builds
docker buildx create --name multiarch --use 2>/dev/null || docker buildx use multiarch

echo "================================"
echo "Building Docker images for linux/amd64,linux/arm64..."
echo "================================"

# ساخت و پوش سرویس Go
echo "Building and pushing Go service..."
docker buildx build --platform linux/amd64,linux/arm64 \
  -t $REGISTRY/tts-go-service:$TAG \
  -f go-service/Dockerfile \
  --push go-service/
if [ $? -ne 0 ]; then
  echo "Failed to build Go service"
  exit 1
fi

# ساخت و پوش سرویس پایتون
echo "Building and pushing Python service..."
docker buildx build --platform linux/amd64,linux/arm64 \
  -t $REGISTRY/tts-python-service:$TAG \
  -f python-service/Dockerfile \
  --push python-service/
if [ $? -ne 0 ]; then
  echo "Failed to build Python service"
  exit 1
fi

# ساخت و پوش فرانت‌اند
echo "Building and pushing Frontend..."
docker buildx build --platform linux/amd64,linux/arm64 \
  -t $REGISTRY/tts-frontend:$TAG \
  -f frontend/Dockerfile \
  --push frontend/
if [ $? -ne 0 ]; then
  echo "Failed to build Frontend"
  exit 1
fi

echo ""
echo "================================"
echo "✓ All images built and pushed successfully!"
echo "================================"
echo ""
echo "Images (multi-platform: linux/amd64, linux/arm64):"
echo "  - $REGISTRY/tts-go-service:$TAG"
echo "  - $REGISTRY/tts-python-service:$TAG"
echo "  - $REGISTRY/tts-frontend:$TAG"