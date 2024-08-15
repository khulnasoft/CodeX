# codex services

Interact with Codex services via process-compose

```bash
codex services <ls|restart|start|stop> [flags]
```

## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
| `-c, --config string` | path to directory containing a codex.json config file |
| `-h, --help` | help for services |
| `-q, --quiet` | Quiet mode: Suppresses logs. |

## Subcommands

* [codex services ls](codex_services_ls.md)	 - List available services
* [codex services restart](codex_services_restart.md)	 - Restarts service. If no service is specified, restarts all services
* [codex services start](codex_services_start.md)	 - Starts service. If no service is specified, starts all services
* [codex services stop](codex_services_stop.md)	 - Stops service. If no service is specified, stops all services

## SEE ALSO

* [codex](codex.md)	 - Instant, easy, predictable development environments
