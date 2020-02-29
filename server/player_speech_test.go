package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xackery/eqcp/pb"
	"github.com/xackery/eqcp/server"
	"github.com/xackery/eqemuconfig"
)

func TestPlayerSpeechSCRUD(t *testing.T) {
	assert := assert.New(t)
	cfg, err := eqemuconfig.GetConfig()
	if !assert.NoError(err) {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	s, err := server.New(ctx, cancel, "127.0.0.1:9090", cfg)
	if !assert.NoError(err) {
		t.Fatal(err)
	}

	//create
	respC, err := s.PlayerSpeechCreate(ctx, &pb.PlayerSpeechCreateRequest{Values: map[string]string{
		"from":    "fromtest",
		"to":      "totest",
		"message": "messageTest",
	}})
	if !assert.NoError(err) {
		t.Fatal(err)
	}

	if respC.Id < 1 {
		t.Fatal("response invalid")
	}
	id := respC.Id

	//read
	respR, err := s.PlayerSpeechRead(ctx, &pb.PlayerSpeechReadRequest{Id: id})
	if err != nil {
		t.Fatal(err)
	}
	if respR.PlayerSpeech.Id != id {
		t.Fatalf("expected %d, got %d", respR.PlayerSpeech.Id, id)
	}

	//search
	respS, err := s.PlayerSpeechSearch(ctx, &pb.PlayerSpeechSearchRequest{Values: map[string]string{
		"minstatus": fmt.Sprintf("%d", respR.PlayerSpeech.Minstatus),
		"id":        fmt.Sprintf("%d", respR.PlayerSpeech.Id),
	}})
	if err != nil {
		t.Fatal(err)
	}
	if respS == nil || len(respS.PlayerSpeechs) < 1 {
		t.Fatal("search failed to get any results")
	}

	//patch
	respP, err := s.PlayerSpeechPatch(ctx, &pb.PlayerSpeechPatchRequest{Id: id, Key: "minstatus", Value: "1"})
	if err != nil {
		t.Fatal(err)
	}
	if respP == nil || respP.Rowsaffected < 1 {
		t.Fatal("patch failed to get any results")
	}

	//update
	respU, err := s.PlayerSpeechUpdate(ctx, &pb.PlayerSpeechUpdateRequest{Values: map[string]string{
		"minstatus": "2",
	}})
	if err != nil {
		t.Fatal(err)
	}
	if respU == nil || respP.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}

	//delete
	respD, err := s.PlayerSpeechDelete(ctx, &pb.PlayerSpeechDeleteRequest{Id: id})
	if err != nil {
		t.Fatal(err)
	}
	if respD == nil || respD.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}
}
