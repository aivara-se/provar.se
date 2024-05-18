# apps/collect

PWA used by clients to collect feedback from their customers.

## Contributing

Create a `.env` file with the following which will be used when running the application locally.

```
PUBLIC_PROVAR_APP_URL=""
PUBLIC_PROVAR_API_URL=""
```

Optionally create a `.env.test` file with the following if you wish to use a separate database when running unit tests.

```
PUBLIC_PROVAR_APP_URL=""
PUBLIC_PROVAR_API_URL=""
```

Start the application locally

```shell
yarn dev
```

## Run tests

```shell
yarn test:unit
```
