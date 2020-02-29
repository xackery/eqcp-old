package server_test

import (
	"context"
	"testing"

	"github.com/xackery/eqcp/server"
	"github.com/xackery/eqemuconfig"
)

var s *server.Server

func serverSetup(t *testing.T) *server.Server {
	if s != nil {
		return s
	}
	cfg, err := eqemuconfig.GetConfig()
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	s, err = server.New(ctx, cancel, "127.0.0.1:9090", cfg)
	if err != nil {
		t.Fatal(err)
	}
	return s
}
