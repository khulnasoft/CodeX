# codex generate direnv

Top level command for generating the .envrc file for your Codex Project. This can be used with [direnv](../ide_configuration/direnv.md) to automatically start your shell when you cd into your codex directory

```bash
codex generate direnv [flags]
```

## Options

<!-- Markdown table of options -->
| Option | Description |
| --- | --- |
| `-c, --config string` | path to directory containing a codex.json config file |
|  `-e, --env stringToString` |  environment variables to set in the codex environment (default []) |
|  `--env-file string` | path to a file containing environment variables to set in the codex environment. If the file does not exist, then this parameter is ignored |
| `-h, --help` | help for direnv |
| `-q, --quiet` | Quiet mode: Suppresses logs. |

## SEE ALSO

* [codex generate](codex_generate.md)	 - Generate supporting files for your project
