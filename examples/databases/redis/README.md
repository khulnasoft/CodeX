# redis-7.0.5

## redis Notes

Running `codex services start redis` will start redis as a daemon in the background.

You can manually start Redis in the foreground by running `redis-server $REDIS_CONF --port $REDIS_PORT`.

Logs, pidfile, and data dumps are stored in `.codex/virtenv/redis`. You can change this by modifying the `dir` directive in `codex.d/redis/redis.conf`

## Services

* redis

Use `codex services start|stop [service]` to interact with services

## This plugin creates the following helper files

* ./codex.d/redis/redis.conf

## This plugin sets the following environment variables

* REDIS_PORT=6379
* REDIS_CONF=./codex.d/redis/redis.conf

To show this information, run `codex info redis`
