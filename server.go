package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (server *Server) Run(port string, handler http.Handler) error {
	server.httpServer = &http.Server{
		Addr:           "127.0.0.1:" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return server.httpServer.ListenAndServe()
}

func (server *Server) ShutDown(ctx context.Context) error {
	return server.httpServer.Shutdown(ctx)
}
