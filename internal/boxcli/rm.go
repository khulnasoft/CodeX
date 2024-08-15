// Copyright 2024 Khulnasoft Inc. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package boxcli

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/khulnasoft/codex/internal/codex"
	"github.com/khulnasoft/codex/internal/codex/devopt"
)

type removeCmdFlags struct {
	config configFlags
}

func removeCmd() *cobra.Command {
	flags := removeCmdFlags{}
	command := &cobra.Command{
		Use:     "rm <pkg>...",
		Short:   "Remove a package from your codex",
		Args:    cobra.MinimumNArgs(1),
		PreRunE: ensureNixInstalled,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRemoveCmd(cmd, args, flags)
		},
	}

	flags.config.register(command)
	return command
}

func runRemoveCmd(cmd *cobra.Command, args []string, flags removeCmdFlags) error {
	box, err := codex.Open(&devopt.Opts{
		Dir:         flags.config.path,
		Environment: flags.config.environment,
		Stderr:      cmd.ErrOrStderr(),
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return box.Remove(cmd.Context(), args...)
}
