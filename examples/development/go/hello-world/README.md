# Go

Go projects can be run in Codex by adding the Go SDK to your project. If your project uses cgo or compiles against C libraries, you should also include them in your packages to ensure Go can compile successfully

[**Example Repo**](https://github.com/khulnasoft/codex/tree/main/examples/development/go/hello-world)

[![Open In Codex.khulnasoft.com](https://www.khulnasoft/img/codex/open-in-codex.svg)](https://codex.khulnasoft.com/open/templates/go)

## Adding Go to your Project

`codex add go`, or add the following to your `codex.json`

```json
  "packages": [
    "go@latest"
  ]
```

This will install the latest version of the Go SDK. You can find other installable versions of Go by running `codex search go`. You can also view the available versions on [Nixhub](https://www.nixhub.io/packages/go)

If you need additional C libraries, you can add them along with `gcc` to your package list. For example, if libcap is required for your project:

```json
"packages": [
    "go",
    "gcc",
    "libcap"
]
```
