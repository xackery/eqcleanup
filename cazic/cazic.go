//This script will disable fabled npcs and quests
package cazic

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/tools/injector"
	"github.com/xackery/eqcleanup/tools/loot"
	npcpkg "github.com/xackery/eqcleanup/tools/npc"
	"github.com/xackery/eqcleanup/tools/quest"
	"github.com/xackery/eqcleanup/tools/spawngroup"
	"github.com/xackery/eqemuconfig"
)

var focus = "cazic"
var zonename = "cazicthule"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	var totalChanged int64
	var result sql.Result
	fmt.Println("Purging out old data about", focus, "...")

	if totalChanged, err = loot.RemoveLootByZone(db, zonename); err != nil {
		return
	}

	if totalChanged, err = spawngroup.RemoveSpawnGroupAndEntryByZone(db, zonename); err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, "spawn group entries")

	if totalChanged, err = npcpkg.RemoveNPCByZone(db, zonename); err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, "npcs")

	if totalChanged, err = quest.RemoveAllQuestsForZone(config, zonename); err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, "quest files")

	fmt.Println("Inserting new data about", focus, "...")

	//Loot
	q := injector.PrepareInsertString(&loottables[0], "loottable")
	totalChanged = 0

	for _, l := range loottables {
		l.Id.Valid = false
		if result, err = db.NamedExec(q, &l); err != nil {
			return
		}
		var lastId int64

		if lastId, err = result.LastInsertId(); err != nil {
			return
		}

		//now that lootable has changed, sync npc ids to use proper loot table ids
		for i, _ := range npctypes {
			if npctypes[i].Loottable_id == int(l.Id.Int64) {
				npctypes[i].Loottable_id = int(lastId)
				return
			}
		}

		for i, _ := range loottableentries {
			if loottableentries[i].Loottable_id == int(l.Id.Int64) {
				loottableentries[i].Loottable_id = int(lastId)
			}
		}

		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "loottable")

	//NPCS
	if len(npctypes) < 1 {
		err = fmt.Errorf("invalid number of npcs to insert")
		return
	}
	q = injector.PrepareInsertString(&npctypes[0], "npc_types")
	//fmt.Println(q)
	totalChanged = 0
	for _, npctype := range npctypes {
		if _, err = db.NamedExec(q, &npctype); err != nil {
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

	q = injector.PrepareInsertString(&spawngroups[0], "spawngroup")

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

	q = injector.PrepareInsertString(&spawns[0], "spawn2")
	totalChanged = 0
	for _, sg := range spawns {
		if _, err = db.NamedExec(q, &sg); err != nil {
			return
		}
		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "spawn2")

	q = injector.PrepareInsertString(&spawnentries[0], "spawnentry")
	totalChanged = 0
	for _, sg := range spawnentries {
		if _, err = db.NamedExec(q, &sg); err != nil {
			return
		}
		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "spawnentry")

	return
}

type M map[string]interface{} // just an alias
