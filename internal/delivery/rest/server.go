package rest

import (
	"context"
	"github.com/martynove/gophermart/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	logger     *logrus.Logger
}

func (s *Server) Run(cfg *config.Config, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           cfg.ServerAddress,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
