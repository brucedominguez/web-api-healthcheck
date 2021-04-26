# web-api-healthcheck

This a simple *GO* web api to test the connection to a Postgres Database.

## Endpoints

*Health check* = /health - Provides a health check of the service.

`curl localhost:8080/health |jq `

```bash
{
  "status": "OK",
  "version": "0.1.0-local",
  "time": "2021-02-15 07:36:02",
  "db_information": "PostgreSQL 13.1 (Debian 13.1-1.pgdg100+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 8.3.0-6) 8.3.0, 64-bit",
  "error": ""
}
```

## Usage

`docker pull brucedominguez/web-api-healthcheck:latest`

or build locally:

`docker build -t <tag name>:latest .`

Add the following environment variables

```bash
POSTGRES_HOST=<ADD THE HOST HERE>
POSTGRES_USER=<ADD THE USER HERE>
POSTGRES_PASSWORD=<ADD THE PASSWORD HERE>
```

An example can be seen [here](./example/) which has a postgres database and webapi

```
docker-compose up
```