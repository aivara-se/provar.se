# Provar.se

Provar is an online open source tool to collect feedback from your customers.

## Code Structure

This repository contains the following directories at the root level:

- **apps:** frontend application, backend services and static websites
- **libs:** code shared between multiple apps and outside this repository
- **docs:** useful documentation for developers who wish to contribute

## Getting Started

This guide describes how to setup a Linux based development environment. But it should also work for MacOS perhaps with minor tweaks to the setup script. If you are developing on Microsoft Windows, please use WSL2. Check the [deployment.md](./docs/deployments.md) and [components.md](./docs/components.md) to learn more about how it is deployed on provar.se

At the moment, the local development environment depends on several online services (eg: GCS bucket) which will be addressed in the future.

- Install Docker, Node, Yarn and Go.

  - [Install docker](https://docs.docker.com/engine/install/)
  - [Install node v20](https://formulae.brew.sh/formula/node)
  - [Install yarn v4](https://yarnpkg.com/getting-started/install)
  - [Install go](https://go.dev/doc/install)

- Clone the repository and install yarn dependencies.

```shell
git clone https://github.com/aivara-se/provar.se.git
cd provar.se
yarn install
```

- Use the `yarn setup-local-env` command to create required `.env` files needed for local development. This wil fill most of the env variables but several environment variables need to be filled manually.

```shell
yarn setup-local-env
```

- Start the local environment with `yarn start-local-env` and run database migrations with `yarn workspace migrations run migration:run`

- Create a MaxMind account, download the "GeoLite2 City" database from [MaxMind](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data) and place it as `apps/webapi/GeoLite2-City.mmdb`. This database will be used to add location data based on feedback source ip address.
