package shenv

type ksh struct{}

// Ksh adds support the korn shell
var Ksh Shell = ksh{}

// um, this is ChatGPT writing it. I need to verify and test
const kshHook = `
_codex_hook() {
  eval "$(codex shellenv --config {{ .ProjectDir }})";
}
if [[ "$(typeset -f precmd)" != *"_codex_hook"* ]]; then
  function precmd {
    codex_hook
  }
fi
`

func (sh ksh) Hook() (string, error) {
	return kshHook, nil
}

func (sh ksh) Export(e ShellExport) (out string) {
	panic("not implemented")
}

func (sh ksh) Dump(env Env) (out string) {
	panic("not implemented")
}
