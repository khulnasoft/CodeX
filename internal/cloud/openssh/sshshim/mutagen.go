// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package sshshim

import (
	"bytes"
	"context"
	"log/slog"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/khulnasoft/codex/internal/cloud/mutagenbox"
)

// EnsureLiveVMOrTerminateMutagenSessions returns true if a liveVM is found, OR sshArgs were connecting to a server that is not a codex-VM.
// EnsureLiveVMOrTerminateMutagenSessions returns false iff the sshArgs were connecting to a codex VM AND a deadVM is found.
func EnsureLiveVMOrTerminateMutagenSessions(ctx context.Context, sshArgs []string) (bool, error) {
	vmAddr := vmAddressIfAny(sshArgs)

	slog.Debug("found vm address", "addr", vmAddr)
	if vmAddr == "" {
		// We support the no Vm scenario, in case mutagen ssh-es into another server
		// TODO savil. Revisit the no VM scenario if we can control the mutagen daemon for codex-only
		// syncing via MUTAGEN_DATA_DIRECTORY.
		return true, nil
	}

	isActive, err := checkActiveVMWithRetries(ctx, vmAddr)
	if err != nil {
		return false, errors.WithStack(err)
	}
	if !isActive {
		slog.Debug("terminating vm mutagen session", "addr", vmAddr)
		// If no vm is active, then we should terminate the running mutagen sessions
		return false, terminateMutagenSessions(vmAddr)
	}
	return true, nil
}

func terminateMutagenSessions(vmAddr string) error {
	username, hostname, found := strings.Cut(vmAddr, "@")
	if !found {
		hostname = username
	}
	machineID, _, found := strings.Cut(hostname, ".")
	if !found {
		return errors.Errorf(
			"expected to find a period (.) in hostname (%s), but did not. "+
				"For completeness, VmAddr is %s", hostname, vmAddr)
	}

	if err := mutagenbox.TerminateSessionsForMachine(machineID, nil /*env*/); err != nil {
		return err
	}

	return mutagenbox.ForwardTerminateByHost(hostname)
}

func checkActiveVMWithRetries(ctx context.Context, vmAddr string) (bool, error) {
	var finalErr error

	// Try 3 times:
	for num := 0; num < 3; num++ {
		isActive, err := checkActiveVM(ctx, vmAddr)
		if err == nil && isActive {
			// found an active VM
			return true, nil
		}
		finalErr = err
		time.Sleep(10 * time.Second)
		slog.Debug("failed to find active vm", "attempt", num, "addr", vmAddr)
	}
	return false, finalErr
}

func checkActiveVM(ctx context.Context, vmAddr string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Minute*2)
	defer cancel()
	cmd := exec.CommandContext(ctx, "ssh", vmAddr, "echo 'alive'")

	var bufErr, bufOut bytes.Buffer
	cmd.Stderr = &bufErr
	cmd.Stdout = &bufOut

	err := cmd.Run()
	if err != nil {
		if e := (&exec.ExitError{}); errors.As(err, &e) && e.ExitCode() == 255 {
			slog.Debug("checkActiveVM: No active VM. returning false for exit status 255")
			return false, nil
		}
		// For now, any error is deemed to indicate a VM that is no longer running.
		// We can tighten this by listening for the specific exit error code (255)
		slog.Debug("Error checking for Active VM: %s. Stdout: %s, Stderr: %s, cmd.Run err: %s\n",
			vmAddr,
			bufOut.String(),
			bufErr.String(),
			err,
		)
		return false, errors.WithStack(err)
	}
	return true, nil
}

// vmAddressIfAny will seek to find the codex-vm hostname if it exists
// in the sshArgs. If not, it returns an empty string.
func vmAddressIfAny(sshArgs []string) string {
	const codexVMAddressSuffix = "codex-vms.internal"
	for _, sshArg := range sshArgs {
		if strings.HasSuffix(sshArg, codexVMAddressSuffix) {
			return sshArg
		}
	}
	slog.Debug("did not find vm address in ssh args", "args", sshArgs)
	return ""
}
