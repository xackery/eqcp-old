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

// TradeSearch implements SCRUD endpoints
func (s *Server) TradeSearch(ctx context.Context, req *pb.TradeSearchRequest) (*pb.TradeSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	resp := new(pb.TradeSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	req.Orderby = strings.ToLower(req.Orderby)
	if req.Orderby == "" {
		req.Orderby = "name"
	}
	orderByFields := []string{"name"}

	query := "SELECT count(id) as total, trade_types.* FROM trade_types WHERE "

	args := map[string]interface{}{}
	if len(req.Name) > 0 {
		query += "name LIKE :name"
		args["name"] = fmt.Sprintf("%%%s%%", req.Name)
	}

	isValid := false
	for _, field := range orderByFields {
		if req.Orderby != field {
			continue
		}
		isValid = true
	}
	if !isValid {
		return nil, fmt.Errorf("invalid orderby. Valid options are: %s", strings.Join(orderByFields, ","))
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

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		trade := new(Trade)
		err = rows.StructScan(trade)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Trades = append(resp.Trades, trade.ToProto())
		resp.Total = trade.Total
	}

	return resp, nil
}

// TradeCreate implements SCRUD endpoints
func (s *Server) TradeCreate(ctx context.Context, req *pb.TradeCreateRequest) (*pb.TradeCreateResponse, error) {

	fmt.Println(req)
	trade := new(Trade)

	st := reflect.TypeOf(*trade)

	args := map[string]interface{}{}
	query := "INSERT INTO trade_types"

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
			insertField += fmt.Sprintf("%s %s", comma, tag)
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

	resp := new(pb.TradeCreateResponse)
	resp.Id = lastID

	return resp, nil
}

// TradeRead implements SCRUD endpoints
func (s *Server) TradeRead(ctx context.Context, req *pb.TradeReadRequest) (*pb.TradeReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	resp := new(pb.TradeReadResponse)

	if req.Id < 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	query := "SELECT * FROM trade_types WHERE "

	args := map[string]interface{}{}
	query += "id = :id"
	args["id"] = req.Id

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		trade := new(Trade)
		err = rows.StructScan(trade)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Trade = trade.ToProto()
	}
	return resp, nil
}

// TradeUpdate implements SCRUD endpoints
func (s *Server) TradeUpdate(ctx context.Context, req *pb.TradeUpdateRequest) (*pb.TradeUpdateResponse, error) {
	trade := new(Trade)

	st := reflect.TypeOf(*trade)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE trade_types SET"

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
	resp := new(pb.TradeUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// TradeDelete implements SCRUD endpoints
func (s *Server) TradeDelete(ctx context.Context, req *pb.TradeDeleteRequest) (*pb.TradeDeleteResponse, error) {
	query := "DELETE FROM trade_types WHERE id = :id LIMIT 1"

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
	resp := new(pb.TradeDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// TradePatch implements SCRUD endpoints
func (s *Server) TradePatch(ctx context.Context, req *pb.TradePatchRequest) (*pb.TradePatchResponse, error) {
	trade := new(Trade)

	st := reflect.TypeOf(*trade)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE trade_types SET"

	comma := ""
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)

		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		if strings.ToLower(tag) != strings.ToLower(req.Key) {
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
	resp := new(pb.TradePatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// Trade represents an TRADE DB binding
type Trade struct {
	Tradeid    int64     `db:"trade_id"`    // int(11) NOT NULL AUTO_INCREMENT,
	Time       time.Time `db:"time"`        // timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
	Char1id    int64     `db:"char1_id"`    // int(11) DEFAULT '0',
	Char1pp    int64     `db:"char1_pp"`    // int(11) DEFAULT '0',
	Char1gp    int64     `db:"char1_gp"`    // int(11) DEFAULT '0',
	Char1sp    int64     `db:"char1_sp"`    // int(11) DEFAULT '0',
	Char1cp    int64     `db:"char1_cp"`    // int(11) DEFAULT '0',
	Char1items int64     `db:"char1_items"` // mediumint(7) DEFAULT '0',
	Char2id    int64     `db:"char2_id"`    // int(11) DEFAULT '0',
	Char2pp    int64     `db:"char2_pp"`    // int(11) DEFAULT '0',
	Char2gp    int64     `db:"char2_gp"`    // int(11) DEFAULT '0',
	Char2sp    int64     `db:"char2_sp"`    // int(11) DEFAULT '0',
	Char2cp    int64     `db:"char2_cp"`    // int(11) DEFAULT '0',
	Char2items int64     `db:"char2_items"` // mediumint(7) DEFAULT '0',
	Total      int64     `db:"total"`
}

// ToProto converts the trade type struct to protobuf
func (t *Trade) ToProto() *pb.Trade {
	trade := &pb.Trade{}
	trade.Tradeid = t.Tradeid
	trade.Time = t.Time.Unix()
	trade.Char1Id = t.Char1id
	trade.Char1Pp = t.Char1pp
	trade.Char1Gp = t.Char1gp
	trade.Char1Sp = t.Char1sp
	trade.Char1Cp = t.Char1cp
	trade.Char1Items = t.Char1items
	trade.Char2Id = t.Char2id
	trade.Char2Pp = t.Char2pp
	trade.Char2Gp = t.Char2gp
	trade.Char2Sp = t.Char2sp
	trade.Char2Cp = t.Char2cp
	trade.Char2Items = t.Char2items
	return trade
}
