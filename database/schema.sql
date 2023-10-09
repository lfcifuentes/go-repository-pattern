CREATE DATABASE IF NOT EXISTS repositorydb;

-- Añade un punto y coma al final de la declaración anterior
-- para separarla de las siguientes sentencias SQL

CREATE TABLE IF NOT EXISTS usuarios (
                                        id INT AUTO_INCREMENT PRIMARY KEY,
                                        nombre VARCHAR(255),
                                        email VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS productos (
                                         id INT AUTO_INCREMENT PRIMARY KEY,
                                         nombre_producto VARCHAR(255),
                                         precio DECIMAL(10, 2),
                                         descripcion TEXT
);