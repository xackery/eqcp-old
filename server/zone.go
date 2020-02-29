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

// ZoneSearch implements SCRUD endpoints
func (s *Server) ZoneSearch(ctx context.Context, req *pb.ZoneSearchRequest) (*pb.ZoneSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	resp := new(pb.ZoneSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	zone := new(Zone)

	st := reflect.TypeOf(*zone)
	sv := reflect.ValueOf(zone)
	se := sv.Elem()

	query := "SELECT {fieldMap} FROM zone WHERE"

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

			if se.Field(i).Kind() == reflect.String ||
				se.Field(i).Kind() == reflect.Struct {
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
		zone := new(Zone)
		err = rows.StructScan(zone)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = zone.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		zone := new(Zone)
		err = rows.StructScan(zone)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Zones = append(resp.Zones, zone.ToProto())
	}

	return resp, nil
}

// ZoneCreate implements SCRUD endpoints
func (s *Server) ZoneCreate(ctx context.Context, req *pb.ZoneCreateRequest) (*pb.ZoneCreateResponse, error) {

	zone := new(Zone)

	st := reflect.TypeOf(*zone)

	args := map[string]interface{}{}
	query := "INSERT INTO zones"

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

	resp := new(pb.ZoneCreateResponse)
	resp.Id = lastID

	return resp, nil
}

// ZoneRead implements SCRUD endpoints
func (s *Server) ZoneRead(ctx context.Context, req *pb.ZoneReadRequest) (*pb.ZoneReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	resp := new(pb.ZoneReadResponse)

	if req.Id < 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	query := "SELECT * FROM zone WHERE "

	args := map[string]interface{}{}
	query += "id = :id"
	args["id"] = req.Id

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		zone := new(Zone)
		err = rows.StructScan(zone)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Zone = zone.ToProto()
	}
	return resp, nil
}

// ZoneUpdate implements SCRUD endpoints
func (s *Server) ZoneUpdate(ctx context.Context, req *pb.ZoneUpdateRequest) (*pb.ZoneUpdateResponse, error) {
	zone := new(Zone)

	st := reflect.TypeOf(*zone)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE zones SET"

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
	resp := new(pb.ZoneUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// ZoneDelete implements SCRUD endpoints
func (s *Server) ZoneDelete(ctx context.Context, req *pb.ZoneDeleteRequest) (*pb.ZoneDeleteResponse, error) {
	query := "DELETE FROM zone WHERE id = :id LIMIT 1"

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
	resp := new(pb.ZoneDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// ZonePatch implements SCRUD endpoints
func (s *Server) ZonePatch(ctx context.Context, req *pb.ZonePatchRequest) (*pb.ZonePatchResponse, error) {
	zone := new(Zone)

	st := reflect.TypeOf(*zone)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE zones SET"

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
	resp := new(pb.ZonePatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// Zone represents an ZONE DB binding
type Zone struct {
	Shortname              sql.NullString `db:"short_name"`                // varchar(32) DEFAULT NULL,
	ID                     int64          `db:"id"`                        // int(10) NOT NULL AUTO_INCREMENT,
	Filename               sql.NullString `db:"file_name"`                 // varchar(16) DEFAULT NULL,
	Longname               string         `db:"long_name"`                 // text NOT NULL,
	Mapfilename            sql.NullString `db:"map_file_name"`             // varchar(100) DEFAULT NULL,
	Safex                  float32        `db:"safe_x"`                    // float NOT NULL DEFAULT '0',
	Safey                  float32        `db:"safe_y"`                    // float NOT NULL DEFAULT '0',
	Safez                  float32        `db:"safe_z"`                    // float NOT NULL DEFAULT '0',
	Graveyardid            float32        `db:"graveyard_id"`              // float NOT NULL DEFAULT '0',
	Minlevel               int64          `db:"min_level"`                 // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Minstatus              int64          `db:"min_status"`                // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Zoneidnumber           int64          `db:"zoneidnumber"`              // int(4) NOT NULL DEFAULT '0',
	Version                int64          `db:"version"`                   // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Timezone               int64          `db:"timezone"`                  // int(5) NOT NULL DEFAULT '0',
	Maxclients             int64          `db:"maxclients"`                // int(5) NOT NULL DEFAULT '0',
	Ruleset                int64          `db:"ruleset"`                   // int(10) unsigned NOT NULL DEFAULT '0',
	Note                   sql.NullString `db:"note"`                      // varchar(80) DEFAULT NULL,
	Underworld             float32        `db:"underworld"`                // float NOT NULL DEFAULT '0',
	Minclip                float32        `db:"minclip"`                   // float NOT NULL DEFAULT '450',
	Maxclip                float32        `db:"maxclip"`                   // float NOT NULL DEFAULT '450',
	Fogminclip             float32        `db:"fog_minclip"`               // float NOT NULL DEFAULT '450',
	Fogmaxclip             float32        `db:"fog_maxclip"`               // float NOT NULL DEFAULT '450',
	Fogblue                int64          `db:"fog_blue"`                  // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Fogred                 int64          `db:"fog_red"`                   // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Foggreen               int64          `db:"fog_green"`                 // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Sky                    int64          `db:"sky"`                       // tinyint(3) unsigned NOT NULL DEFAULT '1',
	Ztype                  int64          `db:"ztype"`                     // tinyint(3) unsigned NOT NULL DEFAULT '1',
	Zoneexpmultiplier      float32        `db:"zone_exp_multiplier"`       // decimal(6,2) NOT NULL DEFAULT '0.00',
	Walkspeed              float32        `db:"walkspeed"`                 // float NOT NULL DEFAULT '0.4',
	Timetype               int64          `db:"time_type"`                 // tinyint(3) unsigned NOT NULL DEFAULT '2',
	Fogred1                int64          `db:"fog_red1"`                  // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Foggreen1              int64          `db:"fog_green1"`                // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Fogblue1               int64          `db:"fog_blue1"`                 // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Fogminclip1            float32        `db:"fog_minclip1"`              // float NOT NULL DEFAULT '450',
	Fogmaxclip1            float32        `db:"fog_maxclip1"`              // float NOT NULL DEFAULT '450',
	Fogred2                int64          `db:"fog_red2"`                  // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Foggreen2              int64          `db:"fog_green2"`                // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Fogblue2               int64          `db:"fog_blue2"`                 // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Fogminclip2            float32        `db:"fog_minclip2"`              // float NOT NULL DEFAULT '450',
	Fogmaxclip2            float32        `db:"fog_maxclip2"`              // float NOT NULL DEFAULT '450',
	Fogred3                int64          `db:"fog_red3"`                  // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Foggreen3              int64          `db:"fog_green3"`                // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Fogblue3               int64          `db:"fog_blue3"`                 // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Fogminclip3            float32        `db:"fog_minclip3"`              // float NOT NULL DEFAULT '450',
	Fogmaxclip3            float32        `db:"fog_maxclip3"`              // float NOT NULL DEFAULT '450',
	Fogred4                int64          `db:"fog_red4"`                  // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Foggreen4              int64          `db:"fog_green4"`                // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Fogblue4               int64          `db:"fog_blue4"`                 // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Fogminclip4            float32        `db:"fog_minclip4"`              // float NOT NULL DEFAULT '450',
	Fogmaxclip4            float32        `db:"fog_maxclip4"`              // float NOT NULL DEFAULT '450',
	Fogdensity             float32        `db:"fog_density"`               // float NOT NULL DEFAULT '0',
	Flagneeded             string         `db:"flag_needed"`               // varchar(128) NOT NULL DEFAULT '',
	Canbind                int64          `db:"canbind"`                   // tinyint(4) NOT NULL DEFAULT '1',
	Cancombat              int64          `db:"cancombat"`                 // tinyint(4) NOT NULL DEFAULT '1',
	Canlevitate            int64          `db:"canlevitate"`               // tinyint(4) NOT NULL DEFAULT '1',
	Castoutdoor            int64          `db:"castoutdoor"`               // tinyint(4) NOT NULL DEFAULT '1',
	Hotzone                int64          `db:"hotzone"`                   // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Insttype               int64          `db:"insttype"`                  // tinyint(1) unsigned zerofill NOT NULL DEFAULT '0',
	Shutdowndelay          int64          `db:"shutdowndelay"`             // bigint(16) unsigned NOT NULL DEFAULT '5000',
	Peqzone                int64          `db:"peqzone"`                   // tinyint(4) NOT NULL DEFAULT '1',
	Expansion              int64          `db:"expansion"`                 // tinyint(3) NOT NULL DEFAULT '0',
	Suspendbuffs           int64          `db:"suspendbuffs"`              // tinyint(1) unsigned NOT NULL DEFAULT '0',
	Rainchance1            int64          `db:"rain_chance1"`              // int(4) NOT NULL DEFAULT '0',
	Rainchance2            int64          `db:"rain_chance2"`              // int(4) NOT NULL DEFAULT '0',
	Rainchance3            int64          `db:"rain_chance3"`              // int(4) NOT NULL DEFAULT '0',
	Rainchance4            int64          `db:"rain_chance4"`              // int(4) NOT NULL DEFAULT '0',
	Rainduration1          int64          `db:"rain_duration1"`            // int(4) NOT NULL DEFAULT '0',
	Rainduration2          int64          `db:"rain_duration2"`            // int(4) NOT NULL DEFAULT '0',
	Rainduration3          int64          `db:"rain_duration3"`            // int(4) NOT NULL DEFAULT '0',
	Rainduration4          int64          `db:"rain_duration4"`            // int(4) NOT NULL DEFAULT '0',
	Snowchance1            int64          `db:"snow_chance1"`              // int(4) NOT NULL DEFAULT '0',
	Snowchance2            int64          `db:"snow_chance2"`              // int(4) NOT NULL DEFAULT '0',
	Snowchance3            int64          `db:"snow_chance3"`              // int(4) NOT NULL DEFAULT '0',
	Snowchance4            int64          `db:"snow_chance4"`              // int(4) NOT NULL DEFAULT '0',
	Snowduration1          int64          `db:"snow_duration1"`            // int(4) NOT NULL DEFAULT '0',
	Snowduration2          int64          `db:"snow_duration2"`            // int(4) NOT NULL DEFAULT '0',
	Snowduration3          int64          `db:"snow_duration3"`            // int(4) NOT NULL DEFAULT '0',
	Snowduration4          int64          `db:"snow_duration4"`            // int(4) NOT NULL DEFAULT '0',
	Gravity                float32        `db:"gravity"`                   // float NOT NULL DEFAULT '0.4',
	Type                   int64          `db:"type"`                      // int(3) NOT NULL DEFAULT '0',
	Skylock                int64          `db:"skylock"`                   // tinyint(4) NOT NULL DEFAULT '0',
	Fastregenhp            int64          `db:"fast_regen_hp"`             // int(11) NOT NULL DEFAULT '180',
	Fastregenmana          int64          `db:"fast_regen_mana"`           // int(11) NOT NULL DEFAULT '180',
	Fastregenendurance     int64          `db:"fast_regen_endurance"`      // int(11) NOT NULL DEFAULT '180',
	Npcmaxaggrodist        int64          `db:"npc_max_aggro_dist"`        // int(11) NOT NULL DEFAULT '600',
	Maxmovementupdaterange int64          `db:"max_movement_update_range"` // int(11) unsigned NOT NULL DEFAULT '600',
	Total                  int64          `db:"total"`
}

// ToProto converts the zone type struct to protobuf
func (z *Zone) ToProto() *pb.Zone {
	zone := &pb.Zone{}
	zone.Shortname = z.Shortname.String
	zone.Id = z.ID
	zone.Filename = z.Filename.String
	zone.Longname = z.Longname
	zone.Mapfilename = z.Mapfilename.String
	zone.Safex = z.Safex
	zone.Safey = z.Safey
	zone.Safez = z.Safez
	zone.Graveyardid = z.Graveyardid
	zone.Minlevel = z.Minlevel
	zone.Minstatus = z.Minstatus
	zone.Zoneidnumber = z.Zoneidnumber
	zone.Version = z.Version
	zone.Timezone = z.Timezone
	zone.Maxclients = z.Maxclients
	zone.Ruleset = z.Ruleset
	zone.Note = z.Note.String
	zone.Underworld = z.Underworld
	zone.Minclip = z.Minclip
	zone.Maxclip = z.Maxclip
	zone.Fogminclip = z.Fogminclip
	zone.Fogmaxclip = z.Fogmaxclip
	zone.Fogblue = z.Fogblue
	zone.Fogred = z.Fogred
	zone.Foggreen = z.Foggreen
	zone.Sky = z.Sky
	zone.Ztype = z.Ztype
	zone.Zoneexpmultiplier = z.Zoneexpmultiplier
	zone.Walkspeed = z.Walkspeed
	zone.Timetype = z.Timetype
	zone.Fogred1 = z.Fogred1
	zone.Foggreen1 = z.Foggreen1
	zone.Fogblue1 = z.Fogblue1
	zone.Fogminclip1 = z.Fogminclip1
	zone.Fogmaxclip1 = z.Fogmaxclip1
	zone.Fogred2 = z.Fogred2
	zone.Foggreen2 = z.Foggreen2
	zone.Fogblue2 = z.Fogblue2
	zone.Fogminclip2 = z.Fogminclip2
	zone.Fogmaxclip2 = z.Fogmaxclip2
	zone.Fogred3 = z.Fogred3
	zone.Foggreen3 = z.Foggreen3
	zone.Fogblue3 = z.Fogblue3
	zone.Fogminclip3 = z.Fogminclip3
	zone.Fogmaxclip3 = z.Fogmaxclip3
	zone.Fogred4 = z.Fogred4
	zone.Foggreen4 = z.Foggreen4
	zone.Fogblue4 = z.Fogblue4
	zone.Fogminclip4 = z.Fogminclip4
	zone.Fogmaxclip4 = z.Fogmaxclip4
	zone.Fogdensity = z.Fogdensity
	zone.Flagneeded = z.Flagneeded
	zone.Canbind = z.Canbind
	zone.Cancombat = z.Cancombat
	zone.Canlevitate = z.Canlevitate
	zone.Castoutdoor = z.Castoutdoor
	zone.Hotzone = z.Hotzone
	zone.Insttype = z.Insttype
	zone.Shutdowndelay = z.Shutdowndelay
	zone.Peqzone = z.Peqzone
	zone.Expansion = z.Expansion
	zone.Suspendbuffs = z.Suspendbuffs
	zone.Rainchance1 = z.Rainchance1
	zone.Rainchance2 = z.Rainchance2
	zone.Rainchance3 = z.Rainchance3
	zone.Rainchance4 = z.Rainchance4
	zone.Rainduration1 = z.Rainduration1
	zone.Rainduration2 = z.Rainduration2
	zone.Rainduration3 = z.Rainduration3
	zone.Rainduration4 = z.Rainduration4
	zone.Snowchance1 = z.Snowchance1
	zone.Snowchance2 = z.Snowchance2
	zone.Snowchance3 = z.Snowchance3
	zone.Snowchance4 = z.Snowchance4
	zone.Snowduration1 = z.Snowduration1
	zone.Snowduration2 = z.Snowduration2
	zone.Snowduration3 = z.Snowduration3
	zone.Snowduration4 = z.Snowduration4
	zone.Gravity = z.Gravity
	zone.Type = z.Type
	zone.Skylock = z.Skylock
	zone.Fastregenhp = z.Fastregenhp
	zone.Fastregenmana = z.Fastregenmana
	zone.Fastregenendurance = z.Fastregenendurance
	zone.Npcmaxaggrodist = z.Npcmaxaggrodist
	zone.Maxmovementupdaterange = z.Maxmovementupdaterange
	return zone
}
