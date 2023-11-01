# apps/webapp

Web application where users can configure organizations, projects and feature flags.

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
```

Optionally create a `.env.test` file with the following if you wish to use a separate database when running unit tests.

```
MONGODB_URL=""
```

Start the application locally

```shell
yarn dev
```
