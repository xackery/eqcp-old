package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xackery/eqcp/pb"
	"github.com/xackery/eqemuconfig"
	"google.golang.org/grpc"
)

// Server represents a general server
type Server struct {
	ctx     context.Context
	cancel  context.CancelFunc
	mux     *runtime.ServeMux
	gserver *grpc.Server
	gconn   net.Listener
	cfg     *eqemuconfig.Config
	db      *sqlx.DB
}

// New creates a new server
func New(ctx context.Context, cancel context.CancelFunc, host string, cfg *eqemuconfig.Config) (*Server, error) {
	var err error
	s := &Server{
		ctx:    ctx,
		cancel: cancel,
		cfg:    cfg,
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Db)
	s.db, err = sqlx.Open("mysql", conn)
	if err != nil {
		return nil, errors.Wrap(err, "sql open")
	}
	log.Debug().Msgf("connected to %s:%s %s", cfg.Database.Host, cfg.Database.Port, cfg.Database.Db)

	log.Debug().Msgf("grpc listening on %s", host)
	s.gconn, err = net.Listen("tcp", host)
	if err != nil {
		return nil, errors.Wrap(err, "net listen")
	}

	opts := []grpc.DialOption{grpc.WithInsecure()}

	s.gserver = grpc.NewServer()
	pb.RegisterAccountServiceServer(s.gserver, s)
	pb.RegisterBugServiceServer(s.gserver, s)
	pb.RegisterCharacterServiceServer(s.gserver, s)
	pb.RegisterHandinServiceServer(s.gserver, s)
	pb.RegisterItemServiceServer(s.gserver, s)
	pb.RegisterNpcServiceServer(s.gserver, s)
	pb.RegisterPetitionServiceServer(s.gserver, s)
	pb.RegisterPlayerSpeechServiceServer(s.gserver, s)
	pb.RegisterTradeServiceServer(s.gserver, s)
	pb.RegisterZoneServiceServer(s.gserver, s)
	s.mux = runtime.NewServeMux()

	err = pb.RegisterAccountServiceHandlerFromEndpoint(ctx, s.mux, host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle bug")
	}
	err = pb.RegisterBugServiceHandlerFromEndpoint(ctx, s.mux, host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle bug")
	}
	err = pb.RegisterCharacterServiceHandlerFromEndpoint(ctx, s.mux, host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle character")
	}
	err = pb.RegisterHandinServiceHandlerFromEndpoint(ctx, s.mux, host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle handin")
	}
	err = pb.RegisterItemServiceHandlerFromEndpoint(ctx, s.mux, host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle item")
	}
	err = pb.RegisterNpcServiceHandlerFromEndpoint(ctx, s.mux, host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle npc")
	}
	err = pb.RegisterPetitionServiceHandlerFromEndpoint(ctx, s.mux, host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle petition")
	}
	err = pb.RegisterPlayerSpeechServiceHandlerFromEndpoint(ctx, s.mux, host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle playerspeech")
	}
	err = pb.RegisterTradeServiceHandlerFromEndpoint(ctx, s.mux, host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle trade")
	}
	err = pb.RegisterZoneServiceHandlerFromEndpoint(ctx, s.mux, host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle zone")
	}

	go s.httpServe()
	go s.grpcServe()
	return s, nil
}

// Close closes the server
func (s *Server) Close() {
	s.db.Close()
}

func (s *Server) httpServe() {

	log.Debug().Str("url", fmt.Sprintf("http://127.0.0.1:8081")).Msgf("rest listening on %s", ":8081")
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
