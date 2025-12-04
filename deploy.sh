#!/bin/bash

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}TTS Application Deployment Script${NC}"
echo "=================================="

# Check if running on server
if [ "$1" == "server" ]; then
    echo -e "${YELLOW}Deploying on server...${NC}"
    
    # Create directories
    mkdir -p /opt/tts/ssl
    mkdir -p /opt/tts/data
    mkdir -p /opt/tts/training_runs
    
    cd /opt/tts
    
    # Pull latest images
    echo -e "${YELLOW}Pulling Docker images...${NC}"
    docker-compose -f docker-compose.prod.yml pull
    
    # Start services
    echo -e "${YELLOW}Starting services...${NC}"
    docker-compose -f docker-compose.prod.yml up -d
    
    # Check status
    echo -e "${YELLOW}Checking service status...${NC}"
    docker-compose -f docker-compose.prod.yml ps
    
    echo -e "${GREEN}Deployment complete!${NC}"
    echo -e "${GREEN}Access your application at: https://clone.bearbyte.cloud${NC}"

elif [ "$1" == "build" ]; then
    echo -e "${YELLOW}Building and pushing images...${NC}"
    
    # Setup buildx for multi-platform builds
    docker buildx create --name multiarch --use 2>/dev/null || docker buildx use multiarch
    
    # Build and push Go service
    echo -e "${YELLOW}Building Go service...${NC}"
    docker buildx build --platform linux/amd64,linux/arm64 \
        -t registry.hamdocker.ir/hesamksdev/tts-go-service:latest \
        -f go-service/Dockerfile \
        --push go-service/
    
    # Build and push Python service
    echo -e "${YELLOW}Building Python service...${NC}"
    docker buildx build --platform linux/amd64,linux/arm64 \
        -t registry.hamdocker.ir/hesamksdev/tts-python-service:latest \
        -f python-service/Dockerfile \
        --push python-service/
    
    # Build and push Frontend
    echo -e "${YELLOW}Building Frontend...${NC}"
    docker buildx build --platform linux/amd64,linux/arm64 \
        -t registry.hamdocker.ir/hesamksdev/tts-frontend:latest \
        -f frontend/Dockerfile \
        --push frontend/
    
    echo -e "${GREEN}All images built and pushed successfully!${NC}"

elif [ "$1" == "local" ]; then
    echo -e "${YELLOW}Starting local development...${NC}"
    docker-compose up -d
    echo -e "${GREEN}Services started!${NC}"
    echo -e "${GREEN}Frontend: http://localhost:3000${NC}"
    echo -e "${GREEN}API: http://localhost:8080${NC}"

else
    echo -e "${RED}Usage: ./deploy.sh [build|server|local]${NC}"
    echo ""
    echo "Commands:"
    echo "  build   - Build and push Docker images to registry"
    echo "  server  - Deploy on production server (run on server)"
    echo "  local   - Start local development environment"
fi
