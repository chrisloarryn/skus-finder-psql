# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN GOARCH=amd64 GOOS=linux go build -o /products-api cmd/main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /products-api /products-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/products-api"]