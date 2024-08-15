// Copyright 2024 Khulnasoft Ltd. and contributors. All rights reserved.
// Use of this source code is governed by the license in the LICENSE file.

package boxcli

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/khulnasoft/codex/internal/build"
	"github.com/khulnasoft/codex/internal/codex"
	"github.com/khulnasoft/codex/internal/codex/devopt"
	"github.com/khulnasoft/codex/internal/codex/providers/identity"
	"github.com/khulnasoft/codex/internal/ux"
	"github.com/khulnasoft/codex/lib/pkg/api"
)

func authCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Codex auth commands",
	}

	cmd.AddCommand(loginCmd())
	cmd.AddCommand(logoutCmd())
	cmd.AddCommand(whoAmICmd())
	cmd.AddCommand(authNewTokenCommand())

	return cmd
}

func loginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login to codex",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := identity.AuthClient(identity.AuthRedirectDefault)
			if err != nil {
				return err
			}
			t, err := c.LoginFlow()
			if err != nil {
				return err
			}
			// TODO: all uses of IDClaims() are broken when using a static
			// non-expiring token (i.e. API_TOKEN)
			fmt.Fprintf(cmd.ErrOrStderr(), "Logged in as: %s\n", t.IDClaims().Email)
			return nil
		},
	}

	return cmd
}

func logoutCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Logout from codex",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := identity.AuthClient(identity.AuthRedirectDefault)
			if err != nil {
				return err
			}
			err = c.LogoutFlow()
			if err == nil {
				fmt.Fprintln(cmd.OutOrStdout(), "Logged out successfully")
			}
			return err
		},
	}

	return cmd
}

type whoAmICmdFlags struct {
	showTokens bool
}

func whoAmICmd() *cobra.Command {
	flags := &whoAmICmdFlags{}
	cmd := &cobra.Command{
		Use:   "whoami",
		Short: "Show the current user",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			box, err := codex.Open(&devopt.Opts{Dir: wd, Stderr: cmd.ErrOrStderr()})
			if err != nil {
				return err
			}
			return box.UninitializedSecrets(cmd.Context()).
				WhoAmI(cmd.Context(), cmd.OutOrStdout(), flags.showTokens)
		},
	}

	cmd.Flags().BoolVar(
		&flags.showTokens,
		"show-tokens",
		false,
		"Show the access, id, and refresh tokens",
	)

	return cmd
}

func authNewTokenCommand() *cobra.Command {
	tokensCmd := &cobra.Command{
		Use:   "tokens",
		Short: "Manage codex auth tokens",
	}

	newCmd := &cobra.Command{
		Use:   "new",
		Short: "Create a new token",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			token, err := identity.GenSession(ctx)
			if err != nil {
				return err
			}
			client := api.NewClient(ctx, build.KhulnasoftAPIHost(), token)
			pat, err := client.CreateToken(ctx)
			if err != nil {
				// This is a hack because errors are not returning with correct code.
				// Once that is fixed, we can switch to use *connect.Error Code() instead.
				if strings.Contains(err.Error(), "permission_denied") {
					ux.Ferror(
						cmd.ErrOrStderr(),
						"You do not have permission to create a token. Please contact your"+
							" administrator.",
					)
					return nil
				}
				return err
			}
			ux.Fsuccess(cmd.OutOrStdout(), "Token created.\n\n")
			table := tablewriter.NewWriter(cmd.OutOrStdout())
			table.SetRowLine(true)
			table.AppendBulk([][]string{
				{"Token ID", pat.GetToken().GetId()},
				{"Secret", pat.GetToken().GetSecret()},
			})
			table.Render()
			return nil
		},
	}

	tokensCmd.AddCommand(newCmd)

	return tokensCmd
}
