package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/xackery/eqcp/pb"
)

func TestAccountSCRUD(t *testing.T) {
	ctx := context.Background()
	s := serverSetup(t)

	//create
	respC, err := s.AccountCreate(ctx, &pb.AccountCreateRequest{Values: map[string]string{"name": "test"}})
	if err != nil {
		t.Fatal(err)
	}

	if respC.Id < 1 {
		t.Fatal("response invalid")
	}
	id := respC.Id

	//read
	respR, err := s.AccountRead(ctx, &pb.AccountReadRequest{Id: id})
	if err != nil {
		t.Fatal(err)
	}
	if respR.Account.Id != id {
		t.Fatalf("expected %d, got %d", respR.Account.Id, id)
	}

	//search
	respS, err := s.AccountSearch(ctx, &pb.AccountSearchRequest{Values: map[string]string{
		"name": respR.Account.Name,
		"id":   fmt.Sprintf("%d", respR.Account.Id),
	}})
	if err != nil {
		t.Fatal(err)
	}
	if respS == nil || len(respS.Accounts) < 1 {
		t.Fatal("search failed to get any results")
	}

	//patch
	respP, err := s.AccountPatch(ctx, &pb.AccountPatchRequest{Id: id, Key: "name", Value: "test2"})
	if err != nil {
		t.Fatal(err)
	}
	if respP == nil || respP.Rowsaffected < 1 {
		t.Fatal("patch failed to get any results")
	}

	//update
	respU, err := s.AccountUpdate(ctx, &pb.AccountUpdateRequest{Values: map[string]string{"name": "test3"}})
	if err != nil {
		t.Fatal(err)
	}
	if respU == nil || respP.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}

	//delete
	respD, err := s.AccountDelete(ctx, &pb.AccountDeleteRequest{Id: id})
	if err != nil {
		t.Fatal(err)
	}
	if respD == nil || respD.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}
}
