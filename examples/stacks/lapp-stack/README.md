# LAPP Stack

This example shows how to build a simple application using Apache, PHP, and PostgreSQL. It uses Codex Plugins for all 3 packages to simplify configuration.

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/open/templates/lapp-stack)

## How to Run

The following steps may be done inside or outside a codex shell.

1. Initialize a database by running `codex run init_db`.
1. Create the database and load the test data by using `codex run create_db`.
1. Start Apache, PHP-FPM, and Postgres in the background by run `codex services start`.
1. You can now test the app using `localhost:8080` to hit the Apache Server. If you want Apache to listen on a different port, you can change the `HTTPD_PORT` environment variable in the Codex init_hook.

### How to Recreate this Example

1. Create a new project with:
    ```bash
    codex create --template lapp-stack
    codex install
    ```

1. Update `codex.d/apache/httpd.conf` to point to the directory with your PHP files. You'll need to update the `DocumentRoot` and `Directory` directives.
1. Follow the instructions above in the How to Run section to initialize your project.

### Related Docs

* [Using PHP with Codex](https://www.khulnasoft/codex/docs/codex_examples/languages/php/)
* [Using Apache with Codex](https://www.khulnasoft/codex/docs/codex_examples/servers/apache/)
* [Using PostgreSQL with Codex](https://www.khulnasoft/codex/docs/codex_examples/databases/postgres/)
