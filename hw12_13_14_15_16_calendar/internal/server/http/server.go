package internalhttp

import (
	"context"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
)

type Server struct {
	logger app.Logger
	app    Application
}

type Application interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

func NewServer(logger app.Logger, app Application) *Server {
	return &Server{
		logger: logger,
		app:    app,
	}
}

func (s *Server) Start(ctx context.Context) error {
	err := s.app.Start(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.app.Stop(ctx)
	if err != nil {
		return err
	}

	return nil
}

// TODO
