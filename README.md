# skus-finder-psql

[![Run tests](https://github.com/chrisloarryn/skus-finder-psql/actions/workflows/test.yaml/badge.svg)](https://github.com/chrisloarryn/skus-finder-psql/actions/workflows/test.yaml)

REST API to manage products by SKU. The service supports create, list, get by SKU, update and delete operations, and it can run with an in-memory repository or PostgreSQL persistence.

## Updated stack

- Go `1.26.1` with `toolchain go1.26.1`
- CI matrix on Go `1.25.x` and `1.26.x`
- Gin `v1.12.0`
- Gorm `v1.31.1`
- PostgreSQL driver `v1.6.0`
- MySQL driver `v1.6.0`
- Testify `v1.11.1`
- Ozzo Validation `v4.3.0`
- PostgreSQL container `18.3`

## Run with Docker Compose

Requirements:

- Docker Desktop or Docker Engine with Compose

Start the API and PostgreSQL:

```bash
docker compose up --build
```

Services:

- API: `http://localhost:8088`
- PostgreSQL: `localhost:65432`

The compose setup builds both Dockerfiles in the repository:

- Root `Dockerfile`: hardened multi-stage build with Go `1.26.1`, distroless runtime and non-root execution
- `sql/Dockerfile`: PostgreSQL `18.3` with `init.sql` and `postgres` user

## Run locally without Docker

By default the app uses the in-memory repository. If `ENVIRONMENT=PRODUCTION`, it connects to PostgreSQL and runs `AutoMigrate` for the `products` table.

1. Copy `.env.example` to `.env` and adjust values if needed.
2. Set the variables for PostgreSQL or leave `ENVIRONMENT` empty to use the in-memory repository.
3. Run the API:

```bash
go run ./cmd/main.go
```

Useful database variables:

```bash
ENVIRONMENT=PRODUCTION
PORT=8088
DB_HOST=localhost
DB_PORT=65432
DB_NAME=postgres
DB_USER=postgres
DB_PASSWORD=postgres
```

## Quality checks

```bash
go vet ./...
go test ./...
```

Static analysis and security checks used in the repository:

```bash
golangci-lint run ./...
gosec ./...
govulncheck ./...
trivy config --severity HIGH,CRITICAL --misconfig-scanners dockerfile .
trivy fs --scanners vuln,secret,misconfig --severity HIGH,CRITICAL .
```

## CI workflow

GitHub Actions runs on pushes and pull requests for `main` and `develop`, and it also supports manual execution with `workflow_dispatch`. Workflow concurrency is enabled per workflow and ref to avoid overlapping runs on the same branch.

## Product rules

| Field | Description | Allowed values | Required |
|---|---|---|---|
| SKU | Candidate identifier of the product | `FAL-1000000` to `FAL-99999999` | Yes |
| Name | Short description | Min length `3`, max length `50` | Yes |
| Brand | Brand name | Min length `3`, max length `50` | Yes |
| Size | Product size | Optional text | No |
| Price | Sell price | `1.00` to `99999999.00` | Yes |
| Principal image | Main catalog image | Valid URL | Yes |
| Other images | Additional images | Valid URL list | No |

Sample products:

| SKU | Name | Brand | Size | Price |
|---|---|---|---|---|
| FAL-8406270 | 500 Zapatilla Urbana Mujer | New Balance | 37 | 42990.00 |
| FAL-881952283 | Bicicleta Baltoro Aro 29 | Jeep | ST | 399990.00 |
| FAL-881898502 | Camisa Manga Corta Hombre | Basement | M | 24990.00 |

## Architecture notes

- HTTP layer built with Gin
- Domain validations implemented with Ozzo Validation
- Persistence abstraction through `products.Repository`
- PostgreSQL persistence with Gorm
- In-memory repository available for local development and tests
- Unit tests based on `gomock` and `testify`

## Endpoints

- `GET /ping`
- `GET /api/v1/products`
- `POST /api/v1/products`
- `GET /api/v1/products/{productSKU}`
- `PATCH /api/v1/products/{productSKU}`
- `DELETE /api/v1/products/{productSKU}`
