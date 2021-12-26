package main

import (
	"context"
	"fmt"
	"net/http"

	"architecture-lab-3/server/virtual-machines"
)

type HttpPortNumber int

type APIServer struct {
	Port    HttpPortNumber
	vmHandler virtualmachines.HttpHandlerFunc
	server  *http.Server
}

func (s *APIServer) Start() error {
	if s.vmHandler == nil {
		return fmt.Errorf("http handler is undefined")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not specified")
	}

	handler := new(http.ServeMux)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	handler.HandleFunc("/virtualmachines", s.vmHandler)
	return s.server.ListenAndServe()
}

func (s *APIServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
