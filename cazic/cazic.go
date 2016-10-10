//This script will disable fabled npcs and quests
package cazic

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/tools/loot"
	npcpkg "github.com/xackery/eqcleanup/tools/npc"
	"github.com/xackery/eqcleanup/tools/quest"
	"github.com/xackery/eqcleanup/tools/spawngroup"
	"github.com/xackery/eqemuconfig"
	"reflect"
)

var focus = "cazic"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	var totalChanged int64
	var result sql.Result
	fmt.Println("Purging out old data about cazic thule...")

	if totalChanged, err = loot.RemoveLootByZone(db, "cazicthule"); err != nil {
		return
	}

	if totalChanged, err = spawngroup.RemoveSpawnGroupAndEntryByZone(db, "cazicthule"); err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, "spawn group entries")

	if totalChanged, err = npcpkg.RemoveNPCByZone(db, "cazicthule"); err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, "npcs")

	if totalChanged, err = quest.RemoveAllQuestsForZone(config, "cazicthule"); err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, "quest files")

	fmt.Println("Inserting new data about cazic thule...")

	//NPCS
	if len(npcids) < 1 {
		err = fmt.Errorf("invalid number of npcs to insert")
		return
	}
	q := "INSERT INTO npc_types ("
	m := mapFields(&npcids[0])
	fields := ""
	for k, _ := range m {
		q += fmt.Sprintf("%s, ", k)
		fields += fmt.Sprintf(":%s, ", k)
	}
	fields = fields[0 : len(fields)-2]
	q = q[0:len(q)-2] + ") VALUES (" + fields + ")"
	//fmt.Println(q)
	totalChanged = 0
	for _, npcid := range npcids {
		if _, err = db.NamedExec(q, &npcid); err != nil {
			return
		}
		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "npcs")

	//SPAWNGROUP
	if len(spawngroups) < 1 {
		err = fmt.Errorf("invalid number of npcs to insert")
		return
	}
	q = "INSERT INTO spawngroup ("
	m = mapFields(&spawngroups[0])
	fields = ""
	for k, _ := range m {
		q += fmt.Sprintf("%s, ", k)
		fields += fmt.Sprintf(":%s, ", k)
	}
	fields = fields[0 : len(fields)-2]
	q = q[0:len(q)-2] + ") VALUES (" + fields + ")"
	//fmt.Println(q)
	totalChanged = 0
	for _, sg := range spawngroups {
		sg.Id.Valid = false //disable spawngroup ids

		if result, err = db.NamedExec(q, &sg); err != nil {
			return
		}
		var lastId int64

		if lastId, err = result.LastInsertId(); err != nil {
			return
		}

		//Now that spawngroup has changed, affect other records that depend on spawngroupid.
		for i, _ := range spawnentries {
			if spawnentries[i].Spawngroupid == int(sg.Id.Int64) {
				spawnentries[i].Spawngroupid = int(lastId)
			}
		}
		for i, _ := range spawns {
			if spawns[i].Spawngroupid == int(sg.Id.Int64) {
				spawns[i].Spawngroupid = int(lastId)
			}
		}

		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "spawngroups")

	//SPAWN2
	if len(spawngroups) < 1 {
		err = fmt.Errorf("invalid number of npcs to insert")
		return
	}
	q = "INSERT INTO spawn2 ("
	m = mapFields(&spawns[0])
	fields = ""
	for k, _ := range m {
		q += fmt.Sprintf("%s, ", k)
		fields += fmt.Sprintf(":%s, ", k)
	}
	fields = fields[0 : len(fields)-2]
	q = q[0:len(q)-2] + ") VALUES (" + fields + ")"
	//fmt.Println(q)
	totalChanged = 0
	for _, sg := range spawns {
		if result, err = db.NamedExec(q, &sg); err != nil {
			return
		}
		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "spawn2")

	//SPAWNENTRY
	if len(spawngroups) < 1 {
		err = fmt.Errorf("invalid number of npcs to insert")
		return
	}
	q = "INSERT INTO spawnentry ("
	m = mapFields(&spawnentries[0])
	fields = ""
	for k, _ := range m {
		q += fmt.Sprintf("%s, ", k)
		fields += fmt.Sprintf(":%s, ", k)
	}
	fields = fields[0 : len(fields)-2]
	q = q[0:len(q)-2] + ") VALUES (" + fields + ")"
	//fmt.Println(q)
	totalChanged = 0
	for _, sg := range spawnentries {
		if result, err = db.NamedExec(q, &sg); err != nil {
			return
		}
		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "spawnentry")

	return
}

type M map[string]interface{} // just an alias

func mapFields(x interface{}) M {
	o := make(M)
	v := reflect.ValueOf(x).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := t.FieldByIndex([]int{i})
		// skip unexported fields
		if f.PkgPath != "" {
			continue
		}
		if f.Tag.Get("db") == "" {
			continue
		}
		o[f.Tag.Get("db")] = v.FieldByIndex([]int{i}).Interface()

	}
	return o
}
