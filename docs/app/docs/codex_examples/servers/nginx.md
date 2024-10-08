---
title: Nginx
---

NGINX can be automatically configured by Codex via the built-in NGINX Plugin. This plugin will activate automatically when you install NGINX using `codex add nginx`

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/servers/nginx)

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/open/templates/nginx)

## Adding NGINX to your Shell

Run `codex add nginx`, or add the following to your `codex.json`

```json
  "packages": [
    "nginx@latest"
  ]
```

This will install the latest version of NGINX. You can find other installable versions of NGINX by running `codex search nginx`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/nginx)

## NGINX Plugin Details

### Services
* nginx

Use `codex services start|stop nginx` to start and stop the NGINX service in the background

### Helper Files
The following helper files will be created in your project directory:

* codex.d/nginx/nginx.conf
* codex.d/nginx/nginx.template
* codex.d/nginx/fastcgi.conf
* codex.d/web/index.html

Codex uses [envsubst](https://www.gnu.org/software/gettext/manual/html_node/envsubst-Invocation.html) to generate `nginx.conf` from the `nginx.template` file every time Codex starts a shell, service, or script. This allows you to create an NGINX config using environment variables by modifying `nginx.template`. To edit your NGINX configuration, you should modify the `nginx.template` file. 

Note that by default, NGINX is configured with `./codex.d/web` as the root directory. To change this, you should modify `./codex.d/nginx/nginx.template`

### Environment Variables
```bash
NGINX_CONFDIR=codex.d/nginx/nginx.conf
NGINX_PATH_PREFIX=.codex/virtenv/nginx
NGINX_TMPDIR=.codex/virtenv/nginx/temp
```

### Notes
You can easily configure NGINX by modifying these env variables in your shell's `init_hook`

To customize:
* Use $NGINX_CONFDIR to change the configuration directory
* Use $NGINX_LOGDIR to change the log directory
* Use $NGINX_PIDDIR to change the pid directory
* Use $NGINX_RUNDIR to change the run directory
* Use $NGINX_SITESDIR to change the sites directory
* Use $NGINX_TMPDIR to change the tmp directory. Use $NGINX_USER to change the user
* Use $NGINX_GROUP to customize.

You can also customize the `nginx.conf` and `fastcgi.conf` stored in `codex.d/nginx`
