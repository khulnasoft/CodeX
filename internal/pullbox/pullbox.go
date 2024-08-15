// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package pullbox

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"runtime/trace"

	"github.com/pkg/errors"

	"github.com/khulnasoft/codex/internal/boxcli/usererr"
	"github.com/khulnasoft/codex/internal/codex/devopt"
	"github.com/khulnasoft/codex/internal/pullbox/git"
	"github.com/khulnasoft/codex/internal/pullbox/s3"
	"github.com/khulnasoft/codex/internal/pullbox/tar"
	"github.com/khulnasoft/codex/internal/ux"
)

type codexProject interface {
	ProjectDir() string
}

type pullbox struct {
	codexProject
	devopt.PullboxOpts
}

func New(codex codexProject, opts devopt.PullboxOpts) *pullbox {
	return &pullbox{codex, opts}
}

// Pull
// This can be rewritten to be more readable and less repetitive. Possibly
// something like:
// puller := getPullerForURL(url)
// return puller.Pull()
func (p *pullbox) Pull(ctx context.Context) error {
	defer trace.StartRegion(ctx, "Pull").End()
	var err error

	notEmpty, err := profileIsNotEmpty(p.ProjectDir())
	if err != nil {
		return err
	} else if notEmpty && !p.Overwrite {
		return fs.ErrExist
	}

	if p.URL != "" {
		ux.Finfo(os.Stderr, "Pulling global config from %s\n", p.URL)
	} else {
		ux.Finfo(os.Stderr, "Pulling global config\n")
	}

	var tmpDir string

	if p.URL == "" {
		if p.Credentials.IDToken == "" {
			return usererr.New("Not logged in")
		}
		profile := "default" // TODO: make this editable
		if tmpDir, err = s3.PullToTmp(ctx, &p.Credentials, profile); err != nil {
			return err
		}
		return p.copyToProfile(tmpDir)
	}

	if git.IsRepoURL(p.URL) {
		if tmpDir, err = git.CloneToTmp(p.URL); err != nil {
			return err
		}
		// Remove the .git directory, we don't want to keep state
		if err := os.RemoveAll(filepath.Join(tmpDir, ".git")); err != nil {
			return errors.WithStack(err)
		}
		return p.copyToProfile(tmpDir)
	}

	if p.IsTextCodexConfig() {
		return p.pullTextCodexConfig(ctx)
	}

	if isArchive, err := urlIsArchive(p.URL); err != nil {
		return err
	} else if isArchive {
		data, err := download(p.URL)
		if err != nil {
			return err
		}

		if tmpDir, err = tar.Extract(data); err != nil {
			return err
		}

		return p.copyToProfile(tmpDir)
	}

	return usererr.New("Could not determine how to pull %s", p.URL)
}

func (p *pullbox) Push(ctx context.Context) error {
	if p.URL != "" {
		ux.Finfo(os.Stderr, "Pushing global config to %s\n", p.URL)
	} else {
		ux.Finfo(os.Stderr, "Pushing global config\n")
	}

	if p.URL == "" {
		profile := "default" // TODO: make this editable
		if p.Credentials.IDToken == "" {
			return usererr.New("Not logged in")
		}
		ux.Finfo(
			os.Stderr,
			"Logged in as %s, pushing to to codex cloud (profile: %s)\n",
			p.Credentials.Email,
			profile,
		)
		return s3.Push(ctx, &p.Credentials, p.ProjectDir(), profile)
	}
	return git.Push(ctx, p.ProjectDir(), p.URL)
}
