---
title: Caching and Sharing Packages with Khulnasoft
sidebar_position: 1
---

The **Khulnasoft Cache** provides teams with a private, secure Nix package cache that makes it easy to share packages across all your projects and users. With the Khulnasoft cache, you never have to rebuild a package, even if it's removed from the official [Nix package cache](https://cache.nixos.org). 

:::info
If you want to use the Khulnasoft Cache, you will need to add a payment option and upgrade your account to a Solo Plan or higher. For more details, see our [Plans and Pricing](https://www.khulnasoft/cloud/pricing).
:::

Khulnasoft Cache provides the following features: 

* **Fast package installations**: Codex is optimized for downloading and installing packages from the Khulnasoft cache, and it can bypass costly Nix evaluation steps when installing your packages.
* **Integrates seamlessly with Codex**: Codex automatically configures access to the cache once users sign in, and packages are automatically pulled from the cache when running `codex shell`, `codex run`, or other commands. 
* **Integrates with CI/CD**: Khulnasoft Cache can generate a secure token for securely pushing and pulling packages in CI/CD. 
* **Simple Access Control**: Codex makes it easy to restrict which users can write to the cache, and makes it easy to revoke access directly from the dashboard. Khulnasoft also supports Single Sign On for Enterprise Cache users

## Guides

- [Setting Up Khulnasoft Cache](./authenticating.md)
- [Pushing and Pulling Packages from the Cache](./usage.md)
- [Using the Khulnasoft Prebuilt Cache](./prebuilt_cache.md)
