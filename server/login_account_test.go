package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/xackery/eqcp/pb"
)

func TestLoginAccountSCRUD(t *testing.T) {
	ctx := context.Background()
	s := serverSetup(t)

	//create
	_, err := s.LoginAccountCreate(ctx, &pb.LoginAccountCreateRequest{Values: map[string]string{"id": "1", "accountname": "test"}})
	if err != nil {
		t.Fatal(err)
	}

	id := int64(1)

	//read
	respR, err := s.LoginAccountRead(ctx, &pb.LoginAccountReadRequest{Id: id})
	if err != nil {
		t.Fatal(err)
	}
	if respR.LoginAccount.Id != id {
		t.Fatalf("expected %d, got %d", respR.LoginAccount.Id, id)
	}

	//search
	respS, err := s.LoginAccountSearch(ctx, &pb.LoginAccountSearchRequest{Values: map[string]string{
		"accountname": respR.LoginAccount.Accountname,
		"id":          fmt.Sprintf("%d", respR.LoginAccount.Id),
	}})
	if err != nil {
		t.Fatal(err)
	}
	if respS == nil || len(respS.LoginAccounts) < 1 {
		t.Fatal("search failed to get any results")
	}

	//patch
	respP, err := s.LoginAccountPatch(ctx, &pb.LoginAccountPatchRequest{Id: id, Key: "accountname", Value: "test2"})
	if err != nil {
		t.Fatal(err)
	}
	if respP == nil || respP.Rowsaffected < 1 {
		t.Fatal("patch failed to get any results")
	}

	//update
	respU, err := s.LoginAccountUpdate(ctx, &pb.LoginAccountUpdateRequest{Values: map[string]string{"accountname": "test3"}})
	if err != nil {
		t.Fatal(err)
	}
	if respU == nil || respP.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}

	//delete
	respD, err := s.LoginAccountDelete(ctx, &pb.LoginAccountDeleteRequest{Id: id})
	if err != nil {
		t.Fatal(err)
	}
	if respD == nil || respD.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}
}
