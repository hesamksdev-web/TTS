#!/bin/bash

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

SERVER="root@46.224.57.222"
REMOTE_PATH="/opt/tts"

echo -e "${YELLOW}Copying TTS files to server...${NC}"
echo "Server: $SERVER"
echo "Path: $REMOTE_PATH"
echo ""

# Check if server is reachable
echo -e "${YELLOW}Checking server connectivity...${NC}"
if ! ping -c 1 46.224.57.222 &> /dev/null; then
    echo -e "${RED}Cannot reach server. Please check connectivity.${NC}"
    exit 1
fi
echo -e "${GREEN}Server is reachable!${NC}"
echo ""

# Copy files
echo -e "${YELLOW}Copying files...${NC}"

# Create remote directories
ssh $SERVER "mkdir -p $REMOTE_PATH/{ssl,data,training_runs}"

# Copy docker-compose
echo "Copying docker-compose.server.yml..."
scp docker-compose.server.yml $SERVER:$REMOTE_PATH/

# Copy nginx config
echo "Copying nginx.conf..."
scp nginx.conf $SERVER:$REMOTE_PATH/

# Copy go-service
echo "Copying go-service..."
scp -r go-service $SERVER:$REMOTE_PATH/

# Copy python-service
echo "Copying python-service..."
scp -r python-service $SERVER:$REMOTE_PATH/

# Copy frontend
echo "Copying frontend..."
scp -r frontend $SERVER:$REMOTE_PATH/

# Copy documentation
echo "Copying documentation..."
scp QUICK_START.md $SERVER:$REMOTE_PATH/
scp SERVER_DEPLOY.md $SERVER:$REMOTE_PATH/

echo ""
echo -e "${GREEN}Files copied successfully!${NC}"
echo ""
echo -e "${YELLOW}Next steps on server:${NC}"
echo "1. ssh root@46.224.57.222"
echo "2. cd /opt/tts"
echo "3. cat > .env << 'EOF'"
echo "   DB_PASSWORD=your_password"
echo "   JWT_SECRET=your_secret"
echo "   EOF"
echo "4. docker-compose -f docker-compose.server.yml build"
echo "5. docker-compose -f docker-compose.server.yml up -d"
