# codex shellenv

Print shell commands that add Codex packages to your PATH

```bash
codex shellenv [flags]
```

## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
| `-c, --config string` | path to directory containing a codex.json config file |
|  `-e, --env stringToString` |  environment variables to set in the codex environment (default []) |
|  `--env-file string` | path to a file containing environment variables to set in the codex environment |
| `--pure` | If this flag is specified, codex creates an isolated environment inheriting almost no variables from the current environment. A few variables, in particular HOME, USER and DISPLAY, are retained. |
| `-h, --help` | help for shellenv |
| `-q, --quiet` | suppresses logs |


### SEE ALSO

* [codex](codex.md)	 - Instant, easy, predictable development environments
