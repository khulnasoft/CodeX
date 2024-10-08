// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package boxcli

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/khulnasoft/codex/internal/boxcli/usererr"
	"github.com/khulnasoft/codex/internal/codex"
	"github.com/khulnasoft/codex/internal/codex/devopt"
	"github.com/khulnasoft/codex/internal/envir"
)

type shellCmdFlags struct {
	envFlag
	config     configFlags
	omitNixEnv bool
	printEnv   bool
	pure       bool
}

// shellFlagDefaults are the flag default values that differ
// from the `codex` command versus `codex global` command.
type shellFlagDefaults struct {
	omitNixEnv bool
}

func shellCmd(defaults shellFlagDefaults) *cobra.Command {
	flags := shellCmdFlags{}
	command := &cobra.Command{
		Use:   "shell",
		Short: "Start a new shell with access to your packages",
		Long: "Start a new shell with access to your packages.\n\n" +
			"If the --config flag is set, the shell will be started using the codex.json found in the --config flag directory. " +
			"If --config isn't set, then codex recursively searches the current directory and its parents.",
		Args:    cobra.NoArgs,
		PreRunE: ensureNixInstalled,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShellCmd(cmd, flags)
		},
	}

	command.Flags().BoolVar(
		&flags.printEnv, "print-env", false, "print script to setup shell environment")
	command.Flags().BoolVar(
		&flags.pure, "pure", false, "if this flag is specified, codex creates an isolated shell inheriting almost no variables from the current environment. A few variables, in particular HOME, USER and DISPLAY, are retained.")
	command.Flags().BoolVar(
		&flags.omitNixEnv, "omit-nix-env", defaults.omitNixEnv,
		"shell environment will omit the env-vars from print-dev-env",
	)
	_ = command.Flags().MarkHidden("omit-nix-env")

	flags.config.register(command)
	flags.envFlag.register(command)
	return command
}

func runShellCmd(cmd *cobra.Command, flags shellCmdFlags) error {
	env, err := flags.Env(flags.config.path)
	if err != nil {
		return err
	}
	// Check the directory exists.
	box, err := codex.Open(&devopt.Opts{
		Dir:         flags.config.path,
		Env:         env,
		Environment: flags.config.environment,
		Stderr:      cmd.ErrOrStderr(),
	})
	if err != nil {
		return errors.WithStack(err)
	}

	if flags.printEnv {
		// false for includeHooks is because init hooks is not compatible with .envrc files generated
		// by versions older than 0.4.6
		script, err := box.EnvExports(cmd.Context(), devopt.EnvExportsOpts{})
		if err != nil {
			return err
		}
		// explicitly print to stdout instead of stderr so that direnv can read the output
		fmt.Fprint(cmd.OutOrStdout(), script)
		return nil // return here to prevent opening a codex shell
	}

	if envir.IsCodexShellEnabled() {
		return shellInceptionErrorMsg("codex shell")
	}

	return box.Shell(cmd.Context(), devopt.EnvOptions{
		OmitNixEnv: flags.omitNixEnv,
		Pure:       flags.pure,
	})
}

func shellInceptionErrorMsg(cmdPath string) error {
	return usererr.New("You are already in an active %[1]s.\nRun `exit` before calling `%[1]s` again."+
		" Shell inception is not supported.", cmdPath)
}
