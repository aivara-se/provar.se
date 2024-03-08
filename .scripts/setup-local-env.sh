#!/bin/bash

# Install global dependencies
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Source apps/webapi/.secrets file if it exists
if [ -f apps/webapi/.secrets ]; then
  source apps/webapi/.secrets
fi

# Create .env file for webapi service
cat << EOF > apps/webapi/.env
AUTH_SECRET=$(openssl rand -base64 40)
AUTH_GITHUB_ID=$APPS_WEBAPI_APP_AUTH_GITHUB_ID
AUTH_GITHUB_SECRET=$APPS_WEBAPI_APP_AUTH_GITHUB_SECRET
AUTH_GOOGLE_ID=$APPS_WEBAPI_APP_AUTH_GOOGLE_ID
AUTH_GOOGLE_SECRET=$APPS_WEBAPI_APP_AUTH_GOOGLE_SECRET
EMAIL_FROM=noreply@provar.se
EMAIL_SERVER=smtp://user:pass@localhost:1025
DATABASE_URI=postgresql://postgres:postgres@localhost/provar?sslmode=disable
GEOLITE2_DB=./GeoLite2-City.mmdb

# Local only environment variables
GOOGLE_APPLICATION_CREDENTIALS=../../credentials.json
PORT=3001
EOF

# Create .env file for webapp service
cat << EOF > apps/webapp/.env
PUBLIC_PROVAR_APP_URL=http://localhost:3002
PUBLIC_PROVAR_API_URL=http://localhost:3001

# Local only environment variables
PORT=3002
EOF
