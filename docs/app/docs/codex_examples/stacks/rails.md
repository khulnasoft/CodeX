---
title: Ruby on Rails
description: An example of a Nix powered Ruby on Rails application with Codex
---

This example demonstrates how to setup a simple Rails application. It makes use of the Ruby Plugin, and installs SQLite to use as a database.

[Example Repo](https://github.com/khulnasoft/codex/tree/main/examples/stacks/rails)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/rails)

## How To Run

Run `codex shell` to install rails and prepare the project.

Once the shell starts, you can start the rails app by running:

```bash
cd blog
bin/rails server
```

## How to Recreate this Example

1. Create a new Codex project with `codex init`
2. Add the packages using

   ```bash
   codex add ruby_3_1 bundler nodejs-19_x yarn sqlite
   ```

3. Start a codex shell, and install the rails CLI with `gem install rails`
4. Create your Rails app by running the following in your Codex Shell

   ```bash
   rails new blog
   ```
