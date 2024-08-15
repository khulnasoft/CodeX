# LEPP Stack

An example Codex shell for NGINX, Postgres, and PHP. This example uses Codex Plugins for all 3 packages to simplify configuration

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/lepp-stack)

## How to Run

### Initializing

In this directory, run: `codex run init_db` to initialize a db.

To start the Servers + Postgres service, run: `codex services start`

### Creating the DB

You can run the creation script using `codex run create_db`. This will create a Postgres DB based on `setup_postgres_db.sql`.

### Testing the Example

You can query Nginx on port 80, which will route to the PHP example.

## How to Recreate this Example

1. Create a new project with:
   ```bash
   codex create --template lapp-stack
   codex install
   ```

2. Update `codex.d/nginx/httpd.conf` to point to the directory with your PHP files. You'll need to update the `root` directive to point to your project folder
3. Follow the instructions above in the How to Run section to initialize your project.

Note that the `.sock` filepath can only be maximum 100 characters long. You can point to a different path by setting the `PGHOST` env variable in your `codex.json` as follows:

```
"env": {
    "PGHOST": "/<some-shorter-path>"
}
```

### Related Docs

* [Using PHP with Codex](https://www.khulnasoft/codex/docs/codex_examples/languages/php/)
* [Using Nginx with Codex](https://www.khulnasoft/codex/docs/codex_examples/servers/nginx/)
* [Using PostgreSQL with Codex](https://www.khulnasoft/codex/docs/codex_examples/databases/postgres/)
