// Copyright 2024 Khulnasoft Inc. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package boxcli

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/khulnasoft/codex/internal/codex"
	"github.com/khulnasoft/codex/internal/codex/devopt"
	"github.com/khulnasoft/codex/internal/devpkg"
	"github.com/khulnasoft/codex/internal/ux"
)

type installCmdFlags struct {
	runCmdFlags
	tidyLockfile bool
}

func installCmd() *cobra.Command {
	flags := installCmdFlags{}
	command := &cobra.Command{
		Use:     "install",
		Short:   "Install all packages mentioned in codex.json",
		Args:    cobra.MaximumNArgs(0),
		PreRunE: ensureNixInstalled,
		RunE: func(cmd *cobra.Command, args []string) error {
			return installCmdFunc(cmd, flags)
		},
	}

	flags.config.register(command)
	command.Flags().BoolVar(
		&flags.tidyLockfile, "tidy-lockfile", false,
		"Fix missing store paths in the codex.lock file.",
		// Could potentially do more in the future.
	)

	return command
}

func installCmdFunc(cmd *cobra.Command, flags installCmdFlags) error {
	// Check the directory exists.
	box, err := codex.Open(&devopt.Opts{
		Dir:         flags.config.path,
		Environment: flags.config.environment,
		Stderr:      cmd.ErrOrStderr(),
	})
	if err != nil {
		return errors.WithStack(err)
	}
	ctx := cmd.Context()
	if flags.tidyLockfile {
		ctx = ux.HideMessage(ctx, devpkg.MissingStorePathsWarning)
	}
	if err = box.Install(ctx); err != nil {
		return errors.WithStack(err)
	}
	if flags.tidyLockfile {
		if err = box.FixMissingStorePaths(ctx); err != nil {
			return errors.WithStack(err)
		}
	}
	fmt.Fprintln(cmd.ErrOrStderr(), "Finished installing packages.")
	return nil
}
