---
title: direnv
---


## direnv
___
[direnv](https://direnv.net) is an open source environment management tool that allows setting unique environment variables per directory in your file system. This guide covers how to configure direnv to seamlessly work with a codex project.

### Prerequisites
* Install direnv and hook it to your shell. Follow [this guide](https://direnv.net/#basic-installation) if you haven't done it.

### Setting up Codex Shell and direnv

#### New Project

If you have direnv installed, Codex will generate an .envrc file when you run `codex generate direnv` and enables it by running `direnv allow` in the background:

```bash
➜  codex generate direnv
Success: generated .envrc file
Success: ran `direnv allow`
direnv: loading ~/src/codex/docs/.envrc
direnv: using codex
```

This will generate a `.envrc` file in your project directory that contains `codex.json`. Run `direnv allow` to activate your shell upon directory navigation. Run `direnv revoke` to stop. Changes to `codex.json` automatically trigger direnv to reset the environment. The generated `.envrc` file doesn't need any further configuration. Just having the generated file along with an installed direnv and Codex is enough to make direnv integrate with Codex.


#### Existing Project

For an existing project, you can add a `.envrc` file by running `codex generate direnv`:

```bash
➜  codex generate direnv
Success: generated .envrc file
Success: ran `direnv allow`
direnv: loading ~/src/codex/docs/.envrc
direnv: using codex
```

The generated `.envrc` file doesn't need any further configuration. Just having the generated file along with installed direnv and Codex, is enough to make direnv integration with Codex work.

#### Adding Custom Env Variables or Env Files to your Direnv Config

In some cases, you may want to override certain environment variables in your Codex config when running it locally. You can add custom environment variables from the command line or from a file using the `--env` and `--env-file` flags.

If you would like to add custom environment variables to your direnv config, you can do so by passing the `--env` flag to `codex generate direnv`. This flag takes a comma-separated list of key-value pairs, where the key is the name of the environment variable and the value is the value of the environment variable. For example, if you wanted to add a `MY_CUSTOM_ENV_VAR` environment variable with a value of `my-custom-value`, you would run the following command:

```bash
codex generate direnv --env MY_CUSTOM_ENV_VAR=my-value
```

The resulting .envrc will have the following:

```bash
# Automatically sets up your codex environment whenever you cd into this
# directory via our direnv integration:

eval "$(codex generate direnv --print-envrc --env MY_CUSTOM_ENV_VAR=my-value)"

# check out https://www.khulnasoft/codex/docs/ide_configuration/direnv/
# for more details
```

You can also tell direnv to read environment variables from a custom `.env` file by passing the `--env-file` flag to `codex generate direnv`. This flag takes a path to a file containing environment variables to set in the codex environment. If the file does not exist, then this parameter is ignored. For example, if you wanted to add a `.env.codex` file located in your project root, you would run the following command:

```bash
codex generate direnv --env-file .env.codex
```

The resulting .envrc will have the following:

```bash
# Automatically sets up your codex environment whenever you cd into this
# directory via our direnv integration:

eval "$(codex generate direnv --print-envrc --env-file .env.codex)"

# check out https://www.khulnasoft/codex/docs/ide_configuration/direnv/
# for more details
```

Note that if Codex cannot find the env file provided to the flag, it will ignore the flag and load your Codex shell environment as normal

### Global settings for direnv

Note that every time changes are made to `codex.json` via `codex add ...`, `codex rm ...` or directly editing the file, requires `direnv allow` to run so that `direnv` can setup the new changes.

Alternatively, a project directory can be whitelisted so that changes will be automatically picked up by `direnv`. This is done by adding following snippet to direnv config file typically at `~/.config/direnv/direnv.toml`. You can create the file and directory if it doesn't exist.

```toml
[whitelist]
prefix = [ "/absolute/path/to/project" ]

```

### Direnv Limitations

Direnv works by creating a sub-shell using your `.envrc` file, your `codex.json`, and other direnv related files, and then exporting the diff in environment variables into your current shell. This imposes some limitations on what it can load into your shell: 

1. Direnv cannot load shell aliases or shell functions that are sourced in your project's `init_hook`. If you want to use direnv and also configure custom aliases, we recommend using [Codex Scripts](../guides/scripts.md). 
2. Direnv does not allow modifications to the $PS1 environment variable. This means `init_hooks` that modify your prompt will not work as expected. For more information, see the [direnv wiki](https://github.com/direnv/direnv/wiki/PS1)

Note that sourcing aliases, functions, and `$PS1` should work as expected when using `codex shell`, `codex run`, and `codex services`

### VSCode setup with direnv

To seamlessly integrate VSCode with a direnv environment, follow these steps:

1. Open a terminal window and activate direnv with `direnv allow`.
2. Launch VSCode from the same terminal window using the command `code .` This ensures that VSCode inherits the direnv environment.

Alternatively, you can use the [direnv VSCode extension](https://marketplace.visualstudio.com/items?itemName=mkhl.direnv) if your VSCode workspace has a .envrc file.

If this guide is missing something, feel free to contribute by opening a [pull request](https://github.com/khulnasoft/codex/pulls) in Github.
