package unionserver

import (
	"context"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/app"
	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server"
	internalgrpc "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server/grpc"
	internalhttp "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server/http"
)

type Server struct {
	httpServer *internalhttp.Server
	grpcServer *internalgrpc.Server
	logger     app.Logger
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		err := s.grpcServer.Start(ctx)
		if err != nil {
			s.logger.Fatal(err.Error())
		}
	}()

	return s.httpServer.Start(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.grpcServer.Stop(ctx)
	if err != nil {
		return err
	}

	return s.httpServer.Stop(ctx)
}

func NewServer(
	logger app.Logger,
	app server.Application,
	httpCfg server.HTTPServerConfig,
	grpcCfg server.GRPCServerConfig,
) *Server {
	return &Server{
		logger:     logger,
		httpServer: internalhttp.NewServer(logger, app, httpCfg.BindAddress, httpCfg.BindPort, httpCfg.Timeout),
		grpcServer: internalgrpc.NewServer(app, logger, grpcCfg.BindAddress, grpcCfg.BindPort),
	}
}
