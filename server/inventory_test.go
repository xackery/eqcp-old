package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/xackery/eqcp/pb"
)

func TestInventorySCRUD(t *testing.T) {
	ctx := context.Background()
	s := serverSetup(t)

	//create
	respC, err := s.InventoryCreate(ctx, &pb.InventoryCreateRequest{Charid: 1, Slotid: 2, Values: map[string]string{"charges": "1"}})
	if err != nil {
		t.Fatal(err)
	}

	if respC.Charid < 1 {
		t.Fatal("response invalid")
	}

	//read
	respR, err := s.InventoryRead(ctx, &pb.InventoryReadRequest{Charid: respC.Charid, Slotid: respC.Slotid})
	if err != nil {
		t.Fatal(err)
	}
	if respR == nil || respR.Inventory == nil {
		t.Fatalf("response empty")
	}
	if respR.Inventory.Charid != respC.Charid {
		t.Fatalf("expected %d, got %d", respR.Inventory.Charid, respC.Charid)
	}

	//search
	respS, err := s.InventorySearch(ctx, &pb.InventorySearchRequest{Values: map[string]string{
		"charid": fmt.Sprintf("%d", respR.Inventory.Charid),
		"slotid": fmt.Sprintf("%d", respR.Inventory.Slotid),
	}})
	if err != nil {
		t.Fatal(err)
	}
	if respS == nil || len(respS.Inventorys) < 1 {
		t.Fatal("search failed to get any results")
	}

	//patch
	respP, err := s.InventoryPatch(ctx, &pb.InventoryPatchRequest{Charid: respC.Charid, Slotid: respC.Slotid, Key: "charges", Value: "2"})
	if err != nil {
		t.Fatal(err)
	}
	if respP == nil || respP.Rowsaffected < 1 {
		t.Fatal("patch failed to get any results")
	}

	//update
	respU, err := s.InventoryUpdate(ctx, &pb.InventoryUpdateRequest{Charid: respC.Charid, Slotid: respC.Slotid, Values: map[string]string{"charges": "3"}})
	if err != nil {
		t.Fatal(err)
	}
	if respU == nil || respP.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}

	//delete
	respD, err := s.InventoryDelete(ctx, &pb.InventoryDeleteRequest{Charid: respC.Charid, Slotid: respC.Slotid})
	if err != nil {
		t.Fatal(err)
	}
	if respD == nil || respD.Rowsaffected < 1 {
		t.Fatal("updated failed to get any results")
	}
}
