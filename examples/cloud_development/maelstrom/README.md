# Maelstrom

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/github.com/khulnasoft/codex-examples?folder=cloud_development/maelstrom)

A Codex for running [Maelstrom](https://github.com/jepsen-io/maelstrom) Tests. Maelstrom is a testing library for toy distributed systems built by @aphyr, useful for learning the basics and principals of building distributed systems

You should also check out the [Fly.io Distributed Systems Challenge](https://fly.io/dist-sys/)

## Prerequisites

If you don't already have [Codex](https://www.khulnasoft/codex/docs/installing_codex/), you can install it by running the following command:

```bash
curl -s https://get.khulnasoft/install.sh | bash
```

You can skip this step if you're running on Codex.sh

## Usage

1. Install Maelstrom by running `codex run install`. This should install Maelstrom 0.2.2 in a `maelstrom` subdirectory

1. cd into the `maelstrom` directory and run `./maelstrom` to verify everything is working

1. You can now follow the docs and run the tests in the Maelstrom Docs + Readme. You can use `glow` from the command line to browse the docs.

This shell includes Ruby 3.10 for running the Ruby Demos. To run demos in other languages, install the appropriate runtimes using `codex add`. For example, to run the Python demos, use `codex add python310`.
