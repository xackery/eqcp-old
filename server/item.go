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

// ItemSearch implements SCRUD endpoints
func (s *Server) ItemSearch(ctx context.Context, req *pb.ItemSearchRequest) (*pb.ItemSearchResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}

	ignores := []string{"total"}

	resp := new(pb.ItemSearchResponse)
	if req.Limit < 1 {
		req.Limit = 10
	}
	if req.Limit > 100 {
		req.Limit = 100
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	item := new(Item)

	st := reflect.TypeOf(*item)
	sv := reflect.ValueOf(item)
	se := sv.Elem()

	query := "SELECT {fieldMap} FROM items WHERE"

	args := map[string]interface{}{}

	comma := ""
	isIgnored := false
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		name := strings.ToLower(field.Name)

		tag, ok := field.Tag.Lookup("db")
		if !ok {
			continue
		}

		if req.Orderby != "" && strings.ToLower(name) == strings.ToLower(req.Orderby) {
			req.Orderby = tag
		}

		for key, value := range req.Values {

			if strings.ToLower(name) != strings.ToLower(key) {
				continue
			}

			isIgnored = false
			for _, ignore := range ignores {
				if name != ignore {
					continue
				}
				isIgnored = true
			}
			if isIgnored {
				continue
			}

			if name == "zone" { //special zone search code
				return s.itemSearchByZone(ctx, req, value)
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
		item := new(Item)
		err = rows.StructScan(item)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = item.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		item := new(Item)
		err = rows.StructScan(item)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Items = append(resp.Items, item.ToProto())
	}

	return resp, nil
}

func (s *Server) itemSearchByZone(ctx context.Context, req *pb.ItemSearchRequest, zone string) (*pb.ItemSearchResponse, error) {
	args := map[string]interface{}{}
	resp := new(pb.ItemSearchResponse)
	query := `
SELECT {fieldMap} FROM items WHERE id IN 
(SELECT item_id FROM lootdrop_entries 
WHERE lootdrop_id IN 
(SELECT lootdrop_id FROM loottable_entries 
WHERE loottable_id IN 
(SELECT loottable_id FROM spawnentry se 
INNER JOIN spawn2 s2 ON s2.spawngroupid = se.spawngroupID 
INNER JOIN npc_types n ON n.id = se.npcid
WHERE s2.enabled = 1 
AND s2.zone = :zone 
AND loottable_id > 0)
)
)`

	args["zone"] = zone

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
		item := new(Item)
		err = rows.StructScan(item)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Total = item.Total
		break
	}

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err = s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		item := new(Item)
		err = rows.StructScan(item)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Items = append(resp.Items, item.ToProto())
	}
	return resp, nil
}

// ItemCreate implements SCRUD endpoints
func (s *Server) ItemCreate(ctx context.Context, req *pb.ItemCreateRequest) (*pb.ItemCreateResponse, error) {

	item := new(Item)

	st := reflect.TypeOf(*item)

	args := map[string]interface{}{}
	query := "INSERT INTO items"

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

	resp := new(pb.ItemCreateResponse)
	resp.Id = lastID

	return resp, nil
}

// ItemRead implements SCRUD endpoints
func (s *Server) ItemRead(ctx context.Context, req *pb.ItemReadRequest) (*pb.ItemReadResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("request nil")
	}
	resp := new(pb.ItemReadResponse)

	if req.Id < 0 {
		return nil, fmt.Errorf("id must be greater than 0")
	}
	query := "SELECT * FROM items WHERE "

	args := map[string]interface{}{}
	query += "id = :id"
	args["id"] = req.Id

	log.Debug().Interface("args", args).Msgf("query: %s", query)
	rows, err := s.db.NamedQueryContext(ctx, query, args)
	if err != nil {
		return nil, errors.Wrap(err, "query failed")
	}

	for rows.Next() {
		item := new(Item)
		err = rows.StructScan(item)
		if err != nil {
			return nil, errors.Wrap(err, "structscan")
		}
		resp.Item = item.ToProto()
	}
	return resp, nil
}

// ItemUpdate implements SCRUD endpoints
func (s *Server) ItemUpdate(ctx context.Context, req *pb.ItemUpdateRequest) (*pb.ItemUpdateResponse, error) {
	item := new(Item)

	st := reflect.TypeOf(*item)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE items SET"

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
	resp := new(pb.ItemUpdateResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// ItemDelete implements SCRUD endpoints
func (s *Server) ItemDelete(ctx context.Context, req *pb.ItemDeleteRequest) (*pb.ItemDeleteResponse, error) {
	query := "DELETE FROM items WHERE id = :id LIMIT 1"

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
	resp := new(pb.ItemDeleteResponse)

	resp.Rowsaffected = rowCount
	return resp, nil
}

// ItemPatch implements SCRUD endpoints
func (s *Server) ItemPatch(ctx context.Context, req *pb.ItemPatchRequest) (*pb.ItemPatchResponse, error) {
	item := new(Item)

	st := reflect.TypeOf(*item)

	args := map[string]interface{}{
		"id": req.Id,
	}
	query := "UPDATE items SET"

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
	resp := new(pb.ItemPatchResponse)
	resp.Rowsaffected = rowCount

	return resp, nil
}

// Item represents an ITEM DB binding
type Item struct {
	ID                  int64          `db:"id"`                  // int(11) NOT NULL DEFAULT '0',
	Minstatus           int64          `db:"minstatus"`           // smallint(5) NOT NULL DEFAULT '0',
	Name                string         `db:"Name"`                // varchar(64) NOT NULL DEFAULT '',
	Aagi                int64          `db:"aagi"`                // int(11) NOT NULL DEFAULT '0',
	Ac                  int64          `db:"ac"`                  // int(11) NOT NULL DEFAULT '0',
	Accuracy            int64          `db:"accuracy"`            // int(11) NOT NULL DEFAULT '0',
	Acha                int64          `db:"acha"`                // int(11) NOT NULL DEFAULT '0',
	Adex                int64          `db:"adex"`                // int(11) NOT NULL DEFAULT '0',
	Aint                int64          `db:"aint"`                // int(11) NOT NULL DEFAULT '0',
	Artifactflag        int64          `db:"artifactflag"`        // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Asta                int64          `db:"asta"`                // int(11) NOT NULL DEFAULT '0',
	Astr                int64          `db:"astr"`                // int(11) NOT NULL DEFAULT '0',
	Attack              int64          `db:"attack"`              // int(11) NOT NULL DEFAULT '0',
	Augrestrict         int64          `db:"augrestrict"`         // int(11) NOT NULL DEFAULT '0',
	Augslot1type        int64          `db:"augslot1type"`        // tinyint(3) NOT NULL DEFAULT '0',
	Augslot1visible     int64          `db:"augslot1visible"`     // tinyint(3) NOT NULL DEFAULT '0',
	Augslot2type        int64          `db:"augslot2type"`        // tinyint(3) NOT NULL DEFAULT '0',
	Augslot2visible     int64          `db:"augslot2visible"`     // tinyint(3) NOT NULL DEFAULT '0',
	Augslot3type        int64          `db:"augslot3type"`        // tinyint(3) NOT NULL DEFAULT '0',
	Augslot3visible     int64          `db:"augslot3visible"`     // tinyint(3) NOT NULL DEFAULT '0',
	Augslot4type        int64          `db:"augslot4type"`        // tinyint(3) NOT NULL DEFAULT '0',
	Augslot4visible     int64          `db:"augslot4visible"`     // tinyint(3) NOT NULL DEFAULT '0',
	Augslot5type        int64          `db:"augslot5type"`        // tinyint(3) NOT NULL DEFAULT '0',
	Augslot5visible     int64          `db:"augslot5visible"`     // tinyint(3) NOT NULL DEFAULT '0',
	Augslot6type        int64          `db:"augslot6type"`        // tinyint(3) NOT NULL DEFAULT '0',
	Augslot6visible     int64          `db:"augslot6visible"`     // tinyint(3) NOT NULL DEFAULT '0',
	Augtype             int64          `db:"augtype"`             // int(11) NOT NULL DEFAULT '0',
	Avoidance           int64          `db:"avoidance"`           // int(11) NOT NULL DEFAULT '0',
	Awis                int64          `db:"awis"`                // int(11) NOT NULL DEFAULT '0',
	Bagsize             int64          `db:"bagsize"`             // int(11) NOT NULL DEFAULT '0',
	Bagslots            int64          `db:"bagslots"`            // int(11) NOT NULL DEFAULT '0',
	Bagtype             int64          `db:"bagtype"`             // int(11) NOT NULL DEFAULT '0',
	Bagwr               int64          `db:"bagwr"`               // int(11) NOT NULL DEFAULT '0',
	Banedmgamt          int64          `db:"banedmgamt"`          // int(11) NOT NULL DEFAULT '0',
	Banedmgraceamt      int64          `db:"banedmgraceamt"`      // int(11) NOT NULL DEFAULT '0',
	Banedmgbody         int64          `db:"banedmgbody"`         // int(11) NOT NULL DEFAULT '0',
	Banedmgrace         int64          `db:"banedmgrace"`         // int(11) NOT NULL DEFAULT '0',
	Bardtype            int64          `db:"bardtype"`            // int(11) NOT NULL DEFAULT '0',
	Bardvalue           int64          `db:"bardvalue"`           // int(11) NOT NULL DEFAULT '0',
	Book                int64          `db:"book"`                // int(11) NOT NULL DEFAULT '0',
	Casttime            int64          `db:"casttime"`            // int(11) NOT NULL DEFAULT '0',
	Casttime2           int64          `db:"casttime_"`           // int(11) NOT NULL DEFAULT '0',
	Charmfile           string         `db:"charmfile"`           // varchar(32) NOT NULL DEFAULT '',
	Charmfileid         string         `db:"charmfileid"`         // varchar(32) NOT NULL DEFAULT '',
	Classes             int64          `db:"classes"`             // int(11) NOT NULL DEFAULT '0',
	Color               int64          `db:"color"`               // int(10) unsigned NOT NULL DEFAULT '0',
	Combateffects       string         `db:"combateffects"`       // varchar(10) NOT NULL DEFAULT '',
	Extradmgskill       int64          `db:"extradmgskill"`       // int(11) NOT NULL DEFAULT '0',
	Extradmgamt         int64          `db:"extradmgamt"`         // int(11) NOT NULL DEFAULT '0',
	Price               int64          `db:"price"`               // int(11) NOT NULL DEFAULT '0',
	Cr                  int64          `db:"cr"`                  // int(11) NOT NULL DEFAULT '0',
	Damage              int64          `db:"damage"`              // int(11) NOT NULL DEFAULT '0',
	Damageshield        int64          `db:"damageshield"`        // int(11) NOT NULL DEFAULT '0',
	Deity               int64          `db:"deity"`               // int(11) NOT NULL DEFAULT '0',
	Delay               int64          `db:"delay"`               // int(11) NOT NULL DEFAULT '0',
	Augdistiller        int64          `db:"augdistiller"`        // int(11) NOT NULL DEFAULT '0',
	Dotshielding        int64          `db:"dotshielding"`        // int(11) NOT NULL DEFAULT '0',
	Dr                  int64          `db:"dr"`                  // int(11) NOT NULL DEFAULT '0',
	Clicktype           int64          `db:"clicktype"`           // int(11) NOT NULL DEFAULT '0',
	Clicklevel2         int64          `db:"clicklevel2"`         // int(11) NOT NULL DEFAULT '0',
	Elemdmgtype         int64          `db:"elemdmgtype"`         // int(11) NOT NULL DEFAULT '0',
	Elemdmgamt          int64          `db:"elemdmgamt"`          // int(11) NOT NULL DEFAULT '0',
	Endur               int64          `db:"endur"`               // int(11) NOT NULL DEFAULT '0',
	Factionamt1         int64          `db:"factionamt1"`         // int(11) NOT NULL DEFAULT '0',
	Factionamt2         int64          `db:"factionamt2"`         // int(11) NOT NULL DEFAULT '0',
	Factionamt3         int64          `db:"factionamt3"`         // int(11) NOT NULL DEFAULT '0',
	Factionamt4         int64          `db:"factionamt4"`         // int(11) NOT NULL DEFAULT '0',
	Factionmod1         int64          `db:"factionmod1"`         // int(11) NOT NULL DEFAULT '0',
	Factionmod2         int64          `db:"factionmod2"`         // int(11) NOT NULL DEFAULT '0',
	Factionmod3         int64          `db:"factionmod3"`         // int(11) NOT NULL DEFAULT '0',
	Factionmod4         int64          `db:"factionmod4"`         // int(11) NOT NULL DEFAULT '0',
	Filename            string         `db:"filename"`            // varchar(32) NOT NULL DEFAULT '',
	Focuseffect         int64          `db:"focuseffect"`         // int(11) NOT NULL DEFAULT '0',
	Fr                  int64          `db:"fr"`                  // int(11) NOT NULL DEFAULT '0',
	Fvnodrop            int64          `db:"fvnodrop"`            // int(11) NOT NULL DEFAULT '0',
	Haste               int64          `db:"haste"`               // int(11) NOT NULL DEFAULT '0',
	Clicklevel          int64          `db:"clicklevel"`          // int(11) NOT NULL DEFAULT '0',
	Hp                  int64          `db:"hp"`                  // int(11) NOT NULL DEFAULT '0',
	Regen               int64          `db:"regen"`               // int(11) NOT NULL DEFAULT '0',
	Icon                int64          `db:"icon"`                // int(11) NOT NULL DEFAULT '0',
	Idfile              string         `db:"idfile"`              // varchar(30) NOT NULL DEFAULT '',
	Itemclass           int64          `db:"itemclass"`           // int(11) NOT NULL DEFAULT '0',
	Itemtype            int64          `db:"itemtype"`            // int(11) NOT NULL DEFAULT '0',
	Ldonprice           int64          `db:"ldonprice"`           // int(11) NOT NULL DEFAULT '0',
	Ldontheme           int64          `db:"ldontheme"`           // int(11) NOT NULL DEFAULT '0',
	Ldonsold            int64          `db:"ldonsold"`            // int(11) NOT NULL DEFAULT '0',
	Light               int64          `db:"light"`               // int(11) NOT NULL DEFAULT '0',
	Lore                string         `db:"lore"`                // varchar(80) NOT NULL DEFAULT '',
	Loregroup           int64          `db:"loregroup"`           // int(11) NOT NULL DEFAULT '0',
	Magic               int64          `db:"magic"`               // int(11) NOT NULL DEFAULT '0',
	Mana                int64          `db:"mana"`                // int(11) NOT NULL DEFAULT '0',
	Manaregen           int64          `db:"manaregen"`           // int(11) NOT NULL DEFAULT '0',
	Enduranceregen      int64          `db:"enduranceregen"`      // int(11) NOT NULL DEFAULT '0',
	Material            int64          `db:"material"`            // int(11) NOT NULL DEFAULT '0',
	Herosforgemodel     int64          `db:"herosforgemodel"`     // int(11) NOT NULL DEFAULT '0',
	Maxcharges          int64          `db:"maxcharges"`          // int(11) NOT NULL DEFAULT '0',
	Mr                  int64          `db:"mr"`                  // int(11) NOT NULL DEFAULT '0',
	Nodrop              int64          `db:"nodrop"`              // int(11) NOT NULL DEFAULT '0',
	Norent              int64          `db:"norent"`              // int(11) NOT NULL DEFAULT '0',
	Pendingloreflag     int64          `db:"pendingloreflag"`     // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Pr                  int64          `db:"pr"`                  // int(11) NOT NULL DEFAULT '0',
	Procrate            int64          `db:"procrate"`            // int(11) NOT NULL DEFAULT '0',
	Races               int64          `db:"races"`               // int(11) NOT NULL DEFAULT '0',
	Range               int64          `db:"range"`               // int(11) NOT NULL DEFAULT '0',
	Reclevel            int64          `db:"reclevel"`            // int(11) NOT NULL DEFAULT '0',
	Recskill            int64          `db:"recskill"`            // int(11) NOT NULL DEFAULT '0',
	Reqlevel            int64          `db:"reqlevel"`            // int(11) NOT NULL DEFAULT '0',
	Sellrate            float64        `db:"sellrate"`            // float NOT NULL DEFAULT '0',
	Shielding           int64          `db:"shielding"`           // int(11) NOT NULL DEFAULT '0',
	Size                int64          `db:"size"`                // int(11) NOT NULL DEFAULT '0',
	Skillmodtype        int64          `db:"skillmodtype"`        // int(11) NOT NULL DEFAULT '0',
	Skillmodvalue       int64          `db:"skillmodvalue"`       // int(11) NOT NULL DEFAULT '0',
	Slots               int64          `db:"slots"`               // int(11) NOT NULL DEFAULT '0',
	Clickeffect         int64          `db:"clickeffect"`         // int(11) NOT NULL DEFAULT '0',
	Spellshield         int64          `db:"spellshield"`         // int(11) NOT NULL DEFAULT '0',
	Strikethrough       int64          `db:"strikethrough"`       // int(11) NOT NULL DEFAULT '0',
	Stunresist          int64          `db:"stunresist"`          // int(11) NOT NULL DEFAULT '0',
	Summonedflag        int64          `db:"summonedflag"`        // tinyint(3) unsigned NOT NULL DEFAULT '0',
	Tradeskills         int64          `db:"tradeskills"`         // int(11) NOT NULL DEFAULT '0',
	Favor               int64          `db:"favor"`               // int(11) NOT NULL DEFAULT '0',
	Weight              int64          `db:"weight"`              // int(11) NOT NULL DEFAULT '0',
	Unk012              int64          `db:"UNK012"`              // int(11) NOT NULL DEFAULT '0',
	Unk013              int64          `db:"UNK013"`              // int(11) NOT NULL DEFAULT '0',
	Benefitflag         int64          `db:"benefitflag"`         // int(11) NOT NULL DEFAULT '0',
	Unk054              int64          `db:"UNK054"`              // int(11) NOT NULL DEFAULT '0',
	Unk059              int64          `db:"UNK059"`              // int(11) NOT NULL DEFAULT '0',
	Booktype            int64          `db:"booktype"`            // int(11) NOT NULL DEFAULT '0',
	Recastdelay         int64          `db:"recastdelay"`         // int(11) NOT NULL DEFAULT '0',
	Recasttype          int64          `db:"recasttype"`          // int(11) NOT NULL DEFAULT '0',
	Guildfavor          int64          `db:"guildfavor"`          // int(11) NOT NULL DEFAULT '0',
	Unk123              int64          `db:"UNK123"`              // int(11) NOT NULL DEFAULT '0',
	Unk124              int64          `db:"UNK124"`              // int(11) NOT NULL DEFAULT '0',
	Attuneable          int64          `db:"attuneable"`          // int(11) NOT NULL DEFAULT '0',
	Nopet               int64          `db:"nopet"`               // int(11) NOT NULL DEFAULT '0',
	Updated             time.Time      `db:"updated"`             // datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
	Comment             string         `db:"comment"`             // varchar(255) NOT NULL DEFAULT '',
	Unk127              int64          `db:"UNK127"`              // int(11) NOT NULL DEFAULT '0',
	Pointtype           int64          `db:"pointtype"`           // int(11) NOT NULL DEFAULT '0',
	Potionbelt          int64          `db:"potionbelt"`          // int(11) NOT NULL DEFAULT '0',
	Potionbeltslots     int64          `db:"potionbeltslots"`     // int(11) NOT NULL DEFAULT '0',
	Stacksize           int64          `db:"stacksize"`           // int(11) NOT NULL DEFAULT '0',
	Notransfer          int64          `db:"notransfer"`          // int(11) NOT NULL DEFAULT '0',
	Stackable           int64          `db:"stackable"`           // int(11) NOT NULL DEFAULT '0',
	Unk134              string         `db:"UNK134"`              // varchar(255) NOT NULL DEFAULT '',
	Unk137              int64          `db:"UNK137"`              // int(11) NOT NULL DEFAULT '0',
	Proceffect          int64          `db:"proceffect"`          // int(11) NOT NULL DEFAULT '0',
	Proctype            int64          `db:"proctype"`            // int(11) NOT NULL DEFAULT '0',
	Proclevel2          int64          `db:"proclevel2"`          // int(11) NOT NULL DEFAULT '0',
	Proclevel           int64          `db:"proclevel"`           // int(11) NOT NULL DEFAULT '0',
	Unk142              int64          `db:"UNK142"`              // int(11) NOT NULL DEFAULT '0',
	Worneffect          int64          `db:"worneffect"`          // int(11) NOT NULL DEFAULT '0',
	Worntype            int64          `db:"worntype"`            // int(11) NOT NULL DEFAULT '0',
	Wornlevel2          int64          `db:"wornlevel2"`          // int(11) NOT NULL DEFAULT '0',
	Wornlevel           int64          `db:"wornlevel"`           // int(11) NOT NULL DEFAULT '0',
	Unk147              int64          `db:"UNK147"`              // int(11) NOT NULL DEFAULT '0',
	Focustype           int64          `db:"focustype"`           // int(11) NOT NULL DEFAULT '0',
	Focuslevel2         int64          `db:"focuslevel2"`         // int(11) NOT NULL DEFAULT '0',
	Focuslevel          int64          `db:"focuslevel"`          // int(11) NOT NULL DEFAULT '0',
	Unk152              int64          `db:"UNK152"`              // int(11) NOT NULL DEFAULT '0',
	Scrolleffect        int64          `db:"scrolleffect"`        // int(11) NOT NULL DEFAULT '0',
	Scrolltype          int64          `db:"scrolltype"`          // int(11) NOT NULL DEFAULT '0',
	Scrolllevel2        int64          `db:"scrolllevel2"`        // int(11) NOT NULL DEFAULT '0',
	Scrolllevel         int64          `db:"scrolllevel"`         // int(11) NOT NULL DEFAULT '0',
	Unk157              int64          `db:"UNK157"`              // int(11) NOT NULL DEFAULT '0',
	Serialized          sql.NullTime   `db:"serialized"`          // datetime DEFAULT NULL,
	Verified            sql.NullTime   `db:"verified"`            // datetime DEFAULT NULL,
	Serialization       sql.NullString `db:"serialization"`       // text,
	Source              string         `db:"source"`              // varchar(20) NOT NULL DEFAULT '',
	Unk033              int64          `db:"UNK033"`              // int(11) NOT NULL DEFAULT '0',
	Lorefile            string         `db:"lorefile"`            // varchar(32) NOT NULL DEFAULT '',
	Unk014              int64          `db:"UNK014"`              // int(11) NOT NULL DEFAULT '0',
	Svcorruption        int64          `db:"svcorruption"`        // int(11) NOT NULL DEFAULT '0',
	Skillmodmax         int64          `db:"skillmodmax"`         // int(11) NOT NULL DEFAULT '0',
	Unk060              int64          `db:"UNK060"`              // int(11) NOT NULL DEFAULT '0',
	Augslot1unk2        int64          `db:"augslot1unk2"`        // int(11) NOT NULL DEFAULT '0',
	Augslot2unk2        int64          `db:"augslot2unk2"`        // int(11) NOT NULL DEFAULT '0',
	Augslot3unk2        int64          `db:"augslot3unk2"`        // int(11) NOT NULL DEFAULT '0',
	Augslot4unk2        int64          `db:"augslot4unk2"`        // int(11) NOT NULL DEFAULT '0',
	Augslot5unk2        int64          `db:"augslot5unk2"`        // int(11) NOT NULL DEFAULT '0',
	Augslot6unk2        int64          `db:"augslot6unk2"`        // int(11) NOT NULL DEFAULT '0',
	Unk120              int64          `db:"UNK120"`              // int(11) NOT NULL DEFAULT '0',
	Unk121              int64          `db:"UNK121"`              // int(11) NOT NULL DEFAULT '0',
	Questitemflag       int64          `db:"questitemflag"`       // int(11) NOT NULL DEFAULT '0',
	Unk132              sql.NullString `db:"UNK132"`              // text CHARACTER SET utf8,
	Clickunk5           int64          `db:"clickunk5"`           // int(11) NOT NULL DEFAULT '0',
	Clickunk6           string         `db:"clickunk6"`           // varchar(32) NOT NULL DEFAULT '',
	Clickunk7           int64          `db:"clickunk7"`           // int(11) NOT NULL DEFAULT '0',
	Procunk1            int64          `db:"procunk1"`            // int(11) NOT NULL DEFAULT '0',
	Procunk2            int64          `db:"procunk2"`            // int(11) NOT NULL DEFAULT '0',
	Procunk3            int64          `db:"procunk3"`            // int(11) NOT NULL DEFAULT '0',
	Procunk4            int64          `db:"procunk4"`            // int(11) NOT NULL DEFAULT '0',
	Procunk6            string         `db:"procunk6"`            // varchar(32) NOT NULL DEFAULT '',
	Procunk7            int64          `db:"procunk7"`            // int(11) NOT NULL DEFAULT '0',
	Wornunk1            int64          `db:"wornunk1"`            // int(11) NOT NULL DEFAULT '0',
	Wornunk2            int64          `db:"wornunk2"`            // int(11) NOT NULL DEFAULT '0',
	Wornunk3            int64          `db:"wornunk3"`            // int(11) NOT NULL DEFAULT '0',
	Wornunk4            int64          `db:"wornunk4"`            // int(11) NOT NULL DEFAULT '0',
	Wornunk5            int64          `db:"wornunk5"`            // int(11) NOT NULL DEFAULT '0',
	Wornunk6            string         `db:"wornunk6"`            // varchar(32) NOT NULL DEFAULT '',
	Wornunk7            int64          `db:"wornunk7"`            // int(11) NOT NULL DEFAULT '0',
	Focusunk1           int64          `db:"focusunk1"`           // int(11) NOT NULL DEFAULT '0',
	Focusunk2           int64          `db:"focusunk2"`           // int(11) NOT NULL DEFAULT '0',
	Focusunk3           int64          `db:"focusunk3"`           // int(11) NOT NULL DEFAULT '0',
	Focusunk4           int64          `db:"focusunk4"`           // int(11) NOT NULL DEFAULT '0',
	Focusunk5           int64          `db:"focusunk5"`           // int(11) NOT NULL DEFAULT '0',
	Focusunk6           string         `db:"focusunk6"`           // varchar(32) NOT NULL DEFAULT '',
	Focusunk7           int64          `db:"focusunk7"`           // int(11) NOT NULL DEFAULT '0',
	Scrollunk1          int64          `db:"scrollunk1"`          // int(11) NOT NULL DEFAULT '0',
	Scrollunk2          int64          `db:"scrollunk2"`          // int(11) NOT NULL DEFAULT '0',
	Scrollunk3          int64          `db:"scrollunk3"`          // int(11) NOT NULL DEFAULT '0',
	Scrollunk4          int64          `db:"scrollunk4"`          // int(11) NOT NULL DEFAULT '0',
	Scrollunk5          int64          `db:"scrollunk5"`          // int(11) NOT NULL DEFAULT '0',
	Scrollunk6          string         `db:"scrollunk6"`          // varchar(32) NOT NULL DEFAULT '',
	Scrollunk7          int64          `db:"scrollunk7"`          // int(11) NOT NULL DEFAULT '0',
	Unk193              int64          `db:"UNK193"`              // int(11) NOT NULL DEFAULT '0',
	Purity              int64          `db:"purity"`              // int(11) NOT NULL DEFAULT '0',
	Evoitem             int64          `db:"evoitem"`             // int(11) NOT NULL DEFAULT '0',
	Evoid               int64          `db:"evoid"`               // int(11) NOT NULL DEFAULT '0',
	Evolvinglevel       int64          `db:"evolvinglevel"`       // int(11) NOT NULL DEFAULT '0',
	Evomax              int64          `db:"evomax"`              // int(11) NOT NULL DEFAULT '0',
	Clickname           string         `db:"clickname"`           // varchar(64) NOT NULL DEFAULT '',
	Procname            string         `db:"procname"`            // varchar(64) NOT NULL DEFAULT '',
	Wornname            string         `db:"wornname"`            // varchar(64) NOT NULL DEFAULT '',
	Focusname           string         `db:"focusname"`           // varchar(64) NOT NULL DEFAULT '',
	Scrollname          string         `db:"scrollname"`          // varchar(64) NOT NULL DEFAULT '',
	Dsmitigation        int64          `db:"dsmitigation"`        // smallint(6) NOT NULL DEFAULT '0',
	Heroicstr           int64          `db:"heroic_str"`          // smallint(6) NOT NULL DEFAULT '0',
	Heroicint           int64          `db:"heroic_int"`          // smallint(6) NOT NULL DEFAULT '0',
	Heroicwis           int64          `db:"heroic_wis"`          // smallint(6) NOT NULL DEFAULT '0',
	Heroicagi           int64          `db:"heroic_agi"`          // smallint(6) NOT NULL DEFAULT '0',
	Heroicdex           int64          `db:"heroic_dex"`          // smallint(6) NOT NULL DEFAULT '0',
	Heroicsta           int64          `db:"heroic_sta"`          // smallint(6) NOT NULL DEFAULT '0',
	Heroiccha           int64          `db:"heroic_cha"`          // smallint(6) NOT NULL DEFAULT '0',
	Heroicpr            int64          `db:"heroic_pr"`           // smallint(6) NOT NULL DEFAULT '0',
	Heroicdr            int64          `db:"heroic_dr"`           // smallint(6) NOT NULL DEFAULT '0',
	Heroicfr            int64          `db:"heroic_fr"`           // smallint(6) NOT NULL DEFAULT '0',
	Heroiccr            int64          `db:"heroic_cr"`           // smallint(6) NOT NULL DEFAULT '0',
	Heroicmr            int64          `db:"heroic_mr"`           // smallint(6) NOT NULL DEFAULT '0',
	Heroicsvcorrup      int64          `db:"heroic_svcorrup"`     // smallint(6) NOT NULL DEFAULT '0',
	Healamt             int64          `db:"healamt"`             // smallint(6) NOT NULL DEFAULT '0',
	Spelldmg            int64          `db:"spelldmg"`            // smallint(6) NOT NULL DEFAULT '0',
	Clairvoyance        int64          `db:"clairvoyance"`        // smallint(6) NOT NULL DEFAULT '0',
	Backstabdmg         int64          `db:"backstabdmg"`         // smallint(6) NOT NULL DEFAULT '0',
	Created             string         `db:"created"`             // varchar(64) NOT NULL DEFAULT '',
	Elitematerial       int64          `db:"elitematerial"`       // smallint(6) NOT NULL DEFAULT '0',
	Ldonsellbackrate    int64          `db:"ldonsellbackrate"`    // smallint(6) NOT NULL DEFAULT '0',
	Scriptfileid        int64          `db:"scriptfileid"`        // smallint(6) NOT NULL DEFAULT '0',
	Expendablearrow     int64          `db:"expendablearrow"`     // smallint(6) NOT NULL DEFAULT '0',
	Powersourcecapacity int64          `db:"powersourcecapacity"` // smallint(6) NOT NULL DEFAULT '0',
	Bardeffect          int64          `db:"bardeffect"`          // smallint(6) NOT NULL DEFAULT '0',
	Bardeffecttype      int64          `db:"bardeffecttype"`      // smallint(6) NOT NULL DEFAULT '0',
	Bardlevel2          int64          `db:"bardlevel2"`          // smallint(6) NOT NULL DEFAULT '0',
	Bardlevel           int64          `db:"bardlevel"`           // smallint(6) NOT NULL DEFAULT '0',
	Bardunk1            int64          `db:"bardunk1"`            // smallint(6) NOT NULL DEFAULT '0',
	Bardunk2            int64          `db:"bardunk2"`            // smallint(6) NOT NULL DEFAULT '0',
	Bardunk3            int64          `db:"bardunk3"`            // smallint(6) NOT NULL DEFAULT '0',
	Bardunk4            int64          `db:"bardunk4"`            // smallint(6) NOT NULL DEFAULT '0',
	Bardunk5            int64          `db:"bardunk5"`            // smallint(6) NOT NULL DEFAULT '0',
	Bardname            string         `db:"bardname"`            // varchar(64) NOT NULL DEFAULT '',
	Bardunk7            int64          `db:"bardunk7"`            // smallint(6) NOT NULL DEFAULT '0',
	Unk214              int64          `db:"UNK214"`              // smallint(6) NOT NULL DEFAULT '0',
	Unk219              int64          `db:"UNK219"`              // int(11) NOT NULL DEFAULT '0',
	Unk220              int64          `db:"UNK220"`              // int(11) NOT NULL DEFAULT '0',
	Unk221              int64          `db:"UNK221"`              // int(11) NOT NULL DEFAULT '0',
	Heirloom            int64          `db:"heirloom"`            // int(11) NOT NULL DEFAULT '0',
	Unk223              int64          `db:"UNK223"`              // int(11) NOT NULL DEFAULT '0',
	Unk224              int64          `db:"UNK224"`              // int(11) NOT NULL DEFAULT '0',
	Unk225              int64          `db:"UNK225"`              // int(11) NOT NULL DEFAULT '0',
	Unk226              int64          `db:"UNK226"`              // int(11) NOT NULL DEFAULT '0',
	Unk227              int64          `db:"UNK227"`              // int(11) NOT NULL DEFAULT '0',
	Unk228              int64          `db:"UNK228"`              // int(11) NOT NULL DEFAULT '0',
	Unk229              int64          `db:"UNK229"`              // int(11) NOT NULL DEFAULT '0',
	Unk230              int64          `db:"UNK230"`              // int(11) NOT NULL DEFAULT '0',
	Unk231              int64          `db:"UNK231"`              // int(11) NOT NULL DEFAULT '0',
	Unk232              int64          `db:"UNK232"`              // int(11) NOT NULL DEFAULT '0',
	Unk233              int64          `db:"UNK233"`              // int(11) NOT NULL DEFAULT '0',
	Unk234              int64          `db:"UNK234"`              // int(11) NOT NULL DEFAULT '0',
	Placeable           int64          `db:"placeable"`           // int(11) NOT NULL DEFAULT '0',
	Unk236              int64          `db:"UNK236"`              // int(11) NOT NULL DEFAULT '0',
	Unk237              int64          `db:"UNK237"`              // int(11) NOT NULL DEFAULT '0',
	Unk238              int64          `db:"UNK238"`              // int(11) NOT NULL DEFAULT '0',
	Unk239              int64          `db:"UNK239"`              // int(11) NOT NULL DEFAULT '0',
	Unk240              int64          `db:"UNK240"`              // int(11) NOT NULL DEFAULT '0',
	Unk241              int64          `db:"UNK241"`              // int(11) NOT NULL DEFAULT '0',
	Epicitem            int64          `db:"epicitem"`            // int(11) NOT NULL DEFAULT '0',
	Zone                string         `db:"zone"`                //used in search, not an actual field but triggers a new query type
	Total               int64          `db:"total"`               //used in search, item total
}

// ToProto converts the item type struct to protobuf
func (i *Item) ToProto() *pb.Item {
	item := &pb.Item{}
	item.Id = i.ID
	item.Minstatus = i.Minstatus
	item.Name = i.Name
	item.Aagi = i.Aagi
	item.Ac = i.Ac
	item.Accuracy = i.Accuracy
	item.Acha = i.Acha
	item.Adex = i.Adex
	item.Aint = i.Aint
	item.Artifactflag = i.Artifactflag
	item.Asta = i.Asta
	item.Astr = i.Astr
	item.Attack = i.Attack
	item.Augrestrict = i.Augrestrict
	item.Augslot1Type = i.Augslot1type
	item.Augslot1Visible = i.Augslot1visible
	item.Augslot2Type = i.Augslot2type
	item.Augslot2Visible = i.Augslot2visible
	item.Augslot3Type = i.Augslot3type
	item.Augslot3Visible = i.Augslot3visible
	item.Augslot4Type = i.Augslot4type
	item.Augslot4Visible = i.Augslot4visible
	item.Augslot5Type = i.Augslot5type
	item.Augslot5Visible = i.Augslot5visible
	item.Augslot6Type = i.Augslot6type
	item.Augslot6Visible = i.Augslot6visible
	item.Augtype = i.Augtype
	item.Avoidance = i.Avoidance
	item.Awis = i.Awis
	item.Bagsize = i.Bagsize
	item.Bagslots = i.Bagslots
	item.Bagtype = i.Bagtype
	item.Bagwr = i.Bagwr
	item.Banedmgamt = i.Banedmgamt
	item.Banedmgraceamt = i.Banedmgraceamt
	item.Banedmgbody = i.Banedmgbody
	item.Banedmgrace = i.Banedmgrace
	item.Bardtype = i.Bardtype
	item.Bardvalue = i.Bardvalue
	item.Book = i.Book
	item.Casttime = i.Casttime
	item.Casttime2 = i.Casttime2
	item.Charmfile = i.Charmfile
	item.Charmfileid = i.Charmfileid
	item.Classes = i.Classes
	item.Color = i.Color
	item.Combateffects = i.Combateffects
	item.Extradmgskill = i.Extradmgskill
	item.Extradmgamt = i.Extradmgamt
	item.Price = i.Price
	item.Cr = i.Cr
	item.Damage = i.Damage
	item.Damageshield = i.Damageshield
	item.Deity = i.Deity
	item.Delay = i.Delay
	item.Augdistiller = i.Augdistiller
	item.Dotshielding = i.Dotshielding
	item.Dr = i.Dr
	item.Clicktype = i.Clicktype
	item.Clicklevel2 = i.Clicklevel2
	item.Elemdmgtype = i.Elemdmgtype
	item.Elemdmgamt = i.Elemdmgamt
	item.Endur = i.Endur
	item.Factionamt1 = i.Factionamt1
	item.Factionamt2 = i.Factionamt2
	item.Factionamt3 = i.Factionamt3
	item.Factionamt4 = i.Factionamt4
	item.Factionmod1 = i.Factionmod1
	item.Factionmod2 = i.Factionmod2
	item.Factionmod3 = i.Factionmod3
	item.Factionmod4 = i.Factionmod4
	item.Filename = i.Filename
	item.Focuseffect = i.Focuseffect
	item.Fr = i.Fr
	item.Fvnodrop = i.Fvnodrop
	item.Haste = i.Haste
	item.Clicklevel = i.Clicklevel
	item.Hp = i.Hp
	item.Regen = i.Regen
	item.Icon = i.Icon
	item.Idfile = i.Idfile
	item.Itemclass = i.Itemclass
	item.Itemtype = i.Itemtype
	item.Ldonprice = i.Ldonprice
	item.Ldontheme = i.Ldontheme
	item.Ldonsold = i.Ldonsold
	item.Light = i.Light
	item.Lore = i.Lore
	item.Loregroup = i.Loregroup
	item.Magic = i.Magic
	item.Mana = i.Mana
	item.Manaregen = i.Manaregen
	item.Enduranceregen = i.Enduranceregen
	item.Material = i.Material
	item.Herosforgemodel = i.Herosforgemodel
	item.Maxcharges = i.Maxcharges
	item.Mr = i.Mr
	item.Nodrop = i.Nodrop
	item.Norent = i.Norent
	item.Pendingloreflag = i.Pendingloreflag
	item.Pr = i.Pr
	item.Procrate = i.Procrate
	item.Races = i.Races
	item.Range = i.Range
	item.Reclevel = i.Reclevel
	item.Recskill = i.Recskill
	item.Reqlevel = i.Reqlevel
	item.Sellrate = float32(i.Sellrate)
	item.Shielding = i.Shielding
	item.Size = i.Size
	item.Skillmodtype = i.Skillmodtype
	item.Skillmodvalue = i.Skillmodvalue
	item.Slots = i.Slots
	item.Clickeffect = i.Clickeffect
	item.Spellshield = i.Spellshield
	item.Strikethrough = i.Strikethrough
	item.Stunresist = i.Stunresist
	item.Summonedflag = i.Summonedflag
	item.Tradeskills = i.Tradeskills
	item.Favor = i.Favor
	item.Weight = i.Weight
	item.Unk012 = i.Unk012
	item.Unk013 = i.Unk013
	item.Benefitflag = i.Benefitflag
	item.Unk054 = i.Unk054
	item.Unk059 = i.Unk059
	item.Booktype = i.Booktype
	item.Recastdelay = i.Recastdelay
	item.Recasttype = i.Recasttype
	item.Guildfavor = i.Guildfavor
	item.Unk123 = i.Unk123
	item.Unk124 = i.Unk124
	item.Attuneable = i.Attuneable
	item.Nopet = i.Nopet
	item.Updated = i.Updated.Unix()
	item.Comment = i.Comment
	item.Unk127 = i.Unk127
	item.Pointtype = i.Pointtype
	item.Potionbelt = i.Potionbelt
	item.Potionbeltslots = i.Potionbeltslots
	item.Stacksize = i.Stacksize
	item.Notransfer = i.Notransfer
	item.Stackable = i.Stackable
	item.Unk134 = i.Unk134
	item.Unk137 = i.Unk137
	item.Proceffect = i.Proceffect
	item.Proctype = i.Proctype
	item.Proclevel2 = i.Proclevel2
	item.Proclevel = i.Proclevel
	item.Unk142 = i.Unk142
	item.Worneffect = i.Worneffect
	item.Worntype = i.Worntype
	item.Wornlevel2 = i.Wornlevel2
	item.Wornlevel = i.Wornlevel
	item.Unk147 = i.Unk147
	item.Focustype = i.Focustype
	item.Focuslevel2 = i.Focuslevel2
	item.Focuslevel = i.Focuslevel
	item.Unk152 = i.Unk152
	item.Scrolleffect = i.Scrolleffect
	item.Scrolltype = i.Scrolltype
	item.Scrolllevel2 = i.Scrolllevel2
	item.Scrolllevel = i.Scrolllevel
	item.Unk157 = i.Unk157
	item.Serialized = i.Serialized.Time.Unix()
	item.Verified = i.Verified.Time.Unix()
	item.Serialization = i.Serialization.String
	item.Source = i.Source
	item.Unk033 = i.Unk033
	item.Lorefile = i.Lorefile
	item.Unk014 = i.Unk014
	item.Svcorruption = i.Svcorruption
	item.Skillmodmax = i.Skillmodmax
	item.Unk060 = i.Unk060
	item.Augslot1Unk2 = i.Augslot1unk2
	item.Augslot2Unk2 = i.Augslot2unk2
	item.Augslot3Unk2 = i.Augslot3unk2
	item.Augslot4Unk2 = i.Augslot4unk2
	item.Augslot5Unk2 = i.Augslot5unk2
	item.Augslot6Unk2 = i.Augslot6unk2
	item.Unk120 = i.Unk120
	item.Unk121 = i.Unk121
	item.Questitemflag = i.Questitemflag
	item.Unk132 = i.Unk132.String
	item.Clickunk5 = i.Clickunk5
	item.Clickunk6 = i.Clickunk6
	item.Clickunk7 = i.Clickunk7
	item.Procunk1 = i.Procunk1
	item.Procunk2 = i.Procunk2
	item.Procunk3 = i.Procunk3
	item.Procunk4 = i.Procunk4
	item.Procunk6 = i.Procunk6
	item.Procunk7 = i.Procunk7
	item.Wornunk1 = i.Wornunk1
	item.Wornunk2 = i.Wornunk2
	item.Wornunk3 = i.Wornunk3
	item.Wornunk4 = i.Wornunk4
	item.Wornunk5 = i.Wornunk5
	item.Wornunk6 = i.Wornunk6
	item.Wornunk7 = i.Wornunk7
	item.Focusunk1 = i.Focusunk1
	item.Focusunk2 = i.Focusunk2
	item.Focusunk3 = i.Focusunk3
	item.Focusunk4 = i.Focusunk4
	item.Focusunk5 = i.Focusunk5
	item.Focusunk6 = i.Focusunk6
	item.Focusunk7 = i.Focusunk7
	item.Scrollunk1 = i.Scrollunk1
	item.Scrollunk2 = i.Scrollunk2
	item.Scrollunk3 = i.Scrollunk3
	item.Scrollunk4 = i.Scrollunk4
	item.Scrollunk5 = i.Scrollunk5
	item.Scrollunk6 = i.Scrollunk6
	item.Scrollunk7 = i.Scrollunk7
	item.Unk193 = i.Unk193
	item.Purity = i.Purity
	item.Evoitem = i.Evoitem
	item.Evoid = i.Evoid
	item.Evolvinglevel = i.Evolvinglevel
	item.Evomax = i.Evomax
	item.Clickname = i.Clickname
	item.Procname = i.Procname
	item.Wornname = i.Wornname
	item.Focusname = i.Focusname
	item.Scrollname = i.Scrollname
	item.Dsmitigation = i.Dsmitigation
	item.Heroicstr = i.Heroicstr
	item.Heroicint = i.Heroicint
	item.Heroicwis = i.Heroicwis
	item.Heroicagi = i.Heroicagi
	item.Heroicdex = i.Heroicdex
	item.Heroicsta = i.Heroicsta
	item.Heroiccha = i.Heroiccha
	item.Heroicpr = i.Heroicpr
	item.Heroicdr = i.Heroicdr
	item.Heroicfr = i.Heroicfr
	item.Heroiccr = i.Heroiccr
	item.Heroicmr = i.Heroicmr
	item.Heroicsvcorrup = i.Heroicsvcorrup
	item.Healamt = i.Healamt
	item.Spelldmg = i.Spelldmg
	item.Clairvoyance = i.Clairvoyance
	item.Backstabdmg = i.Backstabdmg
	item.Created = i.Created
	item.Elitematerial = i.Elitematerial
	item.Ldonsellbackrate = i.Ldonsellbackrate
	item.Scriptfileid = i.Scriptfileid
	item.Expendablearrow = i.Expendablearrow
	item.Powersourcecapacity = i.Powersourcecapacity
	item.Bardeffect = i.Bardeffect
	item.Bardeffecttype = i.Bardeffecttype
	item.Bardlevel2 = i.Bardlevel2
	item.Bardlevel = i.Bardlevel
	item.Bardunk1 = i.Bardunk1
	item.Bardunk2 = i.Bardunk2
	item.Bardunk3 = i.Bardunk3
	item.Bardunk4 = i.Bardunk4
	item.Bardunk5 = i.Bardunk5
	item.Bardname = i.Bardname
	item.Bardunk7 = i.Bardunk7
	item.Unk214 = i.Unk214
	item.Unk219 = i.Unk219
	item.Unk220 = i.Unk220
	item.Unk221 = i.Unk221
	item.Heirloom = i.Heirloom
	item.Unk223 = i.Unk223
	item.Unk224 = i.Unk224
	item.Unk225 = i.Unk225
	item.Unk226 = i.Unk226
	item.Unk227 = i.Unk227
	item.Unk228 = i.Unk228
	item.Unk229 = i.Unk229
	item.Unk230 = i.Unk230
	item.Unk231 = i.Unk231
	item.Unk232 = i.Unk232
	item.Unk233 = i.Unk233
	item.Unk234 = i.Unk234
	item.Placeable = i.Placeable
	item.Unk236 = i.Unk236
	item.Unk237 = i.Unk237
	item.Unk238 = i.Unk238
	item.Unk239 = i.Unk239
	item.Unk240 = i.Unk240
	item.Unk241 = i.Unk241
	item.Epicitem = i.Epicitem
	return item
}
