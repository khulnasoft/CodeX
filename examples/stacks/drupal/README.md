# Drupal Stack

This example shows how to run a Drupal application in Codex. It makes use of the PHP and Apache Plugins, while demonstrating how to configure a MariaDB instance to work with Codex Cloud.

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/open/templates/drupal)

## How to Run the example

In this directory, run:

`codex shell`

To start all your services (PHP, MySQL, and NGINX), run `codex services up`. To stop the services, run `codex services stop`

To create the `codex_drupal` database and example table, you should run:

`mysql -u root < setup_db.sql`

To install Drupal and your dependencies, run `composer install`. The Drupal app will be installed in the `/web` directory, and you can configure your site by visiting `localhost:8000/autoload` in your browser and following the interactive instructions

To exit the shell, use `exit`

## Configuration

Because the Nix Store is immutable, we need to store our configuration, data, and logs in a local project directory. This is stored in the `codex.d` directory, in a subfolder for each of the packages that we will be installing.
