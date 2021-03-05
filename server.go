package course

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s Server) Run(port string, handle http.Handler) error {
	s.server = &http.Server{
		Addr:           ":" + port,
		Handler:        handle,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	return s.server.ListenAndServe()
}

func (s Server) ShutDown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
