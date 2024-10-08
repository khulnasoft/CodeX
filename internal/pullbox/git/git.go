// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package git

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/khulnasoft/codex/internal/cmdutil"
	"github.com/khulnasoft/codex/internal/fileutil"
)

func CloneToTmp(repo string) (string, error) {
	tmpDir, err := fileutil.CreateCodexTempDir()
	if err != nil {
		return "", err
	}

	if err := clone(repo, tmpDir); err != nil {
		return "", err
	}
	return tmpDir, nil
}

func IsRepoURL(url string) bool {
	// For now only support ssh
	return strings.HasPrefix(url, "git@") ||
		(strings.HasPrefix(url, "https://") && strings.HasSuffix(url, ".git"))
}

func clone(repo, dir string) error {
	cmd := cmdutil.CommandTTY("git", "clone", repo, dir)
	cmd.Dir = dir
	err := cmd.Run()
	return errors.WithStack(err)
}
