package internalhttp

import (
	"context"
	"net"
	"net/http"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/storage"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
)

type Server struct {
	server *http.Server
	logger app.Logger
	app    Application
}

type Application interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	GetEvents() storage.EventList
	GetEventById(id string) *storage.Event
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
	eventByIdHandler := NewGetEventById(s.app, s.logger)
	mux.Handle("GET /events/{id}", middleware.loggingMiddleware(eventByIdHandler))

	eventsHandler := NewGetEventsHandler(s.app, s.logger)
	mux.Handle("GET /events", middleware.loggingMiddleware(eventsHandler))

	mux.Handle("/", middleware.loggingMiddleware(&Stub{}))

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
