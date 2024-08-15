// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package envir

const (
	CodexCache         = "CODEX_CACHE"
	CodexFeaturePrefix = "CODEX_FEATURE_"
	CodexGateway       = "CODEX_GATEWAY"
	// CodexLatestVersion is the latest version available of the codex CLI binary.
	// NOTE: it should NOT start with v (like 0.4.8)
	CodexLatestVersion  = "CODEX_LATEST_VERSION"
	CodexRegion         = "CODEX_REGION"
	CodexSearchHost     = "CODEX_SEARCH_HOST"
	CodexShellEnabled   = "CODEX_SHELL_ENABLED"
	CodexShellStartTime = "CODEX_SHELL_START_TIME"
	CodexVM             = "CODEX_VM"

	LauncherVersion = "LAUNCHER_VERSION"
	LauncherPath    = "LAUNCHER_PATH"

	GitHubUsername = "GITHUB_USER_NAME"
	SSHTTY         = "SSH_TTY"

	XDGDataHome   = "XDG_DATA_HOME"
	XDGConfigHome = "XDG_CONFIG_HOME"
	XDGCacheHome  = "XDG_CACHE_HOME"
	XDGStateHome  = "XDG_STATE_HOME"
)

// system
const (
	Env   = "ENV"
	Home  = "HOME"
	Path  = "PATH"
	Shell = "SHELL"
	User  = "USER"
)
