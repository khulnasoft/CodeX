# Website

This website is built using [Docusaurus 2](https://docusaurus.io/), a modern static website generator.

You can also test and contribute to our docs online using Codex Cloud!

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/github.com/khulnasoft/codex?folder=docs/app)

## Installation

```bash
cd docs/app     # from the codex root directory
codex shell    # optional, develop inside a codex
yarn install    # run in codex shell
```

### Local Development

```bash
yarn start
```

This command starts a local development server and opens up a browser window. Most changes are reflected live without having to restart the server.

### Build

```bash
yarn build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

### Deployment

When a pull request is opened, it will automatically deploy via CICD to a preview.
When a pull request is merged, it will automatically deploy to production.
Check https://www.khulnasoft/codex/ after merge to see the latest changes.
