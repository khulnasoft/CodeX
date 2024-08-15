# codex services restart

Restarts service. If no service is specified, restarts all services and process-compose.

```bash
codex services restart [service]... [flags]
```

:::info
  Note: We recommend using `codex services up` if you are starting all your services and process-compose. This command lets you specify your process-compose file and whether to run process-compose in the foreground or background.
:::

## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
|  `-e, --env stringToString` |  environment variables to set in the codex environment (default []) |
|  `--env-file string` | path to a file containing environment variables to set in the codex environment |
| `-h, --help` | help for restart |
| `-q, --quiet` | Quiet mode: Suppresses logs. |

## SEE ALSO

* [codex services](codex_services.md)	 - Interact with codex services

