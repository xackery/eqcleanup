//This script will disable soulbinders by removing their spawnentry and spawngroups with a wildcard '%soulbinder%'
package rodent

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"github.com/xackery/eqcleanup/quest"
	"github.com/xackery/eqcleanup/spawngroup"
)

var focus = "rodent"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	ids, err := spawngroup.GetSpawnGroupIdsByNameWildcard(db, "a_rodent")
	if err != nil {
		err = fmt.Errorf("Error getting ", focus, " Ids: %s", err.Error())
		return
	}
	if len(ids) < 1 {
		fmt.Println("No ", focus, "were found to delete")
		return
	}

	//Exterminators
	extids, err := spawngroup.GetSpawnGroupIdsByNameWildcard(db, "exterminator_")
	if err != nil {
		err = fmt.Errorf("Error getting exterminator Ids: %s", err.Error())
		return
	}
	if len(extids) > 1 {
		for _, extid := range extids {
			ids = append(ids, extid)
		}
	}

	totalChanged, err := spawngroup.RemoveSpawnGroupAndEntryById(db, ids)
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in spawnentry and spawngroup successfully.")

	filePaths := []string{
		"felwithea/Exterminator_Valern.pl",
		"freeporteast/Exterminator_Larkey.lua",
		"freeporteast/Exterminator_Qalantir.lua",
		"freporte/Exterminator_Larkey.lua",
		"freportw/Exterminator_Qalantir.lua",
		"kaladimb/Exterminator_Vin.pl",
		"neriakb/Exterminator_Damasi.pl",
		"neriakc/Exterminator_Gilea.pl",
		"qeynos/Exterminator_Rasmon.lua",
		"qeynos2/Exterminator_Wintloag.lua",
		"rivervale/Exterminator_Sutten.lua",
	}

	delCount, err := quest.Remove(config, filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
