// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package plugin

import (
	"github.com/khulnasoft/codex/internal/lock"
)

type Manager struct {
	codexProject

	lockfile *lock.File
}

type codexProject interface {
	AllPackageNamesIncludingRemovedTriggerPackages() []string
	ProjectDir() string
}

type managerOption func(*Manager)

func NewManager(opts ...managerOption) *Manager {
	m := &Manager{}
	m.ApplyOptions(opts...)
	return m
}

func WithLockfile(lockfile *lock.File) managerOption {
	return func(m *Manager) {
		m.lockfile = lockfile
	}
}

func WithCodex(provider codexProject) managerOption {
	return func(m *Manager) {
		m.codexProject = provider
	}
}

func (m *Manager) ApplyOptions(opts ...managerOption) {
	for _, opt := range opts {
		opt(m)
	}
}
