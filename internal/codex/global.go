// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package codex

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/khulnasoft/codex/internal/xdg"
)

// In the future we will support multiple global profiles
const currentGlobalProfile = "default"

func GlobalDataPath() (string, error) {
	path := xdg.DataSubpath(filepath.Join("codex/global", currentGlobalProfile))
	if err := os.MkdirAll(path, 0o755); err != nil {
		return "", errors.WithStack(err)
	}

	nixProfilePath := filepath.Join(path)
	currentPath := xdg.DataSubpath("codex/global/current")

	// For now default is always current. In the future we will support multiple
	// and allow user to switch. Remove any existing symlink and create a new one
	// because previous versions of codex may have created a symlink to a
	// different profile.
	existing, _ := os.Readlink(currentPath)
	if existing != nixProfilePath {
		_ = os.Remove(currentPath)
	}

	err := os.Symlink(nixProfilePath, currentPath)
	if err != nil && !errors.Is(err, fs.ErrExist) {
		return "", errors.WithStack(err)
	}

	return path, nil
}
