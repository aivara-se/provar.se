# Getting Started

You can use docker for running postgres with a command similar to this.

```shell
docker run -d --name postgres \
  -v /mnt/postgres/postgresql.conf:/etc/postgresql/postgresql.conf \
	-v /mnt/postgres/data:/var/lib/postgresql/data \
  -e POSTGRES_PASSWORD=mysecretpasswordthatnooneelseknowsabout \
  -p 5432:5432 \
  postgres:17 -c 'config_file=/etc/postgresql/postgresql.conf'
```

### Connect with PgAdmin

The database port is not open to the internet. You need to forward a local port with ssh and use that to connect.

```shell
# setup ssh config as given on root readme.md
ssh -L 15432:localhost:5432 postgres.provar.se
```
