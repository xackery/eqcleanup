package eraitem

import (
	"fmt"

	"github.com/dixonwille/wmenu"
	"github.com/xackery/eqcleanup/client"
)

var focus = "era"

// Menu represents the wipe menu options
func Menu(opt wmenu.Opt) (err error) {
	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	ok = true
	fmt.Println("This will remove out of era items from being obtainable in the game")

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

// Clean removes out of era items
func Clean(c *client.Client) (err error) {

	//Mobs
	ids := []int64{
		111050, //Gellrazz Scalerunner
		25429,  //a dark spirit of nektulos
		59438,  //coldwind blackfoot
		48343,  //immug lashtail
		13138,  //frog invasion
		6251,
	}
	//#checkpoint_ten

	totalRemoved, err := c.Database.SpawnGroupAndEntryByIDDelete(ids)
	if err != nil {
		return
	}
	fmt.Println("Removed", totalRemoved, "spawngroup and entries")

	//Items
	ids = loadIds()
	totalChanged, err := c.Database.ItemInstanceDelete(ids)
	if err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in all player-accessible item locations.")

	filePaths := []string{
		//"befallen/Wraps_McGee.lua",
	}

	delCount, err := c.Quest.Delete(filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}

func loadIds() (ids []int64) {
	ids = []int64{
		84186,  //collector's pickaxe
		110916, //Collector's Lava Protection Mask
		110931, //Collector's Sample Cloak
		19762,  //Shimmering Black Pearl
		37848,
		35086,
		41992,  //glowing othni leggings
		50326,  //simple combatant's orb
		123502, //osmium ore
		34240,  //thallium ore
		44546,  //consigned bite items
		44547,  //consigned bite items
		44548,  //consigned bite items
		44549,  //consigned bite items
		44550,  //consigned bite items
		44551,  //consigned bite items
		44552,  //consigned bite items
		44553,  //consigned bite items
		44554,  //consigned bite items
		44555,  //consigned bite items
		44556,  //consigned bite items
		44557,  //consigned bite items
		44558,  //consigned bite items
		44559,  //consigned bite items
		44560,  //consigned bite items
		56907,  //consigned bite items
		124669, //consigned bite items
		97864,  //Raw Fine Runic Hide
		87563,  //Steamfont Geysers
		41983,  //Glowing Athlai Hat
		36286,  //Ornate Binding Powder
		38691,  //Complex Platinum Silvered Rune
		35085,  //Curzon
		87516,  //Santug's Gift
		79639,  //Silverwing Cloak
		44764,  //grade c gormar venom
		32969,  //cracked leather belt
		62243,  //Exquisite Silk Turban
	}

	for i := int64(62250); i < 62587; i++ { //Ldon ish
		ids = append(ids, i)
	}
	for i := int64(41000); i < 48891; i++ {
		ids = append(ids, i)
	}

	for i := int64(62230); i < 99212; i++ {
		ids = append(ids, i)
	}
	return
}
