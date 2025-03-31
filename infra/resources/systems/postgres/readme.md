# Getting Started

## Deploying

Copy the postgres config file and related resources to the target host.

```shell
scp ./systems/postgres/uploads/{*,.*} postgres.provar.se:/root/postgres
```

Log into the system and mount the volume attached to the system on `/mnt/postgres`.
You can do this by by updating the `/etc/fstab` file. It should have a line similar
to what is given below:

```
/dev/disk/by-id/scsi-0HC_Volume_102345560 /mnt/postgres ext4 discard,nofail,defaults 0 0
```

Create a `data` directory under `/mnt/postgres` and copy the postgresql config file.

```shell
mkdir -p /mnt/postgres/data
cp ./postgresql.conf /mnt/postgres/
```

Finally, start it with this command

```shell
docker compose up -d --build
```

### Connect with PgAdmin

The database port is not open to the internet. You need to forward a local port with ssh and use that to connect.

```shell
# setup ssh config as given on root readme.md
ssh -L 15432:localhost:5432 postgres.provar.se
```
