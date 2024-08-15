package multi

import (
	"io/fs"
	"path/filepath"

	"github.com/khulnasoft/codex/internal/debug"
	"github.com/khulnasoft/codex/internal/codex"
	"github.com/khulnasoft/codex/internal/codex/devopt"
	"github.com/khulnasoft/codex/internal/devconfig/configfile"
)

func Open(opts *devopt.Opts) ([]*codex.Codex, error) {
	defer debug.FunctionTimer().End()

	var boxes []*codex.Codex
	err := filepath.WalkDir(
		".",
		func(path string, dirEntry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !dirEntry.IsDir() && filepath.Base(path) == configfile.DefaultName {
				optsCopy := *opts
				optsCopy.Dir = path
				box, err := codex.Open(&optsCopy)
				if err != nil {
					return err
				}
				boxes = append(boxes, box)
			}

			return nil
		},
	)

	return boxes, err
}
