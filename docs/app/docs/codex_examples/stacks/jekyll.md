---
title: Jekyll
---

This example demonstrates how to create and run a Jekyll blog in Codex. It makes use of the Ruby Plugin to configure and setup your project. 

[Example Repo](https://github.com/khulnasoft/codex/tree/main/examples/stacks/jekyll)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/jekyll)

Inspired by https://litchipi.github.io/nix/2023/01/12/build-jekyll-blog-with-nix.html 

## How to Run

1. Install [Codex](https://www.khulnasoft/codex/docs/installing_codex/)
2. Run `codex shell` to install your packages and run the init hook
3. In the root directory, run `codex run generate` to install and package the project with bundler
4. In the root directory, run `codex run server` to start the server. You can access the Jekyll example at `localhost:4000`

## How to Recreate this Example 

1. Install [Codex](https://www.khulnasoft/codex/docs/installing_codex/)
1. In a new directory, run `codex init` to create an empty config
1. Run `codex add ruby_3_1 bundler` to add Ruby and Bundler to your packages
1. Add `"gem install jekyll --version \"~> 3.9.2\""` to your init hook. This will install the Jekyll gem in your local project.
1. Start your `codex shell`, then run `jekyll new myblog` to create the starter project.
1. From here you can install the project using Bundler, and start the server using `jekyll serve`. See the scripts in this example for more details.
