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

// SpellSearch implements SCRUD endpoints
func (s *Server) SpellSearch(ctx context.Context, req *pb.SpellSearchRequest) (*pb.SpellSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	resp := new(pb.SpellSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	spell := new(Spell)

	st := reflect.TypeOf(*spell)
	sv := reflect.ValueOf(spell)
	se := sv.Elem()

	query := "SELECT {fieldMap} FROM spells_new WHERE"

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
			} else if se.Field(i).Kind() == reflect.Struct {
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
		spell := new(Spell)
		err = rows.StructScan(spell)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = spell.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		spell := new(Spell)
		err = rows.StructScan(spell)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Spells = append(resp.Spells, spell.ToProto())
	}

	return resp, nil
}

// SpellCreate implements SCRUD endpoints
func (s *Server) SpellCreate(ctx context.Context, req *pb.SpellCreateRequest) (*pb.SpellCreateResponse, error) {

	ap, err := s.AuthFromContext(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}
	if !ap.hasCommand("mysql") {
		return nil, fmt.Errorf("you do not have permissions to this endpoint")
	}

	spell := new(Spell)

	st := reflect.TypeOf(*spell)

	args := map[string]interface{}{}
	query := "INSERT INTO spell"

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

	resp := new(pb.SpellCreateResponse)
	resp.Id = lastID

	return resp, nil
}

// SpellRead implements SCRUD endpoints
func (s *Server) SpellRead(ctx context.Context, req *pb.SpellReadRequest) (*pb.SpellReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	resp := new(pb.SpellReadResponse)

	if req.Id < 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	query := "SELECT * FROM spells_new WHERE "

	args := map[string]interface{}{}
	query += "id = :id"
	args["id"] = req.Id

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		spell := new(Spell)
		err = rows.StructScan(spell)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Spell = spell.ToProto()
	}
	return resp, nil
}

// SpellUpdate implements SCRUD endpoints
func (s *Server) SpellUpdate(ctx context.Context, req *pb.SpellUpdateRequest) (*pb.SpellUpdateResponse, error) {
	ap, err := s.AuthFromContext(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}
	if !ap.hasCommand("mysql") {
		return nil, fmt.Errorf("you do not have permissions to this endpoint")
	}

	spell := new(Spell)

	st := reflect.TypeOf(*spell)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE spells_new SET"

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
	resp := new(pb.SpellUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// SpellDelete implements SCRUD endpoints
func (s *Server) SpellDelete(ctx context.Context, req *pb.SpellDeleteRequest) (*pb.SpellDeleteResponse, error) {
	ap, err := s.AuthFromContext(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}
	if !ap.hasCommand("mysql") {
		return nil, fmt.Errorf("you do not have permissions to this endpoint")
	}

	query := "DELETE FROM spells_new WHERE id = :id LIMIT 1"

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
	resp := new(pb.SpellDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// SpellPatch implements SCRUD endpoints
func (s *Server) SpellPatch(ctx context.Context, req *pb.SpellPatchRequest) (*pb.SpellPatchResponse, error) {
	ap, err := s.AuthFromContext(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("authfromcontext")
		return nil, fmt.Errorf("permission denied")
	}
	if !ap.hasCommand("mysql") {
		return nil, fmt.Errorf("you do not have permissions to this endpoint")
	}

	spell := new(Spell)

	st := reflect.TypeOf(*spell)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE spells_new SET"

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
	resp := new(pb.SpellPatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// Spell represents an SPELL DB binding
type Spell struct {
	ID                  int64          `db:"id"`                   // int(11) NOT NULL DEFAULT 0,
	Name                sql.NullString `db:"name"`                 // varchar(64) DEFAULT NULL,
	Player1             sql.NullString `db:"player_1"`             // varchar(64) DEFAULT 'BLUE_TRAIL',
	Teleportzone        sql.NullString `db:"teleport_zone"`        // varchar(64) DEFAULT NULL,
	Youcast             sql.NullString `db:"you_cast"`             // varchar(120) DEFAULT NULL,
	Othercasts          sql.NullString `db:"other_casts"`          // varchar(120) DEFAULT NULL,
	Castonyou           sql.NullString `db:"cast_on_you"`          // varchar(120) DEFAULT NULL,
	Castonother         sql.NullString `db:"cast_on_other"`        // varchar(120) DEFAULT NULL,
	Spellfades          sql.NullString `db:"spell_fades"`          // varchar(120) DEFAULT NULL,
	Range               int64          `db:"range"`                // int(11) NOT NULL DEFAULT 100,
	Aoerange            int64          `db:"aoerange"`             // int(11) NOT NULL DEFAULT 0,
	Pushback            int64          `db:"pushback"`             // int(11) NOT NULL DEFAULT 0,
	Pushup              int64          `db:"pushup"`               // int(11) NOT NULL DEFAULT 0,
	Casttime            int64          `db:"cast_time"`            // int(11) NOT NULL DEFAULT 0,
	Recoverytime        int64          `db:"recovery_time"`        // int(11) NOT NULL DEFAULT 0,
	Recasttime          int64          `db:"recast_time"`          // int(11) NOT NULL DEFAULT 0,
	Buffdurationformula int64          `db:"buffdurationformula"`  // int(11) NOT NULL DEFAULT 7,
	Buffduration        int64          `db:"buffduration"`         // int(11) NOT NULL DEFAULT 65,
	Aeduration          int64          `db:"AEDuration"`           // int(11) NOT NULL DEFAULT 0,
	Mana                int64          `db:"mana"`                 // int(11) NOT NULL DEFAULT 0,
	Effectbasevalue1    int64          `db:"effect_base_value1"`   // int(11) NOT NULL DEFAULT 100,
	Effectbasevalue2    int64          `db:"effect_base_value2"`   // int(11) NOT NULL DEFAULT 0,
	Effectbasevalue3    int64          `db:"effect_base_value3"`   // int(11) NOT NULL DEFAULT 0,
	Effectbasevalue4    int64          `db:"effect_base_value4"`   // int(11) NOT NULL DEFAULT 0,
	Effectbasevalue5    int64          `db:"effect_base_value5"`   // int(11) NOT NULL DEFAULT 0,
	Effectbasevalue6    int64          `db:"effect_base_value6"`   // int(11) NOT NULL DEFAULT 0,
	Effectbasevalue7    int64          `db:"effect_base_value7"`   // int(11) NOT NULL DEFAULT 0,
	Effectbasevalue8    int64          `db:"effect_base_value8"`   // int(11) NOT NULL DEFAULT 0,
	Effectbasevalue9    int64          `db:"effect_base_value9"`   // int(11) NOT NULL DEFAULT 0,
	Effectbasevalue10   int64          `db:"effect_base_value10"`  // int(11) NOT NULL DEFAULT 0,
	Effectbasevalue11   int64          `db:"effect_base_value11"`  // int(11) NOT NULL DEFAULT 0,
	Effectbasevalue12   int64          `db:"effect_base_value12"`  // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue1   int64          `db:"effect_limit_value1"`  // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue2   int64          `db:"effect_limit_value2"`  // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue3   int64          `db:"effect_limit_value3"`  // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue4   int64          `db:"effect_limit_value4"`  // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue5   int64          `db:"effect_limit_value5"`  // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue6   int64          `db:"effect_limit_value6"`  // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue7   int64          `db:"effect_limit_value7"`  // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue8   int64          `db:"effect_limit_value8"`  // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue9   int64          `db:"effect_limit_value9"`  // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue10  int64          `db:"effect_limit_value10"` // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue11  int64          `db:"effect_limit_value11"` // int(11) NOT NULL DEFAULT 0,
	Effectlimitvalue12  int64          `db:"effect_limit_value12"` // int(11) NOT NULL DEFAULT 0,
	Max1                int64          `db:"max1"`                 // int(11) NOT NULL DEFAULT 0,
	Max2                int64          `db:"max2"`                 // int(11) NOT NULL DEFAULT 0,
	Max3                int64          `db:"max3"`                 // int(11) NOT NULL DEFAULT 0,
	Max4                int64          `db:"max4"`                 // int(11) NOT NULL DEFAULT 0,
	Max5                int64          `db:"max5"`                 // int(11) NOT NULL DEFAULT 0,
	Max6                int64          `db:"max6"`                 // int(11) NOT NULL DEFAULT 0,
	Max7                int64          `db:"max7"`                 // int(11) NOT NULL DEFAULT 0,
	Max8                int64          `db:"max8"`                 // int(11) NOT NULL DEFAULT 0,
	Max9                int64          `db:"max9"`                 // int(11) NOT NULL DEFAULT 0,
	Max10               int64          `db:"max10"`                // int(11) NOT NULL DEFAULT 0,
	Max11               int64          `db:"max11"`                // int(11) NOT NULL DEFAULT 0,
	Max12               int64          `db:"max12"`                // int(11) NOT NULL DEFAULT 0,
	Icon                int64          `db:"icon"`                 // int(11) NOT NULL DEFAULT 0,
	Memicon             int64          `db:"memicon"`              // int(11) NOT NULL DEFAULT 0,
	Components1         int64          `db:"components1"`          // int(11) NOT NULL DEFAULT -1,
	Components2         int64          `db:"components2"`          // int(11) NOT NULL DEFAULT -1,
	Components3         int64          `db:"components3"`          // int(11) NOT NULL DEFAULT -1,
	Components4         int64          `db:"components4"`          // int(11) NOT NULL DEFAULT -1,
	Componentcounts1    int64          `db:"component_counts1"`    // int(11) NOT NULL DEFAULT 1,
	Componentcounts2    int64          `db:"component_counts2"`    // int(11) NOT NULL DEFAULT 1,
	Componentcounts3    int64          `db:"component_counts3"`    // int(11) NOT NULL DEFAULT 1,
	Componentcounts4    int64          `db:"component_counts4"`    // int(11) NOT NULL DEFAULT 1,
	Noexpendreagent1    int64          `db:"NoexpendReagent1"`     // int(11) NOT NULL DEFAULT -1,
	Noexpendreagent2    int64          `db:"NoexpendReagent2"`     // int(11) NOT NULL DEFAULT -1,
	Noexpendreagent3    int64          `db:"NoexpendReagent3"`     // int(11) NOT NULL DEFAULT -1,
	Noexpendreagent4    int64          `db:"NoexpendReagent4"`     // int(11) NOT NULL DEFAULT -1,
	Formula1            int64          `db:"formula1"`             // int(11) NOT NULL DEFAULT 100,
	Formula2            int64          `db:"formula2"`             // int(11) NOT NULL DEFAULT 100,
	Formula3            int64          `db:"formula3"`             // int(11) NOT NULL DEFAULT 100,
	Formula4            int64          `db:"formula4"`             // int(11) NOT NULL DEFAULT 100,
	Formula5            int64          `db:"formula5"`             // int(11) NOT NULL DEFAULT 100,
	Formula6            int64          `db:"formula6"`             // int(11) NOT NULL DEFAULT 100,
	Formula7            int64          `db:"formula7"`             // int(11) NOT NULL DEFAULT 100,
	Formula8            int64          `db:"formula8"`             // int(11) NOT NULL DEFAULT 100,
	Formula9            int64          `db:"formula9"`             // int(11) NOT NULL DEFAULT 100,
	Formula10           int64          `db:"formula10"`            // int(11) NOT NULL DEFAULT 100,
	Formula11           int64          `db:"formula11"`            // int(11) NOT NULL DEFAULT 100,
	Formula12           int64          `db:"formula12"`            // int(11) NOT NULL DEFAULT 100,
	Lighttype           int64          `db:"LightType"`            // int(11) NOT NULL DEFAULT 0,
	Goodeffect          int64          `db:"goodEffect"`           // int(11) NOT NULL DEFAULT 0,
	Activated           int64          `db:"Activated"`            // int(11) NOT NULL DEFAULT 0,
	Resisttype          int64          `db:"resisttype"`           // int(11) NOT NULL DEFAULT 0,
	Effectid1           int64          `db:"effectid1"`            // int(11) NOT NULL DEFAULT 254,
	Effectid2           int64          `db:"effectid2"`            // int(11) NOT NULL DEFAULT 254,
	Effectid3           int64          `db:"effectid3"`            // int(11) NOT NULL DEFAULT 254,
	Effectid4           int64          `db:"effectid4"`            // int(11) NOT NULL DEFAULT 254,
	Effectid5           int64          `db:"effectid5"`            // int(11) NOT NULL DEFAULT 254,
	Effectid6           int64          `db:"effectid6"`            // int(11) NOT NULL DEFAULT 254,
	Effectid7           int64          `db:"effectid7"`            // int(11) NOT NULL DEFAULT 254,
	Effectid8           int64          `db:"effectid8"`            // int(11) NOT NULL DEFAULT 254,
	Effectid9           int64          `db:"effectid9"`            // int(11) NOT NULL DEFAULT 254,
	Effectid10          int64          `db:"effectid10"`           // int(11) NOT NULL DEFAULT 254,
	Effectid11          int64          `db:"effectid11"`           // int(11) NOT NULL DEFAULT 254,
	Effectid12          int64          `db:"effectid12"`           // int(11) NOT NULL DEFAULT 254,
	Targettype          int64          `db:"targettype"`           // int(11) NOT NULL DEFAULT 2,
	Basediff            int64          `db:"basediff"`             // int(11) NOT NULL DEFAULT 0,
	Skill               int64          `db:"skill"`                // int(11) NOT NULL DEFAULT 98,
	Zonetype            int64          `db:"zonetype"`             // int(11) NOT NULL DEFAULT -1,
	Environmenttype     int64          `db:"EnvironmentType"`      // int(11) NOT NULL DEFAULT 0,
	Timeofday           int64          `db:"TimeOfDay"`            // int(11) NOT NULL DEFAULT 0,
	Classes1            int64          `db:"classes1"`             // int(11) NOT NULL DEFAULT 255,
	Classes2            int64          `db:"classes2"`             // int(11) NOT NULL DEFAULT 255,
	Classes3            int64          `db:"classes3"`             // int(11) NOT NULL DEFAULT 255,
	Classes4            int64          `db:"classes4"`             // int(11) NOT NULL DEFAULT 255,
	Classes5            int64          `db:"classes5"`             // int(11) NOT NULL DEFAULT 255,
	Classes6            int64          `db:"classes6"`             // int(11) NOT NULL DEFAULT 255,
	Classes7            int64          `db:"classes7"`             // int(11) NOT NULL DEFAULT 255,
	Classes8            int64          `db:"classes8"`             // int(11) NOT NULL DEFAULT 255,
	Classes9            int64          `db:"classes9"`             // int(11) NOT NULL DEFAULT 255,
	Classes10           int64          `db:"classes10"`            // int(11) NOT NULL DEFAULT 255,
	Classes11           int64          `db:"classes11"`            // int(11) NOT NULL DEFAULT 255,
	Classes12           int64          `db:"classes12"`            // int(11) NOT NULL DEFAULT 255,
	Classes13           int64          `db:"classes13"`            // int(11) NOT NULL DEFAULT 255,
	Classes14           int64          `db:"classes14"`            // int(11) NOT NULL DEFAULT 255,
	Classes15           int64          `db:"classes15"`            // int(11) NOT NULL DEFAULT 255,
	Classes16           int64          `db:"classes16"`            // int(11) NOT NULL DEFAULT 255,
	Castinganim         int64          `db:"CastingAnim"`          // int(11) NOT NULL DEFAULT 44,
	Targetanim          int64          `db:"TargetAnim"`           // int(11) NOT NULL DEFAULT 13,
	Traveltype          int64          `db:"TravelType"`           // int(11) NOT NULL DEFAULT 0,
	Spellaffectindex    int64          `db:"SpellAffectIndex"`     // int(11) NOT NULL DEFAULT -1,
	Disallowsit         int64          `db:"disallow_sit"`         // int(11) NOT NULL DEFAULT 0,
	Deities0            int64          `db:"deities0"`             // int(11) NOT NULL DEFAULT 0,
	Deities1            int64          `db:"deities1"`             // int(11) NOT NULL DEFAULT 0,
	Deities2            int64          `db:"deities2"`             // int(11) NOT NULL DEFAULT 0,
	Deities3            int64          `db:"deities3"`             // int(11) NOT NULL DEFAULT 0,
	Deities4            int64          `db:"deities4"`             // int(11) NOT NULL DEFAULT 0,
	Deities5            int64          `db:"deities5"`             // int(11) NOT NULL DEFAULT 0,
	Deities6            int64          `db:"deities6"`             // int(11) NOT NULL DEFAULT 0,
	Deities7            int64          `db:"deities7"`             // int(11) NOT NULL DEFAULT 0,
	Deities8            int64          `db:"deities8"`             // int(11) NOT NULL DEFAULT 0,
	Deities9            int64          `db:"deities9"`             // int(11) NOT NULL DEFAULT 0,
	Deities10           int64          `db:"deities10"`            // int(11) NOT NULL DEFAULT 0,
	Deities11           int64          `db:"deities11"`            // int(11) NOT NULL DEFAULT 0,
	Deities12           int64          `db:"deities12"`            // int(12) NOT NULL DEFAULT 0,
	Deities13           int64          `db:"deities13"`            // int(11) NOT NULL DEFAULT 0,
	Deities14           int64          `db:"deities14"`            // int(11) NOT NULL DEFAULT 0,
	Deities15           int64          `db:"deities15"`            // int(11) NOT NULL DEFAULT 0,
	Deities16           int64          `db:"deities16"`            // int(11) NOT NULL DEFAULT 0,
	Field142            int64          `db:"field142"`             // int(11) NOT NULL DEFAULT 100,
	Field143            int64          `db:"field143"`             // int(11) NOT NULL DEFAULT 0,
	Newicon             int64          `db:"new_icon"`             // int(11) NOT NULL DEFAULT 161,
	Spellanim           int64          `db:"spellanim"`            // int(11) NOT NULL DEFAULT 0,
	Uninterruptable     int64          `db:"uninterruptable"`      // int(11) NOT NULL DEFAULT 0,
	Resistdiff          int64          `db:"ResistDiff"`           // int(11) NOT NULL DEFAULT -150,
	Dotstackingexempt   int64          `db:"dot_stacking_exempt"`  // int(11) NOT NULL DEFAULT 0,
	Deleteable          int64          `db:"deleteable"`           // int(11) NOT NULL DEFAULT 0,
	Recourselink        int64          `db:"RecourseLink"`         // int(11) NOT NULL DEFAULT 0,
	Nopartialresist     int64          `db:"no_partial_resist"`    // int(11) NOT NULL DEFAULT 0,
	Field152            int64          `db:"field152"`             // int(11) NOT NULL DEFAULT 0,
	Field153            int64          `db:"field153"`             // int(11) NOT NULL DEFAULT 0,
	Shortbuffbox        int64          `db:"short_buff_box"`       // int(11) NOT NULL DEFAULT -1,
	Descnum             int64          `db:"descnum"`              // int(11) NOT NULL DEFAULT 0,
	Typedescnum         sql.NullInt64  `db:"typedescnum"`          // int(11) DEFAULT NULL,
	Effectdescnum       sql.NullInt64  `db:"effectdescnum"`        // int(11) DEFAULT NULL,
	Effectdescnum2      int64          `db:"effectdescnum2"`       // int(11) NOT NULL DEFAULT 0,
	Npcnolos            int64          `db:"npc_no_los"`           // int(11) NOT NULL DEFAULT 0,
	Field160            int64          `db:"field160"`             // int(11) NOT NULL DEFAULT 0,
	Reflectable         int64          `db:"reflectable"`          // int(11) NOT NULL DEFAULT 0,
	Bonushate           int64          `db:"bonushate"`            // int(11) NOT NULL DEFAULT 0,
	Field163            int64          `db:"field163"`             // int(11) NOT NULL DEFAULT 100,
	Field164            int64          `db:"field164"`             // int(11) NOT NULL DEFAULT -150,
	Ldontrap            int64          `db:"ldon_trap"`            // int(11) NOT NULL DEFAULT 0,
	Endurcost           int64          `db:"EndurCost"`            // int(11) NOT NULL DEFAULT 0,
	Endurtimerindex     int64          `db:"EndurTimerIndex"`      // int(11) NOT NULL DEFAULT 0,
	Isdiscipline        int64          `db:"IsDiscipline"`         // int(11) NOT NULL DEFAULT 0,
	Field169            int64          `db:"field169"`             // int(11) NOT NULL DEFAULT 0,
	Field170            int64          `db:"field170"`             // int(11) NOT NULL DEFAULT 0,
	Field171            int64          `db:"field171"`             // int(11) NOT NULL DEFAULT 0,
	Field172            int64          `db:"field172"`             // int(11) NOT NULL DEFAULT 0,
	Hateadded           int64          `db:"HateAdded"`            // int(11) NOT NULL DEFAULT 0,
	Endurupkeep         int64          `db:"EndurUpkeep"`          // int(11) NOT NULL DEFAULT 0,
	Numhitstype         int64          `db:"numhitstype"`          // int(11) NOT NULL DEFAULT 0,
	Numhits             int64          `db:"numhits"`              // int(11) NOT NULL DEFAULT 0,
	Pvpresistbase       int64          `db:"pvpresistbase"`        // int(11) NOT NULL DEFAULT -150,
	Pvpresistcalc       int64          `db:"pvpresistcalc"`        // int(11) NOT NULL DEFAULT 100,
	Pvpresistcap        int64          `db:"pvpresistcap"`         // int(11) NOT NULL DEFAULT -150,
	Spellcategory       int64          `db:"spell_category"`       // int(11) NOT NULL DEFAULT -99,
	Field181            int64          `db:"field181"`             // int(11) NOT NULL DEFAULT 7,
	Field182            int64          `db:"field182"`             // int(11) NOT NULL DEFAULT 65,
	Pcnpconlyflag       sql.NullInt64  `db:"pcnpc_only_flag"`      // int(11) DEFAULT 0,
	Castnotstanding     sql.NullInt64  `db:"cast_not_standing"`    // int(11) DEFAULT 0,
	Canmgb              int64          `db:"can_mgb"`              // int(11) NOT NULL DEFAULT 0,
	Nodispell           int64          `db:"nodispell"`            // int(11) NOT NULL DEFAULT -1,
	Npccategory         int64          `db:"npc_category"`         // int(11) NOT NULL DEFAULT 0,
	Npcusefulness       int64          `db:"npc_usefulness"`       // int(11) NOT NULL DEFAULT 0,
	Minresist           int64          `db:"MinResist"`            // int(11) NOT NULL DEFAULT 0,
	Maxresist           int64          `db:"MaxResist"`            // int(11) NOT NULL DEFAULT 0,
	Viraltargets        int64          `db:"viral_targets"`        // int(11) NOT NULL DEFAULT 0,
	Viraltimer          int64          `db:"viral_timer"`          // int(11) NOT NULL DEFAULT 0,
	Nimbuseffect        sql.NullInt64  `db:"nimbuseffect"`         // int(11) DEFAULT 0,
	Conestartangle      int64          `db:"ConeStartAngle"`       // int(11) NOT NULL DEFAULT 0,
	Conestopangle       int64          `db:"ConeStopAngle"`        // int(11) NOT NULL DEFAULT 0,
	Sneaking            int64          `db:"sneaking"`             // int(11) NOT NULL DEFAULT 0,
	Notextendable       int64          `db:"not_extendable"`       // int(11) NOT NULL DEFAULT 0,
	Field198            int64          `db:"field198"`             // int(11) NOT NULL DEFAULT 0,
	Field199            int64          `db:"field199"`             // int(11) NOT NULL DEFAULT 1,
	Suspendable         sql.NullInt64  `db:"suspendable"`          // int(11) DEFAULT 0,
	Viralrange          int64          `db:"viral_range"`          // int(11) NOT NULL DEFAULT 0,
	Songcap             sql.NullInt64  `db:"songcap"`              // int(11) DEFAULT 0,
	Field203            sql.NullInt64  `db:"field203"`             // int(11) DEFAULT 0,
	Field204            sql.NullInt64  `db:"field204"`             // int(11) DEFAULT 0,
	Noblock             int64          `db:"no_block"`             // int(11) NOT NULL DEFAULT 0,
	Field206            sql.NullInt64  `db:"field206"`             // int(11) DEFAULT -1,
	Spellgroup          sql.NullInt64  `db:"spellgroup"`           // int(11) DEFAULT 0,
	Rank                int64          `db:"rank"`                 // int(11) NOT NULL DEFAULT 0,
	Field209            sql.NullInt64  `db:"field209"`             // int(11) DEFAULT 0,
	Field210            sql.NullInt64  `db:"field210"`             // int(11) DEFAULT 1,
	Castrestriction     int64          `db:"CastRestriction"`      // int(11) NOT NULL DEFAULT 0,
	Allowrest           sql.NullInt64  `db:"allowrest"`            // int(11) DEFAULT 0,
	Incombat            int64          `db:"InCombat"`             // int(11) NOT NULL DEFAULT 0,
	Outofcombat         int64          `db:"OutofCombat"`          // int(11) NOT NULL DEFAULT 0,
	Field215            sql.NullInt64  `db:"field215"`             // int(11) DEFAULT 0,
	Field216            sql.NullInt64  `db:"field216"`             // int(11) DEFAULT 0,
	Field217            sql.NullInt64  `db:"field217"`             // int(11) DEFAULT 0,
	Aemaxtargets        int64          `db:"aemaxtargets"`         // int(11) NOT NULL DEFAULT 0,
	Maxtargets          sql.NullInt64  `db:"maxtargets"`           // int(11) DEFAULT 0,
	Field220            sql.NullInt64  `db:"field220"`             // int(11) DEFAULT 0,
	Field221            sql.NullInt64  `db:"field221"`             // int(11) DEFAULT 0,
	Field222            sql.NullInt64  `db:"field222"`             // int(11) DEFAULT 0,
	Field223            sql.NullInt64  `db:"field223"`             // int(11) DEFAULT 0,
	Persistdeath        sql.NullInt64  `db:"persistdeath"`         // int(11) DEFAULT 0,
	Field225            int64          `db:"field225"`             // int(11) NOT NULL DEFAULT 0,
	Field226            int64          `db:"field226"`             // int(11) NOT NULL DEFAULT 0,
	Mindist             float32        `db:"min_dist"`             // float NOT NULL DEFAULT 0,
	Mindistmod          float32        `db:"min_dist_mod"`         // float NOT NULL DEFAULT 0,
	Maxdist             float32        `db:"max_dist"`             // float NOT NULL DEFAULT 0,
	Maxdistmod          float32        `db:"max_dist_mod"`         // float NOT NULL DEFAULT 0,
	Minrange            int64          `db:"min_range"`            // int(11) NOT NULL DEFAULT 0,
	Field232            int64          `db:"field232"`             // int(11) NOT NULL DEFAULT 0,
	Field233            int64          `db:"field233"`             // int(11) NOT NULL DEFAULT 0,
	Field234            int64          `db:"field234"`             // int(11) NOT NULL DEFAULT 0,
	Field235            int64          `db:"field235"`             // int(11) NOT NULL DEFAULT 0,
	Field236            int64          `db:"field236"`             // int(11) NOT NULL DEFAULT 0,

	Total int64 `db:"total"`
}

// ToProto converts the spell type struct to protobuf
func (s *Spell) ToProto() *pb.Spell {
	spell := &pb.Spell{}
	spell.Id = s.ID
	spell.Name = s.Name.String
	spell.Player1 = s.Player1.String
	spell.Teleportzone = s.Teleportzone.String
	spell.Youcast = s.Youcast.String
	spell.Othercasts = s.Othercasts.String
	spell.Castonyou = s.Castonyou.String
	spell.Castonother = s.Castonother.String
	spell.Spellfades = s.Spellfades.String
	spell.Range = s.Range
	spell.Aoerange = s.Aoerange
	spell.Pushback = s.Pushback
	spell.Pushup = s.Pushup
	spell.Casttime = s.Casttime
	spell.Recoverytime = s.Recoverytime
	spell.Recasttime = s.Recasttime
	spell.Buffdurationformula = s.Buffdurationformula
	spell.Buffduration = s.Buffduration
	spell.Aeduration = s.Aeduration
	spell.Mana = s.Mana
	spell.Effectbasevalue1 = s.Effectbasevalue1
	spell.Effectbasevalue2 = s.Effectbasevalue2
	spell.Effectbasevalue3 = s.Effectbasevalue3
	spell.Effectbasevalue4 = s.Effectbasevalue4
	spell.Effectbasevalue5 = s.Effectbasevalue5
	spell.Effectbasevalue6 = s.Effectbasevalue6
	spell.Effectbasevalue7 = s.Effectbasevalue7
	spell.Effectbasevalue8 = s.Effectbasevalue8
	spell.Effectbasevalue9 = s.Effectbasevalue9
	spell.Effectbasevalue10 = s.Effectbasevalue10
	spell.Effectbasevalue11 = s.Effectbasevalue11
	spell.Effectbasevalue12 = s.Effectbasevalue12
	spell.Effectlimitvalue1 = s.Effectlimitvalue1
	spell.Effectlimitvalue2 = s.Effectlimitvalue2
	spell.Effectlimitvalue3 = s.Effectlimitvalue3
	spell.Effectlimitvalue4 = s.Effectlimitvalue4
	spell.Effectlimitvalue5 = s.Effectlimitvalue5
	spell.Effectlimitvalue6 = s.Effectlimitvalue6
	spell.Effectlimitvalue7 = s.Effectlimitvalue7
	spell.Effectlimitvalue8 = s.Effectlimitvalue8
	spell.Effectlimitvalue9 = s.Effectlimitvalue9
	spell.Effectlimitvalue10 = s.Effectlimitvalue10
	spell.Effectlimitvalue11 = s.Effectlimitvalue11
	spell.Effectlimitvalue12 = s.Effectlimitvalue12
	spell.Max1 = s.Max1
	spell.Max2 = s.Max2
	spell.Max3 = s.Max3
	spell.Max4 = s.Max4
	spell.Max5 = s.Max5
	spell.Max6 = s.Max6
	spell.Max7 = s.Max7
	spell.Max8 = s.Max8
	spell.Max9 = s.Max9
	spell.Max10 = s.Max10
	spell.Max11 = s.Max11
	spell.Max12 = s.Max12
	spell.Icon = s.Icon
	spell.Memicon = s.Memicon
	spell.Components1 = s.Components1
	spell.Components2 = s.Components2
	spell.Components3 = s.Components3
	spell.Components4 = s.Components4
	spell.Componentcounts1 = s.Componentcounts1
	spell.Componentcounts2 = s.Componentcounts2
	spell.Componentcounts3 = s.Componentcounts3
	spell.Componentcounts4 = s.Componentcounts4
	spell.Noexpendreagent1 = s.Noexpendreagent1
	spell.Noexpendreagent2 = s.Noexpendreagent2
	spell.Noexpendreagent3 = s.Noexpendreagent3
	spell.Noexpendreagent4 = s.Noexpendreagent4
	spell.Formula1 = s.Formula1
	spell.Formula2 = s.Formula2
	spell.Formula3 = s.Formula3
	spell.Formula4 = s.Formula4
	spell.Formula5 = s.Formula5
	spell.Formula6 = s.Formula6
	spell.Formula7 = s.Formula7
	spell.Formula8 = s.Formula8
	spell.Formula9 = s.Formula9
	spell.Formula10 = s.Formula10
	spell.Formula11 = s.Formula11
	spell.Formula12 = s.Formula12
	spell.Lighttype = s.Lighttype
	spell.Goodeffect = s.Goodeffect
	spell.Activated = s.Activated
	spell.Resisttype = s.Resisttype
	spell.Effectid1 = s.Effectid1
	spell.Effectid2 = s.Effectid2
	spell.Effectid3 = s.Effectid3
	spell.Effectid4 = s.Effectid4
	spell.Effectid5 = s.Effectid5
	spell.Effectid6 = s.Effectid6
	spell.Effectid7 = s.Effectid7
	spell.Effectid8 = s.Effectid8
	spell.Effectid9 = s.Effectid9
	spell.Effectid10 = s.Effectid10
	spell.Effectid11 = s.Effectid11
	spell.Effectid12 = s.Effectid12
	spell.Targettype = s.Targettype
	spell.Basediff = s.Basediff
	spell.Skill = s.Skill
	spell.Zonetype = s.Zonetype
	spell.Environmenttype = s.Environmenttype
	spell.Timeofday = s.Timeofday
	spell.Classes1 = s.Classes1
	spell.Classes2 = s.Classes2
	spell.Classes3 = s.Classes3
	spell.Classes4 = s.Classes4
	spell.Classes5 = s.Classes5
	spell.Classes6 = s.Classes6
	spell.Classes7 = s.Classes7
	spell.Classes8 = s.Classes8
	spell.Classes9 = s.Classes9
	spell.Classes10 = s.Classes10
	spell.Classes11 = s.Classes11
	spell.Classes12 = s.Classes12
	spell.Classes13 = s.Classes13
	spell.Classes14 = s.Classes14
	spell.Classes15 = s.Classes15
	spell.Classes16 = s.Classes16
	spell.Castinganim = s.Castinganim
	spell.Targetanim = s.Targetanim
	spell.Traveltype = s.Traveltype
	spell.Spellaffectindex = s.Spellaffectindex
	spell.Disallowsit = s.Disallowsit
	spell.Deities0 = s.Deities0
	spell.Deities1 = s.Deities1
	spell.Deities2 = s.Deities2
	spell.Deities3 = s.Deities3
	spell.Deities4 = s.Deities4
	spell.Deities5 = s.Deities5
	spell.Deities6 = s.Deities6
	spell.Deities7 = s.Deities7
	spell.Deities8 = s.Deities8
	spell.Deities9 = s.Deities9
	spell.Deities10 = s.Deities10
	spell.Deities11 = s.Deities11
	spell.Deities12 = s.Deities12
	spell.Deities13 = s.Deities13
	spell.Deities14 = s.Deities14
	spell.Deities15 = s.Deities15
	spell.Deities16 = s.Deities16
	spell.Field142 = s.Field142
	spell.Field143 = s.Field143
	spell.Newicon = s.Newicon
	spell.Spellanim = s.Spellanim
	spell.Uninterruptable = s.Uninterruptable
	spell.Resistdiff = s.Resistdiff
	spell.Dotstackingexempt = s.Dotstackingexempt
	spell.Deleteable = s.Deleteable
	spell.Recourselink = s.Recourselink
	spell.Nopartialresist = s.Nopartialresist
	spell.Field152 = s.Field152
	spell.Field153 = s.Field153
	spell.Shortbuffbox = s.Shortbuffbox
	spell.Descnum = s.Descnum
	spell.Typedescnum = s.Typedescnum.Int64
	spell.Effectdescnum = s.Effectdescnum.Int64
	spell.Effectdescnum2 = s.Effectdescnum2
	spell.Npcnolos = s.Npcnolos
	spell.Field160 = s.Field160
	spell.Reflectable = s.Reflectable
	spell.Bonushate = s.Bonushate
	spell.Field163 = s.Field163
	spell.Field164 = s.Field164
	spell.Ldontrap = s.Ldontrap
	spell.Endurcost = s.Endurcost
	spell.Endurtimerindex = s.Endurtimerindex
	spell.Isdiscipline = s.Isdiscipline
	spell.Field169 = s.Field169
	spell.Field170 = s.Field170
	spell.Field171 = s.Field171
	spell.Field172 = s.Field172
	spell.Hateadded = s.Hateadded
	spell.Endurupkeep = s.Endurupkeep
	spell.Numhitstype = s.Numhitstype
	spell.Numhits = s.Numhits
	spell.Pvpresistbase = s.Pvpresistbase
	spell.Pvpresistcalc = s.Pvpresistcalc
	spell.Pvpresistcap = s.Pvpresistcap
	spell.Spellcategory = s.Spellcategory
	spell.Field181 = s.Field181
	spell.Field182 = s.Field182
	spell.Pcnpconlyflag = s.Pcnpconlyflag.Int64
	spell.Castnotstanding = s.Castnotstanding.Int64
	spell.Canmgb = s.Canmgb
	spell.Nodispell = s.Nodispell
	spell.Npccategory = s.Npccategory
	spell.Npcusefulness = s.Npcusefulness
	spell.Minresist = s.Minresist
	spell.Maxresist = s.Maxresist
	spell.Viraltargets = s.Viraltargets
	spell.Viraltimer = s.Viraltimer
	spell.Nimbuseffect = s.Nimbuseffect.Int64
	spell.Conestartangle = s.Conestartangle
	spell.Conestopangle = s.Conestopangle
	spell.Sneaking = s.Sneaking
	spell.Notextendable = s.Notextendable
	spell.Field198 = s.Field198
	spell.Field199 = s.Field199
	spell.Suspendable = s.Suspendable.Int64
	spell.Viralrange = s.Viralrange
	spell.Songcap = s.Songcap.Int64
	spell.Field203 = s.Field203.Int64
	spell.Field204 = s.Field204.Int64
	spell.Noblock = s.Noblock
	spell.Field206 = s.Field206.Int64
	spell.Spellgroup = s.Spellgroup.Int64
	spell.Rank = s.Rank
	spell.Field209 = s.Field209.Int64
	spell.Field210 = s.Field210.Int64
	spell.Castrestriction = s.Castrestriction
	spell.Allowrest = s.Allowrest.Int64
	spell.Incombat = s.Incombat
	spell.Outofcombat = s.Outofcombat
	spell.Field215 = s.Field215.Int64
	spell.Field216 = s.Field216.Int64
	spell.Field217 = s.Field217.Int64
	spell.Aemaxtargets = s.Aemaxtargets
	spell.Maxtargets = s.Maxtargets.Int64
	spell.Field220 = s.Field220.Int64
	spell.Field221 = s.Field221.Int64
	spell.Field222 = s.Field222.Int64
	spell.Field223 = s.Field223.Int64
	spell.Persistdeath = s.Persistdeath.Int64
	spell.Field225 = s.Field225
	spell.Field226 = s.Field226
	spell.Mindist = s.Mindist
	spell.Mindistmod = s.Mindistmod
	spell.Maxdist = s.Maxdist
	spell.Maxdistmod = s.Maxdistmod
	spell.Minrange = s.Minrange
	spell.Field232 = s.Field232
	spell.Field233 = s.Field233
	spell.Field234 = s.Field234
	spell.Field235 = s.Field235
	spell.Field236 = s.Field236
	return spell
}
