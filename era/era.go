package era

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/item"
	"github.com/xackery/eqcleanup/quest"
	"github.com/xackery/eqcleanup/spawngroup"
	"github.com/xackery/eqemuconfig"
)

var focus = "era"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
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
		69240,  //Collector's Lightstone
		69241,  //Collector's Mistmoore Granite
		69242,  //Collector's Skunk Scent Gland
		69243,  //Collector's Fire Goblin Skin
		69244,  //Collector's Undead Froglok Tongue
		69245,  //Collector's Scythe
		69246,  //Collector's Water Ring
		69247,  //Collector's Kerran Doll
		69248,  //Collector's Preserved Split Paw Eye
		69249,  //Collector's Shark Tooth
		69250,  //Collector's Cheat Sheet
		84005,  //Collector's Fire Hornet Wing
		84006,  //Collector's Drachnid Web Sack
		84007,  //Collector's Iksar Witch Doll
		84008,  //Collector's Brittle Iksar Skull
		84009,  //Collector's Tump Stump
		84010,  //Collector's Sarnak War Braid
		84011,  //Collector's Bloodgill Scale
		84012,  //Collector's Nohope Moss
		84013,  //Collector's Excavator Claws
		84014,  //Collector's Canine
		84015,  //Collector's Cheat Sheet
		84171,  //Collector's Snow Bunny Foot
		84172,  //Collector's Cougar Tail
		84173,  //Collector's Shardwurm Scale
		84174,  //Collector's Wyvern Claw
		84175,  //Collector's Ice Sculpture
		84176,  //Collector's Brontotherium Hoof
		84177,  //Collector's Chetari Ceremonial Staff
		84178,  //Collector's Velium Trinket
		84179,  //Collector's Tizmak Horn
		84180,  //Collector's Kodiak Fang
		84181,  //Collector's Giant Sea Shell
		84182,  //Collector's Preserved Drake Wing
		84183,  //Collector's Bulthar Tongue
		84184,  //Collector's Sea Pearl
		84185,  //Collector's Lock of Mermaid Hair
		84186,  //Collector's Ry'Gorr Mining Pick
		84187,  //Collector's Terror Tentacle
		84188,  //Collector's Kromrif Signet
		84189,  //Collector's Holgresh Elder Bead
		84262,  //Collector's Crude Stone Idol
		84263,  //Collector's Sonic Wolf Ear
		84264,  //Collector's Grimling Toe
		84265,  //Collector's Elemental Focus
		84266,  //Collector's Polished Stone
		84267,  //Collector's Sun Revenant Veil
		84268,  //Collector's Lightcrawler Shell
		84269,  //Collector's Rockhopper Hide
		84270,  //Collector's Netherbian Claw
		84271,  //Collector's Fungal Fiend Flesh
		84275,  //Collector's Tribal Belt
		84276,  //Collector's Fish Hook
		84277,  //Collector's Ball of Mud
		84278,  //Collector's Scarlet Fish
		84279,  //Collector's Cht'Thk Horn
		84280,  //Collector's Shik`Nar Wing
		84296,  //Collector's Hobgoblin Tusk
		84297,  //Collector's Spider Eye
		84298,  //Collector's Pendant of Marr
		84299,  //Collector's Diaku Blade
		84300,  //Collector's Cloak of Justice
		84301,  //Collector's Frog Tongue
		84302,  //Collector's Tormentor Hide
		84303,  //Collector's Innovative Gear
		84304,  //Collector's Note
		84305,  //Collector's Soul
		110916, //Collector's Lava Protection Mask
		110931, //Collector's Sample Cloak
		41987,  //Glowing Othni Wristband
		19762,  //Shimmering Black Pearl
		64167,
		64168,
		64169,
		64170,
		64171,
		64172,
		64173,
		64174,
		64175,
		64192,
		64193,
		64194,
		64195,
		64196,
		64197,
		64198,
		64199,
		64200,
		64201,
		64202,
		64203,
		64204,
		64205,
		64206,
		64207,
		64208,
		64209,
		64210,
		64211,
		64212,
		64213,
		64214,
		64215,
		64216,
		64217,
		64218,
		64219,
		64220,
		64221,
		64222,
		64223,
		64224,
		64225,
		64226,
		64227,
		64228,
		64231,
		64232,
		64233,
		64234,
		64235,
		64236,
		64237,
		64238,
		64239,
		64240,
		64241,
		64242,
		64243,
		64244,
		64245,
		64246,
		64247,
		64248,
		64249,
		64250,
		64251,
		64252,
		64253,
		64254,
		64255,
		64256,
		64257,
		64258,
		64259,
		64260,
		64261,
		64262,
		64263,
		64264,
		64265,
		64266,
		64267,
		64268,
		64269,
		64270,
		64271,
		64272,
		64273,
		64274,
		64275,
		64276,
		64277,
		64278,
		64279,
		64280,
		64281,
		64282,
		64283,
		64284,
		64285,
		64286,
		64287,
		64288,
		64289,
		64290,
		64291,
		64292,
		64293,
		64294,
		64295,
		64296,
		64297,
		64298,
		64300,
		64301,
		64302,
		64303,
		64304,
		64305,
		64306,
		64307,
		64308,
		64309,
		64310,
		64311,
		64312,
		64313,
		64314,
		64315,
		64316,
		64317,
		64318,
		64319,
		64320,
		64321,
		64322,
		64323,
		64324,
		64325,
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
		41994,  //glowing reis line
		41995,  //glowing reis
		41996,  //glowing reis
		41997,  //glowing reis
		41998,  //glowing reis
		41999,  //glowing reis
		42000,  //glowing reis
		41987,  //glowing othni
		41988,  //glowing othni
		41989,  //glowing othni
		41990,  //glowing othni
		41991,  //glowing othni
		41992,  //glowing othni
		41993,  //glowing othni
		41980,  //glowing athlai
		41981,  //glowing athlai
		41982,  //glowing athlai
		41983,  //glowing athlai
		41984,  //glowing athlai
		41985,  //glowing athlai
		41986,  //glowing athlai
		44921,  //brittle orleander
		44761,  //gormar venom
		44762,  //gormar venom
		44763,  //gormar venom
		44764,  //gormar venom
		44765,  //gormar venom
		44766,  //gormar venom
		44767,  //gormar venom
		44768,  //gormar venom
		44769,  //gormar venom
		44770,  //gormar venom
		44771,  //gormar venom
		44772,  //gormar venom
		44773,  //gormar venom
		44774,  //gormar venom
		44775,  //gormar venom
		124638, //gormar venom
		124653, //gormar venom
		81114,  //shimmering shard
	}
	return
}
