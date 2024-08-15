Test codex using the testscripts framework.

This directory contains testscripts: files ending in `.test.txt` that we
automatically run using the testscripts framework.

For details on how to write these types of files see:
+ https://bitfieldconsulting.com/golang/test-scripts
+ https://pkg.go.dev/github.com/rogpeppe/go-internal/testscript

In addition to the standard testscript commands, we also support running codex
commands. Examples include:
+ `codex init`
+ `codex add <pkg>`
+ ...

We've also added some handy comparison functions
+ `path.len <number>`: verifies that the PATH environment variable has the given number of entries
+ `json.superset <superset.json> <subset.json>`: verifies that `superset.json` has all the keys and values present in `subset.json`
