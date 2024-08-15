// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package lock

type codexProject interface {
	ConfigHash() (string, error)
	NixPkgsCommitHash() string
	AllPackageNamesIncludingRemovedTriggerPackages() []string
	ProjectDir() string
}

type Locker interface {
	Get(string) *Package
	LegacyNixpkgsPath(string) string
	ProjectDir() string
	Resolve(string) (*Package, error)
}
