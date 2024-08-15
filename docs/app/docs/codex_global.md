---
title: Use Codex as your Primary Package Manager
description: Install packages and tools system wide with Codex Global
---

In addition to managing isolated development environments, you can use Codex as a general package manager. Codex Global allows you to add packages to a global `codex.json.` This is useful for installing a standard set of tools you want to use across multiple Codex Projects.

For example — if you use ripgrep as your preferred search tool, you can add it to your global Codex profile with `codex global add ripgrep`. Now whenever you start a Codex shell, you will have ripgrep available, even if it's not in the project's codex.json.

<figure>

![Installing ripgrep using `codex global add ripgrep](../static/img/codex_global.svg)

<figcaption>Installing Packages with Codex Global</figcaption>
</figure>

You can also use `codex global` to replace package managers like `brew` and `apt` by adding the global profile to your path. Because Codex uses Nix to install packages, you can sync your global config to install the same packages on any machine.

Codex saves your global config in a `codex.json` file in your home directory. This file can be shared with other users or checked into source control to synchronize it across machines.


## Adding and Managing Global Packages

You can install a package using `codex global add [<package>]`, where the package names should be a list of [Nix Packages](https://search.nixos.org/packages) you want to install.

For example, if we wanted to install ripgrep, vim, and git to our global profile, we could run:

```bash
codex global add ripgrep vim git

# Output:
ripgrep is now installed
vim is now installed
git is now installed
```

Once installed, the packages will be available whenever you start a Codex Shell, even if it's not included in the project's `codex.json`.

To view a full list of global packages, you can run `codex global list`:

```bash
codex global list

# Output:
* ripgrep
* vim
* git
```

To remove a global package, use:

```bash
codex global rm ripgrep

# Output:
removing 'github:NixOS/nixpkgs/ripgrep'
```

## Using Fleek with Codex Global

[Fleek](https://getfleek.dev/) provides a nicely tuned set of packages and configuration for common tools that is compatible with Codex Global. Configurations are provided at different [levels of bling](https://getfleek.dev/docs/bling), with higher levels adding more packages and opinionated configuration.

To install a Fleek profile, you can use `codex global pull <fleek-url>`, where the Fleek URL indicates the profile you want to install. For example, to install the `high` bling profile, you can run:

```bash
codex global pull https://codex.getfleek.dev/high
```

Fleek profiles also provide a few convenience scripts to automate setting up your profile. You can view the full list of scripts using `codex global run` with no arguments

For more information, see the [Fleek for Codex Docs](https://getfleek.dev/docs/codex)

## Using Global Packages in your Host Shell

If you want to make your global packages available in your host shell, you can add them to your shell PATH. Running `codex global shellenv` will print the command necessary to source the packages.

### Add Global Packages to your Current Host Shell
To temporarily add the global packages to your current shell, run:

```bash
. <(codex global shellenv --init-hook)
```

You can also add a hook to your shell's config to make them available whenever you launch your shell:

### Bash

Add the following command to your `~/.bashrc` file:

```bash
eval "$(codex global shellenv --init-hook)"
```

Make sure to add this hook before any other hooks that use your global packages.

### Zsh
Add the following command to your `~/.zshrc` file:

```bash
eval "$(codex global shellenv --init-hook)"
```

### Fish

Add the following command to your `~/.config/fish/config.fish` file:

```bash
codex global shellenv --init-hook | source
```

## Sharing Your Global Config with Git

You can use Git to synchronize your `codex global` config across multiple machines using `codex global push <remote>` and `codex global pull <remote>`.

Your global `codex.json` and any other files in the Git remote will be stored in `$XDG_DATA_HOME/codex/global/default`. If `$XDG_DATA_HOME` is not set, it will default to `~/.local/share/codex/global/default`. You can view the current global directory by running `codex global path`.

## Next Steps

### Learn more about Codex

* **[Getting Started](quickstart.mdx):** Learn how to install Codex and create your first Codex Shell.
* **[Codex Scripts](guides/scripts.md):** Automate setup steps and configuration for your shell using Codex Scripts.
* **[Configuration Guide](configuration.md):** Learn how to configure your shell and dev environment with `codex.json`.
* **[Browse Examples](https://github.com/khulnasoft/codex-examples):** You can see how to create a development environment for your favorite tools or languages by browsing the Codex Examples repo.
* **[Using Flakes with Codex](guides/using_flakes.md):** Learn how to install packages from Nix Flakes.

### Use Codex with your IDE

* **[Direnv Integration](ide_configuration/direnv.md):** Codex can integrate with [direnv](https://direnv.net/) to automatically activate your shell and packages when you navigate to your project.
* **[Codex for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=khulnasoft.codex):** Install our VS Code extension to speed up common Codex workflows or to use Codex in a devcontainer.

### Get Involved

* **[Join our Discord Community](https://discord.gg/khulnasoft):** Chat with the development team and our growing community of Codex users.
* **[Visit us on Github](https://github.com/khulnasoft/codex):** File issues and provide feedback, or even open a PR to contribute to Codex or our Docs.
