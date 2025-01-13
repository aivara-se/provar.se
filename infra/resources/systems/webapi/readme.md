# Getting Started

Copy the `compose.yaml` files from the `systems/webapi` directory to the server on hetzner (eg: /root/webapi) and start the service with the command below.

```shell
docker compose up -d
```

Use this command to view logs

```shell
docker compose logs -f
```

## Deploying

Create files with environment variables for the reverse proxy and the service.

**.env.caddy**

```shell
CLOUDFLARE_API_TOKEN=
```

**.env.webapi**

```shell
# refer to webapi readme file for required env variables
```

Copy the docker compose file and related resources to the target host.

```shell
scp ./systems/webapi/uploads/{*,.*} webapi.provar.se:/root/webapi
```

Log into the ssytem and start it with this command

```shell
docker compose up --build
```

### Debugging

- Infinite redirects: set SSL/TLS mode on CLoudflare to "Full (strict)".
