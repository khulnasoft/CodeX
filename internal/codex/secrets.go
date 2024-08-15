package codex

import (
	"context"

	"github.com/khulnasoft/codex/internal/build"
	"github.com/khulnasoft/codex/lib/envsec/pkg/envsec"
	"github.com/khulnasoft/codex/lib/envsec/pkg/stores/jetstore"
	"github.com/khulnasoft/codex/lib/pkg/envvar"
)

func (d *Codex) UninitializedSecrets(ctx context.Context) *envsec.Envsec {
	return &envsec.Envsec{
		APIHost: build.KhulnasoftAPIHost(),
		Auth: envsec.AuthConfig{
			ClientID: envvar.Get("ENVSEC_CLIENT_ID", build.ClientID()),
			Issuer:   envvar.Get("ENVSEC_ISSUER", build.Issuer()),
		},
		IsDev:      build.IsDev,
		Stderr:     d.stderr,
		Store:      &KhulnasoftAPIStore{},
		WorkingDir: d.ProjectDir(),
	}
}

func (d *Codex) Secrets(ctx context.Context) (*envsec.Envsec, error) {
	envsecInstance := d.UninitializedSecrets(ctx)

	project, err := envsecInstance.ProjectConfig()
	if err != nil {
		return nil, err
	}

	envsecInstance.EnvID = envsec.EnvID{
		EnvName:   d.environment,
		OrgID:     project.OrgID.String(),
		ProjectID: project.ProjectID.String(),
	}

	if _, err := envsecInstance.InitForUser(ctx); err != nil {
		return nil, err
	}

	return envsecInstance, nil
}
