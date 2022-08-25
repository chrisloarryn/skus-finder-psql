# skus-finder-psql

[![Run tests](https://github.com/chrisloarryn/skus-finder-psql/actions/workflows/test.yaml/badge.svg)](https://github.com/chrisloarryn/skus-finder-psql/actions/workflows/test.yaml)

### Specifications

``To execute, needs:``

### Development

- docker desktop/docker

### Build and execution instructions

```docker-compose up -d --build```

| Command | Description                              |
|---------|------------------------------------------|
| -d      | for detach and run API in the background |
| --build | to force rebuild of api                  |

### Production

``a postgres db will be configured, otherwise take following default values:``

````shell
DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=postgres
DB_USER=postgres
DB_PASSWORD=postgres
````

### Problem

``Designing and implementing an application that allows to store new product, list all of them, retrieve a product by its SKU, update it and delete it.
``

### Description

``The information of a product that we want to store is: ``

| Field           | Description                                                                                                                                             | Data allowed                                                | Required |
|-----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------|----------|
| SKU             | Internal stock-keeping unit. It is the candidate identifier of a product                                                                                | Min: FAL-1000000 Max: FAL-99999999                          | Y        |
| Name            | Short description of the product                                                                                                                        | Must not be blank Min size: 3 Max size: 50                  | Y        |
| Brand           | Name of the brand                                                                                                                                       | Must not be blank Min size: 3 Max size: 50                  | Y        |
| Size            | Size of the product                                                                                                                                     | Must not be blank                                           | N        |
| Price           | Sell price                                                                                                                                              | Min: 1.00                                  Max: 99999999.00 | Y        |
| Principal image | URL of the principal image of the product, which is used in catalogs and is the first image that is showed to customers when access product detail page | URL format                                                  | Y        |
| Other Images    | List of images of the product.                                                                                                                          | URL format                                                  | N        |

### Additional Configurations

- The designed endpoints must use proper HTTP verb, REST naming conventions and return correct HTTP code.
- The application must not expose technology detail, such as language, framework, libraries, and so on, when an
  exception is thrown.

| SKU           | Name                       | Brand       | Size | Price     | Image URLs |
|---------------|----------------------------|-------------|------|-----------|------------|
| FAL-8406270   | 500 Zapatilla Urbana Mujer | New Balance | 37   | 42990.00  | - https:// |
| FAL-881952283 | Bicicleta Baltoro Aro 29   | Jeep        | ST   | 399990.00 | - https:// |
| FAL-881898502 | Camisa Manga Corta Hombre  | Basement    | M    | 24990.00  | - https:// |

### Technology Required

- [x] Golang
- [x] Go mod
- [x] Gin Framework
- [x] Gorm
- [x] Database persistence
- [x] InMemory persistence
- [x] Unit testing
- [ ] Validations of field size/length

#### extras

- [x] gomock (for mocks generation) [[link]](https://github.com/golang/mock)
- [x] testify (for assertions) [[link]](https://github.com/stretchr/testify)

### Deliverables

- The only deliverable is the source code of the solution; it must be published at a GIT version control hosting such
  as: GitHub, GitLab, Bitbucket or other. The repository must have a
  README.md file which contains:
    - Build and execution instructions
    - Brief explanation about architectural and technological decisions made for the application.

### Endpoints detailed (common REST API)

- GET /ping: returns a "pong" message
- GET /api/v1/products: Show all products stored.
- POST /api/v1/products: Allow to create a product.
- GET /api/v1/products/{productSKU}: Get details for a specific product by its product sku.
- PATCH /api/v1/products/{productSKU}: Update a specific product by its product sku.
- DELETE /api/v1/products/{productSKU}: Delete a specific product by its product sku.
    
