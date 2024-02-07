CREATE DATABASE product_db;

USE product_db;

CREATE TABLE products
(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(256) NOT NULL ,
    price DECIMAL(20, 2) NOT NULL DEFAULT 0.00,
    description VARCHAR(1000) NULL ,
    quantity INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

INSERT INTO products (name, price, description, quantity, created_at, updated_at)
VALUES ('iPhone 15 Pro', 22999000, 'iPhone terbaru', 100, '2024-02-07 10:00:00', '2024-02-07 10:00:00'),
       ('iPhone 14 Pro', 16999000, 'iPhone tahun lalu', 100, '2022-09-22 10:00:00', '2022-09-22 10:00:00'),
       ('Macbook Pro M2 Pro 16/1TB', 34999000, 'Macbook tahun lalu', 100, '2022-10-10 10:00:00', '2022-10-10 10:00:00'),
       ('Macbook Pro M3 Pro 18/1TB', 44999000, 'Macbook Pro Terbaru', 100, '2023-10-30 10:00:00', '2023-10-30 10:00:00')