// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package codex

import (
	"context"
	"runtime/trace"

	"github.com/khulnasoft/codex/internal/codex/devopt"
	"github.com/khulnasoft/codex/internal/pullbox"
)

func (d *Codex) Pull(ctx context.Context, opts devopt.PullboxOpts) error {
	ctx, task := trace.NewTask(ctx, "codexPull")
	defer task.End()
	return pullbox.New(d, opts).Pull(ctx)
}

func (d *Codex) Push(ctx context.Context, opts devopt.PullboxOpts) error {
	ctx, task := trace.NewTask(ctx, "codexPush")
	defer task.End()
	return pullbox.New(d, opts).Push(ctx)
}
