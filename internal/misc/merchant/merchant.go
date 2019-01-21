//This script deletes any empty merchants
package merchant

import (
	"fmt"
	"github.com/xackery/eqcleanup/tools/eqdb"
	"github.com/xackery/eqcleanup/tools/spawngroup"
)

var focus = "merchant"

func Clean(args ...string) (err error) {
	db, err := eqdb.Load()
	if err != nil {
		return
	}

	//SpawnGroups
	ids := []int64{}

	mids, err := spawngroup.GetSpawnGroupIdsByEmptyMerchant(db)
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	fmt.Println("Got", len(ids), "Npcs")

	fmt.Println("Removing empty merchant spawn points")

	totalChanged, err := spawngroup.RemoveSpawnGroupAndEntryById(db, ids)
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}

	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in all player-accessible item locations.")
	return
}
