//This script will disable ldon npcs and quests
package ldon

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"github.com/xackery/eqcleanup/quest"
	"github.com/xackery/eqcleanup/spawngroup"
)

var focus = "ldon"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	//Remove Adventure Merchants
	ids, err := spawngroup.GetSpawnGroupIdsByClass(db, 61)
	if err != nil {
		err = fmt.Errorf("Error getting %s Ids: %s", focus, err.Error())
		return
	}

	//Raid Recruiters
	rids := []int64{223, 222, 69987, 69988, 69989, 64252, 64440, 695, 6742, 6743, 983, 982, 67605, 67606, 67607, 50673, 50674, 68307, 4702, 4701, 4698}
	for _, id := range rids {
		ids = append(ids, id)
	}

	//Adventure Recruiters
	rids = []int64{219, 696, 979, 67608, 50677, 4695}
	for _, id := range rids {
		ids = append(ids, id)
	}

	//Teleporter (Magus)
	rids = []int64{38104, 70007, 6745, 216, 976, 46065, 39292, 39723, 67603, 54755, 4693}
	for _, id := range rids {
		ids = append(ids, id)
	}

	filePaths := []string{
		//Teleporter NPCs
		"abysmal/Magus_Pellen.lua",
		"butcher/Magus_Tira.lua",
		"commonlands/Magus_Zeir.pl",
		"ecommons/Magus_Zeir.lua",
		"everfrost/Magus_Delin.lua",
		"guildlobby/Magus_Alaria.lua",
		"natimbi/Magus_Releua.lua",
		"nedaria/Magus_Wenla.lua",
		"northro/Magus_Arindri.pl",
		"nro/Magus_Arindri.lua",
		"southro/Magus_Jerira.pl",
		"sro/Magus_Jerira.lua",

		//Mysterious Voices
		"akanon/player.lua",
		"cabeast/player.pl",
		"cabwest/player.pl",
		"corathus/player.pl",
		"crescent/player.pl",
		"erudnext/player.lua",
		"erudnint/player.lua",
		"felwithea/player.lua",
		"felwitheb/player.lua",
		"freeporteast/player.lua",
		"freeportwest/player.lua",
		"freporte/player.lua",
		"freportn/player.lua",
		"gfaydark/player.pl",
		//"global/global_player.lua",
		"grobb/player.lua",
		"halas/player.lua",
		"neriakb/player.pl",
		"neriakc/player.pl",
		"oggok/player.pl",
		"paineel/player.pl",
		"qey2hh1/player.lua",
		"qeynos/player.lua",
		"qeynos2/player.lua",
		"qrg/player.lua",
		"rathemtn/player.pl",
		"rivervale/player.lua",
		"sharvahl/player.pl",
	}

	delCount, err := quest.Remove(config, filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	fmt.Println("NOTICE: You still need to remove an entry in global/global_player.lua")
	return
}
