package shellgen

import (
	"bytes"
	_ "embed"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/khulnasoft/codex/internal/boxcli/featureflag"
	"github.com/khulnasoft/codex/internal/debug"
	"github.com/khulnasoft/codex/internal/devconfig"
	"github.com/khulnasoft/codex/internal/devpkg"
	"github.com/khulnasoft/codex/internal/lock"
	"github.com/khulnasoft/codex/internal/plugin"
)

//go:embed tmpl/script-wrapper.tmpl
var scriptWrapperTmplString string
var scriptWrapperTmpl = template.Must(template.New("script-wrapper").Parse(scriptWrapperTmplString))

const scriptsDir = ".codex/gen/scripts"

const HooksFilename = ".hooks"

type codexer interface {
	Config() *devconfig.Config
	Lockfile() *lock.File
	InstallablePackages() []*devpkg.Package
	PluginManager() *plugin.Manager
	ProjectDir() string
	SkipInitHookEnvName() string
}

// WriteScriptsToFiles writes scripts defined in codex.json into files inside .codex/gen/scripts.
// Scripts (and hooks) are persisted so that we can easily call them from codex run (inside or outside shell).
func WriteScriptsToFiles(codex codexer) error {
	defer debug.FunctionTimer().End()
	err := os.MkdirAll(filepath.Join(codex.ProjectDir(), scriptsDir), 0o755) // Ensure directory exists.
	if err != nil {
		return errors.WithStack(err)
	}

	// Read dir contents before writing, so we can clean up later.
	entries, err := os.ReadDir(filepath.Join(codex.ProjectDir(), scriptsDir))
	if err != nil {
		return errors.WithStack(err)
	}

	// Write all hooks to a file.
	written := map[string]struct{}{} // set semantics; value is irrelevant
	// always write it, even if there are no hooks, because scripts will source it.
	err = writeRawInitHookFile(codex, codex.Config().InitHook().String())
	if err != nil {
		return errors.WithStack(err)
	}
	written[HooksFilename] = struct{}{}

	// Write scripts to files.
	for name, body := range codex.Config().Scripts() {
		scriptBody, err := ScriptBody(codex, body.String())
		if err != nil {
			return errors.WithStack(err)
		}
		err = WriteScriptFile(codex, name, scriptBody)
		if err != nil {
			return errors.WithStack(err)
		}
		written[name] = struct{}{}
	}

	// Delete any files that weren't written just now.
	for _, entry := range entries {
		scriptName := strings.TrimSuffix(entry.Name(), ".sh")
		if _, ok := written[scriptName]; !ok && !entry.IsDir() {
			err := os.Remove(ScriptPath(codex.ProjectDir(), scriptName))
			if err != nil {
				slog.Debug("failed to clean up script file %s, error = %s", entry.Name(), err) // no need to fail run
			}
		}
	}

	return nil
}

func writeRawInitHookFile(codex codexer, body string) (err error) {
	script, err := createScriptFile(codex, HooksFilename)
	if err != nil {
		return errors.WithStack(err)
	}
	defer script.Close() // best effort: close file

	_, err = script.WriteString(body)
	return errors.WithStack(err)
}

func WriteScriptFile(codex codexer, name, body string) (err error) {
	script, err := createScriptFile(codex, name)
	if err != nil {
		return errors.WithStack(err)
	}
	defer script.Close() // best effort: close file

	if featureflag.ScriptExitOnError.Enabled() {
		// NOTE: Codex scripts run using `sh` for consistency.
		body = fmt.Sprintf("set -e\n\n%s", body)
	}
	_, err = script.WriteString(body)
	return errors.WithStack(err)
}

func createScriptFile(codex codexer, name string) (script *os.File, err error) {
	script, err = os.Create(ScriptPath(codex.ProjectDir(), name))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() {
		// best effort: close file if there was some subsequent error
		if err != nil {
			_ = script.Close()
		}
	}()

	err = script.Chmod(0o755)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return script, nil
}

func ScriptPath(projectDir, scriptName string) string {
	return filepath.Join(projectDir, scriptsDir, scriptName+".sh")
}

func ScriptBody(d codexer, body string) (string, error) {
	var buf bytes.Buffer
	err := scriptWrapperTmpl.Execute(&buf, map[string]string{
		"Body":             body,
		"SkipInitHookHash": d.SkipInitHookEnvName(),
		"InitHookPath":     ScriptPath(d.ProjectDir(), HooksFilename),
	})
	if err != nil {
		return "", errors.WithStack(err)
	}
	return buf.String(), nil
}
