//This script will disable priests of discords by removing their spawnentry and spawngroups with a wildcard '%priest_of_discord%'
package priest

import (
	"fmt"
	"github.com/xackery/eqcleanup/tools/eqconfig"
	"github.com/xackery/eqcleanup/tools/eqdb"
	"github.com/xackery/eqcleanup/tools/quest"
	"github.com/xackery/eqcleanup/tools/spawngroup"
)

var focus = "priests of discord"

func Clean(args ...string) (err error) {
	db, err := eqdb.Load()
	if err != nil {
		return
	}

	config, err := eqconfig.Load()
	if err != nil {
		return
	}
	ids, err := c.Database.SpawnGroupIDsByNameWildcardSearch("priest_of_discord")
	if err != nil {
		err = fmt.Errorf("Error getting %s Ids: %s", focus, err.Error())
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
