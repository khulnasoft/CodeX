# codex VSCode Extension

This is the official VSCode extension for [codex](https://github.com/khulnasoft/codex) open source project by [khulnasoft](https://www.khulnasoft)

## Features

### Open In VSCode button

If a Codex Cloud instance (from [codex.sh](https://codex.sh)) has an `Open In Desktop` button, this extension will make VSCode to be able to connect its workspace to the instance.

### Auto Shell on a codex project

When VSCode Terminal is opened on a codex project, this extension detects `codex.json` and runs `codex shell` so terminal is automatically in codex shell environment. Can be turned off in settings.

### Reopen in Codex shell environment

If the opened workspace in VSCode has a codex.json file, from command palette, invoking the codex command `Reopen in Codex shell environment` will do the following:

1. Installs codex packages if missing.
2. Update workspace settings for MacOS to create terminals without creating a login shell [learn more](https://code.visualstudio.com/docs/terminal/profiles#_why-are-there-duplicate-paths-in-the-terminals-path-environment-variable-andor-why-are-they-reversed-on-macos)
3. Interact with Codex CLI to setup a codex shell.
4. Close current VSCode window and reopen it in a codex shell environment as if VSCode was opened from a codex shell terminal.

NOTE: Requires codex CLI v0.5.5 and above
  installed and in PATH. This feature is in beta. Please report any bugs/issues in [Github](https://github.com/khulnasoft/codex) or our [Discord](https://discord.gg/khulnasoft).

### Run codex commands from command palette

`cmd/ctrl + shift + p` opens vscode's command palette. Typing codex filters all available commands codex extension can run. Those commands are:

- **Init:** Creates a codex.json file
- **Add:** adds a package to codex.json
- **Remove:** Removes a package from codex.json
- **Shell:** Opens a terminal and runs codex shell
- **Run:** Runs a script from codex.json if specified
- **Install** Install packages specified in codex.json
- **Update** Update packages specified in codex.json
- **Search** Search for packages to add to your codex project
- **Generate DevContainer files:** Generates devcontainer.json & Dockerfile inside .devcontainers directory. This allows for running vscode in a container or GitHub Codespaces.
- **Generate a Dockerfile from codex.json:** Generates a Dockerfile a project's root directory. This allows for running the codex project in a container.
- **Reopen in Codex shell environment:** Allows projects with codex.json
  reopen VSCode in codex environment. Note: It requires codex CLI v0.5.5 and above
  installed and in PATH.

### JSON validation when writing a codex.json file

No need to take any action for this feature. When writing a codex.json, if this extension is installed, it will validate and highlight any disallowed fields or values on a codex.json file.

---

### Debug Mode

Enabling debug mode in extension settings will create a seqience of logs in the file `.codex/extension.log`. This feature only tracks the logs for `"Codex: Reopen in Codex Shell environment"` feature.

## Following extension guidelines

Ensure that you've read through the extensions guidelines and follow the best practices for creating your extension.

- [Extension Guidelines](https://code.visualstudio.com/api/references/extension-guidelines)

## Publishing

Steps:

1. Bump the version in `package.json`, and add notes to `CHANGELOG.md`. Sample PR: #951.
2. Manually trigger the [`vscode-ext-release` in Github Actions](https://github.com/khulnasoft/codex/actions/workflows/vscode-ext-release.yaml).
