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

// NpcSearch implements SCRUD endpoints
func (s *Server) NpcSearch(ctx context.Context, req *pb.NpcSearchRequest) (*pb.NpcSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	resp := new(pb.NpcSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	npc := new(Npc)

	st := reflect.TypeOf(*npc)
	sv := reflect.ValueOf(npc)
	se := sv.Elem()

	query := "SELECT {fieldMap} FROM npc_types WHERE"

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
		npc := new(Npc)
		err = rows.StructScan(npc)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = npc.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		npc := new(Npc)
		err = rows.StructScan(npc)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Npcs = append(resp.Npcs, npc.ToProto())
	}

	return resp, nil
}

// NpcCreate implements SCRUD endpoints
func (s *Server) NpcCreate(ctx context.Context, req *pb.NpcCreateRequest) (*pb.NpcCreateResponse, error) {

	npc := new(Npc)

	st := reflect.TypeOf(*npc)

	args := map[string]interface{}{}
	query := "INSERT INTO npc_types"

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

	resp := new(pb.NpcCreateResponse)
	resp.Id = lastID

	return resp, nil
}

// NpcRead implements SCRUD endpoints
func (s *Server) NpcRead(ctx context.Context, req *pb.NpcReadRequest) (*pb.NpcReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	resp := new(pb.NpcReadResponse)

	if req.Id < 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	query := "SELECT * FROM npc_types WHERE "

	args := map[string]interface{}{}
	query += "id = :id"
	args["id"] = req.Id

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		npc := new(Npc)
		err = rows.StructScan(npc)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Npc = npc.ToProto()
	}
	return resp, nil
}

// NpcUpdate implements SCRUD endpoints
func (s *Server) NpcUpdate(ctx context.Context, req *pb.NpcUpdateRequest) (*pb.NpcUpdateResponse, error) {
	npc := new(Npc)

	st := reflect.TypeOf(*npc)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE npc_types SET"

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
	resp := new(pb.NpcUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// NpcDelete implements SCRUD endpoints
func (s *Server) NpcDelete(ctx context.Context, req *pb.NpcDeleteRequest) (*pb.NpcDeleteResponse, error) {
	query := "DELETE FROM npc_types WHERE id = :id LIMIT 1"

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
	resp := new(pb.NpcDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// NpcPatch implements SCRUD endpoints
func (s *Server) NpcPatch(ctx context.Context, req *pb.NpcPatchRequest) (*pb.NpcPatchResponse, error) {
	npc := new(Npc)

	st := reflect.TypeOf(*npc)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE npc_types SET"

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
	resp := new(pb.NpcPatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// Npc represents an NPC DB binding
type Npc struct {
	ID                   int64          `db:"id"`                     //int(11) NOT NULL AUTO_INCREMENT,
	Name                 string         `db:"name"`                   //text NOT NULL,
	Lastname             sql.NullString `db:"lastname"`               //varchar(32) DEFAULT NULL,
	Level                int64          `db:"level"`                  //tinyint(2) unsigned NOT NULL DEFAULT '0',
	Race                 int64          `db:"race"`                   //smallint(5) unsigned NOT NULL DEFAULT '0',
	Class                int64          `db:"class"`                  //tinyint(2) unsigned NOT NULL DEFAULT '0',
	Bodytype             int64          `db:"bodytype"`               //int(11) NOT NULL DEFAULT '1',
	Hp                   int64          `db:"hp"`                     //int(11) NOT NULL DEFAULT '0',
	Mana                 int64          `db:"mana"`                   //int(11) NOT NULL DEFAULT '0',
	Gender               int64          `db:"gender"`                 //tinyint(2) unsigned NOT NULL DEFAULT '0',
	Texture              int64          `db:"texture"`                //tinyint(2) unsigned NOT NULL DEFAULT '0',
	Helmtexture          int64          `db:"helmtexture"`            //tinyint(2) unsigned NOT NULL DEFAULT '0',
	Herosforgemodel      int64          `db:"herosforgemodel"`        //int(11) NOT NULL DEFAULT '0',
	Size                 float32        `db:"size"`                   //float NOT NULL DEFAULT '0',
	Hpregenrate          int64          `db:"hp_regen_rate"`          //int(11) unsigned NOT NULL DEFAULT '0',
	Manaregenrate        int64          `db:"mana_regen_rate"`        //int(11) unsigned NOT NULL DEFAULT '0',
	Loottableid          int64          `db:"loottable_id"`           //int(11) unsigned NOT NULL DEFAULT '0',
	Merchantid           int64          `db:"merchant_id"`            //int(11) unsigned NOT NULL DEFAULT '0',
	Altcurrencyid        int64          `db:"alt_currency_id"`        //int(11) unsigned NOT NULL DEFAULT '0',
	Npcspellsid          int64          `db:"npc_spells_id"`          //int(11) unsigned NOT NULL DEFAULT '0',
	Npcspellseffectsid   int64          `db:"npc_spells_effects_id"`  //int(11) unsigned NOT NULL DEFAULT '0',
	Npcfactionid         int64          `db:"npc_faction_id"`         //int(11) NOT NULL DEFAULT '0',
	Adventuretemplateid  int64          `db:"adventure_template_id"`  //int(10) unsigned NOT NULL DEFAULT '0',
	Traptemplate         int64          `db:"trap_template"`          //int(10) unsigned DEFAULT '0',
	Mindmg               int64          `db:"mindmg"`                 //int(10) unsigned NOT NULL DEFAULT '0',
	Maxdmg               int64          `db:"maxdmg"`                 //int(10) unsigned NOT NULL DEFAULT '0',
	Attackcount          int64          `db:"attack_count"`           //smallint(6) NOT NULL DEFAULT '-1',
	Npcspecialattks      string         `db:"npcspecialattks"`        //varchar(36) NOT NULL DEFAULT '',
	Specialabilities     string         `db:"special_abilities"`      //text,
	Aggroradius          int64          `db:"aggroradius"`            //int(10) unsigned NOT NULL DEFAULT '0',
	Assistradius         int64          `db:"assistradius"`           //int(10) unsigned NOT NULL DEFAULT '0',
	Face                 int64          `db:"face"`                   //int(10) unsigned NOT NULL DEFAULT '1',
	Luclinhairstyle      int64          `db:"luclin_hairstyle"`       //int(10) unsigned NOT NULL DEFAULT '1',
	Luclinhaircolor      int64          `db:"luclin_haircolor"`       //int(10) unsigned NOT NULL DEFAULT '1',
	Luclineyecolor       int64          `db:"luclin_eyecolor"`        //int(10) unsigned NOT NULL DEFAULT '1',
	Luclineyecolor2      int64          `db:"luclin_eyecolor2"`       //int(10) unsigned NOT NULL DEFAULT '1',
	Luclinbeardcolor     int64          `db:"luclin_beardcolor"`      //int(10) unsigned NOT NULL DEFAULT '1',
	Luclinbeard          int64          `db:"luclin_beard"`           //int(10) unsigned NOT NULL DEFAULT '0',
	Drakkinheritage      int64          `db:"drakkin_heritage"`       //int(10) NOT NULL DEFAULT '0',
	Drakkintattoo        int64          `db:"drakkin_tattoo"`         //int(10) NOT NULL DEFAULT '0',
	Drakkindetails       int64          `db:"drakkin_details"`        //int(10) NOT NULL DEFAULT '0',
	Armortintid          int64          `db:"armortint_id"`           //int(10) unsigned NOT NULL DEFAULT '0',
	Armortintred         int64          `db:"armortint_red"`          //tinyint(3) unsigned NOT NULL DEFAULT '0',
	Armortintgreen       int64          `db:"armortint_green"`        //tinyint(3) unsigned NOT NULL DEFAULT '0',
	Armortintblue        int64          `db:"armortint_blue"`         //tinyint(3) unsigned NOT NULL DEFAULT '0',
	Dmeleetexture1       int64          `db:"d_melee_texture1"`       //int(11) NOT NULL DEFAULT '0',
	Dmeleetexture2       int64          `db:"d_melee_texture2"`       //int(11) NOT NULL DEFAULT '0',
	Ammoidfile           string         `db:"ammo_idfile"`            //varchar(30) NOT NULL DEFAULT 'IT10',
	Primmeleetype        int64          `db:"prim_melee_type"`        //tinyint(4) unsigned NOT NULL DEFAULT '28',
	Secmeleetype         int64          `db:"sec_melee_type"`         //tinyint(4) unsigned NOT NULL DEFAULT '28',
	Rangedtype           int64          `db:"ranged_type"`            //tinyint(4) unsigned NOT NULL DEFAULT '7',
	Runspeed             float32        `db:"runspeed"`               //float NOT NULL DEFAULT '0',
	MR                   int64          `db:"MR"`                     //smallint(5) NOT NULL DEFAULT '0',
	CR                   int64          `db:"CR"`                     //smallint(5) NOT NULL DEFAULT '0',
	DR                   int64          `db:"DR"`                     //smallint(5) NOT NULL DEFAULT '0',
	FR                   int64          `db:"FR"`                     //smallint(5) NOT NULL DEFAULT '0',
	PR                   int64          `db:"PR"`                     //smallint(5) NOT NULL DEFAULT '0',
	Corrup               int64          `db:"Corrup"`                 //smallint(5) NOT NULL DEFAULT '0',
	PhR                  int64          `db:"PhR"`                    //smallint(5) unsigned NOT NULL DEFAULT '0',
	Seeinvis             int64          `db:"see_invis"`              //smallint(4) NOT NULL DEFAULT '0',
	Seeinvisundead       int64          `db:"see_invis_undead"`       //smallint(4) NOT NULL DEFAULT '0',
	Qglobal              int64          `db:"qglobal"`                //int(2) unsigned NOT NULL DEFAULT '0',
	AC                   int64          `db:"AC"`                     //smallint(5) NOT NULL DEFAULT '0',
	Npcaggro             int64          `db:"npc_aggro"`              //tinyint(4) NOT NULL DEFAULT '0',
	Spawnlimit           int64          `db:"spawn_limit"`            //tinyint(4) NOT NULL DEFAULT '0',
	Attackspeed          float32        `db:"attack_speed"`           //float NOT NULL DEFAULT '0',
	Attackdelay          int64          `db:"attack_delay"`           //tinyint(3) unsigned NOT NULL DEFAULT '30',
	Findable             int64          `db:"findable"`               //tinyint(4) NOT NULL DEFAULT '0',
	STR                  int64          `db:"STR"`                    //mediumint(8) unsigned NOT NULL DEFAULT '75',
	STA                  int64          `db:"STA"`                    //mediumint(8) unsigned NOT NULL DEFAULT '75',
	DEX                  int64          `db:"DEX"`                    //mediumint(8) unsigned NOT NULL DEFAULT '75',
	AGI                  int64          `db:"AGI"`                    //mediumint(8) unsigned NOT NULL DEFAULT '75',
	INT                  int64          `db:"_INT"`                   //mediumint(8) unsigned NOT NULL DEFAULT '80',
	WIS                  int64          `db:"WIS"`                    //mediumint(8) unsigned NOT NULL DEFAULT '75',
	CHA                  int64          `db:"CHA"`                    //mediumint(8) unsigned NOT NULL DEFAULT '75',
	Seehide              int64          `db:"see_hide"`               //tinyint(4) NOT NULL DEFAULT '0',
	Seeimprovedhide      int64          `db:"see_improved_hide"`      //tinyint(4) NOT NULL DEFAULT '0',
	Trackable            int64          `db:"trackable"`              //tinyint(4) NOT NULL DEFAULT '1',
	Isbot                int64          `db:"isbot"`                  //tinyint(4) NOT NULL DEFAULT '0',
	Exclude              int64          `db:"exclude"`                //tinyint(4) NOT NULL DEFAULT '1',
	ATK                  int64          `db:"ATK"`                    //mediumint(9) NOT NULL DEFAULT '0',
	Accuracy             int64          `db:"Accuracy"`               //mediumint(9) NOT NULL DEFAULT '0',
	Avoidance            int64          `db:"Avoidance"`              //mediumint(9) unsigned NOT NULL DEFAULT '0',
	Slowmitigation       int64          `db:"slow_mitigation"`        //smallint(4) NOT NULL DEFAULT '0',
	Version              int64          `db:"version"`                //smallint(5) unsigned NOT NULL DEFAULT '0',
	Maxlevel             int64          `db:"maxlevel"`               //tinyint(3) NOT NULL DEFAULT '0',
	Scalerate            int64          `db:"scalerate"`              //int(11) NOT NULL DEFAULT '100',
	Privatecorpse        int64          `db:"private_corpse"`         //tinyint(3) unsigned NOT NULL DEFAULT '0',
	Uniquespawnbyname    int64          `db:"unique_spawn_by_name"`   //tinyint(3) unsigned NOT NULL DEFAULT '0',
	Underwater           int64          `db:"underwater"`             //tinyint(3) unsigned NOT NULL DEFAULT '0',
	Isquest              int64          `db:"isquest"`                //tinyint(3) NOT NULL DEFAULT '0',
	Emoteid              int64          `db:"emoteid"`                //int(10) unsigned NOT NULL DEFAULT '0',
	Spellscale           float32        `db:"spellscale"`             //float NOT NULL DEFAULT '100',
	Healscale            float32        `db:"healscale"`              //float NOT NULL DEFAULT '100',
	Notargethotkey       int64          `db:"no_target_hotkey"`       //tinyint(1) unsigned NOT NULL DEFAULT '0',
	Raidtarget           int64          `db:"raid_target"`            //tinyint(1) unsigned NOT NULL DEFAULT '0',
	Armtexture           int64          `db:"armtexture"`             //tinyint(2) NOT NULL DEFAULT '0',
	Bracertexture        int64          `db:"bracertexture"`          //tinyint(2) NOT NULL DEFAULT '0',
	Handtexture          int64          `db:"handtexture"`            //tinyint(2) NOT NULL DEFAULT '0',
	Legtexture           int64          `db:"legtexture"`             //tinyint(2) NOT NULL DEFAULT '0',
	Feettexture          int64          `db:"feettexture"`            //tinyint(2) NOT NULL DEFAULT '0',
	Light                int64          `db:"light"`                  //tinyint(2) NOT NULL DEFAULT '0',
	Walkspeed            int64          `db:"walkspeed"`              //tinyint(2) NOT NULL DEFAULT '0',
	Peqid                int64          `db:"peqid"`                  //int(11) NOT NULL DEFAULT '0',
	Unique               int64          `db:"unique_"`                //tinyint(2) NOT NULL DEFAULT '0',
	Fixed                int64          `db:"fixed"`                  //tinyint(2) NOT NULL DEFAULT '0',
	Ignoredespawn        int64          `db:"ignore_despawn"`         //tinyint(2) NOT NULL DEFAULT '0',
	Showname             int64          `db:"show_name"`              //tinyint(2) NOT NULL DEFAULT '1',
	Untargetable         int64          `db:"untargetable"`           //tinyint(2) NOT NULL DEFAULT '0',
	Charmac              int64          `db:"charm_ac"`               //smallint(5) DEFAULT '0',
	Charmmindmg          int64          `db:"charm_min_dmg"`          //int(10) DEFAULT '0',
	Charmmaxdmg          int64          `db:"charm_max_dmg"`          //int(10) DEFAULT '0',
	Charmattackdelay     int64          `db:"charm_attack_delay"`     //tinyint(3) DEFAULT '0',
	Charmaccuracyrating  int64          `db:"charm_accuracy_rating"`  //mediumint(9) DEFAULT '0',
	Charmavoidancerating int64          `db:"charm_avoidance_rating"` //mediumint(9) DEFAULT '0',
	Charmatk             int64          `db:"charm_atk"`              //mediumint(9) DEFAULT '0',
	Skipgloballoot       int64          `db:"skip_global_loot"`       //tinyint(4) DEFAULT '0',
	Rarespawn            int64          `db:"rare_spawn"`             //tinyint(4) DEFAULT '0',
	Stuckbehavior        int64          `db:"stuck_behavior"`         //tinyint(4) NOT NULL DEFAULT '0',
	Model                int64          `db:"model"`                  //smallint(5) NOT NULL DEFAULT '0',
	Flymode              int64          `db:"flymode"`                //tinyint(4) NOT NULL DEFAULT '-1',
	Total                int64          `db:"total"`
}

// ToProto converts the npc type struct to protobuf
func (n *Npc) ToProto() *pb.Npc {
	npc := &pb.Npc{}
	npc.Id = n.ID
	npc.Name = n.Name
	npc.Lastname = n.Lastname.String
	npc.Level = n.Level
	npc.Race = n.Race
	npc.Class = n.Class
	npc.Bodytype = n.Bodytype
	npc.Hp = n.Hp
	npc.Mana = n.Mana
	npc.Gender = n.Gender
	npc.Texture = n.Texture
	npc.Helmtexture = n.Helmtexture
	npc.Herosforgemodel = n.Herosforgemodel
	npc.Size = n.Size
	npc.Hpregenrate = n.Hpregenrate
	npc.Manaregenrate = n.Manaregenrate
	npc.Loottableid = n.Loottableid
	npc.Merchantid = n.Merchantid
	npc.Altcurrencyid = n.Altcurrencyid
	npc.Npcspellsid = n.Npcspellsid
	npc.Npcspellseffectsid = n.Npcspellseffectsid
	npc.Npcfactionid = n.Npcfactionid
	npc.Adventuretemplateid = n.Adventuretemplateid
	npc.Traptemplate = n.Traptemplate
	npc.Mindmg = n.Mindmg
	npc.Maxdmg = n.Maxdmg
	npc.Attackcount = n.Attackcount
	npc.Npcspecialattks = n.Npcspecialattks
	npc.Specialabilities = n.Specialabilities
	npc.Aggroradius = n.Aggroradius
	npc.Assistradius = n.Assistradius
	npc.Face = n.Face
	npc.Luclinhairstyle = n.Luclinhairstyle
	npc.Luclinhaircolor = n.Luclinhaircolor
	npc.Luclineyecolor = n.Luclineyecolor
	npc.Luclineyecolor2 = n.Luclineyecolor2
	npc.Luclinbeardcolor = n.Luclinbeardcolor
	npc.Luclinbeard = n.Luclinbeard
	npc.Drakkinheritage = n.Drakkinheritage
	npc.Drakkintattoo = n.Drakkintattoo
	npc.Drakkindetails = n.Drakkindetails
	npc.Armortintid = n.Armortintid
	npc.Armortintred = n.Armortintred
	npc.Armortintgreen = n.Armortintgreen
	npc.Armortintblue = n.Armortintblue
	npc.Dmeleetexture1 = n.Dmeleetexture1
	npc.Dmeleetexture2 = n.Dmeleetexture2
	npc.Ammoidfile = n.Ammoidfile
	npc.Primmeleetype = n.Primmeleetype
	npc.Secmeleetype = n.Secmeleetype
	npc.Rangedtype = n.Rangedtype
	npc.Runspeed = n.Runspeed
	npc.Mr = n.MR
	npc.Cr = n.CR
	npc.Dr = n.DR
	npc.Fr = n.FR
	npc.Pr = n.PR
	npc.Corrup = n.Corrup
	npc.Phr = n.PhR
	npc.Seeinvis = n.Seeinvis
	npc.Seeinvisundead = n.Seeinvisundead
	npc.Qglobal = n.Qglobal
	npc.Ac = n.AC
	npc.Npcaggro = n.Npcaggro
	npc.Spawnlimit = n.Spawnlimit
	npc.Attackspeed = n.Attackspeed
	npc.Attackdelay = n.Attackdelay
	npc.Findable = n.Findable
	npc.Str = n.STR
	npc.Sta = n.STA
	npc.Dex = n.DEX
	npc.Agi = n.AGI
	npc.Int = n.INT
	npc.Wis = n.WIS
	npc.Cha = n.CHA
	npc.Seehide = n.Seehide
	npc.Seeimprovedhide = n.Seeimprovedhide
	npc.Trackable = n.Trackable
	npc.Isbot = n.Isbot
	npc.Exclude = n.Exclude
	npc.Atk = n.ATK
	npc.Accuracy = n.Accuracy
	npc.Avoidance = n.Avoidance
	npc.Slowmitigation = n.Slowmitigation
	npc.Version = n.Version
	npc.Maxlevel = n.Maxlevel
	npc.Scalerate = n.Scalerate
	npc.Privatecorpse = n.Privatecorpse
	npc.Uniquespawnbyname = n.Uniquespawnbyname
	npc.Underwater = n.Underwater
	npc.Isquest = n.Isquest
	npc.Emoteid = n.Emoteid
	npc.Spellscale = n.Spellscale
	npc.Healscale = n.Healscale
	npc.Notargethotkey = n.Notargethotkey
	npc.Raidtarget = n.Raidtarget
	npc.Armtexture = n.Armtexture
	npc.Bracertexture = n.Bracertexture
	npc.Handtexture = n.Handtexture
	npc.Legtexture = n.Legtexture
	npc.Feettexture = n.Feettexture
	npc.Light = n.Light
	npc.Walkspeed = n.Walkspeed
	npc.Peqid = n.Peqid
	npc.Unique = n.Unique
	npc.Fixed = n.Fixed
	npc.Ignoredespawn = n.Ignoredespawn
	npc.Showname = n.Showname
	npc.Untargetable = n.Untargetable
	npc.Charmac = n.Charmac
	npc.Charmmindmg = n.Charmmindmg
	npc.Charmmaxdmg = n.Charmmaxdmg
	npc.Charmattackdelay = n.Charmattackdelay
	npc.Charmaccuracyrating = n.Charmaccuracyrating
	npc.Charmavoidancerating = n.Charmavoidancerating
	npc.Charmatk = n.Charmatk
	npc.Skipgloballoot = n.Skipgloballoot
	npc.Rarespawn = n.Rarespawn
	npc.Stuckbehavior = n.Stuckbehavior
	npc.Model = n.Model
	npc.Flymode = n.Flymode
	return npc
}
