package main

import (
	"fmt"
	"os"

	"github.com/atlaskerr/titan/cmd/titan/server"
	"github.com/atlaskerr/titan/cmd/titan/version"

	"github.com/spf13/cobra"
)

func command() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "titan",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	subCommands := []func() *cobra.Command{
		server.Command,
		version.Command,
	}
	for _, subCommand := range subCommands {
		cmd.AddCommand(subCommand())
	}
	return cmd
}

func execute() {
	err := command().Execute()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
