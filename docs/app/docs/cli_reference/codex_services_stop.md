# codex services stop

Stops a service. If no service is specified, stops all your running services and shuts down process-compose.

```bash
codex services stop [service]... [flags]
```

## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
|  `-e, --env stringToString` |  environment variables to set in the codex environment (default []) |
|  `--env-file string` | path to a file containing environment variables to set in the codex environment |
| `-h, --help` | help for stop |
| `-q, --quiet` | Quiet mode: Suppresses logs. |

## SEE ALSO

* [codex services](codex_services.md)	 - Interact with codex services

