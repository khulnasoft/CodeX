---
title: Pushing and Pulling Packages to the Cache
sidebar_position: 3
---
## Pulling Packages from the Cache

Once you have authenticated, Codex will automatically configure your cache for you. You can also manually configure the cache by running:

```bash
codex cache configure 
```

Once configured, Codex will attempt to use the cache whenever you run a command that prompts Codex to install a package in your project. When installing a package, Codex will check for a cached binary in the following locations:

1. Your local Nix Store `/nix/store`
2. The Khulnasoft Cache
3. The Public Nix Cache ([cache.nixos.org](https://cache.nixos.org))

If the package is not available in those locations, then it will build the package from source. 

## Pushing to the Cache

You can push custom packages and project closures to your Khulnasoft Cache directly from the Codex CLI. Push access is currently only available for authenticated users with Admin permissions.

### **Pushing a Codex Project**

You can push the entire closure of a Codex project to the Khulnasoft Cache by navigating to your project root and running

```bash
codex cache upload
```

### Pushing a specific package

You can also push a single package by passing a flake reference to the Codex CLI. 

For example, to push a custom `mongodb` package from a custom flake.nix on your machine, you can run:

```bash
codex cache upload path:./path/to/flake.nix#mongodb
```

You can also cache packages from Github or from Nixpkgs by passing the appropriate Flake reference. This can be useful for caching build artifacts if a package does not exist in the public Nix cache, or if it requires you to build it from source:

```bash
# Cache an installable hosted on Github (process-compose)
codex cache upload github:F1bonnac1/process-compose

# Cache a package from nixpkgs
codex cache upload nixpkgs#mongodb
```