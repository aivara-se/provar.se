# Provar.se

Provar is a tool to collect feedback from users.

## Getting started

Install Node and Yarn

- [Install node v20](https://formulae.brew.sh/formula/node)
- [Install yarn v3](https://yarnpkg.com/getting-started/install)

Create a .env file with the following in apps/webapp directory.

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

Start the application locally

```shell
yarn run dev
```
