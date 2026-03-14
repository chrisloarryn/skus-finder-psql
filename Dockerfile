# syntax=docker/dockerfile:1.7

FROM --platform=$BUILDPLATFORM golang:1.26.1-bookworm AS deps

WORKDIR /src

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
	go mod download

FROM deps AS build

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -trimpath -buildvcs=false -ldflags="-s -w" -o /out/products-api ./cmd/main.go

FROM gcr.io/distroless/base-debian12:nonroot AS runtime

ENV GIN_MODE=release
ENV PORT=8088

WORKDIR /app

COPY --from=build --chown=nonroot:nonroot /out/products-api /app/products-api

USER nonroot:nonroot

EXPOSE 8088

ENTRYPOINT ["/app/products-api"]
