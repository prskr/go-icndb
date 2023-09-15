# go-icndb

This is a simple clone of **I**nternet **C**huck **N**orris **DB**.
The scope of this project is to be a sample project that can be used for:

* trainings
* experiments e.g. with load testint tools
* probably something else

## Getting started

To start the API locally just start the container and publish the default port **3000**.

```shell
docker run --rm -ti -p 3000:3000 ghcr.io/baez90/go-icndb:latest
```

Then start using it: [http://localhost:3000](http://localhost:3000).
When opening the root of the app you're automatically redirected to a Swagger UI describing how to use the API.

## Configuration

`go-icndb` being a nice 'cloud native'-ish application can be configured either via CLI flags or via environment
variables.

| Option             | CLI switch                   | Env variable                     |
|--------------------|------------------------------|----------------------------------|
| Listen address     | `--http.address`             | `ICNDB_HTTP_ADDRESS`             |
| Default first name | `--jokes.default-first-name` | `ICNDB_JOKES_DEFAULT_FIRST_NAME` |
| Default last name  | `--jokes.default-last-name`  | `ICNDB_JOKES_DEFAULT_LAST_NAME`  |

## Observability

`go-icndb` exposes some observability related endpoints:

- `/health` - basic health check, doesn't do much
- `/metrics` - exposes Prometheus compliant metrics