package pkgtype

import (
	"strings"

	"github.com/khulnasoft/codex/nix/flake"
)

func IsFlake(s string) bool {
	if IsRunX(s) {
		return false
	}
	parsed, err := flake.ParseInstallable(s)
	if err != nil {
		return false
	}
	if IsAmbiguous(s, parsed) {
		return false
	}
	return true
}

// IsAmbiguous returns true if a package string could be a Codex package or
// a flake installable. For example, "nixpkgs" is both a Codex package and a
// flake.
func IsAmbiguous(raw string, parsed flake.Installable) bool {
	// Codex package strings never have a #attr_path in them.
	if parsed.AttrPath != "" {
		return false
	}

	// Indirect installables must have a "flake:" scheme to disambiguate
	// them from legacy (unversioned) codex package strings.
	if parsed.Ref.Type == flake.TypeIndirect {
		return !strings.HasPrefix(raw, "flake:")
	}

	// Path installables must have a "path:" scheme, start with "/" or start
	// with "./" to disambiguate them from codex package strings.
	if parsed.Ref.Type == flake.TypePath {
		if raw[0] == '.' || raw[0] == '/' {
			return false
		}
		if strings.HasPrefix(raw, "path:") {
			return false
		}
		return true
	}

	// All other flakeref types must have a scheme, so we know those can't
	// be codex package strings.
	return false
}
