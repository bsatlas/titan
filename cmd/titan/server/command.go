package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Command is the entrypoint for starting the titan registry server.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "server",
		RunE: runE,
	}
	flags := cmd.Flags()
	flags.StringP("config", "c", "config.json", "path to titan server config file")
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	fmt.Println("server command")
	return nil
}
