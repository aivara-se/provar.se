# apps/webapi

Public API to collect feedback from client applications.

## Getting started

Install CLI tools needed to generate swagger documentation.

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

Create a `.env` file with the following which will be used when running the application locally.

```
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
