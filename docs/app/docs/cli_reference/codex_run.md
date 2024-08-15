# codex run

Starts a new interactive shell and runs your target script in it. The shell will exit once your target script is completed or when it is terminated via CTRL-C. Scripts can be defined in your `codex.json`.

You can also run arbitrary commands in your codex shell by passing them as arguments to `codex run`. For example:

```bash
  codex run echo "Hello World"
```
Will print `Hello World` to the console from within your codex shell.

For more details, read our [scripts guide](../guides/scripts.md)

```bash
  codex run <script | command> [flags]
```


## Examples

```bash
# Run a command directly:
  codex add cowsay
  codex run cowsay hello
# Run a command that takes flags:
  codex run cowsay -d hello
# Pass flags to codex while running a command.
# All `codex run` flags must be passed before the command
  codex run -q cowsay -d hello

#Run a script (defined as `"moo": "cowsay moo"`) in your codex.json:
  codex run moo
```

## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
| `-c, --config string` | path to directory containing a codex.json config file |
| `-e, --env stringToString` |  environment variables to set in the codex environment (default []) |
| `--env-file string` | path to a file containing environment variables to set in the codex environment |
| `-h, --help` | help for run |
| `-q, --quiet` | Quiet mode: Suppresses logs. |



## SEE ALSO

* [codex](./codex.md)	 - Instant, easy, predictable shells and containers

