#!/bin/bash

# Install additional dependencies
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/cosmtrek/air@latest

# Create a key pair for SAML
mkdir -p apps/account
openssl req -x509 -newkey rsa:2048 -keyout apps/account/local-saml.pem -out apps/account/local-saml.crt -sha256 -days 365 -nodes -subj "/C=SE/ST=Uppsala/L=Knivsta/O=Security/OU=Operations/CN=provar.se"

# Create a key pair for OpenID
openssl genrsa -out apps/account/local-openid.pem 3072
openssl pkcs8 -topk8 -inform PEM -outform PEM -nocrypt -in apps/account/local-openid.pem -out apps/account/local-openid.pk8.pem
openssl rsa -in apps/account/local-openid.pk8.pem -pubout -out apps/account/local-openid.pub.pem

# Create .env file for accounts service
cat << EOF > apps/account/.env
DB_ENGINE=mongo
DB_URL=mongodb://localhost:27017/accounts
JACKSON_API_KEYS=4X4bPXzeEn7faW1DJ/+CSALMZ4jdZNfn
NEXTAUTH_URL=http://localhost:5225
EXTERNAL_URL=http://localhost:5225
NEXTAUTH_SECRET=ik/egdToHUZTsQ8nbUWjDK0vCziIaj71l/fnNdFICTU=
NEXTAUTH_ADMIN_CREDENTIALS=admin@provar.se:password
SAML_AUDIENCE=http://localhost:5225
PUBLIC_KEY=$(cat apps/account/local-saml.crt | base64)
PRIVATE_KEY=$(cat apps/account/local-saml.pem | base64)
OPENID_RSA_PUBLIC_KEY=$(cat apps/account/local-openid.pub.pem | base64)
OPENID_RSA_PRIVATE_KEY=$(cat apps/account/local-openid.pk8.pem | base64)
EOF

# Create .env file for webapi service
cat << EOF > apps/webapi/.env
MONGO_URI=mongodb://localhost:27017/provar
GEOLITE2_DB=./GeoLite2-City.mmdb
PORT=3001
EOF

# Create .env file for webapp service
cat << EOF > apps/webapp/.env
AUTH_EMAIL_FROM=noreply@provar.se
AUTH_EMAIL_SERVER=smtp://user:pass@localhost:1025
AUTH_GITHUB_ID=
AUTH_GITHUB_SECRET=
AUTH_GOOGLE_ID=
AUTH_GOOGLE_SECRET=
AUTH_SECRET=$(openssl rand -base64 40)
GCS_IMPORT_BUCKET=provar-imports-dev
MONGODB_URL=mongodb://localhost:27017/provar
PROVAR_API_URL=http://localhost:3001
PUBLIC_PROVAR_APP_URL=http://localhost:5173
PUBLIC_PROVAR_API_KEY=placeholder

# Local only
GOOGLE_APPLICATION_CREDENTIALS=../../credentials.json
EOF
