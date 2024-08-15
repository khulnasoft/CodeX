// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package codex

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/khulnasoft/codex/internal/codex/envpath"

	"github.com/khulnasoft/codex/internal/codex/devopt"
	"github.com/khulnasoft/codex/internal/devconfig"
	"github.com/khulnasoft/codex/internal/envir"
	"github.com/khulnasoft/codex/internal/nix"
)

func TestCodex(t *testing.T) {
	t.Setenv("TMPDIR", "/tmp")
	testPaths, err := doublestar.FilepathGlob("../../examples/**/codex.json")
	require.NoError(t, err, "Reading testdata/ should not fail")

	assert.Greater(t, len(testPaths), 0, "testdata/ and examples/ should contain at least 1 test")

	for _, testPath := range testPaths {
		if !strings.Contains(testPath, "/commands/") {
			testShellPlan(t, testPath)
		}
	}
}

func testShellPlan(t *testing.T, testPath string) {
	baseDir := filepath.Dir(testPath)
	testName := fmt.Sprintf("%s_shell_plan", filepath.Base(baseDir))
	t.Run(testName, func(t *testing.T) {
		t.Setenv(envir.XDGDataHome, "/tmp/codex")
		assert := assert.New(t)

		_, err := Open(&devopt.Opts{
			Dir:    baseDir,
			Stderr: os.Stderr,
		})
		assert.NoErrorf(err, "%s should be a valid codex project", baseDir)
	})
}

type testNix struct {
	path string
}

func (n *testNix) PrintDevEnv(ctx context.Context, args *nix.PrintDevEnvArgs) (*nix.PrintDevEnvOut, error) {
	return &nix.PrintDevEnvOut{
		Variables: map[string]nix.Variable{
			"PATH": {
				Type:  "exported",
				Value: n.path,
			},
		},
	}, nil
}

func TestComputeEnv(t *testing.T) {
	d := codexForTesting(t)
	d.nix = &testNix{}
	ctx := context.Background()
	env, err := d.computeEnv(ctx, false /*use cache*/, devopt.EnvOptions{})
	require.NoError(t, err, "computeEnv should not fail")
	assert.NotNil(t, env, "computeEnv should return a valid env")
}

func TestComputeCodexPathIsIdempotent(t *testing.T) {
	codex := codexForTesting(t)
	codex.nix = &testNix{"/tmp/my/path"}
	ctx := context.Background()
	env, err := codex.computeEnv(ctx, false /*use cache*/, devopt.EnvOptions{})
	require.NoError(t, err, "computeEnv should not fail")
	path := env["PATH"]
	assert.NotEmpty(t, path, "path should not be nil")

	t.Setenv("PATH", path)
	t.Setenv(envpath.InitPathEnv, env[envpath.InitPathEnv])
	t.Setenv(envpath.PathStackEnv, env[envpath.PathStackEnv])
	t.Setenv(envpath.Key(codex.ProjectDirHash()), env[envpath.Key(codex.ProjectDirHash())])

	env, err = codex.computeEnv(ctx, false /*use cache*/, devopt.EnvOptions{})
	require.NoError(t, err, "computeEnv should not fail")
	path2 := env["PATH"]

	assert.Equal(t, path, path2, "path should be the same")
}

func TestComputeCodexPathWhenRemoving(t *testing.T) {
	codex := codexForTesting(t)
	codex.nix = &testNix{"/tmp/my/path"}
	ctx := context.Background()
	env, err := codex.computeEnv(ctx, false /*use cache*/, devopt.EnvOptions{})
	require.NoError(t, err, "computeEnv should not fail")
	path := env["PATH"]
	assert.NotEmpty(t, path, "path should not be nil")
	assert.Contains(t, path, "/tmp/my/path", "path should contain /tmp/my/path")

	t.Setenv("PATH", path)
	t.Setenv(envpath.InitPathEnv, env[envpath.InitPathEnv])
	t.Setenv(envpath.PathStackEnv, env[envpath.PathStackEnv])
	t.Setenv(envpath.Key(codex.ProjectDirHash()), env[envpath.Key(codex.ProjectDirHash())])

	codex.nix.(*testNix).path = ""
	env, err = codex.computeEnv(ctx, false /*use cache*/, devopt.EnvOptions{})
	require.NoError(t, err, "computeEnv should not fail")
	path2 := env["PATH"]
	assert.NotContains(t, path2, "/tmp/my/path", "path should not contain /tmp/my/path")

	assert.NotEqual(t, path, path2, "path should not be the same")
}

func codexForTesting(t *testing.T) *Codex {
	path := t.TempDir()
	_, err := devconfig.Init(path)
	require.NoError(t, err, "InitConfig should not fail")
	d, err := Open(&devopt.Opts{
		Dir:    path,
		Stderr: os.Stderr,
	})
	require.NoError(t, err, "Open should not fail")

	return d
}
