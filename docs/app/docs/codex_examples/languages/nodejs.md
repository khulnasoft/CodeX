---
title: NodeJS
---

Most NodeJS Projects will install their dependencies locally using NPM or Yarn, and thus can work with Codex with minimal additional configuration. Per project packages can be managed via NPM or Yarn.

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/nodejs)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/node-npm)

## Adding NodeJS to your Shell

`codex add nodejs`, or in your `codex.json`:

```json
  "packages": [
    "nodejs@18"
  ],
```

This will install NodeJS 18, and comes bundled with `npm`. You can find other installable versions of NodeJS by running `codex search nodejs`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/nodejs)

## Adding Yarn, NPM, or pnpm as your Node Package Manager

We recommend using [Corepack](https://github.com/nodejs/corepack/) to install and manage your Node Package Manager in Codex. Corepack comes bundled with all recent Nodejs versions, and you can tell Codex to automatically configure Corepack using a built-in plugin. When enabled, corepack binaries will be installed in your project's `.codex` directory, and automatically added to your path.

To enable Corepack, set `CODEX_COREPACK_ENABLED` to `true` in your `codex.json`:

```json
{
  "packages": ["nodejs@18"],
  "env": {
    "CODEX_COREPACK_ENABLED": "true"
  }
}
```

To disable Corepack, remove the `CODEX_COREPACK_ENABLED` variable from your codex.json

### Yarn

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/nodejs/nodejs-yarn)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/node-yarn)

### pnpm

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/nodejs/nodejs-pnpm)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/node-pnpm)

## Installing Global Packages

In some situations, you may want to install packages using `npm install --global`. This will fail in Codex since the Nix Store is immutable.

You can instead install these global packages by adding them to the list of packages in your `codex.json`. For example: to add `yalc` and `pm2`:

```json
{
  "packages": [
    "nodejs@18",
    "nodePackages.yalc@latest",
    "nodePackages.pm2@latest"
  ]
}
```
