package internalgrpc

import (
	"context"
	"net"

	"google.golang.org/grpc/reflection"

	pb "github.com/viking311/otus_go/hw12_13_14_15_16_calendar/api"

	"google.golang.org/grpc"

	"github.com/viking311/otus_go/hw12_13_14_15_16_calendar/internal/server"
)

type Server struct {
	bindAddress string
	bindPort    string
	app         server.Application
	grpcServer  *grpc.Server
}

func (s *Server) Start(ctx context.Context) error {
	err := s.app.Start(ctx)
	if err != nil {
		return err
	}

	lsn, err := net.Listen("tcp", net.JoinHostPort(s.bindAddress, s.bindPort))
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer()
	pb.RegisterEventServiceServer(s.grpcServer, NewService(s.app))
	reflection.Register(s.grpcServer)

	return s.grpcServer.Serve(lsn)
}

func (s *Server) Stop(_ context.Context) error {

	s.grpcServer.GracefulStop()
	return nil
}

func NewServer(app server.Application, bindAddress, bindPort string) *Server {
	return &Server{
		app:         app,
		bindAddress: bindAddress,
		bindPort:    bindPort,
	}
}
