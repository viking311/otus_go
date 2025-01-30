package http

import (
	"context"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/infrastructure/repository"
)

type Server struct {
	logger  Logger
	storage repository.EventRepositoryExtInterface
}

type Logger interface { // TODO
}

type Application interface { // TODO
}

func NewServer(logger Logger, storage repository.EventRepositoryExtInterface) *Server {
	return &Server{
		logger:  logger,
		storage: storage,
	}
}

func (s *Server) Start(ctx context.Context) error {
	err := s.storage.Connect(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.storage.Close(ctx)
	if err != nil {
		return err
	}

	return nil
}

// TODO
