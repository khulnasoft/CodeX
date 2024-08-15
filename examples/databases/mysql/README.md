# mysql

## mysql Notes

1. Start the mysql server using `codex services up`
1. Create a database using `"mysql -u root < setup_db.sql"`
1. You can now connect to the database from the command line by running `codex run connect_db`

## Services

* mysql

Use `codex services start|stop [service]` to interact with services

## This plugin sets the following environment variables

* MYSQL_BASEDIR=&lt;projectDir>/.codex/nix/profile/default
* MYSQL_HOME=&lt;projectDir>/.codex/virtenv/mysql/run
* MYSQL_DATADIR=&lt;projectDir>/.codex/virtenv/mysql/data
* MYSQL_UNIX_PORT=&lt;projectDir>/.codex/virtenv/mysql/run/mysql.sock
* MYSQL_PID_FILE=&lt;projectDir>/.codex/virtenv/mysql/run/mysql.pid

To show this information, run `codex info mysql`

Note that the `.sock` filepath can only be maximum 100 characters long. You can point to a different path by setting the `MYSQL_UNIX_PORT` env variable in your `codex.json` as follows:

```json
"env": {
    "MYSQL_UNIX_PORT": "/<some-other-path>/mysql.sock"
}
```
