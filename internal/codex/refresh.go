package codex

import (
	"fmt"
	"os"
	"strings"
)

func (d *Codex) IsDirenvActive() bool {
	return strings.TrimPrefix(os.Getenv("DIRENV_DIR"), "-") == d.projectDir
}

func (d *Codex) isRefreshAliasSet() bool {
	return os.Getenv(d.refreshAliasEnvVar()) == d.refreshCmd()
}

func (d *Codex) refreshAliasEnvVar() string {
	return "CODEX_REFRESH_ALIAS_" + d.ProjectDirHash()
}

func (d *Codex) isGlobal() bool {
	globalPath, _ := GlobalDataPath()
	return d.projectDir == globalPath
}

// In some cases (e.g. 2 non-global projects somehow active at the same time),
// refresh might not match. This is a tiny edge case, so no need to make UX
// great, we just print out the entire command.
func (d *Codex) refreshAliasOrCommand() string {
	if !d.isRefreshAliasSet() {
		// even if alias is not set, it might still be set by the end of this process
		return fmt.Sprintf("`%s` or `%s`", d.refreshAliasName(), d.refreshCmd())
	}
	return d.refreshAliasName()
}

func (d *Codex) refreshAliasName() string {
	if d.isGlobal() {
		return "refresh-global"
	}
	return "refresh"
}

func (d *Codex) refreshCmd() string {
	codexCmd := fmt.Sprintf("shellenv --preserve-path-stack -c %q", d.projectDir)
	if d.isGlobal() {
		codexCmd = "global shellenv --preserve-path-stack -r"
	}
	if isFishShell() {
		return fmt.Sprintf(`eval (codex %s  | string collect)`, codexCmd)
	}
	return fmt.Sprintf(`eval "$(codex %s)" && hash -r`, codexCmd)
}

func (d *Codex) refreshAlias() string {
	if isFishShell() {
		return fmt.Sprintf(
			`if not type %[1]s >/dev/null 2>&1
	export %[2]s='%[3]s'
	alias %[1]s='%[3]s'
end`,
			d.refreshAliasName(),
			d.refreshAliasEnvVar(),
			d.refreshCmd(),
		)
	}
	return fmt.Sprintf(
		`if ! type %[1]s >/dev/null 2>&1; then
	export %[2]s='%[3]s'
	alias %[1]s='%[3]s'
fi`,
		d.refreshAliasName(),
		d.refreshAliasEnvVar(),
		d.refreshCmd(),
	)
}
