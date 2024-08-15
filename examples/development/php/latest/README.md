# PHP

PHP projects can manage most of their dependencies locally with `composer`. Some PHP extensions, however, need to be bundled with PHP at compile time.

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/php/latest)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/php)

## Adding PHP to your Project

Run `codex add php php83Packages.composer`, or add the following to your `codex.json`:

```json
    "packages": [
        "php@latest",
        "php83Packages.composer@latest
    ]
```

If you want a different version of PHP for your project, you can search for available versions by running `codex search php`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/php)

## Installing PHP Extensions

You can compile additional extensions into PHP by adding them to `packages` in your `codex.json`. Codex will automatically ensure that your extensions are included in PHP at compile time.

For example -- to add the `ds` extension, run `codex add php81Extensions.ds`, or update your packages to include the following:

```json
    "packages": [
        "php@latest",
        "php83Packages.composer",
        "php83Extensions.ds"
    ]
```

## PHP Plugin Details

The PHP Plugin will provide the following configuration when you install a PHP runtime with `codex add`. You can also manually add the PHP plugin by adding `plugin:php` to your `include` list in `codex.json`:

```json
    "include": [
        "plugin:php"
    ]
```

### Services

* php-fpm

Use `codex services start|stop php-fpm` to start PHP-FPM in the background.

### Environment Variables

```bash
PHPFPM_PORT=8082
PHPFPM_ERROR_LOG_FILE={PROJECT_DIR}/.codex/virtenv/php/php-fpm.log
PHPFPM_PID_FILE={PROJECT_DIR}/.codex/virtenv/php/php-fpm.pid
PHPRC={PROJECT_DIR}/codex.d/php/php.ini
```

### Helper Files

* {PROJECT_DIR}/codex.d/php81/php-fpm.conf
* {PROJECT_DIR}/codex.d/php81/php.ini

You can modify these files to configure PHP or your PHP-FPM server
