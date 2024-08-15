---
title: Valkey
---

Valkey can be configured automatically using Codex's built in Valkey plugin. This plugin will activate automatically when you install Valkey using `codex add valkey`

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/databases/valkey)

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/open/templates/valkey)

## Adding Valkey to your shell

`codex add valkey`, or in your Codex.json

```json
    "packages": [
        "valkey@latest   "
    ],
```

This will install the latest version of Valkey. You can find other installable versions of Valkey by running `codex search valkey`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/valkey)

## Valkey Plugin Details

The Valkey plugin will automatically create the following configuration when you install Valkey with `codex add`

### Services

* valkey

Use `codex services start|stop [service]` to interact with services

### Helper Files

The following helper files will be created in your project directory:

* \{PROJECT_DIR\}/codex.d/valkey/valkey.conf


### Environment Variables

```bash
VALKEY_PORT=6379
VALKEY_CONF=./codex.d/valkey/valkey.conf
```

### Notes

Running `codex services start valkey` will start valkey as a daemon in the background.

You can manually start Valkey in the foreground by running `valkey-server $VALKEY_CONF --port $VALKEY_PORT`.

Logs, pidfile, and data dumps are stored in `.codex/virtenv/valkey`. You can change this by modifying the `dir` directive in `codex.d/valkey/valkey.conf`
