// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package boxcli

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/khulnasoft/codex/internal/codex"
)

func initCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "init [<dir>]",
		Short: "Initialize a directory as a codex project",
		Long: "Initialize a directory as a codex project. " +
			"This will create an empty codex.json in the current directory. " +
			"You can then add packages using `codex add`",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInitCmd(args)
		},
	}

	return command
}

func runInitCmd(args []string) error {
	path := pathArg(args)

	_, err := codex.InitConfig(path)
	return errors.WithStack(err)
}
