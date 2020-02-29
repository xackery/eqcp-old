package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/xackery/eqcp/pb"
)

func TestHandinSCRUD(t *testing.T) {
	ctx := context.Background()
	s := serverSetup(t)

	//create
	respC, err := s.HandinCreate(ctx, &pb.HandinCreateRequest{Values: map[string]string{"questid": "1"}})
	if err != nil {
		t.Fatal(err)
	}

	if respC.Id < 1 {
		t.Fatal("response invalid")
	}
	id := respC.Id

	//read
	respR, err := s.HandinRead(ctx, &pb.HandinReadRequest{Id: id})
	if err != nil {
		t.Fatal(err)
	}
	if respR.Handin.Id != id {
		t.Fatalf("expected %d, got %d", respR.Handin.Id, id)
	}

	//search
	respS, err := s.HandinSearch(ctx, &pb.HandinSearchRequest{Values: map[string]string{
		"questid": fmt.Sprintf("%d", respR.Handin.Questid),
		"id":      fmt.Sprintf("%d", respR.Handin.Id),
	}})
	if err != nil {
		t.Fatal(err)
	}
	if respS == nil || len(respS.Handins) < 1 {
		t.Fatal("search failed to get any results")
	}

	//patch
	respP, err := s.HandinPatch(ctx, &pb.HandinPatchRequest{Id: id, Key: "questid", Value: "2"})
	if err != nil {
		t.Fatal(err)
	}
	if respP == nil || respP.Rowsaffected < 1 {
		t.Fatal("patch failed to get any results")
	}

	//update
	respU, err := s.HandinUpdate(ctx, &pb.HandinUpdateRequest{Values: map[string]string{"questid": "3"}})
	if err != nil {
		t.Fatal(err)
	}
	if respU == nil || respP.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}

	//delete
	respD, err := s.HandinDelete(ctx, &pb.HandinDeleteRequest{Id: id})
	if err != nil {
		t.Fatal(err)
	}
	if respD == nil || respD.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}
}
