CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

CREATE TABLE tb_user(
  id INT auto_increment primary key,
  name VARCHAR(50) NOT NULL,
  nick VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(100) NOT NULL,
  created_at TIMESTAMP default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE tb_follower(
  user_id INT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES tb_user(id) ON DELETE CASCADE,
  follower_id INT NOT NULL,
  FOREIGN KEY (follower_id) REFERENCES tb_user(id) ON DELETE CASCADE,

  PRIMARY KEY(user_id, follower_id)
) ENGINE=INNODB;
