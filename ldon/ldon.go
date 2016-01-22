//This script will disable ldon npcs and quests
package ldon

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"github.com/xackery/eqcleanup/quest"
)

var focus = "ldon"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {

	filePaths := []string{
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
