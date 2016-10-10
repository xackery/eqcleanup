//This script will disable fabled npcs and quests
package cazic

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/tools/loot"
	"github.com/xackery/eqcleanup/tools/npc"
	"github.com/xackery/eqcleanup/tools/quest"
	"github.com/xackery/eqcleanup/tools/spawngroup"
	"github.com/xackery/eqemuconfig"
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

	if totalChanged, err = npc.RemoveNPCByZone(db, "cazicthule"); err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, "npcs")

	if totalChanged, err = quest.RemoveAllQuestsForZone(config, "cazicthule"); err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, "quest files")

	fmt.Println("Inserting new data about cazic thule...")

	//var rows sql.Rows
	var result sql.Result
	if result, err = db.Exec(npcInsert); err != nil {
		err = fmt.Errorf("Failed to insert npcs: %s", err.Error())
		return
	}

	if totalChanged, err = result.RowsAffected(); err != nil {
		return
	}
	fmt.Println("Inserted", totalChanged, "new npcs")

	return
}
