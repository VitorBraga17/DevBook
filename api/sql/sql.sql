CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios (
    id int auto_increment primary key,
    nome varchar(255) NOT NULL,
    nick varchar(255) NOT NULL unique,
    email varchar(255) NOT NULL unique,
    senha varchar(255) NOT NULL unique,
    criadoEm timestamp default current_timestamp()
) ENGINE=InnoDB;