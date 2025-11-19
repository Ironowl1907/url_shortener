#!/bin/bash

BASE_URL="http://localhost:8080"
EMAIL="test-$(date +%s)@example.com"
PASSWORD="testpass123"

echo "=== Testing URL Shortener API ==="

# Ping
echo -e "\n1. Testing /ping"
curl -s "$BASE_URL/ping"

# Register
echo -e "\n\n2. Registering user"
curl -s -X POST "$BASE_URL/auth/register" \
  -H "Content-Type: application/json" \
  -d "{\"username\":\"testuser\",\"email\":\"$EMAIL\",\"password\":\"$PASSWORD\"}"

# Login
echo -e "\n\n3. Logging in"
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
  -c cookies.txt \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL\",\"password\":\"$PASSWORD\"}")
echo $LOGIN_RESPONSE

# Validate
echo -e "\n\n4. Validating token"
curl -s "$BASE_URL/auth/validate" -b cookies.txt

# Get Me
echo -e "\n\n5. Getting user info"
curl -s "$BASE_URL/auth/me" -b cookies.txt

# Create URL
echo -e "\n\n6. Creating shortened URL"
CREATE_RESPONSE=$(curl -s -X POST "$BASE_URL/urls" \
  -b cookies.txt \
  -H "Content-Type: application/json" \
  -d '{"url":"https://github.com","customShortCode":"gh"}')
echo $CREATE_RESPONSE

# Get all URLs
echo -e "\n\n7. Getting all URLs"
curl -s "$BASE_URL/urls" -b cookies.txt

# Test redirect
echo -e "\n\n8. Testing redirect"
curl -I "$BASE_URL/gh"

# Cleanup
rm -f cookies.txt
echo -e "\n\n=== Tests completed ==="
