# Codex

### Instant, easy, and predictable development environments

[![Join Discord](https://img.shields.io/discord/903306922852245526?color=7389D8&label=discord&logo=discord&logoColor=ffffff&cacheSeconds=1800)](https://discord.gg/khulnasoft) ![License: Apache 2.0](https://img.shields.io/github/license/khulnasoft/codex) [![version](https://img.shields.io/github/v/release/khulnasoft/codex?color=green&label=version&sort=semver)](https://github.com/khulnasoft/codex/releases) [![tests](https://github.com/khulnasoft/codex/actions/workflows/cli-post-release.yml/badge.svg)](https://github.com/khulnasoft/codex/actions/workflows/cli-release.yml?branch=main) [![Built with Codex](https://www.khulnasoft/img/codex/shield_galaxy.svg)](https://www.khulnasoft/codex/docs/contributor-quickstart/)

## What is it?

[Codex](https://www.khulnasoft/codex/) is a command-line tool that lets you easily create isolated shells for development. You start by defining the list of packages required by your development environment, and codex uses that definition to create an isolated environment just for your application.

In practice, Codex works similar to a package manager like `yarn` – except the packages it manages are at the operating-system level (the sort of thing you would normally install with `brew` or `apt-get`). With Codex, you can install over [400,000 package versions](https://www.nixhub.io) from the Nix Package Registry

Codex was originally developed by [Khulnasoft](https://www.khulnasoft) and is internally powered by `nix`. 

## Demo

You can try out Codex in your browser using the button below:

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/new)

The example below creates a development environment with `python 2.7` and `go 1.18`, even though those packages are not installed in the underlying machine:

![screen cast](https://user-images.githubusercontent.com/279789/186491771-6b910175-18ec-4c65-92b0-ed1a91bb15ed.svg)

## Installing Codex

Use the following install script to get the latest version of Codex:

```sh
curl -fsSL https://raw.githubusercontent.com/khulnasoft/codex/master/install.sh | bash
```

Read more on the [Codex docs](https://www.khulnasoft/codex/docs/installing_codex/).

## Benefits

### A consistent shell for everyone on the team

Declare the list of tools needed by your project via a `codex.json` file and run `codex shell`. Everyone working on the project gets a shell environment with the exact same version of those tools.

### Try new tools without polluting your laptop

Development environments created by Codex are isolated from everything else in your laptop. Is there a tool you want to try without making a mess? Add it to a Codex shell, and remove it when you don't want it anymore – all while keeping your laptop pristine.

### Don't sacrifice speed

Codex can create isolated environments right on your laptop, without an extra-layer of virtualization slowing your file system or every command. When you're ready to ship, it'll turn it into an equivalent container – but not before.

### Good-bye conflicting versions

Are you working on multiple projects, all of which need different versions of the same binary? Instead of attempting to install conflicting versions of the same binary on your laptop, create an isolated environment for each project, and use whatever version you want for each.

### Take your environment with you

Codex's dev environments are _portable_. We make it possible to declare your
environment exactly once, and use that single definition in several different ways, including:

+ A local shell created through `codex shell`
+ A devcontainer you can use with VSCode
+ A Dockerfile so you can build a production image with the exact same tools you
  used for development.
+ A remote development environment in the cloud that mirrors your local environment.

## Quickstart: Fast, Deterministic Shell

In this quickstart we’ll create a development shell with specific tools installed. These tools will only be available when using this Codex shell, ensuring we don’t pollute your machine.

1. Open a terminal in a new empty folder.

2. Initialize Codex:

   ```bash
   codex init
   ```

   This creates a `codex.json` file in the current directory. You should commit it to source control.

3. Add command-line tools from Nix. For example, to add Python 3.10:

   ```bash
   codex add python@3.10
   ```

   Search for more packages on [Nixhub.io](https://www.nixhub.io)

4. Your `codex.json` file keeps track of the packages you've added, it should now look like this:

   ```json
   {
      "packages": [
         "python@3.10"
       ]
   }
   ```

5. Start a new shell that has these tools installed:

   ```bash
   codex shell
   ```

   You can tell you’re in a Codex shell (and not your regular terminal) because the shell prompt changed.

6. Use your favorite tools.

   In this example we installed Python 3.10, so let’s use it.

   ```bash
   python --version
   ```

7. Your regular tools are also available including environment variables and config settings.

   ```bash
   git config --get user.name
   ```

8. To exit the Codex shell and return to your regular shell:

   ```bash
   exit
   ```

Read more on the [Codex docs Quickstart](https://www.khulnasoft/codex/docs/quickstart/).

## Additional commands

`codex help` - see all commands

See the [CLI Reference](https://www.khulnasoft/codex/docs/cli_reference/codex/) for the full list of commands.

## Join our Developer Community

+ Chat with us by joining the [Khulnasoft Discord Server](https://discord.gg/khulnasoft) – we have a #codex channel dedicated to this project.
+ File bug reports and feature requests using [Github Issues](https://github.com/khulnasoft/codex/issues)
+ Follow us on [Khulnasoft's Twitter](https://twitter.com/khulnasoft_com) for product updates

## Contributing

Codex is an opensource project so contributions are always welcome. Please read [our contributing guide](CONTRIBUTING.md) before submitting pull requests.

[Codex development readme](codex.md)

## Related Work

Thanks to [Nix](https://nixos.org/) for providing isolated shells.

## Translation

+ [Chinese](./docs/translation/README-zh-CN.md)
+ [Korean](./docs/translation/README-ko-KR.md)

## License

This project is proudly open-source under the [Apache 2.0 License](https://github.com/khulnasoft/codex/blob/main/LICENSE)
