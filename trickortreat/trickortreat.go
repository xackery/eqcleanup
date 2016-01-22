//This script will disable soulbinders by removing their spawnentry and spawngroups with a wildcard '%soulbinder%'
package trickortreat

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"github.com/xackery/eqcleanup/quest"
)

var focus = "trick or treat"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {

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
	}

	delCount, err := quest.Remove(config, filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
