DROP DATABASE jquery_demo;

CREATE USER IF NOT EXISTS 'jq'@'localhost' IDENTIFIED BY 'password';
CREATE DATABASE IF NOT EXISTS jq_demo;
GRANT ALL ON `jq_demo`.* TO 'jq'@'localhost';

use jq_demo;

CREATE TABLE customer(
	id INT NOT NULL AUTO_INCREMENT,
	first_name VARCHAR(255) NOT NULL,
	last_name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	CONSTRAINT PRIMARY KEY(id)
);

INSERT INTO customer (first_name, last_name, email) VALUES ("Mickey", "Mouse", "mmouse@mines.edu"),
	("Donald", "Duck", "dduck@mines.edu"),
	("Eric", "Dattore", "edattore@mines.edu"),
	("John", "Doe", "jdoe@gmail.com"),
	("Jane", "Doe", "jdoe@gmail.com"),
	("Joe", "Schmoe", "jschmoe@gmail.com");
