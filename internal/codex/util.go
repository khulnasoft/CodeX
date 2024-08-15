// Copyright 2024 Khulnasoft Inc. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package codex

import (
	"context"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/khulnasoft/codex/internal/codex/devopt"
	"github.com/khulnasoft/codex/internal/xdg"
)

const processComposeVersion = "1.5.0"

var utilProjectConfigPath string

func initCodexUtilityProject(ctx context.Context, stderr io.Writer) error {
	codexUtilityProjectPath, err := ensureCodexUtilityConfig()
	if err != nil {
		return err
	}

	box, err := Open(&devopt.Opts{
		Dir:    codexUtilityProjectPath,
		Stderr: stderr,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	// Add all utilities here.
	utilities := []string{
		"process-compose@" + processComposeVersion,
	}
	if err = box.Add(ctx, utilities, devopt.AddOpts{}); err != nil {
		return err
	}

	return box.Install(ctx)
}

func ensureCodexUtilityConfig() (string, error) {
	if utilProjectConfigPath != "" {
		return utilProjectConfigPath, nil
	}

	path, err := utilityDataPath()
	if err != nil {
		return "", err
	}

	_, err = InitConfig(path)
	if err != nil {
		return "", err
	}

	// Avoids unnecessarily initializing the config again by caching the path
	utilProjectConfigPath = path

	return path, nil
}

func utilityLookPath(binName string) (string, error) {
	binPath, err := utilityBinPath()
	if err != nil {
		return "", err
	}
	absPath := filepath.Join(binPath, binName)
	_, err = os.Stat(absPath)
	if errors.Is(err, fs.ErrNotExist) {
		return "", err
	}
	return absPath, nil
}

func utilityDataPath() (string, error) {
	path := xdg.DataSubpath("codex/util")
	return path, errors.WithStack(os.MkdirAll(path, 0o755))
}

func utilityNixProfilePath() (string, error) {
	path, err := utilityDataPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(path, ".codex/nix/profile"), nil
}

func utilityBinPath() (string, error) {
	nixProfilePath, err := utilityNixProfilePath()
	if err != nil {
		return "", err
	}

	return filepath.Join(nixProfilePath, "default/bin"), nil
}
