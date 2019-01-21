package halloween

import (
	"fmt"

	"github.com/dixonwille/wmenu"
	"github.com/xackery/eqcleanup/client"
)

var focus = "halloween"

// Menu represents the wipe menu options
func Menu(opt wmenu.Opt) (err error) {
	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	ok = true
	fmt.Println("This will remove fabled NPCs")

	menu := wmenu.NewMenu("Continue operation?")
	menu.Option("Return to main menu", nil, false, func(opt wmenu.Opt) error {
		ok = false
		return nil
	})
	menu.Option("I understand the risks - continue", nil, false, func(opt wmenu.Opt) error {
		return Clean(c)
	})

	for ok {
		err = menu.Run()
	}
	return
}

// Clean removes halloween NPCs
func Clean(c *client.Client) (err error) {

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

	c.Database.SpawnGroupAndEntryByIDDelete(ids)

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

	totalChanged, err := c.Database.ItemInstanceDelete(ids)
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

	delCount, err := c.Quest.Delete(filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
