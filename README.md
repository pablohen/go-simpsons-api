# go-simpsons-api

A small [Gin](https://github.com/gin-gonic/gin) service that exposes the same JSON routes as [The Simpsons API](https://thesimpsonsapi.com/) by **reverse-proxying** to `https://thesimpsonsapi.com`. Responses (including pagination `next` / `prev` links) match the upstream API. [Swagger / OpenAPI](https://github.com/swaggo/swag) documents the expected shapes at `/swagger/index.html`.

## Requirements

- [Go](https://go.dev/dl/) 1.26+ (see `go.mod`)
- Optional: [GNU Make](https://www.gnu.org/software/make/) (for the [`Makefile`](Makefile))

## Quick start

```bash
make run
# or: go run .
```

Then:

- Health: `GET http://localhost:8080/health`
- Proxied API: `GET http://localhost:8080/api`, `/api/characters`, `/api/episodes`, `/api/locations`, plus `/:id` routes as on the upstream service.
- Swagger UI: `http://localhost:8080/swagger/index.html`

### Makefile

| Command | Description |
| ------- | ----------- |
| `make` / `make help` | List targets |
| `make run` | Run the server (`go run .`) |
| `make build` | Build `bin/go-simpsons-api` |
| `make test` | `go test ./...` |
| `make vet` | `go vet ./...` |
| `make tidy` | `go mod tidy` |
| `make swagger` | Regenerate `docs/` from `main.go` and `internal/models` |
| `make dev` | Live reload with [Air](https://github.com/air-verse/air) (`air` if on `PATH`, else `go run …`) |
| `make clean` | Remove `bin/`, `tmp/`, and Air’s `build-errors.log` |

### Environment

| Variable   | Default                         | Description                                      |
| ---------- | ------------------------------- | ------------------------------------------------ |
| `UPSTREAM` | `https://thesimpsonsapi.com`    | Origin used for the reverse proxy (scheme + host). |
| `PORT`     | `8080`                          | Listen port (only the number, e.g. `8080`).      |

Example:

```bash
PORT=3000 UPSTREAM=https://thesimpsonsapi.com make run
```

## Development

### Live reload ([Air](https://github.com/air-verse/air))

```bash
go install github.com/air-verse/air@latest
make dev
```

Configuration is in [`.air.toml`](.air.toml). Builds go to `tmp/` (removed by `make clean`).

### Regenerate Swagger

After changing route annotations in [`main.go`](main.go) or models in [`internal/models`](internal/models):

```bash
make swagger
```
