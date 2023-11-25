# apps/webapp

Web application where users can configure organizations, projects and API keys. This app is deployed to https://app.provar.se

## Getting Started

Create a `.env` file with the following which will be used when running the application locally.

```
AUTH_SECRET=""
AUTH_GITHUB_ID=""
AUTH_GITHUB_SECRET=""
AUTH_GOOGLE_ID=""
AUTH_GOOGLE_SECRET=""
AUTH_EMAIL_SERVER=""
AUTH_EMAIL_FROM=""
MONGODB_URL=""
PUBLIC_PROVAR_API_KEY=""
```

Optionally create a `.env.test` file with the following if you wish to use a separate database when running unit tests.

```
MONGODB_URL=""
PUBLIC_PROVAR_API_KEY=""
```

Start the application locally

```shell
yarn dev
```

## Run tests

```shell
yarn test:unit
```
