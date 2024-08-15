# Change Log

All notable changes to the "codex" extension will be documented in this file.

Check [Keep a Changelog](http://keepachangelog.com/) for recommendations on how to structure this file.

## [0.1.5]

- Rebranding changes from khulnasoft.com to khulnasoft.

## [0.1.4]

- Added debug mode in extension settings (only supports logs for "Reopen in Codex Shell environment" feature).

## [0.1.3]

- Added json validation for codex.json files.

## [0.1.2]

- Fixed error handling when using `Reopen in Codex shell` command in Windows and WSL

## [0.1.1]

- Fixed documentation
- Added codex install command
- Added codex update command
- Added codex search command

## [0.1.0]

- Added reopen in codex shell environment feature that allows projects with codex.json
  reopen vscode in codex environment. Note: It requires codex CLI v0.5.5 and above
  installed and in PATH. This feature is in beta. Please report any bugs/issues in [Github](https://github.com/khulnasoft/codex) or our [Discord](https://discord.gg/khulnasoft).

## [0.0.7]

- Fixed a bug for `Open in VSCode` that ensures the directory in which
  we save the VM's ssh key does exist.

## [0.0.6]

- Fixed a small bug connecting to a remote environment.
- Added better error handling and messages if connecting to codex cloud fails.

## [0.0.5]

- Added handling `Open In VSCode` button with `vscode://` style links.
- Added ability for connecting to Codex Cloud workspace.
- Fixed a bug where codex extension would run `codex shell` when opening
a new terminal in vscode even if there was no codex.json present in the workspace.

## [0.0.4]

- Added `Generate a Dockerfile from codex.json` to the command palette
- Changed `Generate Dev Containers config files` command's logic to use codex CLI.

## [0.0.3]

- Small fix for DevContainers and Github CodeSpaces compatibility.

## [0.0.2]

- Added ability to run codex commands from VSCode command palette
- Added VSCode command to generate DevContainer files to run VSCode in local container or Github CodeSpaces.
- Added customization in settings to turn on/off automatically running `codex shell` when a terminal window is opened.

## [0.0.1]

- Initial release
- When VScode Terminal is opened on a codex project, this extension detects `codex.json` and runs `codex shell` so terminal is automatically in codex shell environment.
