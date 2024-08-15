---
title: MySQL
---
MySQL can be automatically configured for your dev environment by Codex via the built-in MySQL Plugin. This plugin will activate automatically when you install MySQL using `codex add mysql80` or `codex add mysql57`.

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/databases/mysql)

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/github.com/khulnasoft/codex/?folder=examples/databases/mysql)

## Adding MySQL to your Shell

`codex add mysql80`, or in your `codex.json` add

```json
    "packages": [
        "mysql80@latest"
    ]
```

You can also install Mysql 5.7 by using `codex add mysql57`

You can manually add the MySQL Plugin to your `codex.json` by adding it to your `include` list:

```json
    "include": [
        "plugin:mysql"
    ]
```

This will install the latest version of MySQL. You can find other installable versions of MySQL by running `codex search mysql80` or `codex search mysql57`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/mysql80)

## MySQL Plugin Support

Codex will automatically create the following configuration when you run `codex add mysql80` or `codex add mysql57`. You can view the full configuration by running `codex info mysql`


### Services
* mysql

You can use `codex services up|stop mysql` to start or stop the MySQL Server.

### Environment Variables

```bash
MYSQL_BASEDIR=.codex/nix/profile/default
MYSQL_HOME=./.codex/virtenv/mysql/run
MYSQL_DATADIR=./.codex/virtenv/mysql/data
MYSQL_UNIX_PORT=./.codex/virtenv/mysql/run/mysql.sock
MYSQL_PID_FILE=./.codex/mysql/run/mysql.pid
```

### Files

The plugin will also create the following helper files in your project's `.codex/virtenv` folder:

* mysql/flake.nix
* mysql/setup_db.sh
* mysql/process-compose.yaml

These files are used to setup your database and service, and should not be modified

### Notes

* This plugin wraps mysqld to work in your local project. For more information, see the `flake.nix` created in your `.codex/virtenv/mysql` folder.
* This plugin will create a new database for your project in `MYSQL_DATADIR` if one doesn't exist on shell init.
* You can use `mysqld` to manually start the server, and `mysqladmin -u root shutdown` to manually stop it
* `.sock` filepath can only be maximum 100 characters long. You can point to a different path by setting the `MYSQL_UNIX_PORT` env variable in your `codex.json` as follows:

```json
"env": {
    "MYSQL_UNIX_PORT": "/<some-other-path>/mysql.sock"
}
```
