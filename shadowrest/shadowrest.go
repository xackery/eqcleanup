//This script will disable emissary of shadowrest
package shadowrest

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/tools/spawngroup"
	"github.com/xackery/eqemuconfig"
)

var focus = "shadowrest"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {

	//SpawnGroups
	ids := []int64{}

	mids, err := spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Emissary of Shadowrest")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	fmt.Println("Got", len(ids), "Npcs")

	fmt.Println("Removing spell NPC spawn points")

	totalChanged, err := spawngroup.RemoveSpawnGroupAndEntryById(db, ids)
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}

	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in all player-accessible item locations.")
	return
}
