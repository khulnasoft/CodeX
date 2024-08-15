# Jekyll Example

[![Built with Codex](https://www.khulnasoft/img/codex/shield_moon.svg)](https://www.khulnasoft/codex/docs/contributor-quickstart/)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/jekyll)

Inspired by [This Example](https://litchipi.github.io/nix/2023/01/12/build-jekyll-blog-with-nix.html)

## How to Use

1. Install [Codex](https://www.khulnasoft/codex/docs/installing_codex/)
1. Create a new project with:

    ```bash
    codex create --template jekyll
    codex install
    ```

1. Run `codex shell` to install your packages and run the init hook
1. In the root directory, run `codex run generate` to install and package the project with bundler
1. In the root directory, run `codex run serve` to start the server. You can access the Jekyll example at `localhost:4000`

## Related Docs

* [Using Ruby with Codex](https://www.khulnasoft/codex/docs/codex_examples/languages/ruby/)
