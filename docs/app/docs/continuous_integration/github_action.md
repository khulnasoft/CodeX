---
title: Using Codex in CI/CD with GitHub Actions
---

This guide explains how to use Codex in CI/CD using GitHub Actions. The [codex-install-action](https://github.com/marketplace/actions/codex-installer) will install Codex CLI and any packages + configuration defined in your `codex.json` file. You can then run tasks or scripts within `codex shell` to reproduce your environment.

This GitHub Action also supports caching the packages and dependencies installed in your `codex.json`, which can significantly improve CI build times. 

## Usage

`codex-install-action` is available on the [GitHub Marketplace](https://github.com/marketplace/actions/codex-installer) 

In your project's workflow YAML, add the following step: 

```yaml
- name: Install codex
  uses: khulnasoft/codex-install-action@v0.11.0
```

## Example Workflow

The workflow below shows how to use the action to install Codex, and then run arbitrary commands or [Codex Scripts](../guides/scripts.md) in your shell.

```yaml
name: Testing with codex

on: push

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Install codex
        uses: khulnasoft/codex-install-action@v0.11.0

      - name: Run arbitrary commands
        run: codex run -- echo "done!"

      - name: Run a script called test
        run: codex run test
```

## Configuring the Action

See the [GitHub Marketplace page](https://github.com/marketplace/actions/codex-installer) for the latest configuration settings and an example.

For stability over new features and bug fixes, consider pinning `codex-version`. Remember to update this pinned version when you update your local Codex via `codex version update`.
