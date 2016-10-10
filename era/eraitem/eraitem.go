package eraitem

import (
	"fmt"
	"github.com/xackery/eqcleanup/tools/eqconfig"
	"github.com/xackery/eqcleanup/tools/eqdb"
	"github.com/xackery/eqcleanup/tools/item"
	"github.com/xackery/eqcleanup/tools/quest"
	"github.com/xackery/eqcleanup/tools/spawngroup"
)

var focus = "era"

func Clean(args ...string) (err error) {
	db, err := eqdb.Load()
	if err != nil {
		return
	}
	config, err := eqconfig.Load()
	if err != nil {
		return
	}

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

	spawngroup.RemoveSpawnGroupAndEntryById(db, ids)

	//Items
	ids = loadIds()
	totalChanged, err := item.RemoveAllInstancesOfItems(db, ids)
	if err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in all player-accessible item locations.")

	filePaths := []string{
	//"befallen/Wraps_McGee.lua",
	}

	delCount, err := quest.Remove(config, filePaths)
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

	for i := int64(41000); i < 48891; i++ {
		ids = append(ids, i)
	}

	for i := int64(62230); i < 99212; i++ {
		ids = append(ids, i)
	}
	return
}
