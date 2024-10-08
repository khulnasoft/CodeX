# codex shell

Start a new shell or run a command with access to your packages

## Synopsis

Start a new shell or run a command with access to your packages.   
If the --config flag is set, the shell will be started using the codex.json found in the --config flag directory.   
If --config isn't set, then codex recursively searches the current directory and its parents.

```bash
codex shell [flags]
```

## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
|  `-c, --config string`|  path to directory containing a codex.json config file |
|  `-e, --env stringToString` |  environment variables to set in the codex environment (default []) |
|  `--env-file string` | path to a file containing environment variables to set in the codex environment |
|  `--environment string` | environment to use, when supported (e.g.secrets support dev, prod, preview.) (default "dev") |
| `--print-env` | Print a script to setup a codex shell environment |
| `--pure` | If this flag is specified, codex creates an isolated shell inheriting almost no variables from the current environment. A few variables, in particular HOME, USER and DISPLAY, are retained. |
| `-h, --help` | help for shell |
| `-q, --quiet` | Quiet mode: Suppresses logs. |

## SEE ALSO

* [codex](./codex.md)	 - Instant, easy, predictable shells and containers

