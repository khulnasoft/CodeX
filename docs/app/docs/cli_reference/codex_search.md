# codex search

Search for Nix packages

## Synopsis

`codex search` will return a list of packages and versions that match your search query.

You can add a package to your project using `codex add <package>`.

Too add a specific version, use `codex add <package>@<version>`.

```bash
codex search <pkg> [flags]
```

## Example

```bash
$ codex search ripgrep

Warning: Search is experimental and may not work as expected.

Found 8+ results for "ripgrep":

* ripgrep (13.0.0, 12.1.1, 12.0.1)
* ripgrep-all (0.9.6, 0.9.5)

# To add ripgrep 12.1.1 to your project:

$ codex add ripgrep@12.1.1
```

## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
| `-h, --help` | help for shell |
| `-q, --quiet` | Quiet mode: Suppresses logs. |

## SEE ALSO

* [codex](./codex.md)	 - Instant, easy, predictable shells and containers

