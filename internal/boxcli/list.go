// Copyright 2024 Khulnasoft Inc. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package boxcli

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/khulnasoft/codex/internal/codex"
	"github.com/khulnasoft/codex/internal/codex/devopt"
)

type listCmdFlags struct {
	config configFlags
}

func listCmd() *cobra.Command {
	flags := listCmdFlags{}
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List installed packages",
		PreRunE: ensureNixInstalled,
		RunE: func(cmd *cobra.Command, args []string) error {
			box, err := codex.Open(&devopt.Opts{
				Dir:    flags.config.path,
				Stderr: cmd.ErrOrStderr(),
			})
			if err != nil {
				return errors.WithStack(err)
			}
			for _, p := range box.AllPackageNamesIncludingRemovedTriggerPackages() {
				fmt.Fprintf(cmd.OutOrStdout(), "* %s\n", p)
			}
			return nil
		},
	}
	flags.config.register(cmd)
	return cmd
}
