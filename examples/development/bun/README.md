# Bun

Bun projects can be run in Codex by adding the Bun runtime + package manager to your project.

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/bun)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/bun)

## Add Bun to your Project

```bash
codex add bun@latest
```

You can see which versions of `bun` are available using: 

```bash
codex search bun
```

To update bun to the latest version: 

```bash
codex update bun
```

## Scripts

To install dependencies:

```bash
codex run bun install
```

To start + watch your project:

```bash
codex run dev
```

This project was created using `bun init` in bun v1.0.33. [Bun](https://bun.sh) is a fast all-in-one JavaScript runtime.
