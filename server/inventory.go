package server

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xackery/eqcp/pb"

	//mysql db
	_ "github.com/go-sql-driver/mysql"
)

var (
	inventoryPK = "charid"
	inventorySK = "slotid"
)

// InventorySearch implements SCRUD endpoints
func (s *Server) InventorySearch(ctx context.Context, req *pb.InventorySearchRequest) (*pb.InventorySearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	resp := new(pb.InventorySearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	inventory := new(Inventory)

	st := reflect.TypeOf(*inventory)
	sv := reflect.ValueOf(inventory)
	se := sv.Elem()

	query := "SELECT {fieldMap} FROM inventory WHERE"

	args := map[string]interface{}{}
	comma := ""
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		if req.Orderby != "" && strings.ToLower(field.Name) == strings.ToLower(req.Orderby) {
			req.Orderby = tag
		}

		for key, value := range req.Values {
			if strings.ToLower(field.Name) != strings.ToLower(key) {
				continue
			}

			if se.Field(i).Kind() == reflect.String {
				args[tag] = fmt.Sprintf("%%%s%%", value)
				query += fmt.Sprintf("%s %s LIKE :%s", comma, tag, tag)
			} else {
				args[tag] = value
				query += fmt.Sprintf("%s %s = :%s", comma, tag, tag)
			}

			comma = " AND"
		}

	}
	if len(args) < 1 {
		return nil, fmt.Errorf("no valid fields provided")
	}

	if req.Orderby == "" {
		req.Orderby = inventoryPK
	}
	args["orderby"] = req.Orderby
	query += " ORDER BY :orderby"
	if req.Orderdesc {
		query += " DESC"
	} else {
		query += " ASC"
	}

	args["limit"] = req.Limit
	args["offset"] = req.Offset
	query += " LIMIT :limit OFFSET :offset"

	queryTotal := strings.Replace(query, "{fieldMap}", "count(charid) as total", 1)
	query = strings.Replace(query, "{fieldMap}", "*", 1)

	rows, err := s.db.NamedQueryContext(ctx, queryTotal, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	for rows.Next() {
		inventory := new(Inventory)
		err = rows.StructScan(inventory)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = inventory.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		inventory := new(Inventory)
		err = rows.StructScan(inventory)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Inventorys = append(resp.Inventorys, inventory.ToProto())
	}

	return resp, nil
}

// InventoryCreate implements SCRUD endpoints
func (s *Server) InventoryCreate(ctx context.Context, req *pb.InventoryCreateRequest) (*pb.InventoryCreateResponse, error) {

	inventory := new(Inventory)

	st := reflect.TypeOf(*inventory)

	args := map[string]interface{}{}
	query := "INSERT INTO inventory"

	comma := ""
	insertField := ""
	insertValue := ""
	req.Values["charid"] = fmt.Sprintf("%d", req.Charid)
	req.Values["slotid"] = fmt.Sprintf("%d", req.Slotid)

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		for key, value := range req.Values {
			if strings.ToLower(field.Name) != strings.ToLower(key) {
				continue
			}
			args[tag] = value
			insertField += fmt.Sprintf("%s %s", comma, tag)
			insertValue += fmt.Sprintf("%s :%s", comma, tag)
			comma = ","
		}
	}
	if len(args) < 1 {
		return nil, fmt.Errorf("no valid fields provided")
	}

	query += fmt.Sprintf(" (%s) VALUES(%s)", insertField, insertValue)

	log.Debug().Interface("args", args).Msgf("query: %s", query)

	_, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	resp := new(pb.InventoryCreateResponse)
	resp.Charid = req.Charid
	resp.Slotid = req.Slotid

	return resp, nil
}

// InventoryRead implements SCRUD endpoints
func (s *Server) InventoryRead(ctx context.Context, req *pb.InventoryReadRequest) (*pb.InventoryReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	if req.Charid < 1 {
		return nil, fmt.Errorf("charid must be greater than 0")
	}
	if req.Slotid < 1 {
		return nil, fmt.Errorf("slotid must be greater than 0")
	}
	resp := new(pb.InventoryReadResponse)

	if req.Charid < 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	query := "SELECT * FROM inventory WHERE "

	args := map[string]interface{}{}
	query += fmt.Sprintf("%s = :%s AND %s = :%s", inventoryPK, inventoryPK, inventorySK, inventorySK)
	args[inventoryPK] = req.Charid
	args[inventorySK] = req.Slotid

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		inventory := new(Inventory)
		err = rows.StructScan(inventory)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Inventory = inventory.ToProto()
	}
	return resp, nil
}

// InventoryUpdate implements SCRUD endpoints
func (s *Server) InventoryUpdate(ctx context.Context, req *pb.InventoryUpdateRequest) (*pb.InventoryUpdateResponse, error) {
	if req.Charid < 1 {
		return nil, fmt.Errorf("charid must be greater than 0")
	}
	if req.Slotid < 1 {
		return nil, fmt.Errorf("slotid must be greater than 0")
	}
	inventory := new(Inventory)

	st := reflect.TypeOf(*inventory)

	args := map[string]interface{}{
		inventoryPK: req.Charid,
		inventorySK: req.Slotid,
	}
	query := "UPDATE inventory SET"

	comma := ""
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		for key, value := range req.Values {
			if strings.ToLower(field.Name) != strings.ToLower(key) {
				continue
			}
			args[tag] = value

			query += fmt.Sprintf("%s %s = :%s", comma, tag, tag)
			comma = ","
		}

	}
	if len(args) == 1 {
		return nil, fmt.Errorf("no valid fields provided")
	}

	query += fmt.Sprintf(" WHERE %s = :%s AND %s = :%s LIMIT 1", inventoryPK, inventoryPK, inventorySK, inventorySK)

	log.Debug().Interface("args", args).Msgf("query: %s", query)

	result, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rowsaffected")
	}
	resp := new(pb.InventoryUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// InventoryDelete implements SCRUD endpoints
func (s *Server) InventoryDelete(ctx context.Context, req *pb.InventoryDeleteRequest) (*pb.InventoryDeleteResponse, error) {
	if req.Charid < 1 {
		return nil, fmt.Errorf("charid must be greater than 0")
	}
	if req.Slotid < 1 {
		return nil, fmt.Errorf("slotid must be greater than 0")
	}
	query := fmt.Sprintf("DELETE FROM inventory WHERE %s = :%s AND %s = :%s LIMIT 1", inventoryPK, inventoryPK, inventorySK, inventorySK)

	args := map[string]interface{}{
		inventoryPK: req.Charid,
		inventorySK: req.Slotid,
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)

	result, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rowsaffected")
	}
	resp := new(pb.InventoryDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// InventoryPatch implements SCRUD endpoints
func (s *Server) InventoryPatch(ctx context.Context, req *pb.InventoryPatchRequest) (*pb.InventoryPatchResponse, error) {
	inventory := new(Inventory)

	if req.Charid < 1 {
		return nil, fmt.Errorf("charid must be greater than 0")
	}
	if req.Slotid < 1 {
		return nil, fmt.Errorf("slotid must be greater than 0")
	}
	st := reflect.TypeOf(*inventory)

	args := map[string]interface{}{
		inventoryPK: req.Charid,
		inventorySK: req.Slotid,
	}
	query := "UPDATE inventory SET"

	comma := ""
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		if strings.ToLower(field.Name) != strings.ToLower(req.Key) {
			continue
		}
		args[tag] = req.Value

		query += fmt.Sprintf("%s %s = :%s", comma, tag, tag)
		comma = ","
	}
	if len(args) == 2 {
		return nil, fmt.Errorf("no valid fields provided")
	}

	query += fmt.Sprintf(" WHERE %s = :%s AND %s = :%s LIMIT 1", inventoryPK, inventoryPK, inventorySK, inventorySK)
	log.Debug().Interface("args", args).Msgf("query: %s", query)

	result, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rowsaffected")
	}
	resp := new(pb.InventoryPatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// Inventory represents an INVENTORY DB binding
type Inventory struct {
	Charid            int64          `db:"charid"`              // int(11) unsigned NOT NULL DEFAULT '0',
	Slotid            int64          `db:"slotid"`              // mediumint(7) unsigned NOT NULL DEFAULT '0',
	Itemid            sql.NullInt64  `db:"itemid"`              // int(11) unsigned DEFAULT '0',
	Charges           int64          `db:"charges"`             // smallint(3) unsigned DEFAULT '0',
	Color             int64          `db:"color"`               // int(11) unsigned NOT NULL DEFAULT '0',
	Augslot1          int64          `db:"augslot1"`            // mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot2          int64          `db:"augslot2"`            // mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot3          int64          `db:"augslot3"`            // mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot4          int64          `db:"augslot4"`            // mediumint(7) unsigned NOT NULL DEFAULT '0',
	Augslot5          int64          `db:"augslot5"`            // mediumint(7) unsigned DEFAULT '0',
	Augslot6          int64          `db:"augslot6"`            // mediumint(7) NOT NULL DEFAULT '0',
	Instnodrop        int64          `db:"instnodrop"`          // tinyint(1) unsigned NOT NULL DEFAULT '0',
	Customdata        sql.NullString `db:"custom_data"`         // text,
	Ornamenticon      int64          `db:"ornamenticon"`        // int(11) unsigned NOT NULL DEFAULT '0',
	Ornamentidfile    int64          `db:"ornamentidfile"`      // int(11) unsigned NOT NULL DEFAULT '0',
	Ornamentheromodel int64          `db:"ornament_hero_model"` // int(11) NOT NULL DEFAULT '0',

	Total int64 `db:"total"`
}

// ToProto converts the inventory type struct to protobuf
func (i *Inventory) ToProto() *pb.Inventory {
	inventory := &pb.Inventory{}

	inventory.Charid = i.Charid
	inventory.Slotid = i.Slotid
	inventory.Itemid = i.Itemid.Int64
	inventory.Charges = i.Charges
	inventory.Color = i.Color
	inventory.Augslot1 = i.Augslot1
	inventory.Augslot2 = i.Augslot2
	inventory.Augslot3 = i.Augslot3
	inventory.Augslot4 = i.Augslot4
	inventory.Augslot5 = i.Augslot5
	inventory.Augslot6 = i.Augslot6
	inventory.Instnodrop = i.Instnodrop
	inventory.Customdata = i.Customdata.String
	inventory.Ornamenticon = i.Ornamenticon
	inventory.Ornamentidfile = i.Ornamentidfile
	inventory.Ornamentheromodel = i.Ornamentheromodel
	return inventory
}
