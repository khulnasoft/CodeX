package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/lo"
)

// updater is a tool that updates all examples/ in the codex repo.
func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

// run will loop over all examples that have run_test script
// run `codex update` on each such example
func run() error {
	codexRepoDir, err := codexRepoDir()
	if err != nil {
		return errors.WithStack(err)
	}
	examplesDir := filepath.Join(codexRepoDir, "examples")

	err = filepath.WalkDir(
		examplesDir, func(path string, d fs.DirEntry, err error) error {
			return walkExampleDir(codexRepoDir, path, d, err)
		},
	)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// examplesToTry is a counter for the number of examples to try. Useful for debugging.
var examplesToTry = 0

func walkExampleDir(codexRepoDir, path string, dirEntry fs.DirEntry, err error) error {
	if err != nil {
		return errors.WithStack(err)
	}

	// Uncomment to try out changes
	// if examplesToTry > 3 {
	//	return nil
	// }
	_ = examplesToTry // silence linter

	// If it is a directory, then we don't continue.
	if dirEntry.IsDir() {
		// Skip if it is a directory that we don't want to process at all.
		skippedDirs := []string{".codex", "node_modules"}
		if lo.Contains(skippedDirs, dirEntry.Name()) {
			return filepath.SkipDir
		}
		return nil
	}

	// If it is not a codex.json file, then we don't continue.
	if dirEntry.Name() != "codex.json" {
		return nil
	}

	// Read the codex.json file
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return errors.WithStack(err)
	}
	content := string(contentBytes)

	// Skip if it doesn't have a run_test script
	if !strings.Contains(content, "run_test") {
		fmt.Printf("SKIP: config at %s lacks run_test\n", path)
		return nil
	}

	// run `codex update` on this example
	codexExecutable := filepath.Join(codexRepoDir, "dist", "codex")
	cmd := exec.Command(codexExecutable, "update", "-c", filepath.Dir(path))
	if err := cmd.Run(); err != nil {
		return errors.WithStack(err)
	}
	fmt.Printf("Ran `codex update` on %s\n", path)
	examplesToTry++

	return nil
}

func codexRepoDir() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	// This file's directory
	fileDir := filepath.Dir(filename)
	// codex repo directory is 3 levels up
	return filepath.Join(fileDir, "..", "..", ".."), nil
}
