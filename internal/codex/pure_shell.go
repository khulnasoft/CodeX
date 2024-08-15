// Copyright 2024 Khulnasoft Inc. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package codex

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Creates a symlink for codex in .codex/bin
// so that codex can be available inside a pure shell
func createCodexSymlink(d *Codex) error {
	// Get absolute path for where codex is called
	codexPath, err := os.Executable()
	if err != nil {
		return errors.Wrap(err, "failed to create codex symlink. Codex command won't be available inside the shell")
	}
	// ensure .codex/bin directory exists
	binPath := dotcodexBinPath(d)
	if err := os.MkdirAll(binPath, 0o755); err != nil {
		return errors.WithStack(err)
	}
	// Create a symlink between codex and .codex/bin
	err = os.Symlink(codexPath, filepath.Join(binPath, "codex"))
	if err != nil && !errors.Is(err, fs.ErrExist) {
		return errors.Wrap(err, "failed to create codex symlink. Codex command won't be available inside the shell")
	}
	return nil
}

func dotcodexBinPath(d *Codex) string {
	return filepath.Join(d.ProjectDir(), ".codex/bin")
}
