// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package featureflag

import (
	"testing"

	"github.com/khulnasoft/codex/internal/envir"
)

func TestEnabledFeature(t *testing.T) {
	name := "TestEnabledFeature"
	enable(name)
	if !features[name].Enabled() {
		t.Errorf("got %s.Enabled() = false, want true.", name)
	}
}

func TestDisabledFeature(t *testing.T) {
	name := "TestDisabledFeature"
	disable(name)
	if features[name].Enabled() {
		t.Errorf("got %s.Enabled() = true, want false.", name)
	}
}

func TestEnabledFeatureEnv(t *testing.T) {
	name := "TestEnabledFeatureEnv"
	disable(name)
	t.Setenv(envir.CodexFeaturePrefix+name, "1")
	if !features[name].Enabled() {
		t.Errorf("got %s.Enabled() = false, want true.", name)
	}
}

func TestNonExistentFeature(t *testing.T) {
	name := "TestNonExistentFeature"
	if features[name].Enabled() {
		t.Errorf("got %s.Enabled() = true, want false.", name)
	}
}
