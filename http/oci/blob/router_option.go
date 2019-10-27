package blob

import (
	"errors"
	"net/http"
)

// RouterOption applies a parameter to a Server.
type RouterOption func(*Router)

// OptionDownloadHandler sets the http.Handler to use for the unknown endpoint.
func OptionDownloadHandler(handler http.Handler) RouterOption {
	var fn RouterOption = func(s *Router) {
		s.download = handler
	}
	return fn
}

// OptionExistHandler sets the http.Handler to use for the unknown endpoint.
func OptionExistHandler(handler http.Handler) RouterOption {
	var fn RouterOption = func(s *Router) {
		s.exist = handler
	}
	return fn
}

// OptionRemoveHandler sets the http.Handler to use for the unknown endpoint.
func OptionRemoveHandler(handler http.Handler) RouterOption {
	var fn RouterOption = func(s *Router) {
		s.remove = handler
	}
	return fn
}

// OptionUploadHandler sets the http.Handler to use for the unknown endpoint.
func OptionUploadHandler(handler http.Handler) RouterOption {
	var fn RouterOption = func(s *Router) {
		s.upload = handler
	}
	return fn
}

// OptionUndefinedHandler sets the http.Handler to use for the unknown endpoint.
func OptionUndefinedHandler(handler http.Handler) RouterOption {
	var fn RouterOption = func(s *Router) {
		s.undefined = handler
	}
	return fn
}

// NewRouter returns a fully initialized Server.
func NewRouter(options ...RouterOption) (*Router, error) {
	srv := &Router{}
	for _, addOption := range options {
		addOption(srv)
	}
	if srv.undefined == nil {
		return nil, errors.New("no unknown handler defined")
	}
	return srv, nil
}
