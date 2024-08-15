---
title: Frequently Asked Questions
description: Frequently Asked Questions about Codex
---

This doc contains answers to frequently asked questions about Codex that are not covered elsewhere in our documentation. If you have a question that isn't covered here, feel free to ask us on our [Discord](https://discord.gg/khulnasoft), or [open an issue](https://github.com/khulnasoft/codex/issues) on our GitHub repository.

## How does Codex work?

Codex generates isolated, reproducible development environments using the [Nix package manager](https://nixos.org/). Codex uses Nix to install packages, and then creates an isolated shell environment for your project by symlinking the packages you need into your project directory.

## Where does Codex install my packages?

Codex and Nix install your packages in the read-only Nix store, usually located at `/nix/store`. Codex then creates your environment by symlinking the packages you need into the `.codex` directory in your project.

## How do I clean up unused packages from the Nix Store?

You can use `codex run -- nix store gc --extra-experimental-features nix-command` to automatically clean up packages that are no longer needed for your projects.

## Does Codex require Docker or Containers to work?

No. Since Codex uses Nix to install packages and create isolated environments, Docker is not required. If you want to run your Codex project inside a container, you can generate a Dockerfile or devcontainer.json using the `codex generate` command.

## What versions of Nix are supported by Codex?

Codex requires Nix >= 2.12. If Nix is not present on your machine when you first run Codex, it will automatically try to install the latest supported version for you.

## Can I use Codex with NixOS?

Yes! Codex can be installed on any Linux distribution, including NixOS. You can even install Codex via Nixpkgs. See the [installation guide](./installing_codex.mdx) for more details.

## A package I installed is missing header files or libraries I need for development. Where do I find them?

In order to save space, Codex and Nix only install the required components of packages by default. Development header files and libraries are often installed in a separate output of the package (usually `dev`), which can be installed using the `--output` flag on the `codex add` command. 

For example, the command below will install both the default output `out`, and the `cli` output for the prometheus package: 

```bash
codex add prometheus --outputs=out,cli
```

You can also specify non-default outputs in [flake references](./guides/using_flakes.md): 

```bash
codex add github:NixOS/nixpkgs#prometheus^out,cli
```

## One of my project's packages is taking a long time to install. How can I speed up the installation process?

Packages may take a long time to install if they do not have a binary available in the public Nix Cache. If a prebuilt binary is not available, Nix will built the package from source. 

If prebuilt binaries are not available in the public cache, you may want to use the [Khulnasoft Cache](./cloud/cache/index.md) or the [Khulnasoft Prebuilt Cache](./cloud/cache/prebuilt_cache.md) to cache the binaries you build for future use. Using a package cache can reduce package install by up to 90% compared to building from source.

## I'm trying to build a project, but it says that I'm missing `libstdc++`. How do I install this library in my project?

This message means that your project requires an implementation of the C++ Standard Library installed and linked within your shell. You can add the libstdc++ libraries and object files using `codex add stdenv.cc.cc.lib`. 

## I'm seeing a ``GLIBC_X.XX' not found` error when I try to install my packages, or when I install packages from PyPi/RubyGems/NPM/Cargo/other package manager in my shell

This message usually occurs when using older packages, or when mixing different versions of packages within a single shell. The error tends to occur because each Nix package comes bundled with all of it's dependencies, including a version of the C Standard Library, to ensure reproducibility. If your interpreter (Python/Ruby/Node) or runtime is using an older version of `glibc` than what your other packages expect, they will throw this error. 

There are three ways to work around this issue: 
1. You can update your packages to use a newer version (using `codex add`). This newer version will likely come bundled with a newer version of `glibc`. 
2. You can use `codex update` to get the latest Nix derivation for your package. Newer derivations may come bundled with newer dependencies, including `glibc`
3. If you need to use an exact package version, but you still see this error, you can patch it to use a newer version of glibc using `codex add <package>@<version> --patch-glibc`. This will patch your package to use a newer version of glibc, which should resolve any incompatibility issues you might be seeing. **This patch will only affect packages on Linux.**

## How can I use custom Nix packages or overrides with Codex?

You can add customized packages to your Codex environment using our [Flake support](./guides/using_flakes.md). You can use these flakes to modify or override packages from nixpkgs, or to create your own custom packages.

## Can I use Codex if I use [Fish](https://fishshell.com/)?

Yes. In addition to supporting POSIX compliant shells like Zsh and Bash, Codex also works with Fish. 

Note that `init_hooks` in Codex will be run directly in your host shell, so you may have encounter some compatibility issues if you try to start a shell that uses a POSIX-compatible script in the init_hook.  

## How can I rollback to a previous version of Codex?

You can use any previous version of Codex by setting the `CODEX_USE_VERSION` environment variable. For example, to use version 0.8.0, you can run the following or add it to your shell's rcfile: 

```bash
export CODEX_USE_VERSION=0.8.0
```

You can upgrade to the latest version of Codex by unsetting the variable, and running `codex version update`

## How can I prevent Codex from modifying my prompt while inside a shell?

By default, Codex will prefix your prompt with `(codex)` when inside a `codex shell`. You can disable this behavior by setting this environment variable in your shell's rcfile:

```bash
CODEX_NO_PROMPT=true
```

You can now detect being inside a `codex shell` and change your prompt using the method of your choosing.

## How can I uninstall Codex?

To uninstall Codex:

1. Remove the Codex launcher using `rm /usr/local/bin/codex`
2. Remove the Codex binaries using `rm -rf ~/.cache/codex`
3. Remove your Codex global config using `rm -rf .local/share/codex`

If you want to uninstall Nix, you will need to follow the instructions in the Nix Documentation: https://nixos.org/manual/nix/stable/installation/uninstall.
