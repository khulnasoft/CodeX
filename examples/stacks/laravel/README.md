# Laravel

Laravel is a powerful web application framework built with PHP. It's a great choice for building web applications and APIs.

This example shows how to build a simple Laravel application backed by MariaDB and Redis. It uses Codex Plugins for all 3 Nix packages to simplify configuration

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/github.com/khulnasoft/codex/?folder=examples/stacks/laravel)

## How to Run

1. Install [Codex](https://www.khulnasoft/codex/docs/installing_codex/)

1. Create a new Laravel App by running `codex create --template laravel`. This will create a new Laravel project in your current directory.

1. Start your MariaDB and Redis services by running `codex services up`.
   1. This step will also create an empty MariaDB Data Directory and initialize your database with the default settings
   2. This will also start the php-fpm service for serving your PHP project over fcgi. Learn more about [PHP-FPM](https://www.php.net/manual/en/install.fpm.php)

1. Create the laravel database by running `codex run db:create`, and then run Laravel's initial migrations using `codex run db:migrate`

1. You can now start the artisan server by running `codex run serve:dev`. This will start the server on port 8000, which you can access at `localhost:8000`

1. If you're using Laravel on Codex Cloud, you can test the app by appending `/port/8000` to your Codex Cloud URL

1. For more details on building and developing your Laravel project, visit the [Laravel Docs](https://laravel.com/docs/10.x)


## How to Recreate this Example

### Creating the Laravel Project

1. Create a new project with `codex init`

2. Add the packages using the command below. Installing the packages with `codex add` will ensure that the plugins are activated:

    ```bash
    codex add mariadb@latest, php@8.1, nodejs@18, redis@latest, php81Packages.composer@latest
    ```

3. Run `codex shell` to start your shell. This will also initialize your database by running `initdb` in the init hook.

4. Create your laravel project by running:

    ```bash
    composer create-project laravel/laravel tmp

    mv tmp/* tmp/.* .
    ```

### Setting up MariaDB

To use MariaDB, you need to create the default Laravel database. You can do this by running the following commands in your `codex shell`:

```bash
# Start the MariaDB service
codex services up mariadb -b

# Create the database
mysql -u root -e "CREATE DATABASE laravel;"

# Once you're done, stop the MariaDB service
codex services stop mariadb
```
