# apps/webapi

Public API to collect feedback from client applications. This app is deployed to https://api.provar.se

## Getting started

Before you continue, make sure you have followed setup instructions on the root readme.md file and ran the setup script. Open your `.env` file generated by the setup script and fill in missing details.

```
AUTH_SECRET=
AUTH_GITHUB_ID=
AUTH_GITHUB_SECRET=
AUTH_GOOGLE_ID=
AUTH_GOOGLE_SECRET=
EMAIL_FROM=noreply@provar.se
EMAIL_SERVER=smtp://user:pass@localhost:1025
DATABASE_URI=postgresql://postgres:postgres@localhost/provar?sslmode=disable
GEOLITE2_DB=./GeoLite2-City.mmdb

# Local only environment variables
GOOGLE_APPLICATION_CREDENTIALS=../../credentials.json
PORT=3001
```

Create a `.env.test` file with the following if you wish to use a separate database when running unit tests.

```
DATABASE_URI=postgresql://postgres:postgres@localhost/provar-test?sslmode=disable
GEOLITE2_DB=../../GeoLite2-City.mmdb
PORT=4001
```

Start the application locally

```shell
yarn dev
```
