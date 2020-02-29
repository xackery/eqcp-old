package server

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xackery/eqcp/pb"

	//mysql db
	_ "github.com/go-sql-driver/mysql"
)

// PetitionSearch implements SCRUD endpoints
func (s *Server) PetitionSearch(ctx context.Context, req *pb.PetitionSearchRequest) (*pb.PetitionSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	resp := new(pb.PetitionSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	petition := new(Petition)

	st := reflect.TypeOf(*petition)
	sv := reflect.ValueOf(petition)
	se := sv.Elem()

	query := "SELECT {fieldMap} FROM petitions WHERE"

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
		petition := new(Petition)
		err = rows.StructScan(petition)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = petition.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		petition := new(Petition)
		err = rows.StructScan(petition)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Petitions = append(resp.Petitions, petition.ToProto())
	}

	return resp, nil
}

// PetitionCreate implements SCRUD endpoints
func (s *Server) PetitionCreate(ctx context.Context, req *pb.PetitionCreateRequest) (*pb.PetitionCreateResponse, error) {

	fmt.Println(req)
	petition := new(Petition)

	st := reflect.TypeOf(*petition)

	args := map[string]interface{}{}
	query := "INSERT INTO petitions"

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

	resp := new(pb.PetitionCreateResponse)
	resp.Id = lastID

	return resp, nil
}

// PetitionRead implements SCRUD endpoints
func (s *Server) PetitionRead(ctx context.Context, req *pb.PetitionReadRequest) (*pb.PetitionReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	resp := new(pb.PetitionReadResponse)

	if req.Id < 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	query := "SELECT * FROM petitions WHERE "

	args := map[string]interface{}{}
	query += "id = :id"
	args["id"] = req.Id

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		petition := new(Petition)
		err = rows.StructScan(petition)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Petition = petition.ToProto()
	}
	return resp, nil
}

// PetitionUpdate implements SCRUD endpoints
func (s *Server) PetitionUpdate(ctx context.Context, req *pb.PetitionUpdateRequest) (*pb.PetitionUpdateResponse, error) {
	petition := new(Petition)

	st := reflect.TypeOf(*petition)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE petitions SET"

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
	resp := new(pb.PetitionUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// PetitionDelete implements SCRUD endpoints
func (s *Server) PetitionDelete(ctx context.Context, req *pb.PetitionDeleteRequest) (*pb.PetitionDeleteResponse, error) {
	query := "DELETE FROM petitions WHERE id = :id LIMIT 1"

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
	resp := new(pb.PetitionDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// PetitionPatch implements SCRUD endpoints
func (s *Server) PetitionPatch(ctx context.Context, req *pb.PetitionPatchRequest) (*pb.PetitionPatchResponse, error) {
	petition := new(Petition)

	st := reflect.TypeOf(*petition)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE petitions SET"

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
	resp := new(pb.PetitionPatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// Petition reports
type Petition struct {
	ID           int64  `db:"dib"`          //int(10) unsigned NOT NULL AUTO_INCREMENT,
	Petid        int64  `db:"petid"`        //int(10) unsigned NOT NULL DEFAULT '0',
	Charname     string `db:"charname"`     //varchar(32) NOT NULL DEFAULT '',
	Accountname  string `db:"accountname"`  //varchar(32) NOT NULL DEFAULT '',
	Lastgm       string `db:"lastgm"`       //varchar(32) NOT NULL DEFAULT '',
	Petitiontext string `db:"petitiontext"` //text NOT NULL,
	Gmtext       string `db:"gmtext"`       //text,
	Zone         string `db:"zone"`         //varchar(32) NOT NULL DEFAULT '',
	Urgency      int64  `db:"urgency"`      //int(11) NOT NULL DEFAULT '0',
	Charclass    int64  `db:"charclass"`    //int(11) NOT NULL DEFAULT '0',
	Charrace     int64  `db:"charrace"`     //int(11) NOT NULL DEFAULT '0',
	Charlevel    int64  `db:"charlevel"`    //int(11) NOT NULL DEFAULT '0',
	Checkouts    int64  `db:"checkouts"`    //int(11) NOT NULL DEFAULT '0',
	Unavailables int64  `db:"unavailables"` //int(11) NOT NULL DEFAULT '0',
	Ischeckedout int64  `db:"ischeckedout"` //tinyint(4) NOT NULL DEFAULT '0',
	Senttime     int64  `db:"senttime"`     //bigint(11) NOT NULL DEFAULT '0',
	Total        int64  `db:"total"`
}

// ToProto converts the petition type struct to protobuf
func (p *Petition) ToProto() *pb.Petition {
	petition := &pb.Petition{}
	petition.Id = p.ID
	petition.Petid = p.Petid
	petition.Charname = p.Charname
	petition.Accountname = p.Accountname
	petition.Lastgm = p.Lastgm
	petition.Petitiontext = p.Petitiontext
	petition.Gmtext = p.Gmtext
	petition.Zone = p.Zone
	petition.Urgency = p.Urgency
	petition.Charclass = p.Charclass
	petition.Charrace = p.Charrace
	petition.Charlevel = p.Charlevel
	petition.Checkouts = p.Checkouts
	petition.Unavailables = p.Unavailables
	petition.Ischeckedout = p.Ischeckedout
	petition.Senttime = p.Senttime
	return petition
}
