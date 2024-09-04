CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

CREATE TABLE tb_user(
  id INT auto_increment primary key,
  name VARCHAR(50) NOT NULL,
  nick VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(50) NOT NULL UNIQUE,
  created_at TIMESTAMP default current_timestamp()
) ENGINE=INNODB;