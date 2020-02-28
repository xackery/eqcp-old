package server

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xackery/eqcp/pb"
	"google.golang.org/grpc"
)

// Server represents a general server
type Server struct {
	ctx     context.Context
	cancel  context.CancelFunc
	mux     *runtime.ServeMux
	gserver *grpc.Server
	gconn   net.Listener
}

// New creates a new server
func New(ctx context.Context, cancel context.CancelFunc, host string) (*Server, error) {
	var err error
	s := &Server{
		ctx:    ctx,
		cancel: cancel,
	}

	s.gconn, err = net.Listen("tcp", ":9090")
	if err != nil {
		return nil, errors.Wrap(err, "net listen")
	}

	opts := []grpc.DialOption{grpc.WithInsecure()}

	s.gserver = grpc.NewServer()
	pb.RegisterEQCPServer(s.gserver, s)
	s.mux = runtime.NewServeMux()
	err = pb.RegisterEQCPHandlerFromEndpoint(ctx, s.mux, host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle server")
	}
	go s.httpServe()
	go s.grpcServe()
	return s, nil
}

func (s *Server) httpServe() {
	err := http.ListenAndServe(":8081", s.mux)
	if err != nil {
		log.Error().Err(err).Msg("http server died")
	}
}
func (s *Server) grpcServe() {
	err := s.gserver.Serve(s.gconn)
	if err != nil {
		log.Error().Err(err).Msg("grpc server died")
	}
}
