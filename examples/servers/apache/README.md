# Apache

Apache can be automatically configured by Codex via the built-in Apache Plugin. This plugin will activate automatically when you install Apache using `codex add apache`.

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/servers/apache)

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/open/templates/apache)

### Adding Apache to your Shell

Run `codex add apache`, or add the following to your `codex.json`

```json
  "packages": [
    "apache@latest"
  ]
```

This will install the latest version of Apache. You can find other installable versions of Apache by running `codex search apache`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/apache)

## Apache Plugin Details

The Apache plugin will automatically create the following configuration when you install Apache with `codex add`.

### Services

* apache

Use `codex services start|stop apache` to start and stop httpd in the background.

### Helper Files

The following helper files will be created in your project directory:

* {PROJECT_DIR}/codex.d/apacheHttpd/httpd.conf
* {PROJECT_DIR}/codex.d/web/index.html

Note that by default, Apache is configured with `./codex.d/web` as the DocumentRoot. To change this, you should copy and modify the default `./codex.d/apacheHttpd/httpd.conf`.

### Environment Variables

```bash
HTTPD_ACCESS_LOG_FILE={PROJECT_DIR}/.codex/virtenv/apacheHttpd/access.log
HTTPD_ERROR_LOG_FILE={PROJECT_DIR}/.codex/virtenv/apacheHttpd/error.log
HTTPD_PORT=8080
HTTPD_CODEX_CONFIG_DIR={PROJECT_DIR}
HTTPD_CONFDIR={PROJECT_DIR}/codex.d/apacheHttpd
```

### Notes

We recommend copying your `httpd.conf` file to a new directory and updating HTTPD_CONFDIR if you decide to modify it.
