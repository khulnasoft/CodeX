// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package pullbox

import (
	"context"
	"net/url"
	"os"
	"path/filepath"

	"github.com/khulnasoft/codex/internal/cuecfg"
	"github.com/khulnasoft/codex/internal/devconfig"
	"github.com/khulnasoft/codex/internal/fileutil"
)

func (p *pullbox) IsTextCodexConfig() bool {
	if u, err := url.Parse(p.URL); err == nil {
		ext := filepath.Ext(u.Path)
		return cuecfg.IsSupportedExtension(ext)
	}
	// For invalid URLS, just look at the extension
	ext := filepath.Ext(p.URL)
	return cuecfg.IsSupportedExtension(ext)
}

func (p *pullbox) pullTextCodexConfig(ctx context.Context) error {
	if p.isLocalConfig() {
		return p.copyToProfile(p.URL)
	}

	cfg, err := devconfig.LoadConfigFromURL(ctx, p.URL)
	if err != nil {
		return err
	}

	tmpDir, err := fileutil.CreateCodexTempDir()
	if err != nil {
		return err
	}
	if err = cfg.Root.SaveTo(tmpDir); err != nil {
		return err
	}

	return p.copyToProfile(tmpDir)
}

func (p *pullbox) isLocalConfig() bool {
	_, err := os.Stat(p.URL)
	return err == nil
}
