// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package nix

import (
	"io"
	"log/slog"
	"strings"
)

// packageInstallIgnore will skip lines that have the strings in the keys of this map.
// The boolean values inform the writer whether to log the line to debug.Log.
var packageInstallIgnore = map[string]bool{
	`replacing old 'codex-development'`: false,
	`installing 'codex-development'`:    false,
}

type PackageInstallWriter struct {
	io.Writer
}

func (fw *PackageInstallWriter) Write(p []byte) (n int, err error) {
	lines := strings.Split(string(p), "\n")
	for _, line := range lines {
		if line != "" && !fw.ignore(line) {
			_, err = io.WriteString(fw.Writer, "\t"+line+"\n")
			if err != nil {
				return n, err
			}
		}
	}
	return len(p), nil
}

func (*PackageInstallWriter) ignore(line string) bool {
	for filter, shouldLog := range packageInstallIgnore {
		if strings.Contains(line, filter) {
			if shouldLog {
				slog.Debug("hiding output from user", "line", line)
			}
			return true
		}
	}
	return false
}
