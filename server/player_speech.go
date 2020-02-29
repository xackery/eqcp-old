package server

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xackery/eqcp/pb"

	//mysql db
	_ "github.com/go-sql-driver/mysql"
)

// PlayerSpeechSearch implements SCRUD endpoints
func (s *Server) PlayerSpeechSearch(ctx context.Context, req *pb.PlayerSpeechSearchRequest) (*pb.PlayerSpeechSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	resp := new(pb.PlayerSpeechSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	playerSpeech := new(PlayerSpeech)

	st := reflect.TypeOf(*playerSpeech)
	sv := reflect.ValueOf(playerSpeech)
	se := sv.Elem()

	query := "SELECT {fieldMap} FROM qs_player_speech WHERE"

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
		req.Orderby = "id"
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

	queryTotal := strings.Replace(query, "{fieldMap}", "count(id) as total", 1)
	query = strings.Replace(query, "{fieldMap}", "*", 1)

	rows, err := s.db.NamedQueryContext(ctx, queryTotal, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	for rows.Next() {
		playerSpeech := new(PlayerSpeech)
		err = rows.StructScan(playerSpeech)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = playerSpeech.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		playerSpeech := new(PlayerSpeech)
		err = rows.StructScan(playerSpeech)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.PlayerSpeechs = append(resp.PlayerSpeechs, playerSpeech.ToProto())
	}

	return resp, nil
}

// PlayerSpeechCreate implements SCRUD endpoints
func (s *Server) PlayerSpeechCreate(ctx context.Context, req *pb.PlayerSpeechCreateRequest) (*pb.PlayerSpeechCreateResponse, error) {

	fmt.Println(req)
	playerSpeech := new(PlayerSpeech)

	st := reflect.TypeOf(*playerSpeech)

	args := map[string]interface{}{}
	query := "INSERT INTO qs_player_speech"

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
			if strings.ToLower(tag) != strings.ToLower(key) {
				continue
			}
			args[tag] = value
			insertField += fmt.Sprintf("%s `%s`", comma, tag)
			insertValue += fmt.Sprintf("%s :%s", comma, tag)
			comma = ","
		}
	}
	if len(args) == 1 {
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

	resp := new(pb.PlayerSpeechCreateResponse)
	resp.Id = lastID

	return resp, nil
}

// PlayerSpeechRead implements SCRUD endpoints
func (s *Server) PlayerSpeechRead(ctx context.Context, req *pb.PlayerSpeechReadRequest) (*pb.PlayerSpeechReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	resp := new(pb.PlayerSpeechReadResponse)

	if req.Id < 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	query := "SELECT * FROM qs_player_speech WHERE "

	args := map[string]interface{}{}
	query += "id = :id"
	args["id"] = req.Id

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		playerSpeech := new(PlayerSpeech)
		err = rows.StructScan(playerSpeech)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.PlayerSpeech = playerSpeech.ToProto()
	}
	return resp, nil
}

// PlayerSpeechUpdate implements SCRUD endpoints
func (s *Server) PlayerSpeechUpdate(ctx context.Context, req *pb.PlayerSpeechUpdateRequest) (*pb.PlayerSpeechUpdateResponse, error) {
	playerSpeech := new(PlayerSpeech)

	st := reflect.TypeOf(*playerSpeech)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE qs_player_speech SET"

	comma := ""
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		for key, value := range req.Values {
			if strings.ToLower(tag) != strings.ToLower(key) {
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

	query += " WHERE id = :id LIMIT 1"

	log.Debug().Interface("args", args).Msgf("query: %s", query)

	result, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rowsaffected")
	}
	resp := new(pb.PlayerSpeechUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// PlayerSpeechDelete implements SCRUD endpoints
func (s *Server) PlayerSpeechDelete(ctx context.Context, req *pb.PlayerSpeechDeleteRequest) (*pb.PlayerSpeechDeleteResponse, error) {
	query := "DELETE FROM qs_player_speech WHERE id = :id LIMIT 1"

	args := map[string]interface{}{
		"id": req.Id,
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
	resp := new(pb.PlayerSpeechDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// PlayerSpeechPatch implements SCRUD endpoints
func (s *Server) PlayerSpeechPatch(ctx context.Context, req *pb.PlayerSpeechPatchRequest) (*pb.PlayerSpeechPatchResponse, error) {
	playerSpeech := new(PlayerSpeech)

	st := reflect.TypeOf(*playerSpeech)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE qs_player_speech SET"

	comma := ""
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		if strings.ToLower(field.Name) != strings.ToLower(req.Key) {
			continue
		}

		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		args[tag] = req.Value

		query += fmt.Sprintf("%s %s = :%s", comma, tag, tag)
		comma = ","
	}
	if len(args) == 1 {
		return nil, fmt.Errorf("no valid fields provided")
	}

	query += " WHERE id = :id LIMIT 1"
	log.Debug().Interface("args", args).Msgf("query: %s", query)

	result, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rowsaffected")
	}
	resp := new(pb.PlayerSpeechPatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// PlayerSpeech reports
type PlayerSpeech struct {
	ID           int64     `db:"id"`           // int(11) NOT NULL AUTO_INCREMENT,
	From         string    `db:"from"`         // varchar(64) NOT NULL,
	To           string    `db:"to"`           // varchar(64) NOT NULL,
	Frommessage  string    `db:"message"`      // varchar(256) NOT NULL,
	Minstatus    int64     `db:"minstatus"`    // smallint(5) NOT NULL,
	Guilddbid    int64     `db:"guilddbid"`    // int(11) NOT NULL,
	Type         int64     `db:"type"`         // tinyint(3) NOT NULL,
	Timerecorded time.Time `db:"timerecorded"` // timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	Total        int64     `db:"total"`
}

// ToProto converts the playerSpeech type struct to protobuf
func (p *PlayerSpeech) ToProto() *pb.PlayerSpeech {
	playerSpeech := &pb.PlayerSpeech{}
	playerSpeech.Id = p.ID
	playerSpeech.From = p.From
	playerSpeech.To = p.To
	playerSpeech.Frommessage = p.Frommessage
	playerSpeech.Minstatus = p.Minstatus
	playerSpeech.Guilddbid = p.Guilddbid
	playerSpeech.Type = p.Type
	playerSpeech.Timerecorded = p.Timerecorded.Unix()
	return playerSpeech
}
