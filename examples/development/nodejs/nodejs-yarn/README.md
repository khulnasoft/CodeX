# NodeJS

Most NodeJS Projects will install their dependencies locally using NPM or Yarn, and thus can work with Codex with minimal additional configuration. Per project packages can be managed via NPM or Yarn.

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/node-npm)

## Adding NodeJS to your Shell

`codex add nodejs`, or in your `codex.json`:

```json
  "packages": [
    "nodejs@18"
  ],
```

This will install NodeJS 18, and comes bundled with `npm`. You can find other installable versions of NodeJS by running `codex search nodejs`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/nodejs)

## Adding Yarn as your Package Manager

`codex add yarn`, or in your `codex.json` add:

```json
  "packages": [
    "nodejs@18",
    "yarn@latest"
  ],
```

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
