---
title: C# and .NET
---

C# and .NET projects can be easily generated in Codex by adding the dotnet SDK to your project. You can then create new projects using `dotnet new`

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/csharp)

[![Open In Codex.sh](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.sh/open/templates/dotnet)

## Adding .NET to your project

`codex add dotnet-sdk`, or add the following to your `codex.json`:

```json
  "packages": [
    "dotnet-sdk@latest"
  ],
```
This will install the latest version of the dotnet SDK. You can find other installable versions of the dotnet SDK by running `codex search dotnet-sdk`.

If you need a specific version of the .NET SDK, you can search on [Nixhub](https://www.nixhub.io/search?q=dotnet)

## Creating a new C# Project

`dotnet new console -lang "C#" -o <name>`
