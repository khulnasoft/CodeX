# mariadb

## mariadb Notes

1. Start the mariadb server using `codex services up`
1. Create a database using `"mysql -u root < setup_db.sql"`
1. You can now connect to the database from the command line by running `codex run connect_db`

## Services

* mariadb

Use `codex services start|stop [service]` to interact with services

## This plugin sets the following environment variables

* MYSQL_BASEDIR=/<projectDir>/.codex/nix/profile/default
* MYSQL_HOME=/<projectDir>/.codex/virtenv/mariadb/run
* MYSQL_DATADIR=/<projectDir>/.codex/virtenv/mariadb/data
* MYSQL_UNIX_PORT=/<projectDir>/.codex/virtenv/mariadb/run/mysql.sock
* MYSQL_PID_FILE=/<projectDir>/.codex/virtenv/mariadb/run/mysql.pid

To show this information, run `codex info mariadb`

Note that the `.sock` filepath can only be maximum 100 characters long. You can point to a different path by setting the `MYSQL_UNIX_PORT` env variable in your `codex.json` as follows:

```
"env": {
    "MYSQL_UNIX_PORT": "/<some-other-path>/mysql.sock"
}
```
