package http

import (
	"context"
	"net/http"

	"github.com/viciousvs/currencyAPI/config"
)

type Server struct {
	Srv *http.Server
}

func (s *Server) Run(cfg config.ServerConfig, handler http.Handler) error {
	s.Srv = &http.Server{
		Addr:           cfg.Addr,
		Handler:        handler,
		MaxHeaderBytes: cfg.MaxHeaderbytes,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
	}
	return s.Srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Srv.Shutdown(ctx)
}
