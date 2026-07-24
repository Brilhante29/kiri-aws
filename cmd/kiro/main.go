// Package main is the entry point for the kiro CLI.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	kirocli "github.com/kiro-aws/kiro-aws/cli"
	_ "github.com/kiro-aws/kiro-aws/internal/registry" // Register all services via init().
	"github.com/kiro-aws/kiro-aws/internal/server"
)

func main() {
	root := kirocli.NewRootCmd()

	// Root command starts the server when no CLI subcommand is matched.
	// Docker uses `kiro --host 0.0.0.0 --port 4566`, so we accept these flags.
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
		c.Flags().String("host", "", "Server host (overrides KIRO_HOST)")
		c.Flags().Int("port", 0, "Server port (overrides KIRO_PORT)")
	}

	addServerFlags(root)

	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the kiro server",
		RunE:  root.RunE,
	}
	addServerFlags(serveCmd)

	root.AddCommand(serveCmd)

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
