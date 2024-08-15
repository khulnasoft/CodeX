--- You should run this query using psql < setup_db.sql`


DO
$do$
BEGIN
   IF EXISTS (SELECT FROM pg_catalog.pg_roles WHERE  rolname = 'codex_user') THEN
      RAISE NOTICE 'Role "my_user" already exists. Skipping.';
   ELSE
      CREATE USER codex_user WITH PASSWORD 'password';
   END IF;
END
$do$;

DROP TABLE IF EXISTS address_book;
CREATE TABLE address_book (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  phone VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL
);

INSERT INTO address_book (first_name, last_name, phone, email) VALUES ('Jim', 'Hawkins', '555-0100', 'jhawk@khulnasoft.com'), ('Billy', 'Bones', '555-0102', 'bbones@khulnasoft.com');

GRANT ALL PRIVILEGES ON address_book TO codex_user;
