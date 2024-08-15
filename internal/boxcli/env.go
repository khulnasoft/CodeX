// Copyright 2024 Khulnasoft Inc. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package boxcli

import (
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/khulnasoft/codex/internal/codex/devopt"
)

// to be composed into xyzCmdFlags structs
type envFlag devopt.EnvFlags

func (f *envFlag) register(cmd *cobra.Command) {
	cmd.PersistentFlags().StringToStringVarP(
		&f.EnvMap, "env", "e", nil, "environment variables to set in the codex environment",
	)
	cmd.PersistentFlags().StringVar(
		&f.EnvFile, "env-file", "", "path to a file containing environment variables to set in the codex environment",
	)
}

func (f *envFlag) Env(path string) (map[string]string, error) {
	envs := map[string]string{}
	var err error
	if f.EnvFile != "" {
		envPath := f.EnvFile
		if !filepath.IsAbs(envPath) {
			envPath = filepath.Join(path, envPath)
		}
		envs, err = godotenv.Read(envPath)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	for k, v := range f.EnvMap {
		envs[k] = v
	}

	return envs, nil
}
