# Codex Quickstart

This shell includes a basic `codex.json` with a few useful packages installed, and an example init_hook and script

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/github.com/khulnasoft/codex-examples?folder=tutorial)

## Adding New Packages

Run `codex add <package>` to add a new package. Remove it with `codex rm <package>`.

For example: install Python 3.10 by running:

```bash
codex add python310
```

Codex can install over 80,000 packages via the Nix Package Manager. Search for packages at [https://search.nixos.org/packages](https://search.nixos.org/packages)

## Running Codex Scripts

You can add new scripts by editing the `codex.json` file

You can run scripts using `codex run <script>`

For example: you can replay this help text with:

```bash
codex run readme
```

## Next Steps

* Checkout our Docs at [https://www.khulnasoft/codex/docs](https://www.khulnasoft/codex/docs)
* Try out an Example Project at [https://www.khulnasoft/codex/docs/codex-examples](https://www.khulnasoft/codex/docs/codex-examples)
* Report Issues at [https://github.com/khulnasoft/codex/issues/new/choose](https://github.com/khulnasoft/codex/issues/new/choose)
