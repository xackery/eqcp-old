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

// BugSearch implements SCRUD endpoints
func (s *Server) BugSearch(ctx context.Context, req *pb.BugSearchRequest) (*pb.BugSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	ap, err := s.AuthFromContext(ctx, "bug", "search")
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}
	fmt.Println(ap)
	resp := new(pb.BugSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	bug := new(Bug)

	st := reflect.TypeOf(*bug)
	sv := reflect.ValueOf(bug)
	se := sv.Elem()

	query := "SELECT {fieldMap} FROM bug_reports WHERE"

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
		bug := new(Bug)
		err = rows.StructScan(bug)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = bug.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		bug := new(Bug)
		err = rows.StructScan(bug)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Bugs = append(resp.Bugs, bug.ToProto())
	}

	return resp, nil
}

// BugCreate implements SCRUD endpoints
func (s *Server) BugCreate(ctx context.Context, req *pb.BugCreateRequest) (*pb.BugCreateResponse, error) {

	ap, err := s.AuthFromContext(ctx, "bug", "create")
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}
	fmt.Println(ap)

	bug := new(Bug)

	st := reflect.TypeOf(*bug)

	args := map[string]interface{}{}
	query := "INSERT INTO bug_reports"

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

	resp := new(pb.BugCreateResponse)
	resp.Id = lastID

	return resp, nil
}

// BugRead implements SCRUD endpoints
func (s *Server) BugRead(ctx context.Context, req *pb.BugReadRequest) (*pb.BugReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	ap, err := s.AuthFromContext(ctx, "bug", "read")
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}
	fmt.Println(ap)
	resp := new(pb.BugReadResponse)

	if req.Id < 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	query := "SELECT * FROM bug_reports WHERE "

	args := map[string]interface{}{}
	query += "id = :id"
	args["id"] = req.Id

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		bug := new(Bug)
		err = rows.StructScan(bug)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Bug = bug.ToProto()
	}
	return resp, nil
}

// BugUpdate implements SCRUD endpoints
func (s *Server) BugUpdate(ctx context.Context, req *pb.BugUpdateRequest) (*pb.BugUpdateResponse, error) {

	ap, err := s.AuthFromContext(ctx, "bug", "update")
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}
	fmt.Println(ap)
	bug := new(Bug)

	st := reflect.TypeOf(*bug)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE bug_reports SET"

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
	if len(args) < 2 {
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
	resp := new(pb.BugUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// BugDelete implements SCRUD endpoints
func (s *Server) BugDelete(ctx context.Context, req *pb.BugDeleteRequest) (*pb.BugDeleteResponse, error) {
	ap, err := s.AuthFromContext(ctx, "bug", "delete")
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}
	fmt.Println(ap)
	query := "DELETE FROM bug_reports WHERE id = :id LIMIT 1"

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
	resp := new(pb.BugDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// BugPatch implements SCRUD endpoints
func (s *Server) BugPatch(ctx context.Context, req *pb.BugPatchRequest) (*pb.BugPatchResponse, error) {
	bug := new(Bug)

	st := reflect.TypeOf(*bug)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE bug_reports SET"

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
	resp := new(pb.BugPatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// Bug reports
type Bug struct {
	ID                int64     `db:"id"`                  // int(11) unsigned NOT NULL AUTO_INCREMENT,
	Zone              string    `db:"zone"`                // varchar(32) NOT NULL DEFAULT 'Unknown',
	Clientversionid   int64     `db:"client_version_id"`   // int(11) unsigned NOT NULL DEFAULT '0',
	Clientversionname string    `db:"client_version_name"` // varchar(24) NOT NULL DEFAULT 'Unknown',
	Accountid         int64     `db:"account_id"`          // int(11) unsigned NOT NULL DEFAULT '0',
	Characterid       int64     `db:"character_id"`        // int(11) unsigned NOT NULL DEFAULT '0',
	Charactername     string    `db:"character_name"`      // varchar(64) NOT NULL DEFAULT 'Unknown',
	Reporterspoof     int64     `db:"reporter_spoof"`      // tinyint(1) NOT NULL DEFAULT '1',
	Categoryid        int64     `db:"category_id"`         // int(11) unsigned NOT NULL DEFAULT '0',
	Categoryname      string    `db:"category_name"`       // varchar(64) NOT NULL DEFAULT 'Other',
	Reportername      string    `db:"reporter_name"`       // varchar(64) NOT NULL DEFAULT 'Unknown',
	Uipath            string    `db:"ui_path"`             // varchar(128) NOT NULL DEFAULT 'Unknown',
	Posx              int64     `db:"pos_x"`               // float NOT NULL DEFAULT '0',
	Posy              int64     `db:"pos_y"`               // float NOT NULL DEFAULT '0',
	Posz              int64     `db:"pos_z"`               // float NOT NULL DEFAULT '0',
	Heading           int64     `db:"heading"`             // int(11) unsigned NOT NULL DEFAULT '0',
	Timeplayed        int64     `db:"time_played"`         // int(11) unsigned NOT NULL DEFAULT '0',
	Targetid          int64     `db:"target_id"`           // int(11) unsigned NOT NULL DEFAULT '0',
	Targetname        string    `db:"target_name"`         // varchar(64) NOT NULL DEFAULT 'Unknown',
	Optionalinfomask  int64     `db:"optional_info_mask"`  // int(11) unsigned NOT NULL DEFAULT '0',
	Canduplicate      int64     `db:"_can_duplicate"`      // tinyint(1) NOT NULL DEFAULT '0',
	Crashbug          int64     `db:"_crash_bug"`          // tinyint(1) NOT NULL DEFAULT '0',
	Targetinfo        int64     `db:"_target_info"`        // tinyint(1) NOT NULL DEFAULT '0',
	Characterflags    int64     `db:"_character_flags"`    // tinyint(1) NOT NULL DEFAULT '0',
	Unknownvalue      int64     `db:"_unknown_value"`      // tinyint(1) NOT NULL DEFAULT '0',
	Bugreport         string    `db:"bug_report"`          // varchar(1024) NOT NULL DEFAULT '',
	Systeminfo        string    `db:"system_info"`         // varchar(1024) NOT NULL DEFAULT '',
	Reportdatetime    time.Time `db:"report_datetime"`     // datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	Bugstatus         int64     `db:"bug_status"`          // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Lastreview        time.Time `db:"last_review"`         // datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	Lastreviewer      string    `db:"last_reviewer"`       // varchar(64) NOT NULL DEFAULT 'None',
	Reviewernotes     string    `db:"reviewer_notes"`      // varchar(1024) NOT NULL DEFAULT '',
	Total             int64     `db:"total"`
}

// ToProto converts the bug type struct to protobuf
func (b *Bug) ToProto() *pb.Bug {
	bug := &pb.Bug{}
	bug.Id = b.ID
	bug.Zone = b.Zone
	bug.Clientversionid = b.Clientversionid
	bug.Clientversionname = b.Clientversionname
	bug.Accountid = b.Accountid
	bug.Characterid = b.Characterid
	bug.Charactername = b.Charactername
	bug.Reporterspoof = b.Reporterspoof
	bug.Categoryid = b.Categoryid
	bug.Categoryname = b.Categoryname
	bug.Reportername = b.Reportername
	bug.Uipath = b.Uipath
	bug.Posx = b.Posx
	bug.Posy = b.Posy
	bug.Posz = b.Posz
	bug.Heading = b.Heading
	bug.Timeplayed = b.Timeplayed
	bug.Targetid = b.Targetid
	bug.Targetname = b.Targetname
	bug.Optionalinfomask = b.Optionalinfomask
	bug.Canduplicate = b.Canduplicate
	bug.Crashbug = b.Crashbug
	bug.Targetinfo = b.Targetinfo
	bug.Characterflags = b.Characterflags
	bug.Unknownvalue = b.Unknownvalue
	bug.Bugreport = b.Bugreport
	bug.Systeminfo = b.Systeminfo
	bug.Reportdatetime = b.Reportdatetime.Unix()
	bug.Bugstatus = b.Bugstatus
	bug.Lastreview = b.Lastreview.Unix()
	bug.Lastreviewer = b.Lastreviewer
	bug.Reviewernotes = b.Reviewernotes
	return bug
}
