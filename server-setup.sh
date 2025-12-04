#!/bin/bash

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

echo -e "${YELLOW}TTS Server Setup Script${NC}"
echo "=================================="

# Step 1: Install Docker
echo -e "${YELLOW}Step 1: Installing Docker...${NC}"
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
rm get-docker.sh

# Step 2: Install Docker Compose
echo -e "${YELLOW}Step 2: Installing Docker Compose...${NC}"
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" \
  -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# Step 3: Create directories
echo -e "${YELLOW}Step 3: Creating directories...${NC}"
mkdir -p /opt/tts/{ssl,data,training_runs}
cd /opt/tts

# Step 4: Get SSL Certificate
echo -e "${YELLOW}Step 4: Setting up SSL Certificate...${NC}"
apt-get update
apt-get install -y certbot

echo -e "${YELLOW}Getting SSL certificate for clone.bearbyte.cloud...${NC}"
certbot certonly --standalone -d clone.bearbyte.cloud

# Copy certificates
mkdir -p /opt/tts/ssl
cp /etc/letsencrypt/live/clone.bearbyte.cloud/fullchain.pem /opt/tts/ssl/cert.pem
cp /etc/letsencrypt/live/clone.bearbyte.cloud/privkey.pem /opt/tts/ssl/key.pem
chmod 644 /opt/tts/ssl/cert.pem
chmod 644 /opt/tts/ssl/key.pem

echo -e "${GREEN}SSL Certificate setup complete!${NC}"
echo -e "${YELLOW}Next steps:${NC}"
echo "1. Copy the project files to /opt/tts"
echo "2. Create .env file with DB_PASSWORD and JWT_SECRET"
echo "3. Run: docker-compose -f docker-compose.server.yml up -d"
