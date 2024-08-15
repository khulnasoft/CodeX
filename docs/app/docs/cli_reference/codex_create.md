# codex create

Initialize a directory as a codex project using a template

## Synopsis

Initialize a directory as a codex project. This will create an empty codex.json in the current directory. You can then add packages using `codex add`

```bash
codex create [dir] --template <template> [flags]
```

## List of templates

* [**go**](https://github.com/khulnasoft/codex/tree/main/examples/development/go)
* [**node-npm**](https://github.com/khulnasoft/codex/tree/main/examples/development/nodejs/nodejs-npm/)
* [**node-typescript**](https://github.com/khulnasoft/codex/tree/main/examples/development/nodejs/nodejs-typescript/)
* [**node-yarn**](https://github.com/khulnasoft/codex/tree/main/examples/development/nodejs/nodejs-yarn/)
* [**php**](https://github.com/khulnasoft/codex/tree/main/examples/development/php/)
* [**python-pip**](https://github.com/khulnasoft/codex/tree/main/examples/development/python/pip/)
* [**python-pipenv**](https://github.com/khulnasoft/codex/tree/main/examples/development/python/pipenv/)
* [**python-poetry**](https://github.com/khulnasoft/codex/tree/main/examples/development/python/poetry/)
* [**ruby**](https://github.com/khulnasoft/codex/tree/main/examples/development/ruby/)
* [**rust**](https://github.com/khulnasoft/codex/tree/main/examples/development/rust/)


## Options

<!--Markdown Table of Options  -->
| Option | Description |
| --- | --- |
| `-h, --help` | help for init |
| `-t, --template string` | Template to use for the project.|
| `-q, --quiet` | Quiet mode: Suppresses logs. |

## SEE ALSO

* [codex](./codex.md)	 - Instant, easy, predictable shells and containers

