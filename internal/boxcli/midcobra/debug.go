// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package midcobra

import (
	"errors"
	"log/slog"
	"os/exec"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/khulnasoft/codex/internal/boxcli/usererr"
	"github.com/khulnasoft/codex/internal/debug"
	"github.com/khulnasoft/codex/internal/telemetry"
	"github.com/khulnasoft/codex/internal/ux"
)

type DebugMiddleware struct {
	flag *pflag.Flag
}

var _ Middleware = (*DebugMiddleware)(nil)

func (d *DebugMiddleware) AttachToFlag(flags *pflag.FlagSet, flagName string) {
	flags.Bool(
		flagName,
		false,
		"Show full stack traces on errors",
	)
	d.flag = flags.Lookup(flagName)
	d.flag.Hidden = true
}

func (d *DebugMiddleware) preRun(cmd *cobra.Command, args []string) {
	if d == nil {
		return
	}

	if d.flag.Changed {
		strVal := d.flag.Value.String()
		if enabled, _ := strconv.ParseBool(strVal); enabled {
			debug.Enable()
		}
	}
}

func (d *DebugMiddleware) postRun(cmd *cobra.Command, args []string, runErr error) {
	if runErr == nil {
		return
	}
	if userErr, hasUserErr := usererr.Extract(runErr); hasUserErr {
		if usererr.IsWarning(userErr) {
			ux.Fwarning(cmd.ErrOrStderr(), runErr.Error())
			return
		}
		color.New(color.FgRed).Fprintf(cmd.ErrOrStderr(), "\nError: %s\n\n", userErr.Error())
	} else {
		color.New(color.FgRed).Fprintf(cmd.ErrOrStderr(), "Error: %v\n\n", runErr)
	}

	st := debug.EarliestStackTrace(runErr)
	var exitErr *exec.ExitError
	if errors.As(runErr, &exitErr) {
		slog.Error("command error", "stderr", exitErr.Stderr, "execid", telemetry.ExecutionID, "stack", st)
	}
	slog.Error("command error", "execid", telemetry.ExecutionID, "stack", st)
}
