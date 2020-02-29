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

// CharacterSearch implements SCRUD endpoints
func (s *Server) CharacterSearch(ctx context.Context, req *pb.CharacterSearchRequest) (*pb.CharacterSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	resp := new(pb.CharacterSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	character := new(Character)

	st := reflect.TypeOf(*character)
	sv := reflect.ValueOf(character)
	se := sv.Elem()

	query := "SELECT {fieldMap} FROM character_data WHERE"

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
		character := new(Character)
		err = rows.StructScan(character)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = character.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		character := new(Character)
		err = rows.StructScan(character)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Characters = append(resp.Characters, character.ToProto())
	}

	return resp, nil
}

// CharacterCreate implements SCRUD endpoints
func (s *Server) CharacterCreate(ctx context.Context, req *pb.CharacterCreateRequest) (*pb.CharacterCreateResponse, error) {

	character := new(Character)

	st := reflect.TypeOf(*character)

	args := map[string]interface{}{}
	query := "INSERT INTO character_data"

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

	resp := new(pb.CharacterCreateResponse)
	resp.Id = lastID

	return resp, nil
}

// CharacterRead implements SCRUD endpoints
func (s *Server) CharacterRead(ctx context.Context, req *pb.CharacterReadRequest) (*pb.CharacterReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	resp := new(pb.CharacterReadResponse)

	if req.Id < 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	query := "SELECT * FROM character_data WHERE "

	args := map[string]interface{}{}
	query += "id = :id"
	args["id"] = req.Id

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		character := new(Character)
		err = rows.StructScan(character)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Character = character.ToProto()
	}
	return resp, nil
}

// CharacterUpdate implements SCRUD endpoints
func (s *Server) CharacterUpdate(ctx context.Context, req *pb.CharacterUpdateRequest) (*pb.CharacterUpdateResponse, error) {
	character := new(Character)

	st := reflect.TypeOf(*character)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE character_data SET"

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
	resp := new(pb.CharacterUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// CharacterDelete implements SCRUD endpoints
func (s *Server) CharacterDelete(ctx context.Context, req *pb.CharacterDeleteRequest) (*pb.CharacterDeleteResponse, error) {
	query := "DELETE FROM character_data WHERE id = :id LIMIT 1"

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
	resp := new(pb.CharacterDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// CharacterPatch implements SCRUD endpoints
func (s *Server) CharacterPatch(ctx context.Context, req *pb.CharacterPatchRequest) (*pb.CharacterPatchResponse, error) {
	character := new(Character)

	st := reflect.TypeOf(*character)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE character_data SET"

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
	resp := new(pb.CharacterPatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// Character represents an CHARACTER DB binding
type Character struct {
	ID                    int64  `db:"id"`                      // int(11) unsigned NOT NULL AUTO_INCREMENT,
	Accountid             int64  `db:"account_id"`              // int(11) NOT NULL DEFAULT '0',
	Name                  string `db:"name"`                    // varchar(64) NOT NULL DEFAULT '',
	Lastname              string `db:"last_name"`               // varchar(64) NOT NULL DEFAULT '',
	Title                 string `db:"title"`                   // varchar(32) NOT NULL DEFAULT '',
	Suffix                string `db:"suffix"`                  // varchar(32) NOT NULL DEFAULT '',
	Zoneid                int64  `db:"zone_id"`                 // int(11) unsigned NOT NULL DEFAULT '0',
	Zoneinstance          int64  `db:"zone_instance"`           // int(11) unsigned NOT NULL DEFAULT '0',
	Y                     int64  `db:"y"`                       // float NOT NULL DEFAULT '0',
	X                     int64  `db:"x"`                       // float NOT NULL DEFAULT '0',
	Z                     int64  `db:"z"`                       // float NOT NULL DEFAULT '0',
	Heading               int64  `db:"heading"`                 // float NOT NULL DEFAULT '0',
	Gender                int64  `db:"gender"`                  // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Race                  int64  `db:"race"`                    // smallint(11) unsigned NOT NULL DEFAULT '0',
	Class                 int64  `db:"class"`                   // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Level                 int64  `db:"level"`                   // int(11) unsigned NOT NULL DEFAULT '0',
	Deity                 int64  `db:"deity"`                   // int(11) unsigned NOT NULL DEFAULT '0',
	Birthday              int64  `db:"birthday"`                // int(11) unsigned NOT NULL DEFAULT '0',
	Lastlogin             int64  `db:"last_login"`              // int(11) unsigned NOT NULL DEFAULT '0',
	Timeplayed            int64  `db:"time_played"`             // int(11) unsigned NOT NULL DEFAULT '0',
	Level2                int64  `db:"level2"`                  // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Anon                  int64  `db:"anon"`                    // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Gm                    int64  `db:"gm"`                      // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Face                  int64  `db:"face"`                    // int(11) unsigned NOT NULL DEFAULT '0',
	Haircolor             int64  `db:"hair_color"`              // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Hairstyle             int64  `db:"hair_style"`              // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Beard                 int64  `db:"beard"`                   // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Beardcolor            int64  `db:"beard_color"`             // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Eyecolor1             int64  `db:"eye_color_1"`             // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Eyecolor2             int64  `db:"eye_color_2"`             // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Drakkinheritage       int64  `db:"drakkin_heritage"`        // int(11) unsigned NOT NULL DEFAULT '0',
	Drakkintattoo         int64  `db:"drakkin_tattoo"`          // int(11) unsigned NOT NULL DEFAULT '0',
	Drakkindetails        int64  `db:"drakkin_details"`         // int(11) unsigned NOT NULL DEFAULT '0',
	Abilitytimeseconds    int64  `db:"ability_time_seconds"`    // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Abilitynumber         int64  `db:"ability_number"`          // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Abilitytimeminutes    int64  `db:"ability_time_minutes"`    // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Abilitytimehours      int64  `db:"ability_time_hours"`      // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Exp                   int64  `db:"exp"`                     // int(11) unsigned NOT NULL DEFAULT '0',
	Aapointsspent         int64  `db:"aa_points_spent"`         // int(11) unsigned NOT NULL DEFAULT '0',
	Aaexp                 int64  `db:"aa_exp"`                  // int(11) unsigned NOT NULL DEFAULT '0',
	Aapoints              int64  `db:"aa_points"`               // int(11) unsigned NOT NULL DEFAULT '0',
	Groupleadershipexp    int64  `db:"group_leadership_exp"`    // int(11) unsigned NOT NULL DEFAULT '0',
	Raidleadershipexp     int64  `db:"raid_leadership_exp"`     // int(11) unsigned NOT NULL DEFAULT '0',
	Groupleadershippoints int64  `db:"group_leadership_points"` // int(11) unsigned NOT NULL DEFAULT '0',
	Raidleadershippoints  int64  `db:"raid_leadership_points"`  // int(11) unsigned NOT NULL DEFAULT '0',
	Points                int64  `db:"points"`                  // int(11) unsigned NOT NULL DEFAULT '0',
	Curhp                 int64  `db:"cur_hp"`                  // int(11) unsigned NOT NULL DEFAULT '0',
	Mana                  int64  `db:"mana"`                    // int(11) unsigned NOT NULL DEFAULT '0',
	Endurance             int64  `db:"endurance"`               // int(11) unsigned NOT NULL DEFAULT '0',
	Intoxication          int64  `db:"intoxication"`            // int(11) unsigned NOT NULL DEFAULT '0',
	Str                   int64  `db:"str"`                     // int(11) unsigned NOT NULL DEFAULT '0',
	Sta                   int64  `db:"sta"`                     // int(11) unsigned NOT NULL DEFAULT '0',
	Cha                   int64  `db:"cha"`                     // int(11) unsigned NOT NULL DEFAULT '0',
	Dex                   int64  `db:"dex"`                     // int(11) unsigned NOT NULL DEFAULT '0',
	Int                   int64  `db:"int"`                     // int(11) unsigned NOT NULL DEFAULT '0',
	Agi                   int64  `db:"agi"`                     // int(11) unsigned NOT NULL DEFAULT '0',
	Wis                   int64  `db:"wis"`                     // int(11) unsigned NOT NULL DEFAULT '0',
	Zonechangecount       int64  `db:"zone_change_count"`       // int(11) unsigned NOT NULL DEFAULT '0',
	Toxicity              int64  `db:"toxicity"`                // int(11) unsigned NOT NULL DEFAULT '0',
	Hungerlevel           int64  `db:"hunger_level"`            // int(11) unsigned NOT NULL DEFAULT '0',
	Thirstlevel           int64  `db:"thirst_level"`            // int(11) unsigned NOT NULL DEFAULT '0',
	Abilityup             int64  `db:"ability_up"`              // int(11) unsigned NOT NULL DEFAULT '0',
	Ldonpointsguk         int64  `db:"ldon_points_guk"`         // int(11) unsigned NOT NULL DEFAULT '0',
	Ldonpointsmir         int64  `db:"ldon_points_mir"`         // int(11) unsigned NOT NULL DEFAULT '0',
	Ldonpointsmmc         int64  `db:"ldon_points_mmc"`         // int(11) unsigned NOT NULL DEFAULT '0',
	Ldonpointsruj         int64  `db:"ldon_points_ruj"`         // int(11) unsigned NOT NULL DEFAULT '0',
	Ldonpointstak         int64  `db:"ldon_points_tak"`         // int(11) unsigned NOT NULL DEFAULT '0',
	Ldonpointsavailable   int64  `db:"ldon_points_available"`   // int(11) unsigned NOT NULL DEFAULT '0',
	Tributetimeremaining  int64  `db:"tribute_time_remaining"`  // int(11) unsigned NOT NULL DEFAULT '0',
	Careertributepoints   int64  `db:"career_tribute_points"`   // int(11) unsigned NOT NULL DEFAULT '0',
	Tributepoints         int64  `db:"tribute_points"`          // int(11) unsigned NOT NULL DEFAULT '0',
	Tributeactive         int64  `db:"tribute_active"`          // int(11) unsigned NOT NULL DEFAULT '0',
	Pvpstatus             int64  `db:"pvp_status"`              // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Pvpkills              int64  `db:"pvp_kills"`               // int(11) unsigned NOT NULL DEFAULT '0',
	Pvpdeaths             int64  `db:"pvp_deaths"`              // int(11) unsigned NOT NULL DEFAULT '0',
	Pvpcurrentpoints      int64  `db:"pvp_current_points"`      // int(11) unsigned NOT NULL DEFAULT '0',
	Pvpcareerpoints       int64  `db:"pvp_career_points"`       // int(11) unsigned NOT NULL DEFAULT '0',
	Pvpbestkillstreak     int64  `db:"pvp_best_kill_streak"`    // int(11) unsigned NOT NULL DEFAULT '0',
	Pvpworstdeathstreak   int64  `db:"pvp_worst_death_streak"`  // int(11) unsigned NOT NULL DEFAULT '0',
	Pvpcurrentkillstreak  int64  `db:"pvp_current_kill_streak"` // int(11) unsigned NOT NULL DEFAULT '0',
	Pvp2                  int64  `db:"pvp2"`                    // int(11) unsigned NOT NULL DEFAULT '0',
	Pvptype               int64  `db:"pvp_type"`                // int(11) unsigned NOT NULL DEFAULT '0',
	Showhelm              int64  `db:"show_helm"`               // int(11) unsigned NOT NULL DEFAULT '0',
	Groupautoconsent      int64  `db:"group_auto_consent"`      // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Raidautoconsent       int64  `db:"raid_auto_consent"`       // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Guildautoconsent      int64  `db:"guild_auto_consent"`      // tinyint(11) unsigned NOT NULL DEFAULT '0',
	Leadershipexpon       int64  `db:"leadership_exp_on"`       // tinyint(11) unsigned NOT NULL DEFAULT '0',
	RestTimer             int64  `db:"RestTimer"`               // int(11) unsigned NOT NULL DEFAULT '0',
	Airremaining          int64  `db:"air_remaining"`           // int(11) unsigned NOT NULL DEFAULT '0',
	Autosplitenabled      int64  `db:"autosplit_enabled"`       // int(11) unsigned NOT NULL DEFAULT '0',
	Lfp                   int64  `db:"lfp"`                     // tinyint(1) unsigned NOT NULL DEFAULT '0',
	Lfg                   int64  `db:"lfg"`                     // tinyint(1) unsigned NOT NULL DEFAULT '0',
	Mailkey               string `db:"mailkey"`                 // char(16) NOT NULL DEFAULT '',
	Xtargets              int64  `db:"xtargets"`                // tinyint(3) unsigned NOT NULL DEFAULT '5',
	Firstlogon            int64  `db:"firstlogon"`              // tinyint(3) NOT NULL DEFAULT '0',
	Eaaeffects            int64  `db:"e_aa_effects"`            // int(11) unsigned NOT NULL DEFAULT '0',
	Epercenttoaa          int64  `db:"e_percent_to_aa"`         // int(11) unsigned NOT NULL DEFAULT '0',
	Eexpendedaaspent      int64  `db:"e_expended_aa_spent"`     // int(11) unsigned NOT NULL DEFAULT '0',
	Aapointsspentold      int64  `db:"aa_points_spent_old"`     // int(11) unsigned NOT NULL DEFAULT '0',
	Aapointsold           int64  `db:"aa_points_old"`           // int(11) unsigned NOT NULL DEFAULT '0',
	Elastinvsnapshot      int64  `db:"e_last_invsnapshot"`      // int(11) unsigned NOT NULL DEFAULT '0',

	Total int64 `db:"total"`
}

// ToProto converts the character type struct to protobuf
func (c *Character) ToProto() *pb.Character {
	character := &pb.Character{}
	character.Id = c.ID
	character.Accountid = c.Accountid
	character.Name = c.Name
	character.Lastname = c.Lastname
	character.Title = c.Title
	character.Suffix = c.Suffix
	character.Zoneid = c.Zoneid
	character.Zoneinstance = c.Zoneinstance
	character.Y = c.Y
	character.X = c.X
	character.Z = c.Z
	character.Heading = c.Heading
	character.Gender = c.Gender
	character.Race = c.Race
	character.Class = c.Class
	character.Level = c.Level
	character.Deity = c.Deity
	character.Birthday = c.Birthday
	character.Lastlogin = c.Lastlogin
	character.Timeplayed = c.Timeplayed
	character.Level2 = c.Level2
	character.Anon = c.Anon
	character.Gm = c.Gm
	character.Face = c.Face
	character.Haircolor = c.Haircolor
	character.Hairstyle = c.Hairstyle
	character.Beard = c.Beard
	character.Beardcolor = c.Beardcolor
	character.Eyecolor1 = c.Eyecolor1
	character.Eyecolor2 = c.Eyecolor2
	character.Drakkinheritage = c.Drakkinheritage
	character.Drakkintattoo = c.Drakkintattoo
	character.Drakkindetails = c.Drakkindetails
	character.Abilitytimeseconds = c.Abilitytimeseconds
	character.Abilitynumber = c.Abilitynumber
	character.Abilitytimeminutes = c.Abilitytimeminutes
	character.Abilitytimehours = c.Abilitytimehours
	character.Exp = c.Exp
	character.Aapointsspent = c.Aapointsspent
	character.Aaexp = c.Aaexp
	character.Aapoints = c.Aapoints
	character.Groupleadershipexp = c.Groupleadershipexp
	character.Raidleadershipexp = c.Raidleadershipexp
	character.Groupleadershippoints = c.Groupleadershippoints
	character.Raidleadershippoints = c.Raidleadershippoints
	character.Points = c.Points
	character.Curhp = c.Curhp
	character.Mana = c.Mana
	character.Endurance = c.Endurance
	character.Intoxication = c.Intoxication
	character.Str = c.Str
	character.Sta = c.Sta
	character.Cha = c.Cha
	character.Dex = c.Dex
	character.Int = c.Int
	character.Agi = c.Agi
	character.Wis = c.Wis
	character.Zonechangecount = c.Zonechangecount
	character.Toxicity = c.Toxicity
	character.Hungerlevel = c.Hungerlevel
	character.Thirstlevel = c.Thirstlevel
	character.Abilityup = c.Abilityup
	character.Ldonpointsguk = c.Ldonpointsguk
	character.Ldonpointsmir = c.Ldonpointsmir
	character.Ldonpointsmmc = c.Ldonpointsmmc
	character.Ldonpointsruj = c.Ldonpointsruj
	character.Ldonpointstak = c.Ldonpointstak
	character.Ldonpointsavailable = c.Ldonpointsavailable
	character.Tributetimeremaining = c.Tributetimeremaining
	character.Careertributepoints = c.Careertributepoints
	character.Tributepoints = c.Tributepoints
	character.Tributeactive = c.Tributeactive
	character.Pvpstatus = c.Pvpstatus
	character.Pvpkills = c.Pvpkills
	character.Pvpdeaths = c.Pvpdeaths
	character.Pvpcurrentpoints = c.Pvpcurrentpoints
	character.Pvpcareerpoints = c.Pvpcareerpoints
	character.Pvpbestkillstreak = c.Pvpbestkillstreak
	character.Pvpworstdeathstreak = c.Pvpworstdeathstreak
	character.Pvpcurrentkillstreak = c.Pvpcurrentkillstreak
	character.Pvp2 = c.Pvp2
	character.Pvptype = c.Pvptype
	character.Showhelm = c.Showhelm
	character.Groupautoconsent = c.Groupautoconsent
	character.Raidautoconsent = c.Raidautoconsent
	character.Guildautoconsent = c.Guildautoconsent
	character.Leadershipexpon = c.Leadershipexpon
	character.RestTimer = c.RestTimer
	character.Airremaining = c.Airremaining
	character.Autosplitenabled = c.Autosplitenabled
	character.Lfp = c.Lfp
	character.Lfg = c.Lfg
	character.Mailkey = c.Mailkey
	character.Xtargets = c.Xtargets
	character.Firstlogon = c.Firstlogon
	character.Eaaeffects = c.Eaaeffects
	character.Epercenttoaa = c.Epercenttoaa
	character.Eexpendedaaspent = c.Eexpendedaaspent
	character.Aapointsspentold = c.Aapointsspentold
	character.Aapointsold = c.Aapointsold
	character.Elastinvsnapshot = c.Elastinvsnapshot
	return character
}
