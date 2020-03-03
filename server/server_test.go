package server_test

import (
	"context"
	"testing"

	"github.com/xackery/eqcp/config"
	"github.com/xackery/eqcp/server"
)

var s *server.Server

func serverSetup(t *testing.T) *server.Server {
	if s != nil {
		return s
	}

	ctx, cancel := context.WithCancel(context.Background())
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		t.Fatal(err)
	}

	s, err = server.New(ctx, cancel, cfg)
	if err != nil {
		t.Fatal(err)
	}
	return s
}
