package server

import (
	"net"

	"github.com/jasonkwh/coffeed/internal/adapter"
	"github.com/jasonkwh/coffeed/proto"
	"github.com/jasonkwh/gatekeeper"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
)

type DaemonServer struct {
	listener   net.Listener
	grpcServer *grpc.Server
}

func NewDaemonServer(
	hdl adapter.DaemonHandler,
) (*DaemonServer, error) {
	var err error

	srv := &DaemonServer{}

	socketPath, err := getSocketPath()
	if err != nil {
		return nil, err
	}

	srv.listener, err = net.Listen("unix", socketPath)
	if err != nil {
		return nil, err
	}

	proto.RegisterValidators()

	srv.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(gatekeeper.UnaryServerInterceptor()),
		grpc.StreamInterceptor(gatekeeper.StreamServerInterceptor()),
	)

	proto.RegisterCoffeedServiceServer(srv.grpcServer, hdl)

	return srv, nil
}

func (s *DaemonServer) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *DaemonServer) Stop() error {
	var err error

	if s.listener != nil {
		err = multierr.Append(err, s.listener.Close())
	}

	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}

	return err
}
