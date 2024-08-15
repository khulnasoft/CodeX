# codex global

Top level command for managing global packages.

You can use `codex global` to install packages that you want to use across all your local codex projects. For example -- if you usually use `ripgrep` for searching in all your projects, you can use `codex global add ripgrep` to make it available whenever you start a `codex shell` without adding it to each project's `codex.json.` 

You can also use Codex as a global package manager by adding the following line to your shellrc: 

`eval "$(codex global shellenv)"`

For more details, see [Use Codex as your Primary Package Manager](../codex_global.md).

```bash
codex global <subcommand> [flags]
```

## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
| `-c, --config string` | path to directory containing a codex.json config file |
| `-h, --help` | help for generate |
| `-q, --quiet` | Quiet mode: Suppresses logs. |

## Subcommands
* [codex global add](codex_global_add.md)	 - Add a global package to your codex
* [codex global list](codex_global_list.md)	 - List global packages
* [codex global pull](codex_global_pull.md)	 - Pulls a global config from a file or URL.
* [codex global rm](codex_global_rm.md)	 - Remove a global package 
* [codex global shellenv](codex_global_shellenv.md)	 - Print shell commands that add global Codex packages to your PATH

## SEE ALSO

* [codex](codex.md)	 - Instant, easy, predictable development environments
