# codex generate devcontainer

Generate Dockerfile and devcontainer.json files under .devcontainer/ directory

## Synopsis

Generate Dockerfile and devcontainer.json files necessary to run VSCode in remote container environments.

```bash
codex generate devcontainer [flags]
```

### Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
| `-f, --force` | force overwrite on existing files |
| `--root-user` | use `root` as the user for container. Installs nix as single-user mode in Dockerfile |
| `-h, --help` | help for devcontainer |
| `-q, --quiet` | Quiet mode: Suppresses logs. |


### SEE ALSO

* [codex generate](codex_generate.md)	 - Generate supporting files for your project
