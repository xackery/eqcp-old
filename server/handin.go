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

// HandinSearch implements SCRUD endpoints
func (s *Server) HandinSearch(ctx context.Context, req *pb.HandinSearchRequest) (*pb.HandinSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	ap, err := s.AuthFromContext(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}

	if !ap.hasCommand("petitioninfo") {
		return nil, fmt.Errorf("you do not have permissions to this endpoint")
	}

	resp := new(pb.HandinSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	handin := new(Handin)

	st := reflect.TypeOf(*handin)
	sv := reflect.ValueOf(handin)
	se := sv.Elem()

	query := "SELECT {fieldMap} FROM qs_player_handin_record WHERE"

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
		req.Orderby = "handin_id"
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

	queryTotal := strings.Replace(query, "{fieldMap}", "count(handin_id) as total", 1)
	query = strings.Replace(query, "{fieldMap}", "*", 1)

	rows, err := s.db.NamedQueryContext(ctx, queryTotal, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	for rows.Next() {
		handin := new(Handin)
		err = rows.StructScan(handin)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = handin.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		handin := new(Handin)
		err = rows.StructScan(handin)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Handins = append(resp.Handins, handin.ToProto())
	}

	return resp, nil
}

// HandinCreate implements SCRUD endpoints
func (s *Server) HandinCreate(ctx context.Context, req *pb.HandinCreateRequest) (*pb.HandinCreateResponse, error) {
	ap, err := s.AuthFromContext(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}

	if !ap.hasCommand("mysql") {
		return nil, fmt.Errorf("you do not have permissions to this endpoint")
	}
	handin := new(Handin)

	st := reflect.TypeOf(*handin)

	args := map[string]interface{}{}
	query := "INSERT INTO qs_player_handin_record"

	comma := ""
	insertField := ""
	insertValue := ""
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

	result, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "lastinsertedid")
	}

	resp := new(pb.HandinCreateResponse)
	resp.Id = lastID

	return resp, nil
}

// HandinRead implements SCRUD endpoints
func (s *Server) HandinRead(ctx context.Context, req *pb.HandinReadRequest) (*pb.HandinReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	ap, err := s.AuthFromContext(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}

	if !ap.hasCommand("petitioninfo") {
		return nil, fmt.Errorf("you do not have permissions to this endpoint")
	}

	resp := new(pb.HandinReadResponse)

	if req.Id < 0 {
		return nil, fmt.Errorf("Id must be greater than 0")
	}
	query := "SELECT * FROM qs_player_handin_record WHERE "

	args := map[string]interface{}{}
	query += "handin_id = :handin_id"
	args["handin_id"] = req.Id

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		handin := new(Handin)
		err = rows.StructScan(handin)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Handin = handin.ToProto()
	}
	return resp, nil
}

// HandinUpdate implements SCRUD endpoints
func (s *Server) HandinUpdate(ctx context.Context, req *pb.HandinUpdateRequest) (*pb.HandinUpdateResponse, error) {
	handin := new(Handin)
	ap, err := s.AuthFromContext(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}

	if !ap.hasCommand("mysql") {
		return nil, fmt.Errorf("you do not have permissions to this endpoint")
	}

	st := reflect.TypeOf(*handin)

	args := map[string]interface{}{
		"handin_id": req.Id,
	}
	query := "UPDATE qs_player_handin_record SET"

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

	query += " WHERE handin_id = :handin_id LIMIT 1"

	log.Debug().Interface("args", args).Msgf("query: %s", query)

	result, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rowsaffected")
	}
	resp := new(pb.HandinUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// HandinDelete implements SCRUD endpoints
func (s *Server) HandinDelete(ctx context.Context, req *pb.HandinDeleteRequest) (*pb.HandinDeleteResponse, error) {
	ap, err := s.AuthFromContext(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}

	if !ap.hasCommand("mysql") {
		return nil, fmt.Errorf("you do not have permissions to this endpoint")
	}
	query := "DELETE FROM qs_player_handin_record WHERE handin_id = :handin_id LIMIT 1"

	args := map[string]interface{}{
		"handin_id": req.Id,
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
	resp := new(pb.HandinDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// HandinPatch implements SCRUD endpoints
func (s *Server) HandinPatch(ctx context.Context, req *pb.HandinPatchRequest) (*pb.HandinPatchResponse, error) {
	ap, err := s.AuthFromContext(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}

	if !ap.hasCommand("mysql") {
		return nil, fmt.Errorf("you do not have permissions to this endpoint")
	}
	handin := new(Handin)

	st := reflect.TypeOf(*handin)

	args := map[string]interface{}{
		"handin_id": req.Id,
	}
	query := "UPDATE qs_player_handin_record SET"

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
	if len(args) == 1 {
		return nil, fmt.Errorf("no valid fields provided")
	}

	query += " WHERE handin_id = :handin_id LIMIT 1"
	log.Debug().Interface("args", args).Msgf("query: %s", query)

	result, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rowsaffected")
	}
	resp := new(pb.HandinPatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// Handin represents an HANDIN DB binding
type Handin struct {
	ID        int64        `db:"handin_id"`  //int(11) NOT NULL AUTO_INCREMENT,
	Time      sql.NullTime `db:"time"`       //timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
	Questid   int64        `db:"quest_id"`   //int(11) DEFAULT '0',
	Charid    int64        `db:"char_id"`    //int(11) DEFAULT '0',
	Charpp    int64        `db:"char_pp"`    //int(11) DEFAULT '0',
	Chargp    int64        `db:"char_gp"`    //int(11) DEFAULT '0',
	Charsp    int64        `db:"char_sp"`    //int(11) DEFAULT '0',
	Charcp    int64        `db:"char_cp"`    //int(11) DEFAULT '0',
	Charitems int64        `db:"char_items"` //mediumint(7) DEFAULT '0',
	Npcid     int64        `db:"npc_id"`     //int(11) DEFAULT '0',
	Npcpp     int64        `db:"npc_pp"`     //int(11) DEFAULT '0',
	Npcgp     int64        `db:"npc_gp"`     //int(11) DEFAULT '0',
	Npcsp     int64        `db:"npc_sp"`     //int(11) DEFAULT '0',
	Npccp     int64        `db:"npc_cp"`     //int(11) DEFAULT '0',
	Npcitems  int64        `db:"npc_items"`  //mediumint(7) DEFAULT '0',
	Total     int64        `db:"total"`
}

// ToProto converts the handin type struct to protobuf
func (h *Handin) ToProto() *pb.Handin {
	handin := &pb.Handin{}
	handin.Id = h.ID
	handin.Time = h.Time.Time.Unix()
	handin.Questid = h.Questid
	handin.Charid = h.Charid
	handin.Charpp = h.Charpp
	handin.Chargp = h.Chargp
	handin.Charsp = h.Charsp
	handin.Charcp = h.Charcp
	handin.Charitems = h.Charitems
	handin.Npcid = h.Npcid
	handin.Npcpp = h.Npcpp
	handin.Npcgp = h.Npcgp
	handin.Npcsp = h.Npcsp
	handin.Npccp = h.Npccp
	handin.Npcitems = h.Npcitems
	return handin
}
