package server

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Command is the entrypoint for starting the titan registry server.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "server",
		RunE: runE,
	}
	setFlags(cmd.Flags())
	return cmd
}

func setFlags(flags *pflag.FlagSet) {
	// config file flag
	flags.StringP(
		"config",
		"c",
		"config.json",
		"path to titan server config file",
	)
}

func configFromFlags(flags *pflag.FlagSet) {}

func runE(cmd *cobra.Command, args []string) error {
	srv, err := server()
	if err != nil {
		return err
	}

	err = runServer(srv)
	if err != nil {
		return err
	}
	return nil
}
