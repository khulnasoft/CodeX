# codex global shellenv

Print shell commands that add global Codex packages to your PATH

- To add the global packages to the PATH of your current shell, run the following command: 
    
    ```bash
    . <(codex global shellenv)
    ```
    
- To add the global packages to the PATH of all new shells, add the following line to your shell's config file (e.g. `~/.bashrc` or `~/.zshrc`):
    
    ```bash
    eval "$(codex global shellenv)"
    ```

## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
| `--pure` | If this flag is specified, codex creates an isolated environment inheriting almost no variables from the current environment. A few variables, in particular HOME, USER and DISPLAY, are retained. |
| `-h, --help` | help for shellenv |
| `-q, --quiet` | suppresses logs |

## SEE ALSO

* [codex global](codex_global.md)	 - Manages global Codex packages
