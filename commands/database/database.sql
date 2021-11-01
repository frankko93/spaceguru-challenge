CREATE DATABASE IF NOT EXISTS spaceguru;

USE spaceguru;

CREATE TABLE IF NOT EXISTS users (
    id                          INT(11) NOT NULL AUTO_INCREMENT,
    name                        VARCHAR(200) NOT NULL,
    surname                     VARCHAR(200) NOT NULL,
    type                        VARCHAR(50) NOT NULL,
    status                      VARCHAR(50) NOT NULL,
    email                       VARCHAR(200) NOT NULL,
    password                    VARCHAR(200) NOT NULL,
    createdAt                   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt                   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (id)) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS vehicles (
    id                          INT(11) NOT NULL AUTO_INCREMENT,
    type                        VARCHAR(50) NOT NULL,
    status                      VARCHAR(50) NOT NULL,
    description                 MEDIUMTEXT NULL DEFAULT NULL,
    createdAt                   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt                   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (id)) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS travels (
    id                          INT(11) NOT NULL AUTO_INCREMENT,
    type                        VARCHAR(50) NOT NULL,
    status                      VARCHAR(50) NOT NULL,
    route                       VARCHAR(50) NOT NULL,
    createdAt                   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updatedAt                   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (id)) ENGINE=InnoDB;
