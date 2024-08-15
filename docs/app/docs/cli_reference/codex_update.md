# codex update

Updates packages within your project to the latest available version.

## Synopsis

If you provide this command with a list of packages, it will update those packages to the latest available version based on the version tag provided.

For example: if your project has `python@3.11` in your package list, running `codex update` will update your project to the latest patch version of `python 3.11`.

If no packages are provided, this command will update all the versioned packages in your project to the latest acceptable version.

```bash
codex update [pkg]... [flags]
```


## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
| `-c, --config` | Path to codex config file. |
| `-h, --help` | help for shell |
| `-q, --quiet` | Quiet mode: Suppresses logs. |

## SEE ALSO

* [codex](./codex.md)	 - Instant, easy, predictable shells and containers

