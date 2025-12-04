# TTS Application Deployment Guide

## Prerequisites

- Docker and Docker Compose installed
- Access to registry.hamdocker.ir
- Domain: clone.bearbyte.cloud
- Server: 46.224.57.222 (Ubuntu)

## Local Development

```bash
# Start all services locally
./deploy.sh local

# Access:
# Frontend: http://localhost:3000
# API: http://localhost:8080
# Database: localhost:5432
```

## Build and Push to Registry

### On macOS (Apple Silicon)

```bash
# Make script executable
chmod +x deploy.sh

# Build and push all images (multi-platform)
./deploy.sh build

# This will build for:
# - linux/amd64 (for Ubuntu server)
# - linux/arm64 (for Apple Silicon)
```

## Deploy on Server

### Step 1: Prepare Server

```bash
# SSH to server
ssh root@46.224.57.222

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh

# Install Docker Compose
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# Create project directory
mkdir -p /opt/tts
cd /opt/tts
```

### Step 2: Copy Files to Server

```bash
# From your local machine
scp docker-compose.prod.yml root@46.224.57.222:/opt/tts/
scp nginx.conf root@46.224.57.222:/opt/tts/
scp deploy.sh root@46.224.57.222:/opt/tts/
```

### Step 3: Setup SSL Certificate

```bash
# On server
apt-get update && apt-get install -y certbot

# Get SSL certificate
certbot certonly --standalone -d clone.bearbyte.cloud

# Copy certificates
mkdir -p /opt/tts/ssl
cp /etc/letsencrypt/live/clone.bearbyte.cloud/fullchain.pem /opt/tts/ssl/cert.pem
cp /etc/letsencrypt/live/clone.bearbyte.cloud/privkey.pem /opt/tts/ssl/key.pem

# Set permissions
chmod 644 /opt/tts/ssl/cert.pem
chmod 644 /opt/tts/ssl/key.pem
```

### Step 4: Create .env File

```bash
# On server
cat > /opt/tts/.env << EOF
DB_PASSWORD=your_secure_password
JWT_SECRET=your_jwt_secret_key
EOF
```

### Step 5: Deploy

```bash
# On server
cd /opt/tts
chmod +x deploy.sh
./deploy.sh server
```

### Step 6: Verify Deployment

```bash
# Check services
docker-compose -f docker-compose.prod.yml ps

# View logs
docker-compose -f docker-compose.prod.yml logs -f

# Test API
curl https://clone.bearbyte.cloud/api/v1/health
```

## Troubleshooting

### Check service logs
```bash
docker-compose -f docker-compose.prod.yml logs -f [service-name]
```

### Restart services
```bash
docker-compose -f docker-compose.prod.yml restart
```

### Update images
```bash
docker-compose -f docker-compose.prod.yml pull
docker-compose -f docker-compose.prod.yml up -d
```

### SSL Certificate Renewal
```bash
certbot renew
cp /etc/letsencrypt/live/clone.bearbyte.cloud/fullchain.pem /opt/tts/ssl/cert.pem
cp /etc/letsencrypt/live/clone.bearbyte.cloud/privkey.pem /opt/tts/ssl/key.pem
docker-compose -f docker-compose.prod.yml restart nginx
```

## Architecture

```
User → clone.bearbyte.cloud (Nginx + SSL)
         ↓
    ┌────┴────┐
    ↓         ↓
Frontend   API (Go)
(3000)     (8080)
    ↓         ↓
    └────┬────┘
         ↓
    Database (PostgreSQL)
    Python Service
```

## Environment Variables

- `DB_PASSWORD`: Database password
- `JWT_SECRET`: JWT signing secret
- `NEXT_PUBLIC_API_URL`: Frontend API URL (https://clone.bearbyte.cloud/api/v1)

## Support

For issues, check logs:
```bash
docker-compose -f docker-compose.prod.yml logs -f
```
