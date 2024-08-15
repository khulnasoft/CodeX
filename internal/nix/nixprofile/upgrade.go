// Copyright 2024 Khulnasoft Inc. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package nixprofile

import (
	"os"

	"github.com/khulnasoft/codex/internal/devpkg"
	"github.com/khulnasoft/codex/internal/lock"
	"github.com/khulnasoft/codex/internal/nix"
)

func ProfileUpgrade(ProfileDir string, pkg *devpkg.Package, lock *lock.File) error {
	nameOrIndex, err := ProfileListNameOrIndex(
		&ProfileListNameOrIndexArgs{
			Lockfile:   lock,
			Writer:     os.Stderr,
			Package:    pkg,
			ProfileDir: ProfileDir,
		},
	)
	if err != nil {
		return err
	}

	return nix.ProfileUpgrade(ProfileDir, nameOrIndex)
}
