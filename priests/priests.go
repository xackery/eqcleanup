//This script will disable priests of discords by removing their spawnentry and spawngroups with a wildcard '%priest_of_discord%'
package priests

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"github.com/xackery/eqcleanup/quest"
	"github.com/xackery/eqcleanup/spawngroup"
)

var focus = "priests of discord"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	ids, err := spawngroup.GetSpawnGroupIdsByNameWildcard(db, "priest_of_discord")
	if err != nil {
		err = fmt.Errorf("Error getting ", focus, " Ids: %s", err.Error())
		return
	}
	if len(ids) < 1 {
		fmt.Println("No ", focus, "were found to delete")
		return
	}
	totalChanged, err := spawngroup.RemoveSpawnGroupAndEntryById(db, ids)
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in spawnentry and spawngroup successfully.")

	filePaths := []string{
		"global/Priest_of_Discord.lua",
		"kaladima/Priest_of_Discord.pl",
	}

	delCount, err := quest.Remove(config, filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
