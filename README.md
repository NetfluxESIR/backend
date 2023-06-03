# Netflux ESIR - Backend

This repository contains the backend of the Netflux ESIR project used to store metadata and videos representations in DB.

## Install

### Requirements

- [git](https://git-scm.com/downloads)
- [Go](https://golang.org/doc/install)
- [An AWS S3 bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-bucket.html)
- [A PostgreSQL database](https://www.postgresql.org/download/)

### Run from source

#### Build the server

```bash
git clone https://https://github.com/NetfluxESIR/backend.git
cd backend
make build
```

#### Run the server

> Don't forget to set "AWS_ACCESS_KEY_ID" and "AWS_SECRET_ACCESS_KEY" environment variables to access to your S3 bucket.

```bash
Start the Netflux API server.

Usage:
  netflux serve [flags]

Flags:
  -a, --admin-account string    Admin account. (default "admin")
  -P, --admin-password string   Admin password. (default "admin")
  -d, --dsn string              Data source name. (default "netflux:netflux@postgres:5432/netflux?sslmode=disable")
  -h, --help                    help for serve
  -H, --host string             Host to listen on. (default "localhost")
  -l, --loglevel string         Log level. (default "info")
  -p, --port int                Port to listen on. (default 8080)
  -b, --s3-bucket string        S3 bucket name.
  -r, --s3-region string        S3 region.
```

### Run with Docker

##### Requirements

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/) (optional)

> You still have to set "AWS_ACCESS_KEY_ID" and "AWS_SECRET_ACCESS_KEY" environment variables to access to your S3 bucket.

You can find a docker image of the backend [here](https://github.com/NetfluxESIR/backend/pkgs/container/backend).

```bash
docker run -p 8080:8080 ghcr.io/netfluxesir/backend:latest serve --s3-bucket <your-bucket-name> --s3-region <your-bucket-region>
```

> Note: you can also use the [docker-compose.yml](./docker-compose.yaml) file to run the backend and the database locally.


## API Documentation

You can find the API documentation [here](./pkg/api/specs/pkg/video-api/openapi.yaml).