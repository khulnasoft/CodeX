---
title: Using Plugins
---

This doc describes how to use Codex Plugins with your project. **Plugins** provide a default Codex configuration for a Nix package. Plugins make it easier to get started with packages that require additional setup when installed with Nix, and they offer a familiar interface for configuring packages. They also help keep all of your project's configuration within your project directory, which helps maintain portability and isolation.

## Using Plugins

### Built-in Plugins

If you add one of the packages listed above to your project using `codex add <pkg>`, Codex will automatically activate the plugin for that package.

You can also explicitly add a built-in plugin in your project by adding it to the [`include` section](../configuration.md#include) of your `codex.json` file. For example, to explicitly add the plugin for Nginx, you can add the following to your `codex.json` file:

```json
{
  "include": [
    "plugin:nginx"
  ]
}
```

Built-in plugins are available for the following packages. You can activate the plugins for these packages by running `codex add <package_name>`

* [Apache](../codex_examples/servers/apache.md) (apacheHttpd)
* [Caddy](../codex_examples/servers/caddy.md) (caddy)
* [Nginx](../codex_examples/servers/nginx.md) (nginx)
* [Node.js](../codex_examples/languages/nodejs.md) (nodejs, nodejs-slim)
* [MariaDB](../codex_examples/databases/mariadb.md) (mariadb, mariadb_10_6...)
* [MySQL](../codex_examples/databases/mysql.md) (mysql80, mysql57)
* [PostgreSQL](../codex_examples/databases/postgres.md) (postgresql)
* [Redis](../codex_examples/databases/redis.md) (redis)
* [Valkey](../codex_examples/databases/valkey.md) (valkey)
* [PHP](../codex_examples/languages/php.md) (php, php80, php81, php82...)
* [Pip](../codex_examples/languages/python.md) (python39Packages.pip, python310Packages.pip, python311Packages.pip...)
* [Ruby](../codex_examples/languages/ruby.md)(ruby, ruby_3_1, ruby_3_0...)


### Local Plugins

You can also [define your own plugins](./creating_plugins.md) and use them in your project. To use a local plugin, add the following to the `include` section of your codex.json:

```json
  "include": [
    "path:./path/to/plugin.json"
  ]
```

### Github Hosted Plugins

Sometimes, you may want to share a plugin across multiple projects or users. In this case, you provide a Github reference to a plugin hosted on Github. To install a github hosted plugin, add the following to the include section of your codex.json

```json
  "include": [
    "github:<org>/<repo>?dir=<plugin-dir>"
  ]
```

## An Example of a Plugin: Nginx
Let's take a look at the plugin for Nginx. To get started, let's initialize a new codex project, and add the `nginx` package:

```bash
cd ~/my_proj
codex init && codex add nginx
```

Codex will install the package, activate the `nginx` plugin, and print a short explanation of the plugin's configuration

```bash
Installing nix packages. This may take a while... done.

nginx NOTES:
nginx can be configured with env variables

To customize:
* Use $NGINX_CONFDIR to change the configuration directory
* Use $NGINX_LOGDIR to change the log directory
* Use $NGINX_PIDDIR to change the pid directory
* Use $NGINX_RUNDIR to change the run directory
* Use $NGINX_SITESDIR to change the sites directory
* Use $NGINX_TMPDIR to change the tmp directory. Use $NGINX_USER to change the user
* Use $NGINX_GROUP to customize.

Services:
* nginx

Use `codex services start|stop [service]` to interact with services

This plugin creates the following helper files:
* ~/my_project/codex.d/nginx/nginx.conf
* ~/my_project/codex.d/nginx/fastcgi.conf
* ~/my_project/codex.d/web/index.html

This plugin sets the following environment variables:
* NGINX_CONFDIR=~/my_project/codex.d/nginx/nginx.conf
* NGINX_PATH_PREFIX=~/my_project/.codex/virtenv/nginx
* NGINX_TMPDIR=~/my_project/.codex/virtenv/nginx/temp

To show this information, run `codex info nginx`

nginx (nginx-1.22.1) is now installed.
```

Based on this info page, we can see that Codex has created the configuration we need to run `nginx` in our local shell. Let's take a look at the files it created:

```bash
% tree
.
├── codex.d
│   ├── nginx
│   │   ├── fastcgi.conf
│   │   └── nginx.conf
│   └── web
│       └── index.html
└── codex.json
```

These files give us everything we need to run NGINX, and we can modify the `nginx.conf` and `fastcgi.conf` to customize how Nginx works.

We can also see in the info page that Codex has configured an NGINX service for us. Let's start this service with `codex services start nginx`, and then test it with `curl`:

```bash
> codex services start nginx

Installing nix packages. This may take a while... done.
Starting a codex shell...
Service "nginx" started

> curl localhost:80
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <title>Hello World!</title>
  </head>
  <body>
    Hello World!
  </body>
</html>
```

## Plugin Configuration in detail

When Codex detects a plugin for an installed package, it automatically applies its configuration and prints a short explanation. Developers can review this explanation anytime using `codex info <package_name>`.

### Services
If your package can run as a daemon or background service, Codex can configure and manage that service for you with `codex services`.

To learn more, visit our page on [Codex Services](services.md).

### Environment Variables
Codex stores default environment variables for your package in `.codex/virtenv/<package_name>/.env` in your project directory. Codex automatically updates these environment variables whenever you run `codex shell` or `codex run` to match your current project, and developers should not check these `.env` files into source control.

#### Customizing Environment Variables
If you want to customize the environment variables, you can override them in the `init_hook` of your `codex.json`

### Helper Files
Helper files are files that your package may use for configuration purposes, such as NGINX's `nginx.conf` file. When installing a package, Codex will check for helper files in your project's `codex.d` folder and create them if they do not exist. If helper files are already present, Codex will not overwrite them.

#### Customizing Helper Files
Developers should directly edit helper files and check them into source control if needed

## Plugins Source Code

Codex Plugins are written in JSON and stored in the main Codex Repo. You can view the source code of the current plugins [here](https://github.com/khulnasoft/codex/tree/main/plugins)

