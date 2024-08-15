package shenv

type posix struct{}

// Posix adds support for posix-compatible shells
// Specifically, in the context of codex, this includes
// `dash`, `ash`, and `shell`
var Posix Shell = posix{}

// um, this is ChatGPT writing it. I need to verify and test
const posixHook = `
_codex_hook() {
  local previous_exit_status=$?
  trap : INT
  eval "$(codex shellenv --config {{ .ProjectDir }})"
  trap - INT
  return $previous_exit_status
}
if [ -z "$PROMPT_COMMAND" ] || ! printf "%s" "$PROMPT_COMMAND" | grep -q "_codex_hook"; then
  PROMPT_COMMAND="_codex_hook${PROMPT_COMMAND:+;$PROMPT_COMMAND}"
fi
`

func (sh posix) Hook() (string, error) {
	return posixHook, nil
}

func (sh posix) Export(e ShellExport) (out string) {
	panic("not implemented")
}

func (sh posix) Dump(env Env) (out string) {
	panic("not implemented")
}
