// Copyright 2024 Khulnasoft Inc. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package sshshim

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/khulnasoft/codex/internal/cloud/mutagenbox"
	"github.com/khulnasoft/codex/internal/cloud/openssh"
)

// Setup creates the ssh and scp symlinks
func Setup() error {
	shimDir, err := mutagenbox.ShimDir()
	if err != nil {
		return errors.WithStack(err)
	}

	if err := openssh.EnsureDirExists(shimDir, 0o744, true /*chmod*/); err != nil {
		return err
	}

	codexExecutablePath, err := os.Executable()
	if err != nil {
		return errors.WithStack(err)
	}

	// create ssh symlink
	sshSymlink := filepath.Join(shimDir, "ssh")
	if err := makeSymlink(sshSymlink, codexExecutablePath); err != nil {
		return errors.WithStack(err)
	}

	// create scp symlink
	scpSymlink := filepath.Join(shimDir, "scp")
	return errors.WithStack(makeSymlink(scpSymlink, codexExecutablePath))
}

func makeSymlink(from, target string) error {
	err := os.Remove(from)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return errors.WithStack(err)
	}

	err = os.Symlink(target, from)
	if errors.Is(err, fs.ErrExist) {
		err = nil
	}
	return errors.WithStack(err)
}
