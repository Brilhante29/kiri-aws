// Package main is the entry point for the kiri CLI.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	kiri "github.com/Brilhante29/kiri-aws"
	kiricli "github.com/Brilhante29/kiri-aws/cli"
	_ "github.com/Brilhante29/kiri-aws/internal/registry" // Register all services via init().
	"github.com/Brilhante29/kiri-aws/internal/server"
)

// Build metadata. GoReleaser overrides these at link time
// (-X main.version=... -X main.commit=... -X main.date=...); a plain `go build`
// falls back to the version constant baked into the module.
var (
	version = kiri.Version
	commit  = "none"
	date    = "unknown"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the kiri version, commit, and build date",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			_, err := fmt.Fprintf(cmd.OutOrStdout(), "kiri %s (commit %s, built %s)\n", version, commit, date)
			if err != nil {
				return fmt.Errorf("write version: %w", err)
			}

			return nil
		},
	}
}

func main() {
	root := kiricli.NewRootCmd()

	// Root command starts the server when no CLI subcommand is matched.
	// Docker uses `kiri --host 0.0.0.0 --port 4566`, so we accept these flags.
	root.RunE = func(cmd *cobra.Command, _ []string) error {
		cfg := server.DefaultConfig()

		if host, _ := cmd.Flags().GetString("host"); host != "" {
			cfg.Host = host
		}

		if cmd.Flags().Changed("port") {
			cfg.Port, _ = cmd.Flags().GetInt("port")
		}

		srv := server.New(cfg)

		if err := srv.Run(); err != nil {
			return fmt.Errorf("server failed: %w", err)
		}

		return nil
	}

	// Server flags live on each command that actually starts the server, not
	// on root.PersistentFlags, so client subcommands (s3, acm, ...) do not
	// inherit them.
	addServerFlags := func(c *cobra.Command) {
		c.Flags().String("host", "", "Server host (overrides KIRI_HOST)")
		c.Flags().Int("port", 0, "Server port (overrides KIRI_PORT)")
	}

	addServerFlags(root)

	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the kiri server",
		RunE:  root.RunE,
	}
	addServerFlags(serveCmd)

	root.AddCommand(serveCmd)
	root.AddCommand(newVersionCmd())
	root.Version = version

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
