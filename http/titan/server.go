package titan

import (
	"net/http"
)

// Server is the main titan HTTP server.
type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {}
