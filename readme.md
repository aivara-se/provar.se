# Provar.se

Provar is an online open source tool to collect feedback from your customers.

## Code Structure

This repository contains the following directories at the root level:

- **apps:** frontend application, backend services and static websites
- **libs:** code shared between multiple apps and outside this repository
- **infra:** an example infrastructure as code setup using Hetzner Cloud
- **docs:** useful documentation for developers who wish to contribute

## Getting Started

This guide describes how to setup a Linux based development environment. But it should also work for MacOS perhaps with minor tweaks to the setup script. If you are developing on Microsoft Windows, please use WSL2.

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

Next step is to set-up dependencies, docker containers and the database.

- Use the `yarn localenv setup` command to create required `.env` files needed for local development. This wil fill most of the env variables but several environment variables need to be filled manually (eg: client ids and secrets for social login).

```shell
yarn localenv setup
```

- Start the local environment with `yarn localenv start` and run database migrations with `yarn workspace migrations run migration:run`

- Create a MaxMind account, download the "GeoLite2 City" database from [MaxMind](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data) and place it as `apps/webapi/GeoLite2-City.mmdb`. This database will be used to add location data based on feedback source ip address.

Once done, you should be able to access the following developer tools.

- [http://localhost:5080](http://localhost:5080) - PgAdmin Web - This can be used to inspect the database.
- [http://localhost:8025](http://localhost:8025) - Mailpit Web - This can be used to inspect sent emails.

Finally, we can start the website, backend services and frontend app.

```shell
yarn dev
```

You can access them from these localhost urls:

- [http://localhost:3001](http://localhost:3001) - Backend
- [http://localhost:3002](http://localhost:3002) - Frontend
- [http://localhost:3003](http://localhost:3003) - Website
- [http://localhost:3004](http://localhost:3004) - Collect

Here's a diagram of how these services are connected.

![](./docs/_readme/localenv.png)
