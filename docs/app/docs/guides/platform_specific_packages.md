---
title: Installing Platform Specific Packages
---

At times, you may need to install a package or library that is only available on a specific platform. For example, you may want to install a package that is only available on Linux, while still using the same Codex configuration on your Mac.

Codex allows you to specify which platforms a package should be installed on using the `--platform` and `--exclude-platform` flags. When a package is added using these flags, it will be added to your `codex.json`, but will only be installed when you run Codex on a matching platform.

:::info

Specifying platforms for packages will alter your `codex.json` in a way that is only compatible with **Codex 0.5.12** and newer.

If you encounter errors trying to run a Codex project with platform-specific packages, you may need to run `codex version update`
:::

## Installing Platform Specific Packages

To avoid build or installation errors, you can tell Codex to only install a package on specific platforms using the `--platform` flag when you run `codex add`.

For example, to install the `busybox` package only on Linux platforms, you can run:

```bash
codex add busybox --platform x86_64-linux,aarch64-linux
```

This will add busybox to your `codex.json`, but will only install it when use codex on a Linux machine. The packages section in your config will look like the following

```json
{
    "packages": {
        "busybox": {
            "version": "latest",
            "platforms": ["x86_64-linux", "aarch64-linux"]
        }
    }
}
```

## Excluding a Package from Specific Platforms

You can also tell Codex to exclude a package from a specific platform using the `--exclude-platform` flag. For example, to avoid installing `ripgrep` on an ARM-based Mac, you can run:


```bash
codex add ripgrep --exclude-platform aarch64-darwin
```

This will add ripgrep to your `codex.json`, but will not install it when use codex on an ARM-based Mac. The packages section in your config will look like the following:

```json
{
    "packages": {
        "ripgrep": {
            "version": "latest",
            "excluded_platforms": ["aarch64-darwin"]
        }
    }
}
```

## Supported Platforms

Valid Platforms include:

* `aarch64-darwin`
* `aarch64-linux`
* `x86_64-darwin`
* `x86_64-linux`

The platforms below are also supported, but will build packages from source

* `i686-linux`
* `armv7l-linux`