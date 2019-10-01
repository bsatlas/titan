package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/atlaskerr/titan/http/titan"

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

func runE(cmd *cobra.Command, args []string) error {
	s, err := newService()
	if err != nil {
		return fmt.Errorf("failed to initialize service: %s", err)
	}
	err = runServer(s.handlers.titan)
	if err != nil {
		return fmt.Errorf("failed to run server: %s", err)
	}
	return nil
}

func runServer(handler *titan.Server) error {
	ln, err := listener()
	if err != nil {
		return err
	}
	srv := http.Server{
		Handler: handler,
	}
	go srv.Serve(ln)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	sig := <-sigs
	switch sig {
	case syscall.SIGINT:
		srv.Shutdown(context.Background())
	}
	return nil
}

func listener() (net.Listener, error) {
	ln, err := net.Listen("tcp", "0.0.0.0:34557")
	if err != nil {
		return nil, err
	}
	return ln, nil
}

func server() (*titan.Server, error) {
	return &titan.Server{}, nil
}
