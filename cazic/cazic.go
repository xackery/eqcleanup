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
	if len(lootdrops) < 1 {
		err = fmt.Errorf("lootdrops empty! Please download a newer eqcleanup")
		return
	}
	if len(lootdropentries) < 1 {
		err = fmt.Errorf("lootdropentries empty! Please download a newer eqcleanup")
		return
	}
	if len(loottables) < 1 {
		err = fmt.Errorf("loottables empty! Please download a newer eqcleanup")
		return
	}
	if len(loottableentries) < 1 {
		err = fmt.Errorf("loottablentries empty! Please download a newer eqcleanup")
		return
	}
	if len(npctypes) < 1 {
		err = fmt.Errorf("npctypes empty! Please download a newer eqcleanup")
		return
	}
	if len(spawns) < 1 {
		err = fmt.Errorf("spawns empty! Please download a newer eqcleanup")
		return
	}
	if len(spawnentries) < 1 {
		err = fmt.Errorf("spawnentries empty! Please download a newer eqcleanup")
		return
	}
	if len(spawngroups) < 1 {
		err = fmt.Errorf("spawngroups empty! Please download a newer eqcleanup")
		return
	}
	fmt.Println("Purging out old data about", focus, "...")

	if totalChanged, err = loot.RemoveLootByZone(db, zonename); err != nil {
		return
	}

	if totalChanged, err = spawngroup.RemoveGridByZone(db, zonename); err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, "grid related entries")

	if totalChanged, err = spawngroup.RemoveSpawnGroupAndEntryByZone(db, zonename); err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, "spawn related entries")

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
	var q string
	if q, err = injector.PrepareInsertString(&lootdrops[0], "lootdrop"); err != nil {
		return
	}

	totalChanged = 0

	for _, l := range lootdrops {
		if result, err = db.NamedExec(q, map[string]interface{}{"name": l.Name}); err != nil {
			err = fmt.Errorf("with lootdrop insert: %s", err.Error())
			return
		}
		var lastId int64

		if lastId, err = result.LastInsertId(); err != nil {
			err = fmt.Errorf("with lootdrop last insert: %s", err.Error())
			return
		}

		//now that lootdrops has changed, sync up loottables with proper lootdrop ids
		for i, _ := range loottableentries {
			if loottableentries[i].Lootdrop_id == int(l.Id) {
				loottableentries[i].Lootdrop_id = int(lastId)
			}
		}
		for i, _ := range lootdropentries {
			if lootdropentries[i].Lootdrop_id == int(l.Id) {
				lootdropentries[i].Lootdrop_id = int(lastId)
			}
		}
		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "lootdrops")

	if q, err = injector.PrepareInsertString(&lootdropentries[0], "lootdrop_entries"); err != nil {
		return
	}
	totalChanged = 0

	for _, l := range lootdropentries {
		if result, err = db.NamedExec(q, &l); err != nil {
			err = fmt.Errorf("with lootdropentries insert: %s", err.Error())
			return
		}

		if _, err = result.LastInsertId(); err != nil {
			err = fmt.Errorf("with lootdropentries insert: %s", err.Error())
			return
		}
		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "lootdropentries")

	if q, err = injector.PrepareInsertString(&loottables[0], "loottable"); err != nil {
		return
	}

	for _, l := range loottables {
		l.Id.Valid = false
		if result, err = db.NamedExec(q, &l); err != nil {
			err = fmt.Errorf("with loottable insert: %s", err.Error())
			return
		}
		var lastId int64

		if lastId, err = result.LastInsertId(); err != nil {
			err = fmt.Errorf("with loottable lastid insert: %s", err.Error())
			return
		}

		//now that lootable has changed, sync npc ids to use proper loot table ids
		for i, _ := range npctypes {
			if npctypes[i].Loottable_id == int(l.Id.Int64) {
				npctypes[i].Loottable_id = int(lastId)
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

	if q, err = injector.PrepareInsertString(&loottableentries[0], "loottable_entries"); err != nil {
		return
	}
	totalChanged = 0

	for _, l := range loottableentries {
		if result, err = db.NamedExec(q, &l); err != nil {
			return
		}
		if _, err = result.LastInsertId(); err != nil {
			return
		}

		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "loottable entries")

	//GRID
	if q, err = injector.PrepareInsertString(&grids[0], "grid"); err != nil {
		return
	}
	//fmt.Println(q)
	totalChanged = 0
	for _, g := range grids {
		if _, err = db.NamedExec(q, &g); err != nil {
			return
		}
		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "grids")

	if q, err = injector.PrepareInsertString(&gridentries[0], "grid_entries"); err != nil {
		return
	}
	//fmt.Println(q)
	totalChanged = 0
	for _, g := range gridentries {
		if _, err = db.NamedExec(q, &g); err != nil {
			return
		}
		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "grid_entries")

	//NPCS
	if len(npctypes) < 1 {
		err = fmt.Errorf("invalid number of npcs to insert")
		return
	}
	if q, err = injector.PrepareInsertString(&npctypes[0], "npc_types"); err != nil {
		return
	}
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
	if q, err = injector.PrepareInsertString(&spawngroups[0], "spawngroup"); err != nil {
		return
	}

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

	if q, err = injector.PrepareInsertString(&spawns[0], "spawn2"); err != nil {
		return
	}
	totalChanged = 0
	for _, sg := range spawns {
		if _, err = db.NamedExec(q, &sg); err != nil {
			return
		}
		totalChanged++
	}
	fmt.Println("Inserted", totalChanged, "spawn2")

	if q, err = injector.PrepareInsertString(&spawnentries[0], "spawnentry"); err != nil {
		return
	}
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
