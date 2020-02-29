package server

import (
	"context"
	"database/sql"
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

var (
	loginAccountTable = "login_accounts"
	loginAccountPK    = "id"
)

// LoginAccountSearch implements SCRUD endpoints
func (s *Server) LoginAccountSearch(ctx context.Context, req *pb.LoginAccountSearchRequest) (*pb.LoginAccountSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	resp := new(pb.LoginAccountSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	loginAccount := new(LoginAccount)

	st := reflect.TypeOf(*loginAccount)
	sv := reflect.ValueOf(loginAccount)
	se := sv.Elem()

	query := fmt.Sprintf("SELECT {fieldMap} FROM %s WHERE", loginAccountTable)

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
		req.Orderby = loginAccountPK
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

	queryTotal := strings.Replace(query, "{fieldMap}", fmt.Sprintf("count(%s) as total", loginAccountPK), 1)
	query = strings.Replace(query, "{fieldMap}", "*", 1)

	rows, err := s.db.NamedQueryContext(ctx, queryTotal, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	for rows.Next() {
		loginAccount := new(LoginAccount)
		err = rows.StructScan(loginAccount)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = loginAccount.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		loginAccount := new(LoginAccount)
		err = rows.StructScan(loginAccount)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.LoginAccounts = append(resp.LoginAccounts, loginAccount.ToProto())
	}

	return resp, nil
}

// LoginAccountCreate implements SCRUD endpoints
func (s *Server) LoginAccountCreate(ctx context.Context, req *pb.LoginAccountCreateRequest) (*pb.LoginAccountCreateResponse, error) {

	loginAccount := new(LoginAccount)

	st := reflect.TypeOf(*loginAccount)

	args := map[string]interface{}{}
	query := fmt.Sprintf("INSERT INTO %s", loginAccountTable)

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

	resp := new(pb.LoginAccountCreateResponse)
	resp.Id = lastID

	return resp, nil
}

// LoginAccountRead implements SCRUD endpoints
func (s *Server) LoginAccountRead(ctx context.Context, req *pb.LoginAccountReadRequest) (*pb.LoginAccountReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	resp := new(pb.LoginAccountReadResponse)

	if req.Id < 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	query := fmt.Sprintf("SELECT * FROM %s WHERE ", loginAccountTable)

	args := map[string]interface{}{}
	query += fmt.Sprintf("%s = :%s", loginAccountPK, loginAccountPK)
	args[loginAccountPK] = req.Id

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		loginAccount := new(LoginAccount)
		err = rows.StructScan(loginAccount)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.LoginAccount = loginAccount.ToProto()
	}
	return resp, nil
}

// LoginAccountUpdate implements SCRUD endpoints
func (s *Server) LoginAccountUpdate(ctx context.Context, req *pb.LoginAccountUpdateRequest) (*pb.LoginAccountUpdateResponse, error) {
	loginAccount := new(LoginAccount)

	st := reflect.TypeOf(*loginAccount)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := fmt.Sprintf("UPDATE %s SET", loginAccountTable)

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

	query += fmt.Sprintf(" WHERE %s = :%s LIMIT 1", loginAccountPK, loginAccountPK)

	log.Debug().Interface("args", args).Msgf("query: %s", query)

	result, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rowsaffected")
	}
	resp := new(pb.LoginAccountUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// LoginAccountDelete implements SCRUD endpoints
func (s *Server) LoginAccountDelete(ctx context.Context, req *pb.LoginAccountDeleteRequest) (*pb.LoginAccountDeleteResponse, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = :%s LIMIT 1", loginAccountTable, loginAccountPK, loginAccountPK)

	args := map[string]interface{}{
		loginAccountPK: req.Id,
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
	resp := new(pb.LoginAccountDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// LoginAccountPatch implements SCRUD endpoints
func (s *Server) LoginAccountPatch(ctx context.Context, req *pb.LoginAccountPatchRequest) (*pb.LoginAccountPatchResponse, error) {
	loginAccount := new(LoginAccount)

	st := reflect.TypeOf(*loginAccount)

	args := map[string]interface{}{
		loginAccountPK: req.Id,
	}
	query := fmt.Sprintf("UPDATE %s SET", loginAccountTable)

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

	query += fmt.Sprintf(" WHERE %s = :%s LIMIT 1", loginAccountPK, loginAccountPK)
	log.Debug().Interface("args", args).Msgf("query: %s", query)

	result, err := s.db.NamedExecContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rowsaffected")
	}
	resp := new(pb.LoginAccountPatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// LoginAccount represents an LOGIN_ACCOUNT DB binding
type LoginAccount struct {
	ID                int64          `db:"id"`                 // int(11) unsigned NOT NULL,
	Accountname       string         `db:"account_name"`       // varchar(50) NOT NULL,
	Accountpassword   string         `db:"account_password"`   // text NOT NULL,
	Accountemail      string         `db:"account_email"`      // varchar(100) NOT NULL,
	Sourceloginserver sql.NullString `db:"source_loginserver"` // varchar(64) DEFAULT NULL,
	Lastipaddress     string         `db:"last_ip_address"`    // varchar(15) NOT NULL,
	Lastlogindate     time.Time      `db:"last_login_date"`    // datetime NOT NULL,
	Createdat         sql.NullTime   `db:"created_at"`         // datetime DEFAULT NULL,
	Updatedat         sql.NullTime   `db:"updated_at"`         // datetime DEFAULT CURRENT_TIMESTAMP,

	Total int64 `db:"total"`
}

// ToProto converts the loginAccount type struct to protobuf
func (l *LoginAccount) ToProto() *pb.LoginAccount {
	loginAccount := &pb.LoginAccount{}
	loginAccount.Id = l.ID
	loginAccount.Accountname = l.Accountname
	loginAccount.Accountpassword = l.Accountpassword
	loginAccount.Accountemail = l.Accountemail
	loginAccount.Sourceloginserver = l.Sourceloginserver.String
	loginAccount.Lastipaddress = l.Lastipaddress
	loginAccount.Lastlogindate = l.Lastlogindate.Unix()
	loginAccount.Createdat = l.Createdat.Time.Unix()
	loginAccount.Updatedat = l.Updatedat.Time.Unix()
	return loginAccount
}
