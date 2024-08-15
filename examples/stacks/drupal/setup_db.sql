--- You should run this query using `mysql -u root < setup_db.sql`

DROP DATABASE IF EXISTS codex_drupal;
CREATE DATABASE codex_drupal;

USE codex_drupal

CREATE USER IF NOT EXISTS 'codex_user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON codex_drupal.* TO 'codex_user'@'localhost' IDENTIFIED BY 'password';

-- Connect in drupal using:
-- Database: codex_drupal
-- User: codex_user
-- Password: password
-- Host: 127.0.0.1
-- Port: 3306
