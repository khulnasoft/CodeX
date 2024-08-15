---
title: Rust
---

The easiest way to manage Rust with Codex is to install `rustup`, and then configure the channel you wish to install via Codex's `init_hook`. You can also use the `init_hook` to configure `rustup` to install the Rust toolchain locally.

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/rust)

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/open/templates/rust)

```json
{
    "packages": [
        "rustup@latest",
        "libiconv@latest"
    ],
    "shell": {
        "init_hook": [
            "projectDir=$(dirname $(readlink -f \"$0\"))",
            "rustupHomeDir=\"$projectDir\"/.rustup",
            "mkdir -p $rustupHomeDir",
            "export RUSTUP_HOME=$rustupHomeDir",
            "export LIBRARY_PATH=$LIBRARY_PATH:\"$projectDir/nix/profile/default/lib\"",
            "rustup default stable",
            "cargo fetch"
        ],
        "scripts": {
            "test": "cargo test -- --show-output",
            "start" : "cargo run",
            "build-docs": "cargo doc"
        }
    }
}
```

To pin a specific version of Rust with Rustup, you can add a [rust-toolchain.toml](https://rust-lang.github.io/rustup/overrides.html#the-toolchain-file) and check it into source control.
