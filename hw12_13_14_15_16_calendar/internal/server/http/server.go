package internalhttp

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server/http/middleware"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server/http/handler"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
)

type Server struct {
	server *http.Server
	logger app.Logger
	app    server.Application
}

func NewServer(logger app.Logger, app server.Application, bindAddress, bindPort string, timeout time.Duration) *Server {
	return &Server{
		server: &http.Server{
			Addr:         net.JoinHostPort(bindAddress, bindPort),
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
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

	LoggerMiddleware := middleware.NewLoggingMiddleware(s.logger)

	mux := http.NewServeMux()
	eventByIdHandler := handler.NewGetEventById(s.app, s.logger)
	mux.Handle("GET /events/{id}", LoggerMiddleware.LoggingMiddleware(eventByIdHandler))

	deleteEventHandler := handler.NewDeleteEventHandler(s.app, s.logger)
	mux.Handle("DELETE /events/{id}", LoggerMiddleware.LoggingMiddleware(deleteEventHandler))

	eventsHandler := handler.NewGetEventsHandler(s.app, s.logger)
	mux.Handle("GET /events", LoggerMiddleware.LoggingMiddleware(eventsHandler))

	saveEventHandler := handler.NewSaveEventHandler(s.app, s.logger)
	mux.Handle("POST /events", LoggerMiddleware.LoggingMiddleware(saveEventHandler))

	getEventsByUserHandler := handler.NewGetEventsByUserHandler(s.app, s.logger)
	mux.Handle("GET /users/{userId}/events", LoggerMiddleware.LoggingMiddleware(getEventsByUserHandler))

	getEventsByUserAndDatesHandker := handler.NewGetEventsByUserAndDatesHandler(s.app, s.logger)
	mux.Handle("GET /users/{userId}/events/{dateFrom}/{dateTo}", LoggerMiddleware.LoggingMiddleware(getEventsByUserAndDatesHandker))

	mux.Handle("/", LoggerMiddleware.LoggingMiddleware(&handler.Stub{}))

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
