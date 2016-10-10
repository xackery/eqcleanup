//This script will disable fabled npcs and quests
package fabled

import (
	"fmt"
	"github.com/xackery/eqcleanup/tools/eqdb"
	"github.com/xackery/eqcleanup/tools/spawngroup"
)

var focus = "fabled"

func Clean(args ...string) (err error) {
	db, err := eqdb.Load()
	if err != nil {
		return
	}

	//Remove Adventure Merchants
	ids, err := spawngroup.GetSpawnGroupIdsByNameWildcard(db, "the_fabled")
	if err != nil {
		err = fmt.Errorf("Error getting %s Ids: %s", focus, err.Error())
		return
	}

	totalChanged, err := spawngroup.RemoveSpawnGroupAndEntryById(db, ids)
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in spawnentry and spawngroup successfully.")

	return
}
