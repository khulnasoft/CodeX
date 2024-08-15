--- You should run this query using psql < setup_db.sql`

DROP DATABASE IF EXISTS codex_lamp;
CREATE DATABASE codex_lamp;

CREATE USER codex_user WITH PASSWORD 'password';

DROP TABLE IF EXISTS colors;
CREATE TABLE colors (
	id SERIAL NOT NULL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	hex VARCHAR(7) NOT NULL);

INSERT INTO colors (name, hex) VALUES ('red', '#FF0000'), ('blue', '#0000FF'), ('green', '#00FF00');

GRANT ALL PRIVILEGES ON colors TO codex_user;
