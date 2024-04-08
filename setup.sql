CREATE DATABASE IF NOT EXISTS mydatabase;

USE mydatabase;

CREATE TABLE IF NOT EXISTS volunteers (
    id INT AUTO_INCREMENT,
    name VARCHAR(100),
    number VARCHAR(100),
    phone VARCHAR(100),
    email VARCHAR(100),
    PRIMARY KEY(id)
);
