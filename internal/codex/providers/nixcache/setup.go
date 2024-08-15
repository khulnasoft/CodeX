package nixcache

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/khulnasoft/codex/internal/envir"
	"github.com/khulnasoft/codex/internal/nix"
	"github.com/khulnasoft/codex/internal/redact"
	"github.com/khulnasoft/codex/internal/setup"
	"github.com/khulnasoft/codex/internal/ux"
)

const setupKey = "nixcache-setup"

func IsConfigured(ctx context.Context) bool {
	u, err := user.Current()
	if err != nil {
		return false
	}
	task := &setupTask{u.Username}
	status := setup.Status(ctx, setupKey, task)
	return status == setup.TaskDone
}

func Configure(ctx context.Context) error {
	u, err := user.Current()
	if err != nil {
		return redact.Errorf("nixcache: lookup current user: %v", err)
	}

	task := &setupTask{u.Username}

	// This function might be called from other Codex commands
	// (such as codex add), so we need to provide some context in the sudo
	// prompt.
	const sudoPrompt = "You're logged into a Codex account, but Nix isn't setup to use your account's caches. " +
		"Allow sudo to configure Nix?"
	err = setup.ConfirmRun(ctx, setupKey, task, sudoPrompt)
	if err != nil {
		return redact.Errorf("nixcache: run setup: %w", err)
	}
	return nil
}

func ConfigureReprompt(ctx context.Context, username string) error {
	setup.Reset(setupKey)
	task := &setupTask{username}

	// We're reprompting, so the user explicitly asked to configure the
	// cache. We can keep the sudo prompt short.
	err := setup.ConfirmRun(ctx, setupKey, task, "Allow sudo to configure Nix?")
	if err != nil {
		return redact.Errorf("nixcache: run setup: %w", err)
	}
	return nil
}

// setupTask adds the user to Nix's trusted-users list and updates
// ~root/.aws/config so that they can use their Codex cache with the
// Nix daemon.
type setupTask struct {
	// username is the OS username to trust.
	username string
}

func (s *setupTask) NeedsRun(ctx context.Context, lastRun setup.RunInfo) bool {
	if _, err := nix.DaemonVersion(ctx); err != nil {
		// This looks like a single-user install, so no need to
		// configure the daemon or root's AWS credentials.
		slog.Error("nixcache: skipping setup: error connecting to nix daemon, assuming single-user install", "err", err)
		return false
	}

	if lastRun.Time.IsZero() {
		slog.Debug("nixcache: running setup: first time setup")
		return true
	}
	cfg, err := nix.CurrentConfig(ctx)
	if err != nil {
		slog.Error("nixcache: running setup: error getting current nix config, assuming user isn't trusted", "user", s.username)
		return true
	}
	trusted, err := cfg.IsUserTrusted(ctx, s.username)
	if err != nil {
		slog.Error("nixcache: running setup: error checking if user is trusted, assuming they aren't", "user", s.username)
		return true
	}
	if !trusted {
		slog.Debug("nixcache: running setup: user isn't trusted", "user", s.username)
		return true
	}
	return false
}

func (s *setupTask) Run(ctx context.Context) error {
	ran, err := setup.SudoCodex(ctx, "cache", "configure", "--user", s.username)
	if ran || err != nil {
		return err
	}

	// Update the AWS config before configuring and restarting the Nix
	// daemon.
	err = s.updateAWSConfig()
	if err != nil {
		return redact.Errorf("update root aws config: %v", err)
	}

	trusted := false
	cfg, err := nix.CurrentConfig(ctx)
	if err == nil {
		trusted, _ = cfg.IsUserTrusted(ctx, s.username)
	}
	if !trusted {
		err = nix.IncludeCodexConfig(ctx, s.username)
		if errors.Is(err, nix.ErrUnknownServiceManager) {
			ux.Fwarning(os.Stderr, "Codex configured Nix to use a new cache. Please restart the Nix daemon and re-run Codex.\n")
		} else if err != nil {
			return redact.Errorf("update nix config: %v", err)
		}
	}
	return nil
}

func (s *setupTask) updateAWSConfig() error {
	exe, err := codexExecutable()
	if err != nil {
		return err
	}
	sudo, err := sudoExecutable()
	if err != nil {
		return err
	}
	configPath, err := rootAWSConfigPath()
	if err != nil {
		return err
	}

	// Clear out and backup any existing .aws directory. We need to
	// do this with the entire directory and not just .aws/config
	// because there are other files that can affect credentials.
	backup, err := backupDirectory(filepath.Dir(configPath))
	if err != nil {
		return err
	}

	flag := os.O_WRONLY | os.O_CREATE | os.O_EXCL
	perm := fs.FileMode(0o644)
	config, err := os.OpenFile(configPath, flag, perm)
	if errors.Is(err, os.ErrNotExist) {
		// Avoid os.MkdirAll because we shouldn't be creating anything
		// above the user's home directory.
		if err = os.Mkdir(filepath.Dir(configPath), 0o755); err != nil {
			return redact.Errorf("create ~root/.aws directory: %v", err)
		}
		config, err = os.OpenFile(configPath, flag, perm)
	}
	if err != nil {
		return redact.Errorf("open ~root/.aws/config: %v", err)
	}
	defer config.Close()

	// TODO(gcurtis): it would be nice to use a non-default profile
	// if https://github.com/NixOS/nix/issues/5525 gets fixed.
	header := "# This file was generated by Codex.\n"
	if backup != "" {
		header += "# The old .aws directory was moved to " + backup + ".\n"
	}
	_, err = fmt.Fprintf(config, `%s
[default]
# sudo as the configured user so that their cached credential files have the
# correct ownership.
credential_process = %s -u %s -i %s-- %s cache credentials
`, header, sudo, s.username, propagatedEnv(), exe)
	if err != nil {
		return redact.Errorf("write to ~root/.aws/config: %v", err)
	}
	if err := config.Close(); err != nil {
		return redact.Errorf("close ~root/.aws/config: %v", err)
	}
	return nil
}

// propagatedEnv returns a string of space-separated VAR=value pairs of
// environment variables that should be propagated to the credential_process
// command in ~root/.aws/config. This is especially important for CI because the
// Nix daemon won't otherwise see any environment variables set by the job.
func propagatedEnv() string {
	envs := []string{
		"CODEX_API_TOKEN",
		"CODEX_PROD",
		"CODEX_USE_VERSION",
		"XDG_CACHE_HOME",
		"XDG_CONFIG_DIRS",
		"XDG_CONFIG_HOME",
		"XDG_DATA_DIRS",
		"XDG_DATA_HOME",
		"XDG_RUNTIME_DIR",
		"XDG_STATE_HOME",
	}
	strb := strings.Builder{}
	for _, name := range envs {
		val := os.Getenv(name)
		if val == "" {
			continue
		}
		notPrintable := strings.ContainsFunc(val, func(r rune) bool {
			return !unicode.IsPrint(r)
		})
		if notPrintable {
			slog.Debug("nixcache: not including environment variable in ~root/.aws/config because it contains nonprintable runes: %q=%q", name, val)
			continue
		}

		strb.WriteString(name)
		strb.WriteString(`="`)
		for _, r := range val {
			switch r {
			// Special characters inside double quotes:
			// https://pubs.opengroup.org/onlinepubs/009604499/utilities/xcu_chap02.html#tag_02_02_03
			case '$', '`', '"', '\\':
				strb.WriteByte('\\')
			}
			strb.WriteRune(r)
		}
		strb.WriteString(`" `)
	}
	return strb.String()
}

// rootAWSConfigPath returns the default AWS config path for the root user. In a
// shell this is ~root/.aws/config.
func rootAWSConfigPath() (string, error) {
	u, err := user.LookupId("0")
	if err != nil {
		return "", redact.Errorf("lookup root user: %s", err)
	}
	if u.HomeDir == "" {
		return "", redact.Errorf("empty root user home directory: %s", u.Username, err)
	}
	return filepath.Join(u.HomeDir, ".aws", "config"), nil
}

// backupDirectory creates a backup of a directory and then deletes it. Upon
// success, it returns the path to the backup copy.
func backupDirectory(path string) (string, error) {
	// Remember this function is running as root, so be careful when
	// moving/creating/deleting things.

	path = filepath.Clean(path)
	if path == "/" {
		return "", redact.Errorf("refusing to backup root directory")
	}

	backup := fmt.Sprintf("%s-%d.bak", path, time.Now().Unix())
	err := os.Rename(path, backup)
	if errors.Is(err, os.ErrNotExist) {
		// No pre-existing .aws directory.
		return "", nil
	}
	if err != nil {
		return "", redact.Errorf("backup existing directory %s: %v", path, err)
	}
	return backup, nil
}

// codexExecutable returns the path to the Codex launcher script or the
// current binary if the launcher is unavailable.
func codexExecutable() (string, error) {
	if exe := os.Getenv(envir.LauncherPath); exe != "" {
		if abs, err := filepath.Abs(exe); err == nil {
			return abs, nil
		}
	}

	exe, err := os.Executable()
	if err != nil {
		return "", redact.Errorf("get path to codex executable: %v", err)
	}
	return exe, nil
}

// sudoExecutable searches the PATH for sudo.
func sudoExecutable() (string, error) {
	sudo, err := exec.LookPath("sudo")
	if err != nil {
		return "", redact.Errorf("get path to sudo executable: %v", err)
	}
	return sudo, nil
}
