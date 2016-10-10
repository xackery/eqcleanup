//This script will disable nexus teleporting and announcements
package nexus

import (
	"fmt"
	"github.com/xackery/eqcleanup/tools/eqconfig"
	"github.com/xackery/eqcleanup/tools/eqdb"
	"github.com/xackery/eqcleanup/tools/quest"
	"github.com/xackery/eqcleanup/tools/spawngroup"
)

var focus = "nexus"

func Clean(args ...string) (err error) {
	db, err := eqdb.Load()
	if err != nil {
		return
	}
	config, err := eqconfig.Load()
	if err != nil {
		return
	}
	ids, err := spawngroup.GetSpawnGroupIdsByNameWildcard(db, "nexus_scion")
	if err != nil {
		err = fmt.Errorf("Error getting %s Ids: %s", focus, err.Error())
		return
	}
	vids, err := spawngroup.GetSpawnGroupIdsByNameWildcard(db, "a_mystic_voice")
	for _, id := range vids {
		ids = append(ids, id)
	}

	totalChanged, err := spawngroup.RemoveSpawnGroupAndEntryById(db, ids)
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in spawnentry and spawngroup successfully.")

	filePaths := []string{
		"dreadlands/A_Mystic_Voice.lua",
		"gfaydark/A_Mystic_Voice.pl",
		"greatdivide/A_Mystic_Voice.pl",
		"northkarana/A_Mystic_Voice.lua",
		"tox/A_Mystic_Voice.pl",
		"toxxulia/A_Mystic_Voice.pl",
	}

	delCount, err := quest.Remove(config, filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
