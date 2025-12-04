# Quick Start - Server Deployment

## روی سرور (46.224.57.222)

### 1. SSH به سرور
```bash
ssh root@46.224.57.222
```

### 2. نصب Docker و Docker Compose
```bash
# نصب Docker
curl -fsSL https://get.docker.com | sh

# نصب Docker Compose
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" \
  -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# تایید نصب
docker --version
docker-compose --version
```

### 3. ایجاد پوشه‌ها
```bash
mkdir -p /opt/tts/{ssl,data,training_runs}
cd /opt/tts
```

### 4. SSL Certificate
```bash
# نصب certbot
apt-get update && apt-get install -y certbot

# دریافت certificate
certbot certonly --standalone -d clone.bearbyte.cloud

# کپی certificates
mkdir -p /opt/tts/ssl
cp /etc/letsencrypt/live/clone.bearbyte.cloud/fullchain.pem /opt/tts/ssl/cert.pem
cp /etc/letsencrypt/live/clone.bearbyte.cloud/privkey.pem /opt/tts/ssl/key.pem
chmod 644 /opt/tts/ssl/*
```

### 5. کپی فایل‌های پروژه
```bash
# از ماشین محلی (macOS):
scp -r /Users/hesamksdev/TTS/* root@46.224.57.222:/opt/tts/

# یا استفاده از git:
cd /opt/tts
git clone <your-repo-url> .
```

### 6. ایجاد .env فایل
```bash
cd /opt/tts
cat > .env << 'EOF'
DB_PASSWORD=your_secure_password
JWT_SECRET=your_jwt_secret_key
EOF
```

### 7. Build و Start
```bash
cd /opt/tts

# Build تمام services
docker-compose -f docker-compose.server.yml build

# Start services
docker-compose -f docker-compose.server.yml up -d

# بررسی status
docker-compose -f docker-compose.server.yml ps
```

### 8. بررسی Logs
```bash
# تمام logs
docker-compose -f docker-compose.server.yml logs -f

# یا برای service خاص
docker-compose -f docker-compose.server.yml logs -f go-service
```

### 9. تست
```bash
# Health check
curl https://clone.bearbyte.cloud/api/v1/health

# یا بدون SSL verification
curl -k https://clone.bearbyte.cloud/api/v1/health
```

## دسترسی به Application

- **Frontend**: https://clone.bearbyte.cloud
- **API**: https://clone.bearbyte.cloud/api/v1

## Commands مهم

```bash
# بررسی status
docker-compose -f docker-compose.server.yml ps

# مشاهده logs
docker-compose -f docker-compose.server.yml logs -f

# Restart services
docker-compose -f docker-compose.server.yml restart

# Stop services
docker-compose -f docker-compose.server.yml down

# Rebuild
docker-compose -f docker-compose.server.yml build --no-cache
docker-compose -f docker-compose.server.yml up -d
```

## Troubleshooting

### اگر nginx error داد
```bash
# بررسی nginx logs
docker-compose -f docker-compose.server.yml logs nginx

# بررسی SSL files
ls -la /opt/tts/ssl/
```

### اگر database error داد
```bash
# بررسی postgres logs
docker-compose -f docker-compose.server.yml logs postgres

# دسترسی به database
docker-compose -f docker-compose.server.yml exec postgres psql -U admin -d tts_db
```

### اگر Go service error داد
```bash
# بررسی logs
docker-compose -f docker-compose.server.yml logs go-service

# restart
docker-compose -f docker-compose.server.yml restart go-service
```

## Certificate Renewal

```bash
# Renew
certbot renew

# Copy renewed certs
cp /etc/letsencrypt/live/clone.bearbyte.cloud/fullchain.pem /opt/tts/ssl/cert.pem
cp /etc/letsencrypt/live/clone.bearbyte.cloud/privkey.pem /opt/tts/ssl/key.pem

# Restart nginx
docker-compose -f docker-compose.server.yml restart nginx
```
