package server

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/xackery/eqcp/pb"

	//mysql db
	_ "github.com/go-sql-driver/mysql"
)

// AccountSearch implements SCRUD endpoints
func (s *Server) AccountSearch(ctx context.Context, req *pb.AccountSearchRequest) (*pb.AccountSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	ap, err := s.AuthFromContext(ctx, "account", "search")
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}

	resp := new(pb.AccountSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	account := new(Account)

	st := reflect.TypeOf(*account)
	sv := reflect.ValueOf(account)
	se := sv.Elem()

	fieldMap := account.fieldMap(ap.fields)
	if len(fieldMap) == 0 {
		log.Warn().Msg("fields for account search is empty")
		return nil, fmt.Errorf("permission denied")
	}

	query := "SELECT {fieldMap} FROM account WHERE"
	validate := validator.New()
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
			isValid := false
			if !ap.isAllFieldsOK {
				for _, af := range ap.fields {
					if af == field.Name {
						isValid = true
						break
					}
				}
			}
			if !isValid {
				continue
			}
			validTag, ok := field.Tag.Lookup("validate")
			if ok {
				errs := validate.Var(value, validTag)
				if errs != nil {
					return nil, fmt.Errorf("%s: %s", key, errs.Error())
				}
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

	if ap.isSelfOnly {
		if ap == nil {
			return nil, fmt.Errorf("permission denied, not logged in")
		}
		args["lsaccount_id"] = ap.accountID
		query += " AND lsaccount_id = :lsaccount_id"
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
	query = strings.Replace(query, "{fieldMap}", fieldMap, 1)

	rows, err := s.db.NamedQueryContext(ctx, queryTotal, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}
	for rows.Next() {
		account := new(Account)
		err = rows.StructScan(account)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = account.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		account := new(Account)
		err = rows.StructScan(account)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Accounts = append(resp.Accounts, account.ToProto())
	}

	return resp, nil
}

// AccountCreate implements SCRUD endpoints
func (s *Server) AccountCreate(ctx context.Context, req *pb.AccountCreateRequest) (*pb.AccountCreateResponse, error) {

	ap, err := s.AuthFromContext(ctx, "account", "create")
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}

	account := new(Account)

	st := reflect.TypeOf(*account)

	args := map[string]interface{}{}
	query := "INSERT INTO account"
	validate := validator.New()
	comma := ""
	insertField := ""
	insertValue := ""
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fieldName := strings.ToLower(field.Name)

		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		if !ap.isAllFieldsOK {
			isValid := false
			for _, af := range ap.fields {
				if fieldName != af {
					continue
				}
				isValid = true
			}
			if !isValid {
				continue
			}
		}

		for key, value := range req.Values {
			if fieldName != strings.ToLower(key) {
				continue
			}
			validTag, ok := field.Tag.Lookup("validate")
			if ok {
				errs := validate.Var(value, validTag)
				if errs != nil {
					return nil, fmt.Errorf("%s: %s", key, errs.Error())
				}
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

	stmt, err := s.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	result, err := stmt.Exec(args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "lastinsertedid")
	}

	resp := new(pb.AccountCreateResponse)
	resp.Id = lastID
	return resp, nil
}

// AccountRead implements SCRUD endpoints
func (s *Server) AccountRead(ctx context.Context, req *pb.AccountReadRequest) (*pb.AccountReadResponse, error) {

	ap, err := s.AuthFromContext(ctx, "account", "read")
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}

	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	account := new(Account)
	if req.Id < 0 {
		return nil, fmt.Errorf("id must be greater than or equal to 0")
	}

	fieldMap := account.fieldMap(ap.fields)
	if len(fieldMap) == 0 {
		log.Warn().Msg("fields for account read is empty")
		return nil, fmt.Errorf("permission denied")
	}

	resp := new(pb.AccountReadResponse)

	query := "SELECT {fieldMap} FROM account WHERE "

	args := map[string]interface{}{}
	query += "id = :id"
	args["id"] = req.Id

	if ap.isSelfOnly {
		args["lsaccount_id"] = ap.accountID
		query += " AND lsaccount_id = :lsaccount_id"
	}

	query = strings.Replace(query, "{fieldMap}", fieldMap, 1)

	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		err = rows.StructScan(account)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}

		resp.Account = account.ToProto()
		break
	}
	return resp, nil
}

// AccountUpdate implements SCRUD endpoints
func (s *Server) AccountUpdate(ctx context.Context, req *pb.AccountUpdateRequest) (*pb.AccountUpdateResponse, error) {

	ap, err := s.AuthFromContext(ctx, "account", "update")
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}
	if ap.isSelfOnly && req.Id != ap.accountID {
		return nil, fmt.Errorf("permission denied: only permission for your own account")
	}

	/*if s.cfg.Permission.AccountUpdate > 0 {
		if ap.accountID != req.Id && s.cfg.Permission.AccountUpdate > 0 && ap.status < s.cfg.Permission.AccountUpdate {
			return nil, fmt.Errorf("you do not have permissions to this endpoint")
		}
	}*/

	account := new(Account)

	st := reflect.TypeOf(*account)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE account SET"

	comma := ""
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fieldName := strings.ToLower(field.Name)

		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		if !ap.isAllFieldsOK {
			isValid := false
			for _, af := range ap.fields {
				if fieldName != af {
					continue
				}
				isValid = true
			}
			if !isValid {
				continue
			}
		}

		for key, value := range req.Values {
			if fieldName != strings.ToLower(key) {
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
	resp := new(pb.AccountUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// AccountDelete implements SCRUD endpoints
func (s *Server) AccountDelete(ctx context.Context, req *pb.AccountDeleteRequest) (*pb.AccountDeleteResponse, error) {
	ap, err := s.AuthFromContext(ctx, "account", "delete")
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}

	query := "DELETE FROM account WHERE id = :id"

	args := map[string]interface{}{
		"id": req.Id,
	}

	if ap.isSelfOnly {
		args["lsaccount_id"] = ap.accountID
		query += " AND lsaccount_id = :ls_account_id"
	}
	query += " LIMIT 1"
	log.Debug().Interface("args", args).Msgf("query: %s", query)

	result, err := s.db.NamedExecContext(ctx, query, args)

	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rowsaffected")
	}
	resp := new(pb.AccountDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// AccountPatch implements SCRUD endpoints
func (s *Server) AccountPatch(ctx context.Context, req *pb.AccountPatchRequest) (*pb.AccountPatchResponse, error) {
	ap, err := s.AuthFromContext(ctx, "account", "patch")
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}

	account := new(Account)

	st := reflect.TypeOf(*account)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE account SET"

	comma := ""
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fieldName := strings.ToLower(field.Name)
		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		if fieldName != strings.ToLower(req.Key) {
			continue
		}

		isValid := false
		if !ap.isAllFieldsOK {
			for _, af := range ap.fields {
				if af == fieldName {
					isValid = true
					break
				}
			}
		}
		if !isValid {
			continue
		}

		args[tag] = req.Value

		query += fmt.Sprintf("%s %s = :%s", comma, tag, tag)
		comma = ","
	}
	if len(args) == 1 {
		return nil, fmt.Errorf("no valid fields provided")
	}

	query += " WHERE id = :id"
	if ap.isSelfOnly {
		args["lsaccount_id"] = ap.accountID
		query += " AND lsaccount_id = :lsaccount_id"
	}
	query += " LIMIT 1"
	log.Debug().Interface("args", args).Msgf("query: %s", query)

	stmt, err := s.db.PrepareNamedContext(ctx, query)
	result, err := stmt.ExecContext(ctx, args)
	if err != nil {
		return nil, errors.Wrap(err, "query")
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rowsaffected")
	}
	resp := new(pb.AccountPatchResponse)
	resp.Rowsaffected = rowCount
	return resp, nil
}

// Account represents an ACCOUNT DB binding
type Account struct {
	ID             int64          `db:"id" validate:"gte=0"`                    // int(11) NOT NULL AUTO_INCREMENT,
	Name           string         `db:"name" validate:"gte=3,lte=30,alpha"`     // varchar(30) NOT NULL DEFAULT '',
	Charname       string         `db:"charname" validate:"gte=3,lte=64,alpha"` // varchar(64) NOT NULL DEFAULT '',
	Sharedplat     int64          `db:"sharedplat"  validate:"gte=0"`           // int(11) unsigned NOT NULL DEFAULT '0',
	Password       string         `db:"password" validate:"lte=50"`             // varchar(50) NOT NULL DEFAULT '',
	Status         int64          `db:"status" validate:"gte=0,lte=255"`        // int(5) NOT NULL DEFAULT '0',
	Lsid           string         `db:"ls_id" validate:"gte=0"`                 // varchar(64) DEFAULT 'eqemu',
	Lsaccountid    sql.NullInt64  `db:"lsaccount_id" validate:"gte=0"`          // int(11) unsigned DEFAULT NULL,
	Gmspeed        int64          `db:"gmspeed"  validate:"gte=0"`              // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Revoked        int64          `db:"revoked" validate:"gte=0"`               // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Karma          int64          `db:"karma" validate:"gte=0"`                 // int(5) unsigned NOT NULL DEFAULT '0',
	Miniloginip    string         `db:"minilogin_ip" validate:"ipv4"`           // varchar(32) NOT NULL DEFAULT '',
	Hideme         int64          `db:"hideme" validate:"gte=0"`                // tinyint(4) NOT NULL DEFAULT '0',
	Rulesflag      int64          `db:"rulesflag" validate:"gte=0"`             // tinyint(1) unsigned NOT NULL DEFAULT '0',
	Suspendeduntil time.Time      `db:"suspendeduntil"`                         // datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	Timecreation   int64          `db:"time_creation"`                          // int(10) unsigned NOT NULL DEFAULT '0',
	Expansion      int64          `db:"expansion"`                              // tinyint(4) NOT NULL DEFAULT '0',
	Banreason      sql.NullString `db:"ban_reason" validate:"lte=255"`          // text,
	Suspendreason  sql.NullString `db:"suspend_reason" validate:"lte=255"`      // text,

	Total int64 `db:"total"`
}

// returns a comma delimited list for searc/read
func (a *Account) fieldMap(fields []string) string {
	query := ""

	st := reflect.TypeOf(*a)
	isAllFieldsOK := false
	if len(fields) == 1 && fields[1] == "*" {
		isAllFieldsOK = true
	}
	comma := ""
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fieldName := strings.ToLower(field.Name)

		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		if !isAllFieldsOK {
			isValid := false
			for _, af := range fields {
				if fieldName != af {
					continue
				}
				isValid = true
			}
			if !isValid {
				continue
			}
		}
		query += fmt.Sprintf("%s %s", comma, tag)
		comma = ","
	}
	return query
}

// ToProto converts the account type struct to protobuf
func (a *Account) ToProto() *pb.Account {
	account := &pb.Account{}
	account.Id = a.ID
	account.Name = a.Name
	account.Charname = a.Charname
	account.Sharedplat = a.Sharedplat
	account.Password = a.Password
	account.Status = a.Status
	account.Lsid = a.Lsid
	account.Lsaccountid = a.Lsaccountid.Int64
	account.Gmspeed = a.Gmspeed
	account.Revoked = a.Revoked
	account.Karma = a.Karma
	account.Miniloginip = a.Miniloginip
	account.Hideme = a.Hideme
	account.Rulesflag = a.Rulesflag
	account.Suspendeduntil = a.Suspendeduntil.Unix()
	account.Timecreation = a.Timecreation
	account.Expansion = a.Expansion
	account.Banreason = a.Banreason.String
	account.Suspendreason = a.Suspendreason.String
	return account
}
