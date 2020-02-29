package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/xackery/eqcp/pb"
)

func TestBugSCRUD(t *testing.T) {
	ctx := context.Background()
	s := serverSetup(t)

	//create
	respC, err := s.BugCreate(ctx, &pb.BugCreateRequest{Values: map[string]string{"lastreviewer": "test"}})
	if err != nil {
		t.Fatal(err)
	}

	if respC.Id < 1 {
		t.Fatal("response invalid")
	}
	id := respC.Id

	//read
	respR, err := s.BugRead(ctx, &pb.BugReadRequest{Id: id})
	if err != nil {
		t.Fatal(err)
	}
	if respR.Bug.Id != id {
		t.Fatalf("expected %d, got %d", respR.Bug.Id, id)
	}

	//search
	respS, err := s.BugSearch(ctx, &pb.BugSearchRequest{Values: map[string]string{
		"lastreviewer": respR.Bug.Lastreviewer,
		"id":           fmt.Sprintf("%d", respR.Bug.Id),
	}})
	if err != nil {
		t.Fatal(err)
	}
	if respS == nil || len(respS.Bugs) < 1 {
		t.Fatal("search failed to get any results")
	}

	//patch
	respP, err := s.BugPatch(ctx, &pb.BugPatchRequest{Id: id, Key: "lastreviewer", Value: "test2"})
	if err != nil {
		t.Fatal(err)
	}
	if respP == nil || respP.Rowsaffected < 1 {
		t.Fatal("patch failed to get any results")
	}

	//update
	respU, err := s.BugUpdate(ctx, &pb.BugUpdateRequest{Values: map[string]string{"lastreviewer": "test3"}})
	if err != nil {
		t.Fatal(err)
	}
	if respU == nil || respP.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}

	//delete
	respD, err := s.BugDelete(ctx, &pb.BugDeleteRequest{Id: id})
	if err != nil {
		t.Fatal(err)
	}
	if respD == nil || respD.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}
}
