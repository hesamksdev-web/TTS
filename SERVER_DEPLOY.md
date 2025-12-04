# Server Deployment Guide

## Server Details
- **IP**: 46.224.57.222
- **Domain**: clone.bearbyte.cloud
- **OS**: Ubuntu
- **Architecture**: x86_64 (amd64)

## Prerequisites
- SSH access to server
- Domain DNS configured to point to server IP
- Port 80 and 443 open

## Step 1: Initial Setup

### SSH to Server
```bash
ssh root@46.224.57.222
```

### Run Setup Script
```bash
# Download and run setup script
curl -O https://raw.githubusercontent.com/your-repo/server-setup.sh
chmod +x server-setup.sh
./server-setup.sh
```

This will:
- Install Docker
- Install Docker Compose
- Create necessary directories
- Setup SSL certificate with Let's Encrypt

## Step 2: Copy Project Files

### From Your Local Machine
```bash
# Copy all project files
scp -r /Users/hesamksdev/TTS/* root@46.224.57.222:/opt/tts/

# Or use git clone (if you have a git repo)
ssh root@46.224.57.222
cd /opt/tts
git clone <your-repo-url> .
```

## Step 3: Create Environment File

### On Server
```bash
cd /opt/tts
cat > .env << EOF
DB_PASSWORD=your_secure_password_here
JWT_SECRET=your_jwt_secret_key_here
EOF
```

## Step 4: Build and Start Services

### Build Images (First Time)
```bash
cd /opt/tts

# Build all services
docker-compose -f docker-compose.server.yml build

# Start services
docker-compose -f docker-compose.server.yml up -d

# Check status
docker-compose -f docker-compose.server.yml ps
```

### View Logs
```bash
# All services
docker-compose -f docker-compose.server.yml logs -f

# Specific service
docker-compose -f docker-compose.server.yml logs -f go-service
docker-compose -f docker-compose.server.yml logs -f frontend
docker-compose -f docker-compose.server.yml logs -f python-service
```

## Step 5: Verify Deployment

### Check Services
```bash
docker-compose -f docker-compose.server.yml ps
```

### Test API
```bash
# Health check
curl https://clone.bearbyte.cloud/api/v1/health

# Or without SSL verification (for testing)
curl -k https://clone.bearbyte.cloud/api/v1/health
```

### Access Application
- **Frontend**: https://clone.bearbyte.cloud
- **API**: https://clone.bearbyte.cloud/api/v1

## Troubleshooting

### Check Service Logs
```bash
docker-compose -f docker-compose.server.yml logs -f [service-name]
```

### Restart Services
```bash
docker-compose -f docker-compose.server.yml restart
```

### Restart Specific Service
```bash
docker-compose -f docker-compose.server.yml restart go-service
```

### Stop Services
```bash
docker-compose -f docker-compose.server.yml down
```

### Remove All Data (WARNING: Destructive)
```bash
docker-compose -f docker-compose.server.yml down -v
```

### Rebuild Images
```bash
docker-compose -f docker-compose.server.yml build --no-cache
docker-compose -f docker-compose.server.yml up -d
```

## SSL Certificate Renewal

### Automatic Renewal
```bash
# Test renewal
certbot renew --dry-run

# Renew certificates
certbot renew

# Copy renewed certificates
cp /etc/letsencrypt/live/clone.bearbyte.cloud/fullchain.pem /opt/tts/ssl/cert.pem
cp /etc/letsencrypt/live/clone.bearbyte.cloud/privkey.pem /opt/tts/ssl/key.pem

# Restart nginx
docker-compose -f docker-compose.server.yml restart nginx
```

### Setup Auto-Renewal Cron Job
```bash
# Add to crontab
crontab -e

# Add this line (runs daily at 2 AM)
0 2 * * * certbot renew && cp /etc/letsencrypt/live/clone.bearbyte.cloud/fullchain.pem /opt/tts/ssl/cert.pem && cp /etc/letsencrypt/live/clone.bearbyte.cloud/privkey.pem /opt/tts/ssl/key.pem && docker-compose -f /opt/tts/docker-compose.server.yml restart nginx
```

## Database Management

### Access Database
```bash
docker-compose -f docker-compose.server.yml exec postgres psql -U admin -d tts_db
```

### Backup Database
```bash
docker-compose -f docker-compose.server.yml exec postgres pg_dump -U admin tts_db > backup.sql
```

### Restore Database
```bash
docker-compose -f docker-compose.server.yml exec -T postgres psql -U admin tts_db < backup.sql
```

## Performance Monitoring

### Check Resource Usage
```bash
docker stats
```

### Check Disk Space
```bash
df -h
```

### Check Memory
```bash
free -h
```

## Updating Application

### Pull Latest Changes
```bash
cd /opt/tts
git pull origin main
```

### Rebuild and Restart
```bash
docker-compose -f docker-compose.server.yml build
docker-compose -f docker-compose.server.yml up -d
```

## Backup and Recovery

### Full Backup
```bash
cd /opt/tts
tar -czf backup-$(date +%Y%m%d).tar.gz data/ training_runs/ postgres_data/
```

### Restore from Backup
```bash
cd /opt/tts
tar -xzf backup-YYYYMMDD.tar.gz
docker-compose -f docker-compose.server.yml restart
```

## Security

### Change Default Credentials
1. Update `.env` file with strong passwords
2. Restart services: `docker-compose -f docker-compose.server.yml restart`

### Firewall Rules
```bash
# Allow only necessary ports
ufw allow 22/tcp    # SSH
ufw allow 80/tcp    # HTTP
ufw allow 443/tcp   # HTTPS
ufw enable
```

## Support

For issues:
1. Check logs: `docker-compose -f docker-compose.server.yml logs -f`
2. Verify services: `docker-compose -f docker-compose.server.yml ps`
3. Check disk space: `df -h`
4. Check memory: `free -h`
