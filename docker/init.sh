#!/bin/bash
echo "Starting database with automatic migrations..."
docker compose up -d
echo "✓ Database is ready when health check passes"
