# F# and .NET

F# and .NET projects can be easily generated in Codex by adding the dotnet SDK to your project. You can then create new projects using `dotnet new`

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/fsharp)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/fsharp)

## Adding .NET to your project

`codex add dotnet-sdk`, or add the following to your `codex.json`:

```json
  "packages": [
    "dotnet-sdk@latest"
  ],
```

This will install the latest version of the dotnet SDK. You can find other installable versions of the dotnet SDK by running `codex search dotnet-sdk`. You can also view the available versions on [Nixhub](https://www.nixhub.io/search?q=dotnet)

## Creating a new F# Project

`dotnet new console -lang "F#" -o <name>`
