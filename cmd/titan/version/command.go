package version

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is the semantic version of the compiled binary.
	Version = "undefined"

	// Commit is the target commit SHA of the compiled binary.
	Commit = "undefined"
)

// Command is the entrypoint for titan version information.
func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "version",
		RunE: runE,
	}
	flags := cmd.Flags()
	flags.Bool("json", false, "format output as JSON")
	return cmd
}

func runE(cmd *cobra.Command, args []string) error {
	flags := cmd.Flags()
	structured, err := flags.GetBool("json")
	if err != nil {
		return err
	}

	if structured {
		obj := map[string]interface{}{
			"version": Version,
			"commit":  Commit,
		}
		b, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", string(b))
		return nil
	}
	fmt.Printf("version: %s\ncommit:  %s\n", Version, Commit)
	return nil
}
