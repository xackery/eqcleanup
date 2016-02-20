//This script will disable soulbinders by removing their spawnentry and spawngroups with a wildcard '%soulbinder%'
package trickortreat

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/item"
	"github.com/xackery/eqcleanup/quest"
	"github.com/xackery/eqcleanup/spawngroup"
	"github.com/xackery/eqemuconfig"
)

var focus = "trick or treat"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	//Mobs
	ids := []int64{
		20260, //Jack Lanturn
		20259, //Eve_Hallows
		20289, //a_jack_o_lantern
		20269, //an_imp
		20255, //booberella
		20263, //tricksy_treetor
		20272, //a_zombie
		20273,
		20270,
		20271,
		20268,
		20266,
		20267,
		20265,
		20264,
		20257,
		20258,
		20259,
		20261,
		20262,
		20274,
		20256,
		202373,
	}
	//#checkpoint_ten

	spawngroup.RemoveSpawnGroupAndEntryById(db, ids)

	//Items
	ids = []int64{
		84084, //Gummie Kobolds
		84088, //rock candy

		84091, //sand
		84092, //chunk of coal
		84093, //pocket lint
		84094, //Draykey's Codex
		84095, //Trick or treat bag

	}
	totalChanged, err := item.RemoveAllInstancesOfItems(db, ids)
	if err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in all player-accessible item locations.")

	filePaths := []string{
		"befallen/Wraps_McGee.lua",
		"commonlands/Sergeant_Ragus.pl",
		"crescent/#Aragol_Gloomflow.pl",
		"ecommons/Sergeant_Ragus.pl",
		"fieldofbone/Immug_Lashtail.pl",
		"iceclad/Nilham_the_Chef.pl",
		"kithicor/Old_Man_Draykey.pl",
		"mistmoore/Nate.pl",
		"quests/nektulos/Grom_Shives.pl",
		"quests/netherbian/Poil_Lolp.pl",
		"poknowledge/Grand_Librarian_Maelin.pl",
		"qey2hh1/Scary_Miller.lua",
		"rivervale/Laryen_Lycanthrope.pl",
		"tox/Fuzz_Selppa.pl",
		"toxxulia/Fuzz_Selppa.pl",
		"unrest/Crabby_the_Rotten.pl",
		"nektulos/#checkpoint_ten.pl", //trick or treat stop
	}

	delCount, err := quest.Remove(config, filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
