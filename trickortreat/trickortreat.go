//This script will disable soulbinders by removing their spawnentry and spawngroups with a wildcard '%soulbinder%'
package trickortreat

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/tools/item"
	"github.com/xackery/eqcleanup/tools/quest"
	"github.com/xackery/eqcleanup/tools/spawngroup"
	"github.com/xackery/eqemuconfig"
)

var focus = "trick or treat"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	//Mobs
	ids := []int64{
		19151, //Laryn lycanthrope
		202373,
		202384, //wicked winnie
		202385, //a wizeened hermit
		20255,
		20255, //booberella
		20256,
		20256,
		20257,
		20257,
		20258,
		20258,
		20259,
		20259, //Eve_Hallows
		20260, //Jack Lanturn
		20261,
		20261,
		20262,
		20262,
		20263, //tricksy_treetor
		20264,
		20264,
		20265,
		20265,
		20266,
		20266,
		20267,
		20267,
		20268,
		20268,
		20269,
		20269, //an_imp
		20270,
		20270,
		20271,
		20271,
		20272,
		20272, //a_zombie
		20273,
		20273,
		20274,
		20274,
		20275,  //Mippie Digs
		20279,  //Old man draykey
		20281,  //Syxa
		20285,  //Crazy Charlie
		20288,  //lurgh
		20289,  //a_jack_o_lantern
		25436,  //grom shives
		38178,  //marta stalwart
		394263, //aragol gloomflow
		48350,
		48352,
		48353,
		48354,
		63099,
		98641,
		98643,
		98989, //lurgh
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
		"unrest/Halloween_Trigger.pl", //trick or treat stop
		"unrest/Casper.pl",
		"unrest/Candy_Man.pl",
		"unrest/Evil_Brain_Eater.pl",
		"unrest/Eviler_Brain_Eater.pl",
		"unrest/Evilerer_Brain_Eater.pl",
		"unrest/Imp-ossible.pl",
		"unrest/Jack_o_Lantern.pl",
		"unrest/Not_So_Evil_Brain_Eater.pl",
		"unrest/Super_Ghoul_of_Unlimited_Power.pl",
		"unrest/Werewolf_of_DOOOOOOOM.pl",
	}

	delCount, err := quest.Remove(config, filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
