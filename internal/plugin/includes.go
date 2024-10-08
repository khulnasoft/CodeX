package plugin

import (
	"strings"

	"github.com/khulnasoft/codex/internal/devpkg"
	"github.com/khulnasoft/codex/internal/lock"
)

func LoadConfigFromInclude(include string, lockfile *lock.File, workingDir string) (*Config, error) {
	var includable Includable
	var err error
	if t, name, _ := strings.Cut(include, ":"); t == "plugin" {
		includable = devpkg.PackageFromStringWithDefaults(
			name,
			lockfile,
		)
	} else {
		includable, err = parseIncludable(include, workingDir)
		if err != nil {
			return nil, err
		}
	}
	return getConfigIfAny(includable, lockfile.ProjectDir())
}
