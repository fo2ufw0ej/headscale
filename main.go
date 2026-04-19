package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"
	// Commit is set at build time via ldflags
	Commit = "none"
	// Date is set at build time via ldflags
	Date = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "headscale",
	Short: "headscale - A self-hosted Tailscale control server",
	Long: `headscale is an open source implementation of the Tailscale
control server that can be self-hosted.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("headscale version %s\n", Version)
		fmt.Printf("  commit : %s\n", Commit)
		fmt.Printf("  date   : %s\n", Date)
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the headscale server",
	RunE:  runServer,
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringP(
		"config",
		"c",
		"",
		"Path to configuration file",
	)

	rootCmd.PersistentFlags().Bool(
		"verbose",
		false,
		"Enable verbose/debug logging",
	)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(serveCmd)
}

func runServer(cmd *cobra.Command, args []string) error {
	// Configure logging
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return err
	}

	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Info().
		Str("version", Version).
		Str("commit", Commit).
		Str("date", Date).
		Msg("Starting headscale")

	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}

	if configFile == "" {
		// Check XDG config dir before falling back to /etc, friendlier for local dev
		if xdgConfig := os.Getenv("XDG_CONFIG_HOME"); xdgConfig != "" {
			configFile = xdgConfig + "/headscale/config.yaml"
		} else {
			configFile = "/etc/headscale/config.yaml"
		}
		log.Info().
			Str("config", configFile).
			Msg("No config file specified, using default")
	}

	log.Info().
		Str("config", configFile).
		Msg("Loading configuration")

	// TODO: Initialize and start the server
	log.Info().Msg("headscale server started")

	return nil
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Failed to execute command")
		os.Exit(1)
	}
}
