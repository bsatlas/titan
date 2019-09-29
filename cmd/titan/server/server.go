package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/atlaskerr/titan/http/titan"
)

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
