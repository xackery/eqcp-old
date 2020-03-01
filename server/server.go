package server

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xackery/eqcp/config"
	"github.com/xackery/eqcp/pb"
	"google.golang.org/grpc"
)

// Server represents a general server
type Server struct {
	ctx        context.Context
	cancel     context.CancelFunc
	mux        *runtime.ServeMux
	gserver    *grpc.Server
	gconn      net.Listener
	cfg        *config.Config
	db         *sqlx.DB
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

// New creates a new server
func New(ctx context.Context, cancel context.CancelFunc, cfg *config.Config) (*Server, error) {
	var err error
	s := &Server{
		ctx:    ctx,
		cancel: cancel,
		cfg:    cfg,
	}

	signBytes, err := ioutil.ReadFile(cfg.Jwt.PrivateKeyPath)
	if err != nil {
		return nil, errors.Wrapf(err, "read jwt.private_key_path")
	}

	s.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return nil, errors.Wrapf(err, "parse jwt.private_key_path")
	}

	signBytes, err = ioutil.ReadFile(cfg.Jwt.PublicKeyPath)
	if err != nil {
		return nil, errors.Wrapf(err, "read jwt.public_key_path")
	}

	s.publicKey, err = jwt.ParseRSAPublicKeyFromPEM(signBytes)
	if err != nil {
		return nil, errors.Wrapf(err, "parse jwt.public_key_path")
	}

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Db)
	s.db, err = sqlx.Open("mysql", conn)
	if err != nil {
		return nil, errors.Wrap(err, "sql open")
	}
	log.Debug().Msgf("sql connected to %s:%s %s", cfg.Database.Host, cfg.Database.Port, cfg.Database.Db)

	log.Debug().Msgf("grpc listening on %s", cfg.Grpc.Host)
	s.gconn, err = net.Listen("tcp", cfg.Grpc.Host)
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
	pb.RegisterLoginAccountServiceServer(s.gserver, s)
	pb.RegisterNpcServiceServer(s.gserver, s)
	pb.RegisterPetitionServiceServer(s.gserver, s)
	pb.RegisterPlayerSpeechServiceServer(s.gserver, s)
	pb.RegisterLoginServerServiceServer(s.gserver, s)
	pb.RegisterTradeServiceServer(s.gserver, s)
	pb.RegisterZoneServiceServer(s.gserver, s)
	s.mux = runtime.NewServeMux()

	err = pb.RegisterAccountServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle bug")
	}
	err = pb.RegisterBugServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle bug")
	}
	err = pb.RegisterCharacterServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle character")
	}
	err = pb.RegisterHandinServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle handin")
	}
	err = pb.RegisterItemServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle item")
	}
	err = pb.RegisterLoginAccountServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle loginaccount")
	}
	err = pb.RegisterNpcServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle npc")
	}
	err = pb.RegisterPetitionServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle petition")
	}
	err = pb.RegisterPlayerSpeechServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle playerspeech")
	}
	err = pb.RegisterLoginServerServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle server")
	}
	err = pb.RegisterTradeServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
	if err != nil {
		return nil, errors.Wrap(err, "handle trade")
	}
	err = pb.RegisterZoneServiceHandlerFromEndpoint(ctx, s.mux, cfg.Grpc.Host, opts)
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
	log.Debug().Msgf("api listening on %s", s.cfg.API.Host)
	err := http.ListenAndServe(s.cfg.API.Host, s.mux)
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
