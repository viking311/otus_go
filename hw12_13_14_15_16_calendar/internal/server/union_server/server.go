package unionserver

import (
	"context"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	internalhttp "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server/http"
)

type Server struct {
	httpServer *internalhttp.Server
}

func (s *Server) Start(ctx context.Context) error {
	return s.httpServer.Start(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Stop(ctx)
}

func NewServer(logger app.Logger, app server.Application, httpCfg server.HTTPServerConfig) *Server {
	return &Server{
		httpServer: internalhttp.NewServer(logger, app, httpCfg.BindAddress, httpCfg.BindPort, httpCfg.Timeout),
	}
}
