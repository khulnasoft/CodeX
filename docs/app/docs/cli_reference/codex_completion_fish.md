# codex completion fish

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

```bash
codex completion fish | source
```

To load completions for every new session, execute once:

```bash
codex completion fish > ~/.config/fish/completions/codex.fish
```

You will need to start a new shell for this setup to take effect.


```bash
codex completion fish [flags]
```

## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
| `-h, --help` | help for fish |
| `--no-descriptions` | disable completion descriptions |
| `-q, --quiet` | suppresses logs |

## SEE ALSO

* [codex completion](codex_completion.md)	 - Generate the autocompletion script for the specified shell

