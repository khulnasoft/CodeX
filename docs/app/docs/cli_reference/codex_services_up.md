# codex services up

Starts process-compose and runs all the services in your project. If a list of services is specified in the arguments, only those services will be started.

```bash
codex services up [services]... [flags]
```

This command will launch the process-compose TUI in the foreground. To run process-compose and your services in the background, use the `-b` flag.

Once your services are running, you can manage them using `services start`, `services stop`, and `services restart`.

## Examples
```bash
# Start all services with process compose in the foreground
codex services up

#Start all services with process compose in the background
codex services up -b

# Start only the web service with process compose in the foreground
codex services up web
```

## Options

| Option | Description |
| --- | --- |
| `-b, --background` | Run service in background |
| `-c, --config string` | path to directory containing a codex.json config file |
|  `-e, --env stringToString` |  environment variables to set in the codex environment (default []) |
|  `--env-file string` | path to a file containing environment variables to set in the codex environment |
| `-h, --help` | help for up |
| `--process-compose-file string` | path to process compose file or directory  containing process compose-file.yaml\|yml. Default is directory containing codex.json |
| `-q, --quiet` | Quiet mode: Suppresses logs. |

## SEE ALSO

* [codex services](codex_services.md)	 - Interact with codex services

