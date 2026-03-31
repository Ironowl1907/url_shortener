# url_shortener

Simple url shortening service with users.

Tech Stack:
- Golang
- Gin-gonic
- Gorm
- Sveltekit
- Postgres
- Docker

## Steps for testing

Clone the repo:
```
git clone https://github.com/Ironowl1907/url_shortener.git
```
Move to it:
```
cd url_shortener
```
## .env files (Required!)
Example for `frontend/.env`:
```
API_BASE_URL="http://backend:8080"
```
Example for `backend/.env`:
```
PORT=8080
DB="host=db user=postgres dbname=url_shortener password=verySecurePassword port=5432"
SECRET="K9ERKFzBJ4am4MxcMwJKujcEsx42pT0w"

```
# Deploying
Make sure to be inside url_shotener/
```
docker compose up -d
```
