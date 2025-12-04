#!/bin/bash

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

SERVER="root@Automations"
REMOTE_PATH="/opt/tts"

echo -e "${YELLOW}================================${NC}"
echo -e "${YELLOW}TTS Server Deployment${NC}"
echo -e "${YELLOW}================================${NC}"
echo ""

# Copy docker-compose.prod.yml
echo -e "${YELLOW}Copying docker-compose.prod.yml...${NC}"
scp docker-compose.prod.yml $SERVER:$REMOTE_PATH/

# Copy nginx.conf
echo -e "${YELLOW}Copying nginx.conf...${NC}"
scp nginx.conf $SERVER:$REMOTE_PATH/

# Create .env file on server if it doesn't exist
echo -e "${YELLOW}Creating .env file on server...${NC}"
ssh $SERVER "cat > $REMOTE_PATH/.env << 'ENVEOF'
DB_PASSWORD=password123
JWT_SECRET=my_super_secret_key_2025
ENVEOF"

echo ""
echo -e "${GREEN}✓ Files copied successfully!${NC}"
echo ""
echo -e "${YELLOW}Next steps on server:${NC}"
echo "1. ssh root@Automations"
echo "2. cd /opt/tts"
echo "3. docker-compose -f docker-compose.prod.yml pull"
echo "4. docker-compose -f docker-compose.prod.yml up -d"
echo ""
