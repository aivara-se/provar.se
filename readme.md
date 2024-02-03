# Provar.se

Provar is a tool to collect feedback from users.

## Getting Started

This guide is written for a Linux but it should also work for MacOS. If you are developing on Microsoft Windows, please use WSL2.

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

- Create a Google Cloud Services project, a Google Cloud Storage bucket and a service account that can upload files to the bucket. Download credentials to the service account as `credentials.json` and place it on the root of the repository. This will be used to temporarily upload files when importing feedback data. Set the `GCS_IMPORT_BUCKET` env variable on `apps/webapp/.env` file to the Google Cloud Storage bucket name.

- Create a MaxMind account, download the "GeoLite2 City" database from [MaxMind](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data) and place it as `apps/webapi/GeoLite2-City.mmdb`. This database will be used to add location data based on feedback source ip address.

- _TODO: explain how to set `AUTH_GITHUB_ID` and `AUTH_GITHUB_SECRET` env variables._
- _TODO: explain how to set `AUTH_GOOGLE_ID` and `AUTH_GOOGLE_SECRET` env variables._
- _TODO: explain how to set the `PUBLIC_PROVAR_API_KEY` env variable._
