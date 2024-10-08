package codex

import (
	"context"
	"errors"
	"io"

	"github.com/samber/lo"
	"github.com/khulnasoft/codex/internal/boxcli/usererr"
	"github.com/khulnasoft/codex/internal/build"
	"github.com/khulnasoft/codex/internal/debug"
	"github.com/khulnasoft/codex/internal/codex/providers/identity"
	"github.com/khulnasoft/codex/internal/codex/providers/nixcache"
	"github.com/khulnasoft/codex/internal/devpkg"
	"github.com/khulnasoft/codex/internal/nix"
	"github.com/khulnasoft/codex/internal/ux"
	"github.com/khulnasoft/codex/lib/pkg/auth"
)

func (d *Codex) UploadProjectToCache(
	ctx context.Context,
	cacheURI string,
) error {
	defer debug.FunctionTimer().End()
	if cacheURI == "" {
		var err error
		cacheURI, err = getWriteCacheURI(ctx, d.stderr)
		if err != nil {
			return err
		}
	}

	creds, err := nixcache.CachedCredentials(ctx)
	if err != nil && !errors.Is(err, auth.ErrNotLoggedIn) {
		return err
	}

	packages := lo.Filter(d.InstallablePackages(), devpkg.IsNix)
	if err != nil || len(packages) == 0 {
		return err
	}

	for _, pkg := range packages {
		inCache, err := pkg.AreAllOutputsInCache(ctx, d.stderr, cacheURI)
		if err != nil {
			return err
		}
		if inCache {
			ux.Finfo(d.stderr, "Package %s is already in cache, skipping\n", pkg.Raw)
			continue
		}
		ux.Finfo(d.stderr, "Uploading package %s to cache\n", pkg.Raw)
		installables, err := pkg.Installables()
		if err != nil {
			return err
		}
		for _, installable := range installables {
			err := nix.CopyInstallableToCache(ctx, d.stderr, cacheURI, installable, creds.Env())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func UploadInstallableToCache(
	ctx context.Context,
	stderr io.Writer,
	cacheURI, installable string,
) error {
	if cacheURI == "" {
		var err error
		cacheURI, err = getWriteCacheURI(ctx, stderr)
		if err != nil {
			return err
		}
	}

	creds, err := nixcache.CachedCredentials(ctx)
	if err != nil && !errors.Is(err, auth.ErrNotLoggedIn) {
		return err
	}
	return nix.CopyInstallableToCache(ctx, stderr, cacheURI, installable, creds.Env())
}

func getWriteCacheURI(
	ctx context.Context,
	w io.Writer,
) (string, error) {
	_, err := identity.GenSession(ctx)
	if errors.Is(err, auth.ErrNotLoggedIn) {
		return "",
			usererr.New("You must be logged in to upload to a Nix cache.")
	}
	caches, err := nixcache.WriteCaches(ctx)
	if err != nil {
		return "", err
	}

	if len(caches) == 0 {
		slug, err := identity.GetOrgSlug(ctx)
		if err != nil {
			return "", err
		}
		return "",
			usererr.New(
				"You don't have permission to write to any Nix caches. To configure cache, go to "+
					"%s/teams/%s/codex",
				build.DashboardHostname(),
				slug,
			)
	}
	if len(caches) > 1 {
		ux.Fwarning(w, "Multiple caches available, using %s.\n", caches[0].GetUri())
	}
	return caches[0].GetUri(), nil
}
