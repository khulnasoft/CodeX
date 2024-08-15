---
title: "Starting a Dev Environment with Codex"
sidebar_position: 3
---
## Background

Codex is a command-line tool that lets you easily create reproducible, reliable dev environments. You start by defining the list of packages required by your development environment, and codex uses that definition to create an isolated environment just for your application. Developers can start a dev environment for their project by running `codex shell`.

To learn more about how Codex works, you can read our [introduction](index.md)

This quickstart shows you how to install Codex, and use it to start a development environment for a project that is configured to use Codex via `codex.json`


## Install Codex

Use the following install script to get the latest version of Codex:

```bash
curl -fsSL https://raw.githubusercontent.com/khulnasoft/codex/master/install.sh | bash
```

Codex requires the [Nix Package Manager](https://nixos.org/download.html). If Nix is not detected on your machine when running a command, Codex will automatically install it for you with the default settings for your OS. Don't worry: You can use Codex without needing to learn the Nix Language.

## Start your development shell

1. Open a terminal in the project. The project should contain a `codex.json` that specifies how to create your development environment

1. Start a codex shell for your project:

    To get started, all we have to do is run:
    ```bash
    codex shell
    ```

    **Output:**
    ```bash
    Installing nix packages. This may take a while... done.
    Starting a codex shell...
    (codex) $
    ```

    :::info
    The first time you run `codex shell` may take a while to complete due to Codex downloading prerequisites and package catalogs required by Nix. This delay is a one-time cost, and future invocations and package additions should resolve much faster.
    :::

1. Use the packages provided in your development environment

    The packages listed in your project's `codex.json` should now be available for you to use. For example, if the project's `codex.json` contains `python@3.10`, you should now have `python` in your path:

    ```bash
    $ python --version
    Python 3.10.9
    ```

1. Your host environment's packages and tools are also available, including environment variables and config settings.

    ```bash
    git config --get user.name
    ```

1. You can search for additional packages using `codex search <pkg>`. You can then add them to your Codex shell by running `codex add [pkgs]`

1. To exit the Codex shell and return to your regular shell:

    ```bash
    exit
    ```

## Next Steps

### Learn more about Codex
* **[Codex Global](codex_global.md):** Learn how to use the codex as a global package manager
* **[Codex Scripts](guides/scripts.md):** Automate setup steps and configuration for your shell using Codex Scripts.
* **[Configuration Guide](configuration.md):** Learn how to configure your shell and dev environment with `codex.json`.
* **[Browse Examples](https://github.com/khulnasoft/codex-examples):** You can see how to create a development environment for your favorite tools or languages by browsing the Codex Examples repo.

### Use Codex with your IDE
* **[Direnv Integration](ide_configuration/direnv.md):** Codex can integrate with [direnv](https://direnv.net/) to automatically activate your shell and packages when you navigate to your project.
* **[Codex for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=khulnasoft.codex):** Install our VS Code extension to speed up common Codex workflows or to use Codex in a devcontainer.

### Boost your dev environment with Khulnasoft Cloud

* **[Khulnasoft Secrets](./cloud/secrets/index.md):** Securely store and access your secrets and environment variables in your Codex projects.
* **[Khulnasoft Deploys](./cloud/deploys/index.md):** Deploy your Codex projects as autoscaling services with a single command.
* **[Khulnasoft Cache](./cloud/cache/index.md):** Share and cache packages across all your Codex projects and environments.
* **[Khulnasoft Prebuilt Cache](./cloud/cache/prebuilt_cache.md):** Use the Khulnasoft Public Cache to speed up your Codex builds and share packages with the community.

### Get Involved
* **[Join our Discord Community](https://discord.gg/khulnasoft):** Chat with the development team and our growing community of Codex users.
* **[Visit us on Github](https://github.com/khulnasoft/codex):** File issues and provide feedback, or even open a PR to contribute to Codex or our Docs.
