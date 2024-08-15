---
title: Redis
---

Redis can be configured automatically using Codex's built in Redis plugin. This plugin will activate automatically when you install Redis using `codex add redis`

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/databases/redis)

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/open/templates/redis)

## Adding Redis to your shell

`codex add redis`, or in your Codex.json

```json
    "packages": [
        "redis@latest   "
    ],
```

This will install the latest version of Redis. You can find other installable versions of Redis by running `codex search redis`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/redis)

## Redis Plugin Details

The Redis plugin will automatically create the following configuration when you install Redis with `codex add`

### Services

* redis

Use `codex services start|stop [service]` to interact with services

### Helper Files

The following helper files will be created in your project directory:

* \{PROJECT_DIR\}/codex.d/redis/redis.conf


### Environment Variables

```bash
REDIS_PORT=6379
REDIS_CONF=./codex.d/redis/redis.conf
```

### Notes

Running `codex services start redis` will start redis as a daemon in the background.

You can manually start Redis in the foreground by running `redis-server $REDIS_CONF --port $REDIS_PORT`.

Logs, pidfile, and data dumps are stored in `.codex/virtenv/redis`. You can change this by modifying the `dir` directive in `codex.d/redis/redis.conf`
