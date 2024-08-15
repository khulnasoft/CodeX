# Caddy

Caddy can be configured automatically using Codex's built in Caddy plugin. This plugin will activate automatically when you install Caddy using `codex add caddy`

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/servers/caddy)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/caddy)

### Adding Caddy to your Shell

Run `codex add caddy`, or add the following to your `codex.json`

```json
  "packages": [
    "caddy@latest"
  ]
```

This will install the latest version of Caddy. You can find other installable versions of Caddy by running `codex search caddy`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/caddy)

## Caddy Plugin Details

The Caddy plugin will automatically create the following configuration when you install Caddy with `codex add`

### Services

* caddy

Use `codex services start|stop caddy` to start and stop httpd in the background

### Helper Files

The following helper files will be created in your project directory:

* {PROJECT_DIR}/codex.d/caddy/Caddyfile
* {PROJECT_DIR}/codex.d/web/index.html

Note that by default, Caddy is configured with `./codex.d/web` as the root. To change this, you should modify the default `./codex.d/caddy/Caddyfile` or change the `CADDY_ROOT_DIR` environment variable

### Environment Variables

```bash
* CADDY_CONFIG={PROJECT_DIR}/codex.d/caddy/Caddyfile
* CADDY_LOG_DIR={PROJECT_DIR}/.codex/virtenv/caddy/log
* CADDY_ROOT_DIR={PROJECT_DIR}/codex.d/web
```

### Notes

You can customize the config used by the caddy service by modifying the Caddyfile in codex.d/caddy, or by changing the CADDY_CONFIG environment variable to point to a custom config. The custom config must be either JSON or Caddyfile format.
