// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package codex

import (
	"os"
	"slices"
	"strings"

	"github.com/khulnasoft/codex/internal/codex/envpath"
	"github.com/khulnasoft/codex/internal/envir"
)

const codexSetPrefix = "__CODEX_SET_"

// exportify formats vars as a line-separated string of shell export statements.
// Each line is of the form `export key="value";` with any special characters in
// value escaped. This means that the shell will always interpret values as
// literal strings; no variable expansion or command substitution will take
// place.
func exportify(vars map[string]string) string {
	keys := make([]string, len(vars))
	i := 0
	for k := range vars {
		keys[i] = k
		i++
	}
	slices.Sort(keys) // for reproducibility

	strb := strings.Builder{}
	for _, k := range keys {
		strb.WriteString("export ")
		strb.WriteString(k)
		strb.WriteString(`="`)
		for _, r := range vars[k] {
			switch r {
			// Special characters inside double quotes:
			// https://pubs.opengroup.org/onlinepubs/009604499/utilities/xcu_chap02.html#tag_02_02_03
			case '$', '`', '"', '\\', '\n':
				strb.WriteRune('\\')
			}
			strb.WriteRune(r)
		}
		strb.WriteString("\";\n")
	}
	return strings.TrimSpace(strb.String())
}

// addEnvIfNotPreviouslySetByCodex adds the key-value pairs from new to existing,
// but only if the key was not previously set by codex
// Caveat, this won't mark the values as set by codex automatically. Instead,
// you need to call markEnvAsSetByCodex when you are done setting variables.
// This is so you can add variables from multiple sources (e.g. plugin, codex.json)
// that may build on each other (e.g. PATH=$PATH:...)
func addEnvIfNotPreviouslySetByCodex(existing, new map[string]string) {
	for k, v := range new {
		if _, alreadySet := existing[codexSetPrefix+k]; !alreadySet {
			existing[k] = v
		}
	}
}

func markEnvsAsSetByCodex(envs ...map[string]string) {
	for _, env := range envs {
		for key := range env {
			env[codexSetPrefix+key] = "1"
		}
	}
}

// IsEnvEnabled checks if the codex environment is enabled.
// This allows us to differentiate between global and
// individual project shells.
func (d *Codex) IsEnvEnabled() bool {
	fakeEnv := map[string]string{}
	// the Stack is initialized in the fakeEnv, from the state in the real os.Environ
	pathStack := envpath.Stack(fakeEnv, envir.PairsToMap(os.Environ()))
	return pathStack.Has(d.ProjectDirHash())
}

func (d *Codex) SkipInitHookEnvName() string {
	return "__CODEX_SKIP_INIT_HOOK_" + d.ProjectDirHash()
}
