package internalhttp

import (
	"context"
	"net"
	"net/http"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server/http/handler"
)

type Server struct {
	server *http.Server
	logger app.Logger
	app    Application
}

type Application interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

func NewServer(logger app.Logger, app Application, cfg HTTPServerConfig) *Server {
	return &Server{
		server: &http.Server{
			Addr:         net.JoinHostPort(cfg.BindAddress, cfg.BindPort),
			ReadTimeout:  cfg.Timeout,
			WriteTimeout: cfg.Timeout,
		},
		logger: logger,
		app:    app,
	}
}

func (s *Server) Start(ctx context.Context) error {
	err := s.app.Start(ctx)
	if err != nil {
		return err
	}

	middleware := NewLoggingMiddleware(s.logger)

	mux := http.NewServeMux()
	mux.Handle("/", middleware.loggingMiddleware(&handler.Stub{}))

	s.server.Handler = mux

	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.app.Stop(ctx)
	if err != nil {
		return err
	}
	return nil
}
