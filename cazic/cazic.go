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
	for k, _ := range m {
		q += fmt.Sprintf("%s, ", k)
	}
	q = q[0:len(q)-2] + ") VALUES "
	fmt.Println(q)

	//	if totalChanged, err = result.RowsAffected(); err != nil {
	//		return
	//	}
	fmt.Println("Inserted", totalChanged, "new npcs")

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
		o[f.Name] = v.FieldByIndex([]int{i}).Interface()
	}
	return o
}
