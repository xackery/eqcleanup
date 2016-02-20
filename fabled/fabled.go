//This script will disable fabled npcs and quests
package fabled

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/spawngroup"
	"github.com/xackery/eqemuconfig"
)

var focus = "fabled"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
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

	delCount, err := quest.Remove(config, filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
