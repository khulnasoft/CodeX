
# Python

Python by default will attempt to install your packages globally, or in the Nix Store (which it does not have permissions to modify). To use Python with Codex, we recommend setting up a Virtual Environment using pipenv or Poetry (see below).

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/python)

## Adding Python to your Project

`codex add python@3.10`, or in your `codex.json` add:

```json
  "packages": [
    "python@3.10"
  ],
```

This will install Python 3.10 in your shell. You can find other versions of Python by running `codex search python`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/python)

## Pipenv

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/python/pipenv)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/python-pipenv)

[pipenv](https://pipenv.pypa.io/en/latest/) is a tool that will automatically set up a virtual environment for installing your PyPi packages.

You can install `pipenv` by adding it to the packages in your `codex.json`. You can then manage your packages and virtual environment via a Pipfile

```json
{
    "packages": [
        "python310",
        "pipenv"
    ],
    "shell": {
        "init_hook": "pipenv shell"
    }
}
```

This init_hook will automatically start your virtualenv when you run `codex shell`.
