# apps/webapi

Public API to collect feedback from client applications. This app is deployed to https://api.provar.se

## Getting started

Install CLI tools needed to generate swagger documentation and to live-reload on changes.

```shell
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/cosmtrek/air@latest
```

Create a `.env` file with the following which will be used when running the application locally.

```
MONGO_URI=""
PORT=

# Local only
GOOGLE_APPLICATION_CREDENTIALS=""
```

Optionally create a `.env.test` file with the following if you wish to use a separate database when running unit tests.

```
MONGO_URI=""
```

Start the application locally

```shell
yarn dev
```

## Contributing Code

- OpenAPI spec is built for API endpoints using comments. Read https://github.com/swaggo/swag#declarative-comments-format for more info.
