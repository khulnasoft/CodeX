// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package openssh

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/pkg/errors"
	"github.com/khulnasoft/codex/internal/fileutil"
)

// These must match what's in sshConfigTmpl. We should eventually make the hosts
// a template variable.
const (
	gatewayProdHost = "gateway.codex.khulnasoft.com"
	gatewayDevHost  = "gateway.dev.codex.khulnasoft.com"
)

//go:embed sshconfig.tmpl
var sshConfigText string
var sshConfigTmpl = template.Must(template.New("sshconfig").Parse(sshConfigText))

//go:embed known_hosts
var sshKnownHosts []byte

// SetupCodex updates the user's OpenSSH configuration so that they can connect
// to Codex Cloud hosts. It does nothing if Codex Cloud is already
// configured.
func SetupCodex() error {
	return setupCodex("", 0)
}

// SetupInsecureDebug is like SetupCodex, but also configures an additional
// gateway with host key checking disabled. If gatewayAddr is a
// well-known *.codex.khulnasoft.com gateway, then SetupInsecureDebug doesn't add any
// extra hosts and acts identically to SetupCodex.
func SetupInsecureDebug(gatewayAddr string) error {
	host, port := splitHostPort(gatewayAddr)
	if host != gatewayProdHost && host != gatewayDevHost {
		return setupCodex(host, port)
	}
	return setupCodex("", 0)
}

func setupCodex(debugHost string, debugPort int) error {
	codexSSHDir, err := codexSSHDir()
	if err != nil {
		return err
	}

	// Ensure ~/.config/codex/ssh/sockets exists.
	if _, err := codexSocketsDir(); err != nil {
		return err
	}

	// Try to remove any old debug host keys. It's okay if this fails.
	codexKnownHostsDebug := filepath.Join(codexSSHDir, "known_hosts_debug")
	_ = os.Remove(codexKnownHostsDebug)

	codexKnownHostsPath := filepath.Join(codexSSHDir, "known_hosts")
	codexKnownHosts, err := editFile(codexKnownHostsPath, 0o644)
	if err != nil {
		return err
	}
	defer codexKnownHosts.Close()
	if _, err := codexKnownHosts.Write(sshKnownHosts); err != nil {
		return err
	}
	if err := codexKnownHosts.Commit(); err != nil {
		return err
	}

	codexIncludePath := filepath.Join(codexSSHDir, "config")
	codexSSHConfig, err := editFile(codexIncludePath, 0o644)
	if err != nil {
		return err
	}
	defer codexSSHConfig.Close()

	tmplData := struct {
		ConfigVersion string
		ConfigDir     string
		DebugGateway  struct {
			Host string
			Port int
		}
	}{
		ConfigVersion: "0.0.1",
		ConfigDir:     codexSSHDir,
	}
	tmplData.DebugGateway.Host = debugHost
	tmplData.DebugGateway.Port = debugPort
	err = errors.WithStack(sshConfigTmpl.Execute(codexSSHConfig, tmplData))
	if err != nil {
		return errors.WithStack(err)
	}
	if err := codexSSHConfig.Commit(); err != nil {
		return err
	}
	if err := updateUserSSHConfig(codexIncludePath); err != nil {
		return err
	}

	// Create the known_hosts_debug file with the correct permissions if a
	// debug gateway is configured. It's okay if this fails because it's
	// only used for debugging.
	if debugHost != "" {
		f, err := os.OpenFile(codexKnownHostsDebug, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err == nil {
			f.Close()
		}
	}
	return nil
}

// AddVMKey sets the private SSH key for the given Codex VM host. If a key was
// previously set for the host, AddVMKey replaces it with the new key. The old
// key is not recoverable.
//
// AddVMKey only manages keys specific to Codex Cloud. It will not touch any of
// the user's keys in ~/.ssh.
func AddVMKey(hostname, key string) error {
	keysDir, err := codexKeysDir()
	if err != nil {
		return err
	}
	keyFile, err := editFile(filepath.Join(keysDir, hostname), 0o600)
	if err != nil {
		return err
	}
	defer keyFile.Close()

	if _, err := io.WriteString(keyFile, key); err != nil {
		return err
	}
	return keyFile.Commit()
}

func updateUserSSHConfig(codexIncludePath string) (err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return errors.WithStack(err)
	}
	dotSSH := filepath.Join(home, ".ssh")
	if err := EnsureDirExists(dotSSH, 0o700, true); err != nil {
		return err
	}

	sshConfig, err := editFile(filepath.Join(dotSSH, "config"), 0o644)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := sshConfig.Close()
		if err == nil {
			err = closeErr
		}
	}()

	bufw := bufio.NewWriter(sshConfig)
	_, err = fmt.Fprintf(bufw, "Include \"%s\"\n", codexIncludePath)
	if err != nil {
		return err
	}

	// Look for an existing Include directive, copying the file contents as
	// we read.
	if containsCodexInclude(io.TeeReader(sshConfig, bufw)) {
		// We found an existing Include - don't save and return.
		return nil
	}
	// We didn't find an existing Include - copy the rest of the user's SSH
	// config and then commit the changes.
	if _, err := bufw.ReadFrom(sshConfig); err != nil {
		return errors.WithStack(err)
	}
	if err := bufw.Flush(); err != nil {
		return errors.WithStack(err)
	}
	return sshConfig.Commit()
}

var (
	reCodexInclude = regexp.MustCompile(`(?i)^[ \t]*"?Include"?[ \t=][^#]*codex/ssh/config`)
	reHostOrMatch   = regexp.MustCompile(`(?i)[ \t]*"?(Host|Match) `)
)

func containsCodexInclude(r io.Reader) bool {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		if reCodexInclude.Match(line) {
			return true
		}

		// Unconditional Include directives must come before any Host or
		// Match blocks. If we found one of those blocks then we've gone
		// too far.
		if reHostOrMatch.Match(line) {
			return false
		}
	}
	return false
}

func EnsureDirExists(path string, perm fs.FileMode, chmod bool) error {
	return fileutil.EnsureDirExists(path, perm, chmod)
}

// returns path to ~/.config/codex/ssh
func codexSSHDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", errors.WithStack(err)
	}

	// Ensure ~/.config exists but don't touch existing permissions.
	dotConfig := filepath.Join(home, ".config")
	if err := EnsureDirExists(dotConfig, 0o755, false); err != nil {
		return "", err
	}

	// Ensure ~/.config/codex exists and force permissions to 0755.
	codexConfigDir := filepath.Join(dotConfig, "codex")
	if err := EnsureDirExists(codexConfigDir, 0o755, true); err != nil {
		return "", err
	}

	// Ensure ~/.config/codex/ssh exists and force permissions to 0700.
	codexSSHDir := filepath.Join(codexConfigDir, "ssh")
	if err := EnsureDirExists(codexSSHDir, 0o700, true); err != nil {
		return "", err
	}
	return codexSSHDir, nil
}

func codexKeysDir() (string, error) {
	sshDir, err := codexSSHDir()
	if err != nil {
		return "", err
	}
	keysDir := filepath.Join(sshDir, "keys")
	if err := EnsureDirExists(keysDir, 0o700, true); err != nil {
		return "", err
	}
	return keysDir, nil
}

func codexSocketsDir() (string, error) {
	sshDir, err := codexSSHDir()
	if err != nil {
		return "", err
	}
	sockets := filepath.Join(sshDir, "sockets")
	if err := EnsureDirExists(sockets, 0o700, true); err != nil {
		return "", err
	}
	return sockets, nil
}

// atomicEdit reads from a source file and writes changes to a separate
// temporary file. Upon a call to Commit, it atomically overwrites the source
// file with the temp file, guaranteeing that all of the file Writes succeed or
// none at all. Calling Close before calling Commit discards any written data,
// leaving the source file untouched.
type atomicEdit struct {
	path     string
	editFile *os.File
	tmpFile  *os.File

	closed bool
	err    error
}

// editFile opens the file at path for editing. Writes to atomicEdit will not
// modify the file until Commit is called. If the file doesn't exist, calls to
// Read immediately return io.EOF and Commit will create it with permissions
// perm. If the file does exist, Commit atomically applies any written data and
// changes its permissions to perm.
//
// Calling Close without calling Commit discards all written data. It is
// unnecessary but valid to call Close after Commit. This makes it easier to
// defer closing the file.
func editFile(path string, perm os.FileMode) (*atomicEdit, error) {
	// editFile will be nil when creating a new file.
	editFile, err := os.Open(path)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, errors.WithStack(err)
	}

	// Atomic file renames require that both files are on the same volume.
	// Putting the tmp file in the same directory is the best way to ensure
	// that happens.
	tmp, err := os.CreateTemp(filepath.Dir(path), ".codex")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Make sure to set permissions before writing anything. This also means
	// perm must be user-writeable.
	if err := tmp.Chmod(perm); err != nil {
		return nil, errors.WithStack(err)
	}
	return &atomicEdit{
		path:     path,
		editFile: editFile,
		tmpFile:  tmp,
	}, nil
}

func (a *atomicEdit) Read(p []byte) (n int, err error) {
	if a.editFile == nil {
		return 0, io.EOF
	}
	n, err = a.editFile.Read(p)

	// Don't use `errors.Is` here because we only want to avoid wrapping
	// io.EOF directly. This is for compatibility with the io.Writer
	// interface.
	// nolint:errorlint
	if err != nil && err != io.EOF {
		err = errors.WithStack(err)
	}
	return n, err
}

func (a *atomicEdit) Write(p []byte) (n int, err error) {
	n, err = a.tmpFile.Write(p)

	// Don't use `errors.Is` here because we only want to avoid wrapping
	// io.EOF directly. This is for compatibility with the io.Writer
	// interface.
	// nolint:errorlint
	if err != nil && err != io.EOF {
		err = errors.WithStack(err)
	}
	return n, err
}

func (a *atomicEdit) Commit() error {
	if a.closed {
		return a.err
	}
	a.closed = true

	if a.editFile != nil {
		// Ignore close errors - we only ever read from the original
		// file.
		a.editFile.Close()
	}
	if a.err = errors.WithStack(a.tmpFile.Close()); a.err != nil {
		return a.err
	}
	if a.err = errors.WithStack(os.Rename(a.tmpFile.Name(), a.path)); a.err != nil {
		return a.err
	}
	return a.err
}

func (a *atomicEdit) Close() error {
	if a.closed {
		return a.err
	}
	a.closed = true

	// Ignore close errors - we're throwing away any changes.
	if a.editFile != nil {
		a.editFile.Close()
	}
	a.tmpFile.Close()
	a.err = errors.WithStack(os.Remove(a.tmpFile.Name()))
	return a.err
}
