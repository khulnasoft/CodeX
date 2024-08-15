// Copyright 2024 Khulnasoft Inc. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package shellgen

import "path/filepath"

func genPath(d codexer) string {
	return filepath.Join(d.ProjectDir(), ".codex/gen")
}

func FlakePath(d codexer) string {
	return filepath.Join(genPath(d), "flake")
}
