---
title: MariaDB
---
MariaDB can be automatically configured for your dev environment by Codex via the built-in MariaDB Plugin. This plugin will activate automatically when you install MariaDB using `codex add mariadb`, or when you use a versioned Nix package like `codex add mariadb_1010`

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/databases/mariadb)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/mariadb)

## Adding MariaDB to your Shell

`codex add mariadb`, or in your `codex.json` add

```json
    "packages": [
        "mariadb@latest"
    ]
```

You can manually add the MariaDB Plugin to your `codex.json` by adding it to your `include` list:

```json
    "include": [
        "plugin:mariadb"
    ]
```
This will install the latest version of MariaDB. You can find other installable versions of MariaDB by running `codex search mariadb`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/mariadb)

## MariaDB Plugin Support

Codex will automatically create the following configuration when you run `codex add mariadb`. You can view the full configuration by running `codex info mariadb`


### Services
* mariadb

You can use `codex services up|stop mariadb` to start or stop the MariaDB Server.

### Environment Variables

```bash
MYSQL_BASEDIR=.codex/nix/profile/default
MYSQL_HOME=./.codex/virtenv/mariadb/run
MYSQL_DATADIR=./.codex/virtenv/mariadb/data
MYSQL_UNIX_PORT=./.codex/virtenv/mariadb/run/mysql.sock
MYSQL_PID_FILE=./.codex/mariadb/run/mysql.pid
```

### Files

The plugin will also create the following helper files in your project's `.codex/virtenv` folder:

* mariadb/flake.nix
* mariadb/setup_db.sh
* mariadb/process-compose.yaml

These files are used to setup your database and service, and should not be modified

### Notes

* This plugin wraps mysqld and mysql_install_db to work in your local project. For more information, see the `flake.nix` created in your `.codex/virtenv/mariadb` folder.
* This plugin will create a new database for your project in MYSQL_DATADIR if one doesn't exist on shell init.
* You can use `mysqld` to manually start the server, and `mysqladmin -u root shutdown` to manually stop it
* `.sock` filepath can only be maximum 100 characters long. You can point to a different path by setting the `MYSQL_UNIX_PORT` env variable in your `codex.json` as follows:

```json
"env": {
    "MYSQL_UNIX_PORT": "/<some-other-path>/mysql.sock"
}
```
