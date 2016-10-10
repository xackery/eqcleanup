//This script will disable fabled npcs and quests
package cazic

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/tools/loot"
	npcpkg "github.com/xackery/eqcleanup/tools/npc"
	"github.com/xackery/eqcleanup/tools/quest"
	"github.com/xackery/eqcleanup/tools/spawngroup"
	"github.com/xackery/eqemuconfig"
	"github.com/xackery/goeq/npc"
	"reflect"
)

var focus = "cazic"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	var totalChanged int64

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

	q := "INSERT INTO npc_types ("
	m := mapFields(&npcids[0])
	fmt.Println(npcids[0].Id, npcids[0].Name)
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
		fmt.Println(npcid)
		npcid.Id.Valid = true
		if _, err = db.NamedExec(q, &npcid); err != nil {
			return
		}
		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "npcs")

	return
}

type M map[string]interface{} // just an alias

func mapFields(x *npc.NpcTypes) M {
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
