# Production Deployment Guide

## Prerequisites

- Docker and Docker Compose installed on server
- Access to `registry.hamdocker.ir` (Docker registry)
- Domain configured and pointing to server IP
- SSL certificate ready (or use Let's Encrypt)

## Step 1: Prepare Server

```bash
# SSH to server
ssh root@your_server_ip

# Create project directory
mkdir -p /opt/tts
cd /opt/tts

# Clone repository
git clone https://github.com/hesamksdev-web/TTS.git .
```

## Step 2: Create Environment File

```bash
# Create .env file with production values
cat > .env << EOF
DB_PASSWORD=your_secure_password_here
JWT_SECRET=your_secure_jwt_secret_here
DOMAIN=your-domain.com
EOF

# Secure the file
chmod 600 .env
```

### Generate Secure Values

**For DB_PASSWORD and JWT_SECRET, use:**

```bash
# Generate a 32-character random string
openssl rand -base64 32
```

## Step 3: Setup SSL Certificates

### Option A: Using Let's Encrypt (Recommended)

```bash
# Install certbot
apt-get update && apt-get install -y certbot

# Get certificate
certbot certonly --standalone -d your-domain.com

# Copy certificates
mkdir -p ssl
cp /etc/letsencrypt/live/your-domain.com/fullchain.pem ssl/cert.pem
cp /etc/letsencrypt/live/your-domain.com/privkey.pem ssl/key.pem

# Set permissions
chmod 644 ssl/cert.pem
chmod 644 ssl/key.pem
```

### Option B: Using Existing Certificates

```bash
mkdir -p ssl
cp /path/to/cert.pem ssl/cert.pem
cp /path/to/key.pem ssl/key.pem
chmod 644 ssl/cert.pem ssl/key.pem
```

## Step 4: Login to Docker Registry

```bash
docker login registry.hamdocker.ir
# Enter your credentials
```

## Step 5: Start Services

```bash
# Pull latest images
docker-compose -f docker-compose.prod.yml pull

# Start all services
docker-compose -f docker-compose.prod.yml up -d

# Check status
docker-compose -f docker-compose.prod.yml ps
```

## Step 6: Verify Deployment

```bash
# Check logs
docker-compose -f docker-compose.prod.yml logs -f

# Test API endpoint
curl https://your-domain.com/api/v1/health

# Test frontend
# Open https://your-domain.com in browser
```

## Monitoring & Maintenance

### View Logs

```bash
# All services
docker-compose -f docker-compose.prod.yml logs -f

# Specific service
docker-compose -f docker-compose.prod.yml logs -f go-service
docker-compose -f docker-compose.prod.yml logs -f python-service
docker-compose -f docker-compose.prod.yml logs -f frontend
```

### Restart Services

```bash
# Restart all
docker-compose -f docker-compose.prod.yml restart

# Restart specific service
docker-compose -f docker-compose.prod.yml restart go-service
```

### Update Images

```bash
# Pull latest images
docker-compose -f docker-compose.prod.yml pull

# Restart with new images
docker-compose -f docker-compose.prod.yml up -d
```

### SSL Certificate Renewal

```bash
# Renew certificate
certbot renew

# Copy new certificates
cp /etc/letsencrypt/live/your-domain.com/fullchain.pem ssl/cert.pem
cp /etc/letsencrypt/live/your-domain.com/privkey.pem ssl/key.pem

# Restart nginx
docker-compose -f docker-compose.prod.yml restart nginx
```

## Backup & Recovery

### Backup Database

```bash
# Create backup
docker-compose -f docker-compose.prod.yml exec postgres pg_dump -U admin tts_db > backup.sql

# Restore from backup
docker-compose -f docker-compose.prod.yml exec -T postgres psql -U admin tts_db < backup.sql
```

### Backup Data Volumes

```bash
# Backup training data
tar -czf training_runs_backup.tar.gz training_runs/
tar -czf data_backup.tar.gz data/

# Restore
tar -xzf training_runs_backup.tar.gz
tar -xzf data_backup.tar.gz
```

## Troubleshooting

### Services not starting

```bash
# Check logs
docker-compose -f docker-compose.prod.yml logs

# Verify .env file exists and has correct values
cat .env

# Check if ports are available
netstat -tlnp | grep -E ':(80|443|8080|5432)'
```

### Database connection issues

```bash
# Check database service
docker-compose -f docker-compose.prod.yml logs postgres

# Test connection
docker-compose -f docker-compose.prod.yml exec postgres psql -U admin -d tts_db -c "SELECT 1"
```

### SSL certificate issues

```bash
# Check certificate validity
openssl x509 -in ssl/cert.pem -text -noout

# Check nginx logs
docker-compose -f docker-compose.prod.yml logs nginx
```

## Environment Variables Reference

| Variable | Description | Example |
|----------|-------------|---------|
| `DB_PASSWORD` | PostgreSQL password | `SecurePass123!@#` |
| `JWT_SECRET` | JWT signing secret | `YourSecretKey123!@#$%` |
| `DOMAIN` | Domain name | `clone.bearbyte.cloud` |

## Security Checklist

- [ ] `.env` file created with strong passwords
- [ ] `.env` file is NOT committed to Git
- [ ] SSL certificates installed and valid
- [ ] Docker registry credentials configured
- [ ] Firewall rules configured (80, 443 open)
- [ ] Database backups scheduled
- [ ] Log rotation configured
- [ ] SSL certificate renewal automated

## Support

For issues, check:
1. Docker logs: `docker-compose -f docker-compose.prod.yml logs -f`
2. Service status: `docker-compose -f docker-compose.prod.yml ps`
3. Environment variables: `cat .env`
