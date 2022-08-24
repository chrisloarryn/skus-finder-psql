DROP TABLE IF EXISTS products;

CREATE TABLE products
(
    id              serial PRIMARY KEY NOT NULL,
    name            VARCHAR(50)        NOT NULL,
    brand           VARCHAR(50)        NOT NULL,
    size            VARCHAR(255),
    price           double precision   NOT NULL,
    principal_image VARCHAR(255)       NOT NULL,
    other_images    TEXT[]
);