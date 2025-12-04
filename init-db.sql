-- Create database if not exists
SELECT 'CREATE DATABASE tts_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'tts_db')\gexec
