# codex generate

Top level command for generating Devcontainers,  Dockerfiles, and other useful files for your Codex Project. 

```bash
codex generate <devcontainer|dockerfile|direnv> [flags]
```

## Options

<!-- Markdown table of options -->
| Option | Description |
| --- | --- |
| `-c, --config string` | path to directory containing a codex.json config file |
| `-h, --help` | help for generate |
| `-q, --quiet` | Quiet mode: Suppresses logs. |

## Subcommands

* [codex generate devcontainer](codex_generate_devcontainer.md)	 - Generate Dockerfile and devcontainer.json files under .devcontainer/ directory
* [codex generate direnv](codex_generate_direnv.md)  - Generate a .envrc file to use with direnv
* [codex generate dockerfile](codex_generate_dockerfile.md)	 - Generate a Dockerfile that replicates codex shell
* [codex generate readme](codex_generate_readme.md)	 -  Generate markdown readme file for your project

## SEE ALSO

* [codex](codex.md)	 - Instant, easy, predictable development environments

