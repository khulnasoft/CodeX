-- You should run this query using `mysql -u root < setup_db.sql`

DROP DATABASE IF EXISTS codex_lamp;
CREATE DATABASE codex_lamp;

USE codex_lamp;

CREATE USER 'codex_user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON codex_lamp.* TO 'codex_user'@'localhost';

DROP TABLE IF EXISTS colors;
CREATE TABLE colors (
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(100) NOT NULL,
	hex VARCHAR(7) NOT NULL,
	PRIMARY KEY (id));

INSERT INTO colors (name, hex) VALUES ('red', '#FF0000'), ('blue', '#0000FF'), ('green', '#00FF00');


