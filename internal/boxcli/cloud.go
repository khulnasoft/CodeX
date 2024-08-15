// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package boxcli

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/khulnasoft/codex/internal/boxcli/usererr"
	"github.com/khulnasoft/codex/internal/cloud"
	"github.com/khulnasoft/codex/internal/codex"
	"github.com/khulnasoft/codex/internal/codex/devopt"
	"github.com/khulnasoft/codex/internal/envir"
)

type cloudShellCmdFlags struct {
	config configFlags

	githubUsername string
}

func cloudCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "cloud",
		Short: "[Preview] Remote development environments on the cloud",
		Long: "Remote development environments on the cloud. All cloud commands " +
			"are currently in developer preview and may have some rough edges. " +
			"Please report any issues to https://github.com/khulnasoft/codex/issues",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	command.AddCommand(cloudShellCmd())
	command.AddCommand(cloudInitCmd())
	command.AddCommand(cloudPortForwardCmd())
	return command
}

func cloudInitCmd() *cobra.Command {
	flags := cloudShellCmdFlags{}
	command := &cobra.Command{
		Use:    "init",
		Hidden: true,
		Short:  "Create a Cloud VM without connecting to its shell",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCloudInit(cmd, &flags)
		},
	}
	flags.config.register(command)
	return command
}

func cloudShellCmd() *cobra.Command {
	flags := cloudShellCmdFlags{}

	command := &cobra.Command{
		Use:   "shell",
		Short: "[Preview] Shell into a cloud environment that matches your local codex environment",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCloudShellCmd(cmd, &flags)
		},
	}

	flags.config.register(command)
	command.Flags().StringVarP(
		&flags.githubUsername, "username", "u", "", "Github username to use for ssh",
	)
	return command
}

func cloudPortForwardCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "forward <local-port>:<remote-port> | :<remote-port> | stop | list",
		Short: "[Preview] Port forward a local port to a remote codex cloud port",
		Long: "Port forward a local port to a remote codex cloud port. If 0 or " +
			"no local port is specified, we find a suitable local port. Use 'stop' " +
			"to stop all port forwards.",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ports := strings.Split(args[0], ":")

			if len(ports) != 2 {
				return usererr.New("Invalid port format. Expected <local-port>:<remote-port>")
			}
			localPort, err := cloud.PortForward(ports[0], ports[1])
			if err != nil {
				return errors.WithStack(err)
			}
			cmd.PrintErrf(
				"Port forwarding %s:%s\nTo view in browser, visit http://localhost:%[1]s\n",
				localPort,
				ports[1],
			)
			return nil
		},
	}
	command.AddCommand(cloudPortForwardList())
	command.AddCommand(cloudPortForwardStopCmd())
	return command
}

func cloudPortForwardStopCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stop all port forwards managed by codex",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cloud.PortForwardTerminateAll()
		},
	}
}

func cloudPortForwardList() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all port forwards managed by codex",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			l, err := cloud.PortForwardList()
			if err != nil {
				return errors.WithStack(err)
			}
			for _, p := range l {
				cmd.Println(p)
			}
			return nil
		},
	}
}

func runCloudShellCmd(cmd *cobra.Command, flags *cloudShellCmdFlags) error {
	// calling `codex cloud shell` when already in the VM is not allowed.
	if envir.IsCodexCloud() {
		return shellInceptionErrorMsg("codex cloud shell")
	}

	box, err := codex.Open(&devopt.Opts{
		Dir:         flags.config.path,
		Environment: flags.config.environment,
		Stderr:      cmd.ErrOrStderr(),
	})
	if err != nil {
		return errors.WithStack(err)
	}
	return cloud.Shell(cmd.Context(), cmd.ErrOrStderr(), box.ProjectDir(), flags.githubUsername)
}

func runCloudInit(cmd *cobra.Command, flags *cloudShellCmdFlags) error {
	// calling `codex cloud init` when already in the VM is not allowed.
	if envir.IsCodexCloud() {
		return shellInceptionErrorMsg("codex cloud init")
	}

	box, err := codex.Open(&devopt.Opts{
		Dir:         flags.config.path,
		Environment: flags.config.environment,
		Stderr:      cmd.ErrOrStderr(),
	})
	if err != nil {
		return errors.WithStack(err)
	}
	_, vmhostname, _, err := cloud.InitVM(cmd.Context(), cmd.ErrOrStderr(), box.ProjectDir(), flags.githubUsername)
	if err != nil {
		return err
	}
	// printing vmHostname so that the output of codex cloud init can be read by
	// codex extension
	fmt.Fprintln(cmd.ErrOrStderr(), vmhostname)
	return nil
}
