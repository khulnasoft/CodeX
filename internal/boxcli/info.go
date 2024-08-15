// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package boxcli

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/khulnasoft/codex/internal/codex"
	"github.com/khulnasoft/codex/internal/codex/devopt"
)

type infoCmdFlags struct {
	config   configFlags
	markdown bool
}

func infoCmd() *cobra.Command {
	flags := infoCmdFlags{}
	command := &cobra.Command{
		Use:     "info <pkg>",
		Short:   "Display package info",
		Args:    cobra.ExactArgs(1),
		PreRunE: ensureNixInstalled,
		RunE: func(cmd *cobra.Command, args []string) error {
			return infoCmdFunc(cmd, args[0], flags)
		},
	}

	flags.config.register(command)
	command.Flags().BoolVar(&flags.markdown, "markdown", false, "output in markdown format")
	return command
}

func infoCmdFunc(cmd *cobra.Command, pkg string, flags infoCmdFlags) error {
	box, err := codex.Open(&devopt.Opts{
		Dir:         flags.config.path,
		Environment: flags.config.environment,
		Stderr:      cmd.ErrOrStderr(),
	})
	if err != nil {
		return errors.WithStack(err)
	}

	info, err := box.Info(cmd.Context(), pkg, flags.markdown)
	if err != nil {
		return errors.WithStack(err)
	}
	fmt.Fprint(cmd.OutOrStdout(), info)
	return nil
}
