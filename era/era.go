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
		45881,  //	Mangled Rose of Peace
		45882,  //	Ceramic Weapon Ornament Sketchbook
		45883,  //	Black Acrylia Weapon Ornament Sketchbook
		45884,  //	A Worn Note
		45885,  //	Left part of a Worn Note
		45886,  //	Right part of a Worn Note
		45887,  //	A Box of Reagents
		45888,  //	Chemical Analysis Report
		45889,  //	Magical Analysis Report
		45890,  //	Silk Napkin
		45891,  //	Scale Ore
		45892,  //	Prismatic Palladium Bar
		45893,  //	Palladium Malachite Bracelet
		45894,  //	Palladium Lapis Lazuli Earring
		45895,  //	Palladium Turquoise Ring
		45896,  //	Palladium Hematite Necklace
		45897,  //	Wolf's Eye Palladium Bracelet
		45898,  //	Cat Eye Palladium Bracelet
		45899,  //	Palladium Bloodstone Necklace
		45900,  //	Palladium Onyx Necklace
		45901,  //	Palladium Jasper Earring
		45902,  //	Palladium Carnelian Wedding Ring
		45903,  //	Rose Palladium Necklace
		45904,  //	Palladium Amber Earring
		45905,  //	Jaded Palladium Bracelet
		45906,  //	Palladium Pearl Necklace
		45907,  //	Palladium Topaz Earring
		45908,  //	Palladium Peridot Bracelet
		45909,  //	Palladium Emerald Bracelet
		45910,  //	Palladium Opal Necklace
		45911,  //	Black Pearl Palladium Necklace
		45912,  //	Palladium Ruby Engagement Ring
		45913,  //	Sapphire Palladium Earring
		45914,  //	Palladium Star Ruby Anniversary Ring
		45915,  //	Fire Emerald Palladium Bracelet
		45916,  //	Palladium Fire Opal Ring
		45917,  //	Black Sapphire Palladium Earring
		45918,  //	Palladium Jacinth Wedding Ring
		45919,  //	Palladium Diamond Veil
		45920,  //	Palladium Blue Diamond Tiara
		45923,  //	Book Binding
		45924,  //	Ceramic Mace Ornament Mold Sketch
		45925,  //	Ceramic Great Mace Ornament Mold Sketch
		45926,  //	Ceramic Blade Ornament Mold Sketch
		45927,  //	Ceramic Great Blade Ornament Mold Sketch
		45928,  //	Ceramic Dagger Ornament Mold Sketch
		45929,  //	Ceramic Spear Ornament Mold Sketch
		45930,  //	Ceramic Fist Ornament Mold Sketch
		45931,  //	Black Acrylia War Hammer Ornament Mold Sketch
		45932,  //	Black Acrylia Great Mace Ornament Mold Sketch
		45933,  //	Black Acrylia Blade Ornament Mold Sketch
		45934,  //	Black Acrylia Halberd Ornament Mold Sketch
		45935,  //	Black Acrylia Dagger Ornament Mold Sketch
		45936,  //	Black Acrylia Spear Ornament Mold Sketch
		45937,  //	Full Lucidic Siphon
		45938,  //	Cosgrove Infused Clay
		45939,  //	Unfired Ceramic Mace Ornament Mold
		45940,  //	Unfired Ceramic Great Mace Ornament Mold
		45941,  //	Unfired Ceramic Blade Ornament Mold
		45942,  //	Unfired Ceramic Great Blade Ornament Mold
		45943,  //	Unfired Ceramic Dagger Ornament Mold
		45944,  //	Unfired Ceramic Spear Ornament Mold
		45945,  //	Unfired Ceramic Fist Ornament Mold
		45946,  //	Unfired Black Acrylia War Hammer Ornament Mold
		45947,  //	Unfired Black Acrylia Great Mace Ornament Mold
		45948,  //	Unfired Black Acrylia Blade Ornament Mold
		45949,  //	Unfired Black Acrylia Halberd Ornament Mold
		45950,  //	Unfired Black Acrylia Dagger Ornament Mold
		45951,  //	Unfired Black Acrylia Spear Ornament Mold
		45952,  //	Cosgrite Bar
		45953,  //	Enchanted Cosgrite Bar
		45954,  //	Cosgrite Solo Ring
		45955,  //	Cosgrite Duo Ring
		45956,  //	Cosgrite Trio Ring
		45957,  //	Cosgrite Solo Bracelet
		45958,  //	Cosgrite Duo Bracelet
		45959,  //	Cosgrite Trio Bracelet
		45960,  //	Cosgrite Solo Earring
		45961,  //	Cosgrite Duo Earring
		45962,  //	Cosgrite Trio Earring
		45963,  //	Cosgrite Solo Pendant
		45964,  //	Cosgrite Duo Pendant
		45965,  //	Cosgrite Trio Pendant
		45966,  //	Cosgrite Solo Veil
		45967,  //	Cosgrite Duo Veil
		45968,  //	Cosgrite Trio Veil
		45969,  //	Compartmented Cosgrite Trio Ring
		45970,  //	Compartmented Cosgrite Trio Bracelet
		45971,  //	Compartmented Cosgrite Trio Earring
		45972,  //	Compartmented Cosgrite Trio Pendant
		45973,  //	Compartmented Cosgrite Trio Veil
		45974,  //	Ability: Enchant Cosgrite
		45975,  //	Ability: Mass Enchant Cosgrite
		45976,  //	Tonic of Resonant Fire XII
		45977,  //	Tonic of Resonant Fire XIII
		45978,  //	Tonic of Resonant Fire XIV
		45979,  //	Tonic of Resonant Frost XII
		45980,  //	Tonic of Resonant Frost XIII
		45981,  //	Tonic of Resonant Frost XIV
		45982,  //	Tonic of Resonant Magic XII
		45983,  //	Tonic of Resonant Magic XIII
		45984,  //	Tonic of Resonant Magic XIV
		45985,  //	Tonic of Resonant Toxin XII
		45986,  //	Tonic of Resonant Toxin XIII
		45987,  //	Tonic of Resonant Toxin XIV
		45988,  //	Tonic of Resonant Disease XII
		45989,  //	Tonic of Resonant Disease XIII
		45990,  //	Tonic of Resonant Disease XIV
		45991,  //	Tonic of Resonant Elemental XII
		45992,  //	Tonic of Resonant Elemental XIII
		45993,  //	Tonic of Resonant Elemental XIV
		45994,  //	Tonic of Resonant Chromal XII
		45995,  //	Tonic of Resonant Chromal XIII
		45996,  //	Tonic of Resonant Chromal XIV
		46001,  //	Fabled Cloak of Crystalline Waters
		46002,  //	Fabled Crest of the Drixie
		46003,  //	Fabled Chrysolite
		46004,  //	Fabled Crystalline Robes
		46005,  //	Fabled Frozen Long Sword
		46006,  //	Fabled Foreman's Skull Cap
		46007,  //	Fabled Chipped Velium Amulet
		46008,  //	Fabled Blackened Crystalline Robe
		46009,  //	Fabled Scroll of Insight
		46010,  //	Fabled Scroll of Enlightenment
		46011,  //	Fabled Scroll of Knowledge
		46012,  //	Fabled Queen's Carapace
		46013,  //	Fabled Chetari Wardstaff
		46014,  //	Fabled Willsapper
		46015,  //	Fabled Infestation
		46016,  //	Fabled Cowl of Mortality
		46017,  //	Fabled Gauntlets of the Black
		46018,  //	Fabled Cloak of the Ry'Gorr Oracles
		46019,  //	Fabled Ry'Gorr Battle Mail
		46020,  //	Fabled Captain Ulmogs Brooch
		46021,  //	Fabled Snow Bunny Hat
		46022,  //	Fabled Blizzent's Fang
		46024,  //	Fabled Runebranded Stone Buckler
		46025,  //	Fabled Icetooth's Claws
		46026,  //	Fabled Bracer of Scale
		46027,  //	Fabled Wurm Meat
		46028,  //	Fabled Crystalline Eye
		46029,  //	Fabled Lodizal Shell Shield
		46030,  //	Fabled Eyepatch of Plunder
		46031,  //	Fabled Gown of the Chamberlain
		46032,  //	Fabled Frostreaver's Embroidered Cloak
		46033,  //	Fabled Incandescent Bracer
		46034,  //	Fabled Grizznot's Claws
		46035,  //	Fabled Eyepatch of the Shadows
		46036,  //	Fabled Black Blade of Tormenting
		46037,  //	Fabled Orb of the Infinite Void
		46038,  //	Fabled Horn of Hsagra
		46039,  //	Fabled Ring of Destruction
		46040,  //	Fabled Swiftblade of Zek
		46041,  //	Fabled Chestplate of Vindication
		46042,  //	Fabled Boots of the Ancients
		46043,  //	Fabled Lungi of the Forbidden
		46044,  //	Fabled Clawed Coat of Tyro
		46045,  //	Fabled Tenuous Dragonscale Slippers
		46046,  //	Fabled Ancient Prismatic Warsword
		46047,  //	Fabled Ancient Prismatic Brawl Stick
		46048,  //	Fabled Priceless Velium Claidhmore
		46049,  //	Fabled Tome of Innovation
		46050,  //	Fabled Sorcerer's Robe
		46051,  //	Fabled Assassin's Chestguard
		46052,  //	Fabled Flute of the Sacred Glade
		46053,  //	Fabled Chestplate of Stability
		46054,  //	Fabled Resplendent Robe
		46055,  //	Fabled Girdle of Living Thorns
		46058,  //	Fabled Fellspine's Tail
		46060,  //	Fabled Coral Encrusted Vambraces
		46061,  //	Fabled Cloak of the Seacaller
		46064,  //	Fabled Mask of the Dragon Slayer
		46065,  //	Fabled Aegis of the Shrine
		46066,  //	Fabled Amulet of the Shrine
		46067,  //	Fabled Rod of the Healers
		46068,  //	Fabled Signet of the Shrine
		46069,  //	Fabled Silver Girdle of Stability
		46070,  //	Fabled Silver Bracelet of Speed
		46071,  //	Fabled Black Rock Maul of Crushing
		46072,  //	Fabled Ikatiar's Stinger
		46073,  //	Fabled Claw of Lightning
		46074,  //	Fabled Shovel of the Harvest
		46075,  //	Fabled Katana of Pain
		46076,  //	Fabled Boots of the Dead Dream
		46077,  //	Fabled Cloak of Thorns
		46078,  //	Fabled Net of the Deep Sea
		46079,  //	Fabled Great Spear of Doom
		46080,  //	Fabled Belt of the Destroyer
		46081,  //	Fabled Silver Disk
		46082,  //	Fabled Blackstar, Mace of Night
		46083,  //	Fabled Baton of Flame
		46084,  //	Fabled Shawl of Perception
		46085,  //	Fabled Vyemm's Left Eye
		46086,  //	Fabled Sal'Varae's Robe of Darkness
		46087,  //	Fabled Earring of the Icecaster
		46088,  //	Fabled Embalmers Skinning Knife
		46089,  //	Fabled Abram's Axe of the Stoic
		46090,  //	Fabled Fingerbone Necklace
		46091,  //	Fabled Shield of Shadows
		46092,  //	Fabled Beer Stained Coldain Tunic
		46093,  //	Fabled Crystallized Shadow Belt
		46094,  //	Fabled Tserrina's Robe
		46095,  //	Fabled Crystallized Shadow Tunic
		46096,  //	Fabled Crystal Spider Eyes
		46097,  //	Fabled Crystal Chitin Shield
		46098,  //	Fabled Icicle Tunic
		46099,  //	Fabled Vermilion Robe of Torrefaction
		46100,  //	Fabled Barbed Ringmail Shoulderpads
		46101,  //	Fabled Scimitar of the Emerald Dawn
		46102,  //	Fabled Helm of the Tracker
		46103,  //	Fabled Holgresh Grand Vizier Beads
		46104,  //	Fabled Gelistials Horn
		46105,  //	Fabled Gloves of Earthcrafting
		46106,  //	Fabled Tigeraptor Hide
		46107,  //	Fabled Del Sapara's Talisman
		46108,  //	Fabled Ionat's Talisman
		46109,  //	Fabled Karkona's Talisman
		46110,  //	Fabled Entariz's Talisman
		46111,  //	Fabled Esorpa's Talisman
		46112,  //	Fabled Derasinel's Talisman
		46113,  //	Fabled Ayillish's Talisman
		46114,  //	Fabled Exquisite Velium Warsword
		46115,  //	Fabled Myga's Talisman
		46116,  //	Fabled Harla Dar's Talisman
		46117,  //	Fabled Tantor's Tusk
		46118,  //	Fabled Stronghorn's Horn
		46119,  //	Fabled Klandicar's Talisman
		46120,  //	Fabled Sontalak's Talisman
		46121,  //	Bounty Tea
		46122,  //	Bounty Coffee
		46123,  //	Bounty Juice
		46124,  //	Bounty Ale
		46125,  //	Bounty Spirits
		46126,  //	Cosgrove Ale
		46127,  //	Cosgrove Spirits
		46128,  //	Cosgrove Nostrum
		46129,  //	Brell's Best
		46130,  //	Brewing Brell's Best
		46131,  //	Coldain Pikestaff of the Steadfast
		46132,  //	Kromzekian Walking Stick of Allegiance
		46133,  //	Draconic Cane of the Everlasting
		46134,  //	Mug of the Endless Fan Faire
		46135,  //	Summoned: Party Mug
		46136,  //	Miraculous Cure-All Potion
		46137,  //	Rune of the Ages
		46138,  //	Fairy Wing
		46139,  //	Lens of the Evil Eye
		46140,  //	Lokin's Greaves
		46141,  //	Skullcap of Contemplation
		46142,  //	Ring of the Tun
		46143,  //	Hammer of Falling Stars
		46144,  //	Froglok Totem
		46145,  //	Slime-Covered Earring
		46146,  //	Gloves of the Overthere
		46147,  //	Cockatrice Leather Boots
		46148,  //	Shattered Stone Ring
		46149,  //	Chipped Sarnak Great Axe
		46150,  //	Chipped Sarnak Shield
		46151,  //	Shattered Stone Earring
		46152,  //	Spiroc Feather Spaulders
		46153,  //	Spiroc Feather Choker
		46154,  //	Pirate's Booty
		46155,  //	Stained Driftwood Bow
		46156,  //	Stained Driftwood Stave
		46157,  //	Ornate Pirate Sash
		46158,  //	Stone Runed Spaulders
		46159,  //	Creeping Vine Necklace
		46160,  //	Mud-Splattered Mask
		46161,  //	Dragon Bone Shard
		46162,  //	Chipped Onyx Sphere
		46163,  //	Cloak of Whispers
		46164,  //	Yeti Fur-Lined Sleeves
		46165,  //	Drolvarg Skullcap
		46166,  //	Mantle of the Brae
		46167,  //	Charred Wooden Rod
		46168,  //	Lost Staff of the Combine
		46169,  //	Crag Giant Choker
		46170,  //	Calcified Bone Gauntlets
		46171,  //	Red Silk Lined Boots
		46172,  //	Golem Essence Spaulders
		46173,  //	Sword of the Mist
		46174,  //	Shield of the Mist
		46175,  //	Golem Essence Necklace
		46176,  //	Wurmfire Ring
		46177,  //	Wurmfire Earring
		46178,  //	Wurmfire Totem
		46179,  //	Suspended Wyvern Blood
		46180,  //	Blackened Lava Rock
		46181,  //	Suspended Wurm Blood
		46182,  //	Golden Drolvarg Boots
		46183,  //	White Silk Gloves
		46184,  //	Ruined Leather Mask
		46185,  //	Axe of Destruction
		46186,  //	Reinforced Round Buckler
		46187,  //	Midnight Drolvarg Cloak
		46188,  //	Bloodgill Cap
		46189,  //	Waterlogged Armwraps
		46190,  //	Bloodgill Belt
		46191,  //	Hammer of the Forgotten City
		46192,  //	Bloodgill Shaman's Staff
		46193,  //	Bloodgill Shaman's Effigy
		46195,  //	Chunk of Basalt
		46196,  //	Living Lava
		46197,  //	Molten Skin
		46198,  //	Sessiloid Shard
		46199,  //	Lichenash
		46200,  //	A Cryptic Scroll
		46201,  //	Crystallized Spore Choker
		46202,  //	Drake Tail Belt
		46203,  //	Piranha Jawbone Ring
		46204,  //	Rat Fang Earring
		46205,  //	Crystalwing Rattan Shield
		46206,  //	Chewed Mucktail Basher
		46207,  //	Petrified Spinewood Staff
		46208,  //	Ceremonial Nokk Sword
		46209,  //	Tarnished Nokk Helm
		46210,  //	Tarnished Nokk Chain Sleeves
		46211,  //	Stained Leather Gloves
		46212,  //	Tattered Silk Sandals
		46213,  //	Glowing Carbuncle
		46214,  //	Ceremonial Nokk Morningstar
		46215,  //	Preserved Nokk Bow
		46216,  //	Battered Ceremonial Drum
		46217,  //	Tarnished Nokk Breastplate
		46218,  //	Mucktail Chain Leggings
		46219,  //	Mucktail Leather Cuirass
		46220,  //	Tattered Silk Pantaloons
		46221,  //	Nokk Guardsman Sigil
		46222,  //	Chewed Nokk Leather Belt
		46223,  //	Mucktail Cloak
		46224,  //	Mucktail Mantle
		46225,  //	Tarnished Nokk Splintmail Mantle
		46226,  //	Tattered Nokk Cloak
		46227,  //	Shaded Silk Mask
		46228,  //	Ceremonial Nokk Mace
		46229,  //	Ceremonial Nokk Athame
		46230,  //	Shaded Silk Doll
		46231,  //	Wanderlust Scout Pack
		46232,  //	Tarnished Nokk Bracelet
		46233,  //	Broken Nokk Manacles
		46234,  //	Ichor Encrusted Boots
		46235,  //	Ichor Encrusted Chain Gloves
		46236,  //	Ichor Encrusted Hide Sleeves
		46239,  //	Treantwood Mallet
		46240,  //	Witchwood Staff
		46241,  //	Ethershard Glass
		46244,  //	Mask of Snake Leather
		46245,  //	Snake Spine Collar
		46246,  //	Witch's Skull Mask
		46249,  //	Bonehilt Dagger
		46250,  //	Spear of Nokk
		46251,  //	Fist of Two Blades
		46255,  //	Darkwater Mask
		46257,  //	Witch's Ritual Mask
		46260,  //	Bow of Blightfire
		46262,  //	Ruby of Screaming Facets
		46264,  //	Gauntlets of the Watchtower
		46269,  //	Mossy Stone
		46274,  //	Ring of Six Moons
		46279,  //	Shield of the Dark Tower
		46284,  //	Symbiotic Blightfire Moss
		46285,  //	Belt of Blightfire
		46286,  //	Blightfire Bone Shard
		46287,  //	Mantle of Emerald Poison
		46288,  //	Shoulderguard of the Swarmer
		46289,  //	Girdle of the Foreboding Heart
		46290,  //	Girdle of Defoliation
		46291,  //	Royal Ruby Orb of the Hive
		46292,  //	Barbed Stone Hive Defender's Necklace
		46293,  //	Jade Scimitar of Turina
		46294,  //	Forged Maul of Bixie Bashing
		46295,  //	Reinforced Stone Hive Compound Bow
		46296,  //	Spiked Buckler of Defoliation
		46297,  //	Laced Bixie Veil
		46298,  //	Stormy Shroud of Deceit
		46299,  //	Hive Defender's Plate Armguards
		46300,  //	Hive Poisonmaker's Coif
		46301,  //	Hive Worker's Leather Gloves
		46302,  //	Hive Overseer's Silk Boots
		46303,  //	Shroud of Greater Transmutation
		46304,  //	Hive Defender's Woven Belt
		46305,  //	Orb of Forgiveness
		46306,  //	Hive Defender's Plate Wristguard
		46307,  //	Hive Poisonmaker's Chain Wristguard
		46308,  //	Hive Worker's Leather Wristguard
		46309,  //	Hive Overseer's Silk Wristguard
		46310,  //	Hive Worker's Orb of Health
		46311,  //	Orb of Flowing Pollen
		46312,  //	Hive Defender's Orb of Stinging Fury
		46313,  //	Hive Defender's Legplates
		46314,  //	Hive Poisonmaker's Cuirass
		46320,  //	Crippling Mace of Divinity
		46321,  //	Scourge of Dismemberment
		46327,  //	Bauble of Infused Subterfuge
		46328,  //	Bauble of Greater Subterfuge
		46329,  //	Singed Orb of the Hive
		46330,  //	Petrified Windwillow Burl
		46331,  //	Minohten Chain Arms
		46332,  //	Stained Snakeskin Boots
		46334,  //	Snakeskin Choker
		46336,  //	Stained Minohten Ceremonial Panpipes
		46337,  //	Windwillow Branch Bow
		46338,  //	Dromrek Billy Club
		46339,  //	Nemarsarpe's Fang
		46340,  //	Minohten Guard Sabatons
		46341,  //	Minohten Chain Gloves
		46342,  //	Craita's Dress
		46343,  //	Aurelia's Dress
		46344,  //	Wisp Essence Cloak
		46345,  //	Craita's Saber
		46346,  //	Minohten Hero Tonfa
		46347,  //	Minohten Hero Maul
		46348,  //	Aurelia's Saber
		46349,  //	Wisp Essence Ring
		46350,  //	Centaur Plate Wrist Guard
		46351,  //	Centaur Chain Bracelet
		46352,  //	Tuffein Leather Guard
		46353,  //	Florenta's Dress
		46354,  //	Nestor's Cloak
		46355,  //	Tuffein Pacifier
		46356,  //	Windwillow Lute
		46357,  //	Fieldstrider Ceremonial Spear
		46358,  //	Fieldstrider Brand
		46359,  //	Fieldstrider Legate Shield
		46360,  //	Ooze Encrusted Visor
		46361,  //	Harpy Feather Mantle
		46362,  //	Dromrek Toothpick
		46363,  //	Dromrek Bracelet
		46364,  //	Ionela's Belt
		46365,  //	Dromrek Censor
		46366,  //	Whitby's Back Scratcher
		46367,  //	Ooze Coated Dromrek Short Sword
		46368,  //	Driftwood Wand
		46369,  //	Pristine Harpy Claw
		46370,  //	Goru`Kar Steel Spaulders
		46371,  //	Incinspaianjen Carapace Mask
		46372,  //	Tangleweed Tangle Belt
		46373,  //	Lingering Dryad Arm
		46374,  //	Uriasarpe Fang
		46375,  //	Spider Fang Sword
		46376,  //	Rotwood Bow
		46377,  //	Lingering Dryad Rib Bone
		46378,  //	Ternsmochin Heartwood
		46379,  //	Snake Rattle Ring
		46380,  //	Goru`kar Nose Ring
		46381,  //	Griffon Claw Ring
		46382,  //	Goru`kar Finger Bone
		46383,  //	Blackfeather Good luck Charm
		46384,  //	Griffon Down Cloak
		46385,  //	Harpy Trinket
		46386,  //	Griffon Bone Earring
		46387,  //	Griffon Tamer's Mask
		46388,  //	Griffon Tamer's Shawl
		46389,  //	Serrated Griffon Claws
		46390,  //	Satyr Legbone
		46391,  //	Blackfeather War Drum
		46392,  //	Royal Harpy Guard Shield
		46393,  //	Crystallized Griffon Eye
		46394,  //	Harpy Fury
		46395,  //	Harpy Slave Vambraces
		46396,  //	Griffon Tender Boots
		46398,  //	Blackfeather Noble's Gloves
		46399,  //	Blackfeather Noble's Ring
		46400,  //	Blackfeather Noble's Earring
		46401,  //	Studded Harpy's Belt
		46402,  //	Blackfeather Crested Wristguard
		46403,  //	Blackfeather Chain Wristguard
		46405,  //	Feather Bracelet
		46406,  //	Griffon Tooth Mace
		46407,  //	Razor Sharp Griffon Quill
		46408,  //	Enchanted Harpy Blade
		46409,  //	Cloak of Lost Essence
		46410,  //	Eletyl's Charm
		46411,  //	Earring of the Roost
		46412,  //	Ring of the Roost
		46413,  //	Griffon Tamer's Toolkit
		46414,  //	Bloom of the Elddar
		46415,  //	Griffon Fang
		46416,  //	Blackfeather Stone of Sufficiency
		46417,  //	Lost Essence of the Elddar
		46418,  //	Rage of the Harpies
		46419,  //	Gaze of the Griffon
		46420,  //	Blackfeather Protective Stone
		46421,  //	Blackfeather Rock
		46422,  //	Harpy Powerstone
		46423,  //	Bracer of the Roost
		46424,  //	Down Padded Wrist Wrap
		46425,  //	Feather Down Shirt
		46426,  //	Leadfoot Boots
		46427,  //	Evergleam Sleeves
		46428,  //	Gravelskin Gloves
		46429,  //	Wirewick Cap
		46430,  //	Dark Whisper
		46431,  //	Mask of False Tales
		46432,  //	Wirewick Mantle
		46433,  //	Firebelly's Unclean Cleaver
		46434,  //	Darkfell Bow
		46435,  //	Smash
		46436,  //	Warbound Gauntlets
		46437,  //	Evergleam Coif
		46438,  //	Gravelskin Boots
		46439,  //	Wirewick Sleeves
		46440,  //	Barking Stone
		46441,  //	Warbound Breastplate
		46442,  //	Evergleam Leggings
		46443,  //	Gravelskin Tunic
		46444,  //	Wirewick Leggings
		46445,  //	Starshine
		46446,  //	Coiled Serpentskin Belt
		46447,  //	Wirewick Cape
		46448,  //	Spotted Cape
		46449,  //	Wirewick Belt
		46450,  //	Spirit Mantle
		46451,  //	Bloodspear
		46452,  //	Mammoth Tusk Flute
		46453,  //	Granite Blade of War
		46454,  //	Oldbone Skull
		46455,  //	Veil of Summer
		46456,  //	Earthen Tear Earring
		46457,  //	Stone Hoop
		46458,  //	Worn Rathe Totem
		46459,  //	Mineral Encrusted Plate Sleeves
		46460,  //	Mineral Encrusted Chain Boots
		46461,  //	Mineral Encrusted Hood
		46464,  //	Blade of the Last Knight
		46465,  //	Sandswirl Shield
		46466,  //	Crystalline Staff of the Springs
		46469,  //	Sand Guard Plate Leggings
		46472,  //	Sand Guard Robes
		46473,  //	Sand Guard Mantle
		46474,  //	Goad of the High Desert
		46477,  //	Sunderock Mineral Claw
		46478,  //	Signet of the Mistdragon
		46479,  //	Sunderspring Helm
		46480,  //	Sunderspring Gloves
		46481,  //	Sunderspring Sleeves
		46482,  //	Sunderspring Boots
		46483,  //	Sunderspring Belt
		46484,  //	Lifeseed Emerald
		46485,  //	Petrified Stonefin Egg
		46486,  //	Gemstone of Scale
		46487,  //	Rage Moss
		46488,  //	Ruby of Destructive Runes
		46489,  //	Geothermal Pumice
		46490,  //	Hotsprings Pumice
		46491,  //	Steamvent Ruby
		46492,  //	Dyla's Monocle of Exploration
		46493,  //	Reinforced Green Legion Neckguard
		46494,  //	Flowing Gloompetal Cloak
		46495,  //	Girdle of Rotting Flesh
		46496,  //	Ring of Blood
		46497,  //	Sturdy Bauble of Deceit
		46498,  //	Dyla's Staff of Focus
		46499,  //	Runic Green Legion Longsword
		46500,  //	Pristine Green Legion Warhammer
		46501,  //	Bonecruncher's Bastion
		46502,  //	Orb of the Enslaved
		46503,  //	Rislocha's Mantle of the Protection
		46505,  //	Tadaie's Mask of Blood
		46506,  //	Girdle of Vengeful Fury
		46507,  //	Malevolent Cloak of Eternity
		46508,  //	Watchman Orak's Brilliant Bauble
		46509,  //	Khalosk's Ring of Fury
		46510,  //	Rumos' Regal Rhinestone
		46512,  //	Tadaie's Vengeful Restitution
		46513,  //	Jezrak's Ornate Shield
		46514,  //	Green Legion Compound Bow
		46515,  //	Watchman Orak's Horn of Blood
		46516,  //	Resonant Wand of the Diviner
		46517,  //	Rumos' Resiliant Spaulders
		46518,  //	Green Legion Wristguard
		46519,  //	Wristguard of the Ancient
		46520,  //	Shadowfang's Armguard
		46521,  //	Stone Guardian Wristband of Smashing
		46522,  //	Iico's Bauble of Druidic Contortion
		46523,  //	Velosk's Ring of Bone
		46524,  //	Xacine's Orb of Wizardry
		46525,  //	Truncheon of Blood
		46526,  //	Resiliant Dagger of the Ancient
		46527,  //	Shadowfang's Fist
		46528,  //	Stone Guardian Cudgel of Crushing
		46529,  //	Iico's Orb of the Ancients
		46530,  //	Velosk's Orb of Bone
		46531,  //	Xacine's Wondrous Trinket
		46532,  //	Ancient Orb of the Green Legion
		46533,  //	Decaying Orb of the Green Legion
		46534,  //	Bloody Orb of the Green Legion
		46535,  //	Direwrought Plate Gauntlets
		46536,  //	Direwrought Chain Boots
		46537,  //	Direwrought Leather Sleeves
		46538,  //	Direwrought Silk Hood
		46539,  //	Direwrought Mask
		46540,  //	Club of Scathing Branches
		46541,  //	Ashen-Hewn Bow
		46542,  //	Twin-Edged Razorfist
		46543,  //	Emerald of Forgotten Dreams
		46544,  //	Choker of Battles
		46546,  //	Gloves of the Black Cliff
		46547,  //	Mask of the Black Cliff
		46549,  //	Belt of the Black Cliff
		46551,  //	Windstill Chain Cloak
		46552,  //	Direwind's Beckon
		46554,  //	Ring of Plagueward
		46555,  //	Blackiron Belt
		46558,  //	Skull Mask of the Warlord
		46559,  //	Book of Dire Incantations
		46560,  //	Hoop of Regrowth
		46563,  //	Belt of the Elementalist
		46564,  //	Earring of Poison Clouds
		46565,  //	Direwind Corpse Stone
		46566,  //	Direwind Plague Stone
		46567,  //	Deathcalm Stone
		46568,  //	Warbound Boots
		46569,  //	Evergleam Gauntlets
		46570,  //	Gravelskin Cap
		46571,  //	Ironice Sleeves
		46572,  //	Stonefrost Belt
		46573,  //	Bloodfoot Chopper
		46574,  //	Blackfoot Fang
		46575,  //	Sabertooth Horn
		46576,  //	Snowblind Globe
		46577,  //	Frostpaw Cloak
		46578,  //	Hearthfire Coif
		46579,  //	Wirewick Slippers
		46580,  //	Cape of the Cave Bear
		46581,  //	Belt of the Falls
		46582,  //	Iceborn Diamond Necklace
		46583,  //	Cape of Echoes
		46584,  //	Chillbone Finger
		46585,  //	Ironice Earring
		46586,  //	Hunter's Rage
		46587,  //	Farwalker's Hoop
		46588,  //	Heart of Ice
		46589,  //	Iceborn Band
		46590,  //	Earring of the Cold Heart
		46591,  //	Ring of Severed Power
		46592,  //	Totem of the Lost
		46593,  //	Lowfire Belt
		46594,  //	Mammoth Hide Spaulders
		46595,  //	Necklace of Stars
		46596,  //	Direwolf Hide Spaulders
		46597,  //	Ring of the Glacier
		46598,  //	Iceborn Diamond
		46599,  //	Iceborn Granite
		46600,  //	Stone Kanabo
		46601,  //	Holmguard Helm
		46602,  //	Warforged Sleeves
		46603,  //	Boots of Bounding
		46604,  //	Will of Many Hands
		46605,  //	Ring of Warded Stone
		46606,  //	Ancestral Hoop
		46607,  //	Valdeholm Honorguard Bracer
		46608,  //	Commander's Arm Guard
		46609,  //	Commander's Crested Bracer
		46610,  //	Bracelet of Ancient Stone
		46611,  //	Steamforged Shield of Runes
		46612,  //	Hailstone Heart
		46613,  //	Holmguard Cloak
		46614,  //	Lorekeeper Belt
		46615,  //	Pitted Mask of Sorrows
		46616,  //	Necklace of Ages
		46617,  //	Mantle of Crypts
		46618,  //	Ring Rituals
		46619,  //	Earring of the Arts
		46620,  //	Baelon's Ice Knife
		46621,  //	Belgar's Blade
		46622,  //	Deepbelly
		46623,  //	Goil's Guardian Aura
		46624,  //	The Guardian's Pouch
		46625,  //	Hero's Promise
		46626,  //	Dentrik's Dementia
		46627,  //	Krithgor Ceremonial Mask
		46628,  //	Lorekeeper Charm Necklace
		46629,  //	Cloak of Weighted Chains
		46630,  //	Serpentine Belt of Deftness
		46631,  //	Essence of Stone
		46632,  //	Orb of Chilling Grace
		46633,  //	Jewel of the Stalwart
		46634,  //	House Harfange Heraldic Gorget
		46635,  //	Deathless Mask
		46636,  //	Warbound Belt
		46637,  //	Jester's Mask
		46638,  //	Pulsing Emerald of Life
		46639,  //	Warpower Stone
		46640,  //	Fellstone
		46641,  //	Icehaven Pauldron
		46642,  //	Frostmantle
		46643,  //	Starlight Crystal
		46644,  //	Silver Ring of Numbing Cold
		46645,  //	Earring of Shattered Souls
		46646,  //	Enameled Froststeel Stud
		46647,  //	Sapphire Ring of Bromal Fury
		46648,  //	Rime-Coated Steel Cincture
		46649,  //	Rallosian-Adorned Iron Choker
		46650,  //	Ice Glazed Platinum Mask
		46651,  //	Frost Runed Icy War Hammer
		46652,  //	Jagged Froststeel Lance
		46653,  //	Shimmering Ice-Wrapped Fists
		46654,  //	Rime-Coated Drum of the Avalanche
		46655,  //	Orb of Swirling Shadows
		46656,  //	Spiked Froststeel Pauldron
		46657,  //	Glowing Pendant of Preserved Ice
		46658,  //	Pristine Granite Stone
		46659,  //	Etched Pearl of Ever-Ice
		46660,  //	Forged Froststeel Ingot
		46661,  //	Frosted Diamond of Elemental Fury
		46662,  //	Black Sapphire of Eldritch Focus
		46663,  //	Cracked Ruby of Bloodletting
		46664,  //	Verdant Emerald of Wounding
		46665,  //	Polished Opal of Vigor
		46666,  //	Bloodless Rimed Iron Ring
		46667,  //	Braided Iceweave Cord
		46668,  //	Frostweave Sash of the Muse
		46669,  //	Hail-Studded Collar of Savagery
		46670,  //	Iron Club of False Conviction
		46671,  //	Weighted Emerald Die
		46672,  //	Brittle Frostvine Veil
		46673,  //	Exquisite Cloak of Conjured Flame
		46674,  //	Fire-Runed Icepearl
		46675,  //	Frozen Diamond of Purification
		46676,  //	Frostflake of the All-Enduring
		46677,  //	Black Legion Marker
		46678,  //	Black Legion Templar Band
		46679,  //	Polished Krithgor Bone Ring
		46680,  //	Pyrewood Earring
		46681,  //	Amulet of the Scarlet Legion
		46682,  //	Girdle of the Magma Priest
		46683,  //	Splintering Spinestone Wand
		46684,  //	Weighty Jagged Einhander
		46685,  //	Thrice-Banded Warclub
		46686,  //	Tainted Darksteel Battleaxe
		46687,  //	Spellstring, the Bard's Muse
		46688,  //	Magmaraug's Burning Essence
		46689,  //	Faceless Guise of the Reliquary
		46690,  //	Scaleguard Spaulders
		46691,  //	Soldier's Stone of Purity
		46692,  //	Clergyman's Stone of Purity
		46693,  //	Crest of the Scarlet Legions
		46694,  //	Stone of Dragon Scales
		46695,  //	Barbed Dragon Bones
		46696,  //	Lava-Drenched Gnoll Heart
		46697,  //	Purified Lava Vein
		46698,  //	Eggshell Fragment
		46699,  //	Crystalwing Stud of the Fallen
		46700,  //	Mask of the Black Legion Scouts
		46701,  //	Magmamail Collar
		46702,  //	Band of the Scholars
		46703,  //	Spiritwhisper Cloak
		46704,  //	Mantle of Hanging Flesh
		46705,  //	Ro-Touched Ear Stud
		46706,  //	Umbracite Spiked Shoulderguards
		46707,  //	Uncut Ruby of Purity
		46708,  //	Hardened Coalstone of Purity
		46709,  //	Lava-Seared Drakkin Heart
		46710,  //	Scrying Device
		46711,  //	Direbarbed Bronze Bracer
		46712,  //	Beer-Soaked Bracer
		46713,  //	Boulderfist Family Signet Ring
		46714,  //	Band of Roiling Sands
		46715,  //	Fiddleback's Spinnerette
		46716,  //	Mark of the Dunes
		46717,  //	Steeloak Mantle
		46718,  //	Cloak of Wintersting
		46719,  //	Disk of the Black Oak
		46720,  //	Pauldrons of the Snowdrifter
		46721,  //	All-Seers Belt of Sinew
		46722,  //	Frozen Mask
		46723,  //	Kellak's Electrified Ring of Voracity
		46724,  //	Clouded Ring of Restoration
		46725,  //	Electrified Bauble of Rigor
		46726,  //	Idol of Kangur Vafta Veor
		46727,  //	Veil of the Snow Spider
		46728,  //	Frozen Heart of the Snow Spider Queen
		46729,  //	A Marvelous Elixir
		46730,  //	Gift: Child's Play Anklebiter Doll
		46731,  //	Veil of Dreams
		46732,  //	Dreamweaver's Sash
		46733,  //	Ring of Deep Dreams
		46734,  //	Warskin Shield
		46735,  //	Stormsword
		46736,  //	Mistwalker's Collar
		46737,  //	Shroud of Mist
		46738,  //	Mistwalker's Scale
		46739,  //	Shard of Frozen Mist
		46740,  //	Staff of the Mist Dragon
		46741,  //	Narseekin's Hood
		46742,  //	Pyrekeeper's Mask
		46743,  //	Ashen Orb of the Blightpyre
		46744,  //	Ash of the Fallen
		46745,  //	Staff of the Pyrekeeper
		46751,  //	Goru's Girdle of the Deceased
		46752,  //	Neckguard of the Reincarnate
		46753,  //	Diseased Ring of the Reincarnate
		46754,  //	Goru's Trinket of Fidelity
		46755,  //	Warfists of the Reincarnate
		46756,  //	Holmguard Pauldrons
		46757,  //	Lorekeeper's Veil
		46758,  //	Earring of the Ageless
		46759,  //	Loreward
		46760,  //	Veinsplitter
		46761,  //	Frost-Etched Ruby of Vitality
		46762,  //	Fragment of Corrupted Ice
		46763,  //	Shaped Turquoise Fragment
		46764,  //	Rune-Etched Granite Fragment
		46765,  //	Iridescent Pearl of Icy Precision
		46766,  //	Molded Agate Trinket
		46767,  //	Mossy Tundra Stone
		46768,  //	Etched Obsidian Shard
		46769,  //	Magma Lacquer
		46770,  //	Prism of Captured Energy
		46771,  //	Clipped Griffon Wings
		46772,  //	Arcane Wavebender Prism
		46773,  //	Arcane Actuator
		46774,  //	Krithgor Ribshard
		46775,  //	Flowing Magma Soulstone
		46776,  //	Sharpening Stone of the Sky
		46777,  //	Underfoot Trout
		46778,  //	Freshwater Flounder
		46779,  //	Bayle Mark
		46780,  //	Mark of Brell
		46781,  //	Stoneplate Headguard
		46782,  //	Barkskin Helm
		46783,  //	Vinelink Gloves
		46784,  //	Webweaver's Silk Gloves
		46785,  //	Barkskin Scarf
		46786,  //	Earthen Idol of Rathe
		46787,  //	Barkskin Sash
		46788,  //	Acidsheen Chain Hood
		46789,  //	Hood of Oblivion
		46790,  //	Gauntlets of the Ooze Lord
		46791,  //	Gloves of Ten Hides
		46792,  //	Corroded Artifact Ring
		46793,  //	Ooze Pearl
		46794,  //	Necklace of White Sands
		46795,  //	Windcaller's War Gear
		46796,  //	Windcaller's Under Armor
		46797,  //	Ringmail Ritual Boots
		46798,  //	Plaguehide Boots
		46799,  //	Ring of the Windcaller
		46800,  //	Corpse Bone Ring
		46801,  //	Earring of Pyre Ash
		46802,  //	Sleeves of the Savage
		46803,  //	Snowfall Sleeves
		46804,  //	Sabatons of the Guardian
		46805,  //	Swiftblade Boots
		46806,  //	Earring of Starshine
		46807,  //	Snowshine
		46808,  //	Band of Avalanche Ice
		46809,  //	Gnomework Model XVII's Experimental Chestguard
		46810,  //	Nurtha's Leggings of Focus
		46811,  //	Vermin Model IV's Diseased Bauble
		46812,  //	Vermin Model IV's Diseased Harness
		46813,  //	Nurtha's Shoulderpads of Concentration
		46814,  //	Gnomework Model XVII's Experimental Plate Backing
		46815,  //	Arachnid Model XII's Protective Hull
		46817,  //	Serric's Blood-Encrusted Anorak
		46818,  //	Vergalid's Leggings of Corruption
		46819,  //	Serric's Virulent Augmentation
		46820,  //	Ring of Breeding Corruption
		46821,  //	Spaulders of Dooming Exile
		46822,  //	Vergalid's Infectious Cloak
		46823,  //	Vergalid's Blood Mace of Dismemberment
		46825,  //	Loyalist Coat of Icemail
		46826,  //	Warbound Leggings
		46827,  //	High Lorekeeper Necklace of Office
		46828,  //	Udengar's Thornmesh of War
		46829,  //	Icecracker
		46830,  //	Warsorrow
		46831,  //	Icedancer
		46833,  //	Robe of the Driving Rains
		46834,  //	Sleetstorm Leggings
		46835,  //	Highkeeper's Haven
		46836,  //	Girdle of Stone Strength
		46837,  //	Lute of Longing
		46838,  //	Everlasting Forge Coal
		46839,  //	Gefia's Gift
		46841,  //	Bladed Silver Pendant of Zek
		46842,  //	Engraved Petrified Figurine
		46843,  //	Frostwoven Platinum Band
		46844,  //	Lambent Earring of Melded Iron
		46845,  //	Polished Pauldrons of Pearlescent Ice
		46846,  //	Torn Shawl of the Ice Priests
		46847,  //	Iceweave Cloak of Blackest Night
		46848,  //	Tattered Cryptrobber's Cloak
		46849,  //	Shaped Sapphire Ice Sculpture
		46850,  //	Pitchfork of the Town Rebels
		46851,  //	Ice-Striated Iron Orb
		46852,  //	Ornate Mask of Fallen Dreams
		46853,  //	Battered Silver Runed War Shield
		46854,  //	Luminous Froststeel Warclub
		46855,  //	Frost Runed Blade of the Wraithguard
		46856,  //	Luminous Froststeel Warhammer
		46857,  //	Silver-Wrapped Turquoise Earring
		46858,  //	Frosted Wrap of the Calm Spirit
		46859,  //	Iceweave Shawl of the Elements
		46860,  //	Runed Necklace of Living Bone
		46861,  //	Shimmering Veil of Frozen Mana
		46862,  //	Shattered Mask of Ancestral Rage
		46863,  //	Jagged Blood Iron Ring
		46864,  //	Scorched Froststeel Ring
		46865,  //	Glowing Earring of Purified Ice
		46866,  //	Blood-Infused Iron Stud
		46867,  //	Ice Enameled Gorget of Battle
		46868,  //	Carved Blood Ice Choker
		46869,  //	Iceweave Sash of Lost Faith
		46870,  //	Slashed Leather Belt of Frostbite
		46871,  //	Frost-Carved Granite Statue
		46872,  //	Worn Iron Rallosian Relic
		46873,  //	Engraved Froststeel Visor
		46874,  //	Sleetcaller's Mask of Icy Precision
		46875,  //	Riven Krithgor Battle Shield
		46876,  //	Frost Inlaid Lute of Boreal Wind
		46877,  //	Sekv, Blade of Glacial Fury
		46878,  //	Brumal Fists of Whirling Sleet
		46879,  //	Jagged Bone-Handled Ice Reaver
		46880,  //	Algid Bag of Frozen Tears
		46881,  //	Icy Ring of Boreal Essence
		46882,  //	Faceted Shard of Living Ice
		46883,  //	Crystalline Shoulder Pads of Gelid Winds
		46884,  //	Thick Ice Studded Collar
		46885,  //	Wraithguard Lord Commander's Cloak
		46886,  //	Embroidered Sash of the Icecaller
		46887,  //	Ancient Royal Froststeel Bracer
		46888,  //	Frost-Scarred Giant Hide Bracer
		46889,  //	Corrupted Orb of Living Ice
		46890,  //	Fjorlask, War Maul of Shadows
		46891,  //	Brightflame, Pride of the Lifebringer
		46892,  //	Earring of the Depths
		46893,  //	Slain Noble's Ring
		46894,  //	Crescent Courier's Earring
		46895,  //	Sothgar's Ghostly Mantle
		46896,  //	Spiked Dragongut Shoulderpads
		46897,  //	Cloak of the Scarlet Legion
		46898,  //	Flickering Cloak of Embers
		46899,  //	Everflame Torch of Ashengate
		46900,  //	Burning Heart of Sothgar
		46901,  //	Burned Remains of Shyleen
		46902,  //	Crimson Mask of Triumph
		46903,  //	Cinderskin Buckler
		46904,  //	Arcekor, the Instigator
		46905,  //	Spiritsender's Amberpike
		46906,  //	Bale and Brimstone
		46907,  //	Magmaband of Cestus Dei
		46908,  //	Scarlet Harbinger's Cloak
		46909,  //	Belt of the Vengeful Orphan
		46910,  //	Cape of the Lavaborn
		46911,  //	Mindbender's Sash
		46912,  //	Band of the Wildspirits
		46913,  //	Green Dragonscale Earring
		46914,  //	Hoop of Elemental Mastery
		46915,  //	Fetish of Unlife
		46916,  //	Slayer's Skullband
		46917,  //	Darkfell Fingerbone Necklace
		46918,  //	Haunting Ruby Amulet
		46919,  //	Buckled Belt of the Commandant
		46920,  //	Dyn`leth's Elegant Belt
		46921,  //	Scarlet Legion Command Stone
		46922,  //	Dyn`leth's Spy Glass
		46923,  //	Scaleguardian's Guise
		46924,  //	Daggerscale Visor
		46925,  //	Shield of Fire and Fury
		46926,  //	Trumpet of the Dragon Goddess
		46927,  //	Embershank, the Malady
		46928,  //	Cudgel of Unified Vision
		46929,  //	Spiked Brimstone Warhammer
		46930,  //	Vae`Aender, Stitch's Bow
		46931,  //	Earstud of a Mother's Love
		46932,  //	Wyrmband of the Devoted
		46933,  //	Gore-Splattered Shoulders
		46934,  //	Golden Amulet of the Elddar
		46935,  //	Chord of Lethar's Sinew
		46936,  //	Obsidian Scale Cape
		46937,  //	Spiked Bracelet of the Betrayer
		46938,  //	Embroidered Wool Band of Poetry
		46939,  //	Staff of Final Rites
		46940,  //	Fleshmelter, Lethar's Maul
		46941,  //	Tharkis' Select Reserve Catnip
		46942,  //	Acceber's Steamy Cinnamon Minotaur Cobbler
		46943,  //	Jyve's Savoury Eggs
		46944,  //	Merlu's Majestic Moonpie
		46945,  //	Souldor's Scrumptious Sausage
		46946,  //	Straha's Kunark Fried Cockatrice
		46947,  //	Last Meal of Terminat Hora'Diem
		46948,  //	Oriaas' Timeless Omelet
		46949,  //	Rouan's Draught of the Wanderer
		46950,  //	Tnexus' Terrible Tasting Tea
		46951,  //	Blacktalon Charm Whistle
		46952,  //	Blacktalon Feather Cloak
		46953,  //	Blacktalon Collar
		46954,  //	Blackfeather Raider Bauble
		46955,  //	Blackfeather Raider Mask
		46956,  //	Blackfeather Mantle
		46957,  //	Soulfed Green Legion Lifestone
		46958,  //	Unsated Green Legion Lifestone
		46959,  //	Weak Green Legion Lifestone
		46960,  //	Feeble Green Legion Lifestone
		46961,  //	Fragile Green Legion Lifestone
		46962,  //	Cloak of Icerime
		46963,  //	Lorekeeper Ritual Cope
		46964,  //	Mantle of the Deep Ice
		46965,  //	Icerage Ring
		46966,  //	Earring of Resonant Power
		46967,  //	Krithgor Heirloom Cape
		46968,  //	Heatstone of War
		46969,  //	Mask of Trials
		46970,  //	Blightstone
		46971,  //	Guardian Blade
		46972,  //	Frozen Tears
		46973,  //	Food and Drink of Timeless Heroes
		46974,  //	Faintly Pulsating Crystallized Cosgrove Shard
		46975,  //	Dimly Pulsating Crystallized Cosgrove Shard
		46976,  //	Pulsating Crystallized Cosgrove Shard
		46977,  //	Brightly Pulsating Crystallized Cosgrove Shard
		46978,  //	Brilliantly Pulsating Crystallized Cosgrove Shard
		46979,  //	Dream Dust
		46980,  //	Eidolon Plate Helm
		46981,  //	Eidolon Plate Vambraces
		46982,  //	Eidolon Plate Gauntlets
		46983,  //	Eidolon Plate Boots
		46984,  //	Eidolon Plate Bracers
		46985,  //	Eidolon Plate Breastplate
		46986,  //	Eidolon Plate Greaves
		46987,  //	Summoned: Muzzle of Mowcha
		46988,  //	Summoned: Lucid Belt
		46990,  //	Warweaver
		46991,  //	Runed Cold Iron Warhammer
		46992,  //	Gruesome Package
		46993,  //	Gruesome Package
		46994,  //	Sickly Smelling Package
		46995,  //	Severed Pirate Finger
		46996,  //	Kerran Fishing Pole
		46997,  //	Leapin Lizard Meat Pie
		46998,  //	Bolt of Checkered Cloth
		46999,  //	Bottle of Rum
		47000,  //	Gravity Core
		47001,  //	Class I Augmentation Distiller
		47002,  //	Class II Augmentation Distiller
		47003,  //	Class III Augmentation Distiller
		47004,  //	Class IV Augmentation Distiller
		47005,  //	Class V Augmentation Distiller
		47006,  //	Class VI Augmentation Distiller
		47007,  //	Class VII Augmentation Distiller
		47008,  //	Class VIII Augmentation Distiller
		47009,  //	Class IX Augmentation Distiller
		47010,  //	Class X Augmentation Distiller
		47011,  //	Class XI Augmentation Distiller
		47012,  //	Class XII Augmentation Distiller
		47013,  //	Class XIII Augmentation Distiller
		47014,  //	Class XIV Augmentation Distiller
		47015,  //	Class XV Augmentation Distiller
		47016,  //	Class XVI Augmentation Distiller
		47017,  //	Class XVII Augmentation Distiller
		47018,  //	Class XVIII Augmentation Distiller
		47019,  //	Class XIX Augmentation Distiller
		47020,  //	Class XX Augmentation Distiller
		47021,  //	Class XXI Augmentation Distiller
		47022,  //	Risen Drake Claw Earring
		47023,  //	Draco Heretic Mask
		47024,  //	Skaramatra's Token
		47025,  //	Drakeskin Draping
		47026,  //	Skaramatra Servant Cloak
		47027,  //	Necrocrystal Fragment
		47028,  //	Necrotic Drake Blood Ring
		47029,  //	Forgotten Crypt Key
		47030,  //	Head of the First Creation
		47031,  //	Class I Augmentation Solvent
		47032,  //	Class II Augmentation Solvent
		47033,  //	Class III Augmentation Solvent
		47034,  //	Class IV Augmentation Solvent
		47035,  //	Class V Augmentation Solvent
		47036,  //	Class VI Augmentation Solvent
		47037,  //	Class VII Augmentation Solvent
		47038,  //	Class VIII Augmentation Solvent
		47039,  //	Class IX Augmentation Solvent
		47040,  //	Class X Augmentation Solvent
		47041,  //	Class XI Augmentation Solvent
		47042,  //	Class XII Augmentation Solvent
		47043,  //	Class XIII Augmentation Solvent
		47044,  //	Class XIV Augmentation Solvent
		47045,  //	Class XV Augmentation Solvent
		47046,  //	Class XVI Augmentation Solvent
		47047,  //	Class XVII Augmentation Solvent
		47048,  //	Class XVIII Augmentation Solvent
		47049,  //	Class XIX Augmentation Solvent
		47050,  //	Class XX Augmentation Solvent
		47051,  //	Class XXI Augmentation Solvent
		47052,  //	Corrupted Stellite Shard
		47053,  //	Corrupted Celestrium Shard
		47054,  //	Corrupted Vitallium Shard
		47055,  //	Corrupted Damascite Shard
		47056,  //	Corrupted Palladium Shard
		47057,  //	Corrupted Iridium Shard
		47058,  //	Corrupted Rhodium Shard
		47060,  //	Platinum Plated Worm Pinion Gear Model I
		47061,  //	Platinum Plated Worm Pinion Gear Model II
		47062,  //	Platinum Plated Worm Pinion Gear Model III
		47063,  //	Platinum Plated Worm Pinion Gear Model IV
		47064,  //	Platinum Plated Worm Pinion Gear Model V
		47065,  //	Platinum Plated Worm Pinion Gear Model VI
		47066,  //	Platinum Plated Worm Pinion Gear Model VII
		47067,  //	Platinum Plated Worm Pinion Gear Model VIII
		47068,  //	Platinum Plated Worm Pinion Gear Model IX
		47069,  //	Platinum Plated Worm Pinion Gear Model X
		47070,  //	Platinum Plated Worm Pinion Gear Model XI
		47071,  //	Platinum Plated Worm Pinion Gear Model XII
		47072,  //	Platinum Plated Worm Pinion Gear Model XIII
		47073,  //	Platinum Plated Worm Pinion Gear Model XIV
		47074,  //	Platinum Plated Worm Pinion Gear Model XV
		47075,  //	Platinum Plated Worm Pinion Gear Model XVI
		47076,  //	Platinum Plated Worm Pinion Gear Model XVII
		47077,  //	Platinum Plated Worm Pinion Gear Model XVIII
		47078,  //	Platinum Plated Worm Pinion Gear Model XIX
		47079,  //	Platinum Plated Worm Pinion Gear Model XX
		47081,  //	Scrindite's Soil Shaker
		47082,  //	Rowdy Rock Dispenser
		47083,  //	Rowdy Rock
		47084,  //	Ever-growing Crystal
		47085,  //	Handful of Mud
		47086,  //	Gears of Plenty
		47087,  //	Title of the Earth Warden
		47088,  //	Stone Companion
		47089,  //	Clicking Beetle
		47090,  //	Ticking Companion
		47091,  //	Springwork Spider
		47092,  //	Square Cut Underfoot Diamond
		47093,  //	Round Cut Underfoot Diamond
		47094,  //	Oval Cut Underfoot Diamond
		47095,  //	Half-Moon Cut Underfoot Diamond
		47096,  //	Trilion Cut Underfoot Diamond
		47097,  //	Pear Cut Underfoot Diamond
		47098,  //	Marquise Cut Underfoot Diamond
		47099,  //	Extruded Underfoot Diamond
		47100,  //	Globe of Discordant Energy
		47102,  //	Brell's Bounty
		47103,  //	Prepared Brell's Bounty
		47104,  //	Dried Brell's Bounty
		47105,  //	Grilled Brell's Bounty
		47106,  //	Spiced Brell's Bounty
		47107,  //	Smoked Brell's Bounty
		47108,  //	Buttered Brell's Bounty
		47109,  //	Garlic-Buttered Brell's Bounty
		47110,  //	Lemon-Buttered Brell's Bounty
		47111,  //	Barbequed Brell's Bounty
		47112,  //	Beer-braised Brell's Bounty
		47113,  //	Beer-battered Brell's Bounty
		47114,  //	Grilled Brell's Bounty Stew
		47115,  //	Grilled Brell's Bounty Soup
		47116,  //	Grilled Brell's Bounty Jerky
		47117,  //	Grilled Brell's Bounty Souffle
		47118,  //	Grilled Brell's Bounty Steaks
		47119,  //	Grilled Brell's Bounty Kabobs
		47120,  //	Grilled Brell's Bounty Pie
		47121,  //	Grilled Brell's Bounty Sandwich
		47122,  //	Spiced Brell's Bounty Stew
		47123,  //	Spiced Brell's Bounty Soup
		47124,  //	Spiced Brell's Bounty Jerky
		47125,  //	Spiced Brell's Bounty Souffle
		47126,  //	Spiced Brell's Bounty Steaks
		47127,  //	Spiced Brell's Bounty Kabobs
		47128,  //	Spiced Brell's Bounty Pie
		47129,  //	Spiced Brell's Bounty Sandwich
		47130,  //	Smoked Brell's Bounty Stew
		47131,  //	Smoked Brell's Bounty Soup
		47132,  //	Smoked Brell's Bounty Jerky
		47133,  //	Smoked Brell's Bounty Souffle
		47134,  //	Smoked Brell's Bounty Steaks
		47135,  //	Smoked Brell's Bounty Kabobs
		47136,  //	Smoked Brell's Bounty Pie
		47137,  //	Smoked Brell's Bounty Sandwich
		47138,  //	Buttered Brell's Bounty Stew
		47139,  //	Buttered Brell's Bounty Soup
		47140,  //	Buttered Brell's Bounty Jerky
		47141,  //	Buttered Brell's Bounty Souffle
		47142,  //	Buttered Brell's Bounty Steaks
		47143,  //	Buttered Brell's Bounty Kabobs
		47144,  //	Buttered Brell's Bounty Pie
		47145,  //	Buttered Brell's Bounty Sandwich
		47146,  //	Garlic-Buttered Brell's Bounty Stew
		47147,  //	Garlic-Buttered Brell's Bounty Soup
		47148,  //	Garlic-Buttered Brell's Bounty Jerky
		47149,  //	Garlic-Buttered Brell's Bounty Souffle
		47150,  //	Garlic-Buttered Brell's Bounty Steaks
		47151,  //	Garlic-Buttered Brell's Bounty Kabobs
		47152,  //	Garlic-Buttered Brell's Bounty Pie
		47153,  //	Garlic-Buttered Brell's Bounty Sandwich
		47154,  //	Lemon-Buttered Brell's Bounty Stew
		47155,  //	Lemon-Buttered Brell's Bounty Soup
		47156,  //	Lemon-Buttered Brell's Bounty Jerky
		47157,  //	Lemon-Buttered Brell's Bounty Souffle
		47158,  //	Lemon-Buttered Brell's Bounty Steaks
		47159,  //	Lemon-Buttered Brell's Bounty Kabobs
		47160,  //	Lemon-Buttered Brell's Bounty Pie
		47161,  //	Lemon-Buttered Brell's Bounty Sandwich
		47162,  //	Barbequed Brell's Bounty Stew
		47163,  //	Barbequed Brell's Bounty Soup
		47164,  //	Barbequed Brell's Bounty Jerky
		47165,  //	Barbequed Brell's Bounty Souffle
		47166,  //	Barbequed Brell's Bounty Steaks
		47167,  //	Barbequed Brell's Bounty Kabobs
		47168,  //	Barbequed Brell's Bounty Pie
		47169,  //	Barbequed Brell's Bounty Sandwich
		47170,  //	Beer-braised Brell's Bounty Stew
		47171,  //	Beer-braised Brell's Bounty Soup
		47172,  //	Beer-braised Brell's Bounty Jerky
		47173,  //	Beer-braised Brell's Bounty Souffle
		47174,  //	Beer-braised Brell's Bounty Steaks
		47175,  //	Beer-braised Brell's Bounty Kabobs
		47176,  //	Beer-braised Brell's Bounty Pie
		47177,  //	Beer-braised Brell's Bounty Sandwich
		47178,  //	Beer-battered Brell's Bounty Stew
		47179,  //	Beer-battered Brell's Bounty Soup
		47180,  //	Beer-battered Brell's Bounty Jerky
		47181,  //	Beer-battered Brell's Bounty Souffle
		47182,  //	Beer-battered Brell's Bounty Steaks
		47183,  //	Beer-battered Brell's Bounty Kabobs
		47184,  //	Beer-battered Brell's Bounty Pie
		47185,  //	Beer-battered Brell's Bounty Sandwich
		47186,  //	Brell's Bounty Omelet
		47187,  //	Brell's Bounty Omelet with Cheese
		47188,  //	Deluxe Brell's Bounty Omelet
		47189,  //	Grilled Stuffed Brell's Bounty
		47190,  //	Stuffed Sausage Brell's Bounty
		47191,  //	Brell's Bounty Covered Mammoth Steaks
		47192,  //	Crab and Prawn Stuffed Brell's Bounty
		47193,  //	Brell's Bounty Rounds
		47194,  //	Baking Brell's Bounty
		47195,  //	Sealed New Year's Party Pack
		47197,  //	Tome of Friends
		47198,  //	The Magic Lute
		47200,  //	Greaves of Seething Rage
		47201,  //	Legplates of Abhorrent Memories
		47202,  //	Leggings of Fearsome Deeds
		47203,  //	Envenomed Leggings of Enmity
		47204,  //	Armplates of Endless Fortitude
		47205,  //	Vambraces of Perseverance
		47206,  //	Sleeves of the Steadfast
		47207,  //	Silken Sleeves of Persistence
		47208,  //	Hammered Helm of Foresight
		47209,  //	Helm of Inventive Perception
		47210,  //	Headband of Heightened Cognizance
		47211,  //	Twisted Crown of Consciousness
		47212,  //	Gauntlets of Singular Mastery
		47213,  //	Chain Gauntlets of Adroitness
		47214,  //	Gloves of the Adept
		47215,  //	Laced Gloves of Superiority
		47216,  //	Boots of Shifting Time
		47217,  //	Boots of Altered Perception
		47218,  //	Sandals of Constant Change
		47219,  //	Supple Slippers of Transformation
		47220,  //	Bracer of the Debauched
		47221,  //	Bracer of Corrupted Souls
		47222,  //	Wristband of Spectral Corruption
		47223,  //	Bracelet of the Corrupter
		47224,  //	Deathlace
		47225,  //	Pendant of Stilled Time
		47226,  //	Sapphire Choker of Adaptation
		47227,  //	Mantle of Iron Will
		47228,  //	Brute's Gaudy Shoulderspines
		47229,  //	Shawl of Nefarious Favor
		47230,  //	Belt of Intangible Clairvoyance
		47231,  //	Cord of the Malcontent
		47232,  //	Belt of Inevitable Conversion
		47233,  //	Earring of Mental Incursion
		47234,  //	Clawed Earring of Determination
		47235,  //	Specialist's Green Earstone
		47236,  //	Stud of Focused Aptitude
		47237,  //	Bloodband of Malice
		47238,  //	Band of Eternal Gaze
		47239,  //	Ring of Fluid Perception
		47240,  //	Ring of Ire Intent
		47241,  //	Brazier of Endless Flame
		47242,  //	Screaming Skull of Discontent
		47243,  //	Misty Globe
		47244,  //	Warped Mask of Animosity
		47245,  //	Mask of Uncanny Sight
		47246,  //	Veil of Intense Evolution
		47247,  //	Cloak of Flinty Resolve
		47248,  //	Two-Toned Fur Cape
		47249,  //	Cloak of the Faithless
		47250,  //	Totem of Shattered Hope
		47251,  //	Wand of Twisted Fate
		47252,  //	Battleworn Dented Aegis
		47253,  //	Howling Blood-Stained Bulwark
		47254,  //	Muramite Aggressor's Bulwark
		47255,  //	Relentless Guardian's Staff
		47256,  //	Thunderclap
		47257,  //	Jagged Pureblade
		47258,  //	Soulskive
		47259,  //	Fangs of the Serpent
		47260,  //	Seething Fists of Slaughter
		47261,  //	Worldwalker's Staff
		47262,  //	Cudgel of the Watchful Dragorn
		47263,  //	Bonecleaver
		47264,  //	Matchless Silvery Claymore
		47265,  //	Darkiron Bow of Betrayal
		47273,  //	Necklace of Stability
		47274,  //	Choker of Imprisoned Visions
		47275,  //	Necklace of the Steadfast Spirit
		47276,  //	Warbeads of the Magus
		47277,  //	Pauldrons of the Legion
		47278,  //	Shroud of Eternal Agony
		47279,  //	Shoulderpads of Warfare
		47280,  //	Amice of Ill-Will
		47281,  //	Belt of Contempt
		47282,  //	Chains of Anguish
		47283,  //	Girdle of the Fleet
		47284,  //	Belt of the Stagnant
		47285,  //	Stud of Chilling Precision
		47286,  //	Hanvar's Hoop
		47287,  //	Beaded Hoop of Demise
		47288,  //	Earring of Dragonkin
		47289,  //	Earring of Dark Conflict
		47290,  //	Ring of Persecution
		47291,  //	Ring of Deterrence
		47292,  //	Ring of the Beast
		47293,  //	Ring of Disdain
		47294,  //	Rigid Ring of Prowess
		47295,  //	Symbol of the Overlord
		47296,  //	Totem of the Chimera
		47297,  //	Sorrowmourn Stone
		47298,  //	Globe of Voltage
		47299,  //	Mask of Forbidden Rites
		47300,  //	Mask of Lament
		47301,  //	Golem Stone Face Guard
		47302,  //	Mask of the Crackling Energy
		47303,  //	Cloak of Wailing Woes
		47304,  //	Cloak of Regretful Transgressions
		47305,  //	Cape of Catastrophe
		47306,  //	Plagueborn Cape
		47307,  //	Infernal Staff of Fiery Fate
		47308,  //	Bazu Claw Hammer
		47309,  //	Shield of the Planar Assassin
		47310,  //	Shield of the Lightning Lord
		47311,  //	Aegis of the Dragorn Elders
		47312,  //	Notched Blade of Bloodletting
		47313,  //	Mace of Tortured Nightmares
		47314,  //	Hammer of Rancorous Thoughts
		47315,  //	Blade of Forgotten Faith
		47316,  //	Morguecaller
		47317,  //	Aneuk Dagger of Eye Gouging
		47318,  //	Flayed Flesh Handwraps
		47319,  //	Mace of Grim Tidings
		47320,  //	Blood-Polished Hammer
		47321,  //	Girplan Hammer of Carnage
		47322,  //	Bone Staff of Wickedness
		47323,  //	Greatsword of Mortification
		47324,  //	Warspear of Vexation
		47325,  //	Plaguebreeze
		47326,  //	Rune of Grim Portents
		47327,  //	Rune of Living Lightning
		47328,  //	Gem of Unnatural Regrowth
		47329,  //	Stone of Horrid Transformation
		47330,  //	Rune of Futile Resolutions
		47331,  //	Stone of Planar Protection
		47332,  //	Rune of Astral Celerity
		47333,  //	Abhorrent Brimstone of Charring
		47334,  //	Orb of Forbidden Laughter
		47335,  //	Petrified Girplan Heart
		47336,  //	Kyv Eye of Marksmanship
		47337,  //	Rune of Grim Portents
		47338,  //	Rune of Living Lightning
		47339,  //	Gem of Unnatural Regrowth
		47340,  //	Stone of Horrid Transformation
		47341,  //	Rune of Futile Resolutions
		47342,  //	Stone of Planar Protection
		47343,  //	Rune of Astral Celerity
		47344,  //	Abhorrent Brimstone of Charring
		47345,  //	Orb of Forbidden Laughter
		47346,  //	Petrified Girplan Heart
		47347,  //	Kyv Eye of Marksmanship
		47353,  //	Simple Recurve Bow
		47354,  //	Ornate Recurve Bow
		47355,  //	Intricate Recurve Bow
		47356,  //	Elaborate Recurve Bow
		47357,  //	Inured Recurve Bow
		47358,  //	Elegant Recurve Bow
		47359,  //	Stalwart Recurve Bow
		47360,  //	Blessed Reaching Symbol of the Skeptic
		47361,  //	Blessed Reaching Symbol of Infestation
		47362,  //	Blessed Reaching Symbol of Below
		47363,  //	Blessed Reaching Symbol of the Farceur
		47364,  //	Blessed Reaching Symbol of Terror
		47365,  //	Blessed Reaching Symbol of Thaumaturgy
		47366,  //	Blessed Reaching Symbol of Devotion
		47367,  //	Blessed Reaching Symbol of the Fire Lord
		47368,  //	Blessed Reaching Symbol of Hate
		47369,  //	Blessed Reaching Symbol of Storms
		47370,  //	Blessed Reaching Symbol of Integrity
		47371,  //	Blessed Reaching Symbol of the Depths
		47372,  //	Blessed Reaching Symbol of Stillness
		47373,  //	Blessed Reaching Symbol of the Warmonger
		47374,  //	Blessed Reaching Symbol of Compassion
		47375,  //	Blessed Reaching Symbol of Malignancy
		47376,  //	Blessed Reaching Symbol of Wildfire
		47377,  //	Blessed Reaching Symbol of Decay
		47378,  //	Blessed Reaching Symbol of Judgement
		47379,  //	Blessed Reaching Symbol of Growth
		47380,  //	Blessed Reaching Symbol of the Brood
		47381,  //	Revered Reaching Symbol of the Skeptic
		47382,  //	Revered Reaching Symbol of Infestation
		47383,  //	Revered Reaching Symbol of Below
		47384,  //	Revered Reaching Symbol of the Farceur
		47385,  //	Revered Reaching Symbol of Terror
		47386,  //	Revered Reaching Symbol of Thaumaturgy
		47387,  //	Revered Reaching Symbol of Devotion
		47388,  //	Revered Reaching Symbol of the Fire Lord
		47389,  //	Revered Reaching Symbol of Hate
		47390,  //	Revered Reaching Symbol of Storms
		47391,  //	Revered Reaching Symbol of Integrity
		47392,  //	Revered Reaching Symbol of the Depths
		47393,  //	Revered Reaching Symbol of Stillness
		47394,  //	Revered Reaching Symbol of the Warmonger
		47395,  //	Revered Reaching Symbol of Compassion
		47396,  //	Revered Reaching Symbol of Malignancy
		47397,  //	Revered Reaching Symbol of Wildfire
		47398,  //	Revered Reaching Symbol of Decay
		47399,  //	Revered Reaching Symbol of Judgement
		47400,  //	Revered Reaching Symbol of Growth
		47401,  //	Revered Reaching Symbol of the Brood
		47402,  //	Sacred Reaching Symbol of the Skeptic
		47403,  //	Sacred Reaching Symbol of Infestation
		47404,  //	Sacred Reaching Symbol of Below
		47405,  //	Sacred Reaching Symbol of the Farceur
		47406,  //	Sacred Reaching Symbol of Terror
		47407,  //	Sacred Reaching Symbol of Thaumaturgy
		47408,  //	Sacred Reaching Symbol of Devotion
		47409,  //	Sacred Reaching Symbol of the Fire Lord
		47410,  //	Sacred Reaching Symbol of Hate
		47411,  //	Sacred Reaching Symbol of Storms
		47412,  //	Sacred Reaching Symbol of Integrity
		47413,  //	Sacred Reaching Symbol of the Depths
		47414,  //	Sacred Reaching Symbol of Stillness
		47415,  //	Sacred Reaching Symbol of the Warmonger
		47416,  //	Sacred Reaching Symbol of Compassion
		47417,  //	Sacred Reaching Symbol of Malignancy
		47418,  //	Sacred Reaching Symbol of Wildfire
		47419,  //	Sacred Reaching Symbol of Decay
		47420,  //	Sacred Reaching Symbol of Judgement
		47421,  //	Sacred Reaching Symbol of Growth
		47422,  //	Sacred Reaching Symbol of the Brood
		47423,  //	Eminent Reaching Symbol of the Skeptic
		47424,  //	Eminent Reaching Symbol of Infestation
		47425,  //	Eminent Reaching Symbol of Below
		47426,  //	Eminent Reaching Symbol of the Farceur
		47427,  //	Eminent Reaching Symbol of Terror
		47428,  //	Eminent Reaching Symbol of Thaumaturgy
		47429,  //	Eminent Reaching Symbol of Devotion
		47430,  //	Eminent Reaching Symbol of the Fire Lord
		47431,  //	Eminent Reaching Symbol of Hate
		47432,  //	Eminent Reaching Symbol of Storms
		47433,  //	Eminent Reaching Symbol of Integrity
		47434,  //	Eminent Reaching Symbol of the Depths
		47435,  //	Eminent Reaching Symbol of Stillness
		47436,  //	Eminent Reaching Symbol of the Warmonger
		47437,  //	Eminent Reaching Symbol of Compassion
		47438,  //	Eminent Reaching Symbol of Malignancy
		47439,  //	Eminent Reaching Symbol of Wildfire
		47440,  //	Eminent Reaching Symbol of Decay
		47441,  //	Eminent Reaching Symbol of Judgement
		47442,  //	Eminent Reaching Symbol of Growth
		47443,  //	Eminent Reaching Symbol of the Brood
		47444,  //	Exalted Reaching Symbol of the Skeptic
		47445,  //	Exalted Reaching Symbol of Infestation
		47446,  //	Exalted Reaching Symbol of Below
		47447,  //	Exalted Reaching Symbol of the Farceur
		47448,  //	Exalted Reaching Symbol of Terror
		47449,  //	Exalted Reaching Symbol of Thaumaturgy
		47450,  //	Exalted Reaching Symbol of Devotion
		47451,  //	Exalted Reaching Symbol of the Fire Lord
		47452,  //	Exalted Reaching Symbol of Hate
		47453,  //	Exalted Reaching Symbol of Storms
		47454,  //	Exalted Reaching Symbol of Integrity
		47455,  //	Exalted Reaching Symbol of the Depths
		47456,  //	Exalted Reaching Symbol of Stillness
		47457,  //	Exalted Reaching Symbol of the Warmonger
		47458,  //	Exalted Reaching Symbol of Compassion
		47459,  //	Exalted Reaching Symbol of Malignancy
		47460,  //	Exalted Reaching Symbol of Wildfire
		47461,  //	Exalted Reaching Symbol of Decay
		47462,  //	Exalted Reaching Symbol of Judgement
		47463,  //	Exalted Reaching Symbol of Growth
		47464,  //	Exalted Reaching Symbol of the Brood
		47465,  //	Sublime Reaching Symbol of the Skeptic
		47466,  //	Sublime Reaching Symbol of Infestation
		47467,  //	Sublime Reaching Symbol of Below
		47468,  //	Sublime Reaching Symbol of the Farceur
		47469,  //	Sublime Reaching Symbol of Terror
		47470,  //	Sublime Reaching Symbol of Thaumaturgy
		47471,  //	Sublime Reaching Symbol of Devotion
		47472,  //	Sublime Reaching Symbol of the Fire Lord
		47473,  //	Sublime Reaching Symbol of Hate
		47474,  //	Sublime Reaching Symbol of Storms
		47475,  //	Sublime Reaching Symbol of Integrity
		47476,  //	Sublime Reaching Symbol of the Depths
		47477,  //	Sublime Reaching Symbol of Stillness
		47478,  //	Sublime Reaching Symbol of the Warmonger
		47479,  //	Sublime Reaching Symbol of Compassion
		47480,  //	Sublime Reaching Symbol of Malignancy
		47481,  //	Sublime Reaching Symbol of Wildfire
		47482,  //	Sublime Reaching Symbol of Decay
		47483,  //	Sublime Reaching Symbol of Judgement
		47484,  //	Sublime Reaching Symbol of Growth
		47485,  //	Sublime Reaching Symbol of the Brood
		47486,  //	Venerable Reaching Symbol of the Skeptic
		47487,  //	Venerable Reaching Symbol of Infestation
		47488,  //	Venerable Reaching Symbol of Below
		47489,  //	Venerable Reaching Symbol of the Farceur
		47490,  //	Venerable Reaching Symbol of Terror
		47491,  //	Venerable Reaching Symbol of Thaumaturgy
		47492,  //	Venerable Reaching Symbol of Devotion
		47493,  //	Venerable Reaching Symbol of the Fire Lord
		47494,  //	Venerable Reaching Symbol of Hate
		47495,  //	Venerable Reaching Symbol of Storms
		47496,  //	Venerable Reaching Symbol of Integrity
		47497,  //	Venerable Reaching Symbol of the Depths
		47498,  //	Venerable Reaching Symbol of Stillness
		47499,  //	Venerable Reaching Symbol of the Warmonger
		47500,  //	Venerable Reaching Symbol of Compassion
		47501,  //	Venerable Reaching Symbol of Malignancy
		47502,  //	Venerable Reaching Symbol of Wildfire
		47503,  //	Venerable Reaching Symbol of Decay
		47504,  //	Venerable Reaching Symbol of Judgement
		47505,  //	Venerable Reaching Symbol of Growth
		47506,  //	Venerable Reaching Symbol of the Brood
		47507,  //	Pyrilen Burn I
		47508,  //	Pyrilen Burn II
		47509,  //	Pyrilen Burn III
		47510,  //	Pyrilen Burn IV
		47511,  //	Pyrilen Burn V
		47512,  //	Pyrilen Burn VI
		47513,  //	Pyrilen Burn VII
		47514,  //	Pyrilen Burn VIII
		47515,  //	Pyrilen Burn IX
		47516,  //	Pyrilen Burn X
		47517,  //	Pyrilen Burn XI
		47518,  //	Pyrilen Burn XII
		47519,  //	Pyrilen Burn XIII
		47520,  //	Pyrilen Burn XIV
		47521,  //	Pyrilen Burn XV
		47522,  //	Gelidran Lament I
		47523,  //	Gelidran Lament II
		47524,  //	Gelidran Lament III
		47525,  //	Gelidran Lament IV
		47526,  //	Gelidran Lament V
		47527,  //	Gelidran Lament VI
		47528,  //	Gelidran Lament VII
		47529,  //	Gelidran Lament VIII
		47530,  //	Gelidran Lament IX
		47531,  //	Gelidran Lament X
		47532,  //	Gelidran Lament XI
		47533,  //	Gelidran Lament XII
		47534,  //	Gelidran Lament XIII
		47535,  //	Gelidran Lament XIV
		47536,  //	Gelidran Lament XV
		47537,  //	Hate of the Shissar I
		47538,  //	Hate of the Shissar II
		47539,  //	Hate of the Shissar III
		47540,  //	Hate of the Shissar IV
		47541,  //	Hate of the Shissar V
		47542,  //	Hate of the Shissar VI
		47543,  //	Hate of the Shissar VII
		47544,  //	Hate of the Shissar VIII
		47545,  //	Hate of the Shissar IX
		47546,  //	Hate of the Shissar X
		47547,  //	Hate of the Shissar XI
		47548,  //	Hate of the Shissar XII
		47549,  //	Hate of the Shissar XIII
		47550,  //	Hate of the Shissar XIV
		47551,  //	Hate of the Shissar XV
		47552,  //	Anger of the Shissar I
		47553,  //	Anger of the Shissar II
		47554,  //	Anger of the Shissar III
		47555,  //	Anger of the Shissar IV
		47556,  //	Anger of the Shissar V
		47557,  //	Anger of the Shissar VI
		47558,  //	Anger of the Shissar VII
		47559,  //	Anger of the Shissar VIII
		47560,  //	Anger of the Shissar IX
		47561,  //	Anger of the Shissar X
		47562,  //	Anger of the Shissar XI
		47563,  //	Anger of the Shissar XII
		47564,  //	Anger of the Shissar XIII
		47565,  //	Anger of the Shissar XIV
		47566,  //	Anger of the Shissar XV
		47567,  //	Tallon's Tactic XIV
		47568,  //	Tallon's Tactic XV
		47569,  //	Vallon's Tactic XIV
		47570,  //	Vallon's Tactic XV
		47571,  //	Innoruuk Emulsifier
		47572,  //	Granting More Peace
		47573,  //	Stellite Encrusted Phalangeal Clay
		47574,  //	Stellite Encrusted Carpal Clay
		47575,  //	Stellite Encrusted Brachial Clay
		47576,  //	Stellite Encrusted Tarsal Clay
		47577,  //	Stellite Encrusted Crural Clay
		47578,  //	Stellite Encrusted Thoracic Clay
		47579,  //	Stellite Encrusted Cephalic Clay
		47580,  //	Celestrium Encrusted Phalangeal Clay
		47581,  //	Celestrium Encrusted Carpal Clay
		47582,  //	Celestrium Encrusted Brachial Clay
		47583,  //	Celestrium Encrusted Tarsal Clay
		47584,  //	Celestrium Encrusted Crural Clay
		47585,  //	Celestrium Encrusted Thoracic Clay
		47586,  //	Celestrium Encrusted Cephalic Clay
		47587,  //	Vitallium Encrusted Phalangeal Clay
		47588,  //	Vitallium Encrusted Carpal Clay
		47589,  //	Vitallium Encrusted Brachial Clay
		47590,  //	Vitallium Encrusted Tarsal Clay
		47591,  //	Vitallium Encrusted Crural Clay
		47592,  //	Vitallium Encrusted Thoracic Clay
		47593,  //	Vitallium Encrusted Cephalic Clay
		47594,  //	Damascite Encrusted Phalangeal Clay
		47595,  //	Damascite Encrusted Carpal Clay
		47596,  //	Damascite Encrusted Brachial Clay
		47597,  //	Damascite Encrusted Tarsal Clay
		47598,  //	Damascite Encrusted Crural Clay
		47599,  //	Damascite Encrusted Thoracic Clay
		47600,  //	Damascite Encrusted Cephalic Clay
		47601,  //	Palladium Encrusted Phalangeal Clay
		47602,  //	Palladium Encrusted Carpal Clay
		47603,  //	Palladium Encrusted Brachial Clay
		47604,  //	Palladium Encrusted Tarsal Clay
		47605,  //	Palladium Encrusted Crural Clay
		47606,  //	Palladium Encrusted Thoracic Clay
		47607,  //	Palladium Encrusted Cephalic Clay
		47608,  //	Iridium Encrusted Phalangeal Clay
		47609,  //	Iridium Encrusted Carpal Clay
		47610,  //	Iridium Encrusted Brachial Clay
		47611,  //	Iridium Encrusted Tarsal Clay
		47612,  //	Iridium Encrusted Crural Clay
		47613,  //	Iridium Encrusted Thoracic Clay
		47614,  //	Iridium Encrusted Cephalic Clay
		47615,  //	Rhodium Encrusted Phalangeal Clay
		47616,  //	Rhodium Encrusted Carpal Clay
		47617,  //	Rhodium Encrusted Brachial Clay
		47618,  //	Rhodium Encrusted Tarsal Clay
		47619,  //	Rhodium Encrusted Crural Clay
		47620,  //	Rhodium Encrusted Thoracic Clay
		47621,  //	Rhodium Encrusted Cephalic Clay
		47622,  //	Martialism Bonding Agent
		47623,  //	Intensity Bonding Agent
		47624,  //	Purification Bonding Agent
		47625,  //	Dream Bonding Agent
		47626,  //	Necrotic Bonding Agent
		47627,  //	Consort Bonding Agent
		47628,  //	Dedicant Bonding Agent
		47629,  //	Druadic Bonding Agent
		47630,  //	Box of Latent Energies
		47631,  //	Bonding of Cosgrove Clay and Armor
		47632,  //	Forgeborn Mace Ornament
		47633,  //	Forgeborn Stave Ornament
		47634,  //	Forgeborn Blade Ornament
		47635,  //	Forgeborn Great Blade Ornament
		47636,  //	Forgeborn Dagger Ornament
		47637,  //	Forgeborn Spear Ornament
		47638,  //	Forgeborn Knuckles Ornament
		47639,  //	Forgeborn Shield Ornament
		47640,  //	Forgeborn Bow Ornament
		47641,  //	Aura of Lava
		47642,  //	Raex's Chestplate of Destruction
		47643,  //	Rizlona's Fiery Chestplate
		47644,  //	Trydan's Chestplate of Nobility
		47645,  //	Grimror's Guard of the Plagues
		47646,  //	Ultor's Chestguard of Faith
		47647,  //	Askr's Thunderous Chainmail
		47648,  //	Bidilis' Hauberk of the Elusive
		47649,  //	Rosrak's Hauberk of the Primal
		47650,  //	Ton Po's Chestwraps of Composure
		47651,  //	Kerasha's Sylvan Tunic
		47652,  //	Dumul's Chestwraps of the Brute
		47653,  //	Miragul's Shroud of Risen Souls
		47654,  //	Maelin's Robe of Lore
		47655,  //	Magi`Kot's Robe of Convergence
		47656,  //	Romar's Robe of Visions
		47657,  //	Chalandria's Bite I
		47658,  //	Chalandria's Bite II
		47659,  //	Chalandria's Bite III
		47660,  //	Chalandria's Bite IV
		47661,  //	Chalandria's Bite V
		47662,  //	Chalandria's Bite VI
		47663,  //	Chalandria's Bite VII
		47664,  //	Chalandria's Bite VIII
		47665,  //	Chalandria's Bite IX
		47666,  //	Chalandria's Bite X
		47667,  //	Chalandria's Bite XI
		47668,  //	Chalandria's Bite XII
		47669,  //	Chalandria's Bite XIII
		47670,  //	Chalandria's Bite XIV
		47671,  //	Chalandria's Bite XV
		47672,  //	Concentrated Chalandria's Bite I
		47673,  //	Concentrated Chalandria's Bite II
		47674,  //	Concentrated Chalandria's Bite III
		47675,  //	Concentrated Chalandria's Bite IV
		47676,  //	Concentrated Chalandria's Bite V
		47677,  //	Concentrated Chalandria's Bite VI
		47678,  //	Concentrated Chalandria's Bite VII
		47679,  //	Concentrated Chalandria's Bite VIII
		47680,  //	Concentrated Chalandria's Bite IX
		47681,  //	Concentrated Chalandria's Bite X
		47682,  //	Concentrated Chalandria's Bite XI
		47683,  //	Concentrated Chalandria's Bite XII
		47684,  //	Concentrated Chalandria's Bite XIII
		47685,  //	Concentrated Chalandria's Bite XIV
		47686,  //	Concentrated Chalandria's Bite XV
		47687,  //	Gem Steel War Hammer Ornament
		47688,  //	Gem Steel Stave Ornament
		47689,  //	Gem Steel Blade Ornament
		47690,  //	Gem Steel Great Blade Ornament
		47691,  //	Gem Steel Dagger Ornament
		47692,  //	Gem Steel Spear Ornament
		47693,  //	Gem Steel Spikefist Ornament
		47694,  //	Gem Steel Shield Ornament
		47695,  //	Gem Steel Bow Ornament
		47696,  //	Bag of Plenty
		47697,  //	Title of the Glutton
		47698,  //	Token of Reclamation
		47699,  //	Golem Bones
		47700,  //	Antheia Bloom Seed
		47701,  //	Jumpin Lizard Meat
		47702,  //	Shark Tooth
		47703,  //	Path of Skulls
		47704,  //	Hammer of Bones Ornament
		47705,  //	Staff of Bones Ornament
		47706,  //	Axe of Bones Ornament
		47707,  //	Great Blade of Bones Ornament
		47708,  //	Dagger of Skulls Ornament
		47709,  //	Spear of Bones Ornament
		47710,  //	Clawed Skull Ornament
		47711,  //	Frozen Skull of the Cursed
		47712,  //	Dancing Dagger of Thunder
		47713,  //	Dancing Blade of Fire
		47714,  //	Dancing Shard of Ice
		47715,  //	Plagued Skull Fragment
		47716,  //	Vanishing Point
		47717,  //	Razors Edge
		47718,  //	Blunt Force
		47719,  //	Archaeologist's Runic Belt
		47720,  //	Dragon Brood Crypt Key
		47721,  //	Hammer of Bones Ornamentation
		47722,  //	Staff of Bones Ornamentation
		47723,  //	Axe of Bones Ornamentation
		47724,  //	Great Blade of Bones Ornamentation
		47725,  //	Dagger of Skulls Ornamentation
		47726,  //	Spear of Bones Ornamentation
		47727,  //	Clawed Skull Ornamentation
		47728,  //	Hammer of Bones Ornament Pack
		47729,  //	Staff of Bones Ornament Pack
		47730,  //	Axe of Bones Ornament Pack
		47731,  //	Great Blade of Bones Ornament Pack
		47732,  //	Dagger of Skulls Ornament Pack
		47733,  //	Spear of Bones Ornament Pack
		47734,  //	Clawed Skull Ornament Pack
		47735,  //	Plagued Skull Fragment Pack
		47736,  //	Path of Skulls Pack
		47737,  //	Empty Glowing Skull
		47738,  //	Frost Covered Gem
		47739,  //	Trapped Soul Gem
		47740,  //	Bloodshard Earring
		47741,  //	Tarnished Blood Drop Accent Veil
		47742,  //	Dragoncrypt Token
		47743,  //	Mildew Covered Shawl
		47744,  //	Royal Dragon Hunter's Cloak
		47745,  //	Disempowered Phylactery
		47746,  //	Tarnished Dracoform Ring
		47747,  //	Unwritten Glyph
		47748,  //	Jagged Shard
		47749,  //	Pulsing Sphere
		47750,  //	Chipped Claw
		47751,  //	Crumbling Head
		47752,  //	Broken Torso
		47753,  //	Cracking Legs
		47754,  //	Worn Arms
		47755,  //	Eroding Feet
		47756,  //	Destructive Golem
		47757,  //	Vault Key
		47758,  //	Keystone Mold
		47759,  //	Keystone
		47760,  //	Unstable Glob of Clay
		47761,  //	Vial of Purified Water
		47762,  //	Unstable Mass of Clay
		47765,  //	Bottle of the Avian Pack
		47766,  //	Bottle of the Avian I
		47767,  //	Bottle of the Avian II
		47768,  //	Bottle of the Avian III
		47769,  //	Bottle of the Avian IV
		47770,  //	Bottle of the Avian V
		47771,  //	Bottle of the Avian VI
		47772,  //	Bottle of the Avian VII
		47773,  //	Bottle of the Avian VIII
		47774,  //	Bottle of the Avian IX
		47775,  //	Bottle of the Avian X
		47776,  //	Bottle of the Avian XI
		47777,  //	Bottle of the Avian XII
		47778,  //	Bottle of the Avian XIII
		47779,  //	Bottle of the Avian XIV
		47780,  //	Bottle of the Cetacean Pack
		47781,  //	Bottle of the Cetacean I
		47782,  //	Bottle of the Cetacean II
		47783,  //	Bottle of the Cetacean III
		47784,  //	Bottle of the Cetacean IV
		47785,  //	Bottle of the Cetacean V
		47786,  //	Bottle of the Cetacean VI
		47787,  //	Bottle of the Cetacean VII
		47788,  //	Bottle of the Cetacean VIII
		47789,  //	Bottle of the Cetacean IX
		47790,  //	Bottle of the Cetacean X
		47791,  //	Bottle of the Cetacean XI
		47792,  //	Bottle of the Cetacean XII
		47793,  //	Bottle of the Cetacean XIII
		47794,  //	Bottle of the Cetacean XIV
		47795,  //	Skull of the Amused
		47797,  //	Hand of Stone
		47798,  //	Vampire Mace Ornament
		47799,  //	Vampire Great Hammer Ornament
		47800,  //	Vampire Blade Ornament
		47801,  //	Vampire Great Blade Ornament
		47802,  //	Vampire Dagger Ornament
		47803,  //	Vampire Spear Ornament
		47804,  //	Vampire Spikefist Ornament
		47805,  //	Vampire Shield Ornament
		47806,  //	Aura of the Belfry
		47807,  //	Title of the Hunt
		47808,  //	Title of the Lion Tamer
		47809,  //	Title of the Stone Breaker
		47810,  //	Beden's Collar
		47811,  //	Collar of the King
		47812,  //	Captured Bone of the Lost
		47813,  //	Blessed Underfoot Prayer Shawl
		47814,  //	Revered Underfoot Prayer Shawl
		47815,  //	Sacred Underfoot Prayer Shawl
		47816,  //	Blessed Sanctified Underfoot Prayer Shawl
		47817,  //	Blessed Sanctified Underfoot Prayer Shawl
		47818,  //	Revered Sanctified Underfoot Prayer Shawl
		47819,  //	Revered Sanctified Underfoot Prayer Shawl
		47820,  //	Sacred Sanctified Underfoot Prayer Shawl
		47821,  //	Sacred Sanctified Underfoot Prayer Shawl
		47822,  //	Blessed Serilis Prayer Shawl
		47823,  //	Blessed Serilis Prayer Shawl
		47824,  //	Blessed Serilis Prayer Shawl
		47825,  //	Blessed Serilis Prayer Shawl
		47826,  //	Sacred Serilis Prayer Shawl
		47827,  //	Sacred Serilis Prayer Shawl
		47828,  //	Sacred Serilis Prayer Shawl
		47829,  //	Sacred Serilis Prayer Shawl
		47830,  //	Speckled Golem Stone
		47831,  //	Midnight Gem
		47832,  //	Hunk of Crystal
		47833,  //	Sun Ruby
		47834,  //	Underfoot Topaz
		47835,  //	Brell's Gold
		47836,  //	Hardened Golem's Tear
		47837,  //	Chunk of Floating Ore
		47838,  //	Worthless Rock
		47839,  //	Chunk of Coal
		47840,  //	Tainted Lichen Sample
		47841,  //	Regurgitated Lichen
		47842,  //	Full Purified Lichen Potion
		47844,  //	Slightly Full Purified Lichen Potion
		47845,  //	Half Full Purified Lichen Potion
		47846,  //	Half Empty Purified Lichen Potion
		47847,  //	Almost Empty Purified Lichen Potion
		47848,  //	Treant Bark
		47849,  //	Faernoc's Fang
		47850,  //	Brellian Ore
		47851,  //	Eternal Water
		47852,  //	Thick Kobold Hide
		47853,  //	Raw Gemstone
		47854,  //	Raw Wood
		47855,  //	Sweet Sap
		47856,  //	Strands of Vines
		47857,  //	Waxy Clay
		47858,  //	Brellian Sheet Metal
		47859,  //	Block of Brellian Metal
		47860,  //	Sparkling Gemstones
		47861,  //	Gemless Bracers
		47862,  //	Leather Hilted Handle
		47863,  //	Mallet Head
		47864,  //	Brew Keg
		47865,  //	Brellian Brew
		47866,  //	Bundle of Candles
		47867,  //	Gemless Brellian Crown
		47868,  //	Jeweled Brellian Crown
		47869,  //	Muddite Head
		47870,  //	Muddite Torso
		47871,  //	Pair of Muddite Limbs
		47872,  //	Kobold Leather Hood
		47873,  //	Hood Clasp
		47874,  //	Kobold Leather Cloak
		47875,  //	Brell's Jeweled Bracers
		47876,  //	Brell's Broad-headed Hammer
		47877,  //	Keg of Brellian Brew
		47878,  //	Brell's Candled Crown
		47879,  //	Brell's Basic Companion
		47880,  //	Brell's Hooded Cloak
		47881,  //	Kobold Leather Straps
		47882,  //	Kobold Leather Padding
		47883,  //	Unfired Muddite Companion
		47884,  //	Brellian Metal Bands
		47885,  //	Bundle of Wooden Staves
		47886,  //	Candle Wicks
		47887,  //	Strands of Thread
		47888,  //	Spool of Thread
		47889,  //	Spool of Yarn
		47890,  //	Sticky Sap
		47891,  //	Kobold Hide Robe
		47892,  //	Weak Brellian Brew
		47893,  //	Keg of Weak Brellian Brew
		47894,  //	Transforming Surprise of The Stone Warden
		47895,  //	Transforming Surprise of Unstable Creation
		47896,  //	Transforming Surprise of Borzaloth
		47897,  //	Transforming Surprise of Kilreck the Cleaner
		47898,  //	Transforming Surprise of Magus Sisters
		47900,  //	Brellium Token
		47901,  //	Brell's Bounty Steamed Steaks
		47902,  //	Essence of Steam
		47903,  //	Divine Snow Griffin Souffle
		47904,  //	Neutralized Goo Acid
		47905,  //	Ale Marinated Cliknar Claws
		47906,  //	Intact Cliknar Claw
		47907,  //	Turuff's Favorite Meat
		47908,  //	Clockwork Stew
		47909,  //	Organic Clockwork Oil
		47910,  //	Clockwork Balance Gear
		47911,  //	Gribblebitz Clay Sample
		47912,  //	Foundation Clay Sample
		47913,  //	Golem Clay Sample
		47914,  //	Cliknar Claimed Clay Sample
		47916,  //	Clay of Cosgrove Sample
		47917,  //	Clay Sampling Device
		47918,  //	Pure Water Collecting Tool
		47919,  //	Sterilized Pure Water Collecting Tool
		47920,  //	Vial of Purest Water
		47921,  //	Unfired Clay Statue
		47922,  //	Unfired Clay Statue
		47923,  //	Unfired Clay Statue
		47924,  //	Unfired Clay Statue
		47925,  //	Unfired Clay Statue
		47926,  //	Clay Statue
		47927,  //	Clay Statue
		47928,  //	Clay Statue
		47929,  //	Clay Statue
		47930,  //	Clay Statue
		47931,  //	Glacial Slush Water Collector
		47932,  //	Glacial Slush Water
		47933,  //	Downy Gnoll Fur
		47934,  //	Downy Kobold Fur
		47935,  //	Pristine Mephit Wing Veins
		47936,  //	Loam Tainted Spider Silk
		47937,  //	Loam Tainted Slime
		47938,  //	Woven Fur Padding
		47939,  //	Wing Vein Boning
		47940,  //	Underfoot Silk Thread
		47941,  //	Underfoot Silk Swatch
		47942,  //	Underfoot Silk Panel
		47943,  //	Frost Shocked Downy Gnoll Fur
		47944,  //	Frost Shocked Downy Kobold Fur
		47945,  //	Loam Tainted Slime
		47946,  //	Tangled Lump of Black Spider Silk
		47947,  //	Cosgrove Amalgam
		47948,  //	Essence of Reason
		47949,  //	Essence of Will
		47950,  //	Essence of Passion
		47951,  //	Essence of Substance
		47952,  //	Essence of Instinct
		47953,  //	Essence of Vitality
		47954,  //	Essence of Chaos
		47955,  //	Raw Gem of Reason
		47956,  //	Raw Gem of Will
		47957,  //	Raw Gem of Passion
		47958,  //	Raw Gem of Substance
		47959,  //	Raw Gem of Instinct
		47960,  //	Raw Gem of Vitality
		47961,  //	Raw Gem of Chaos
		47962,  //	Brilliant Gem of Reason
		47963,  //	Brilliant Gem of Will
		47964,  //	Brilliant Gem of Passion
		47965,  //	Brilliant Gem of Substance
		47966,  //	Brilliant Gem of Instinct
		47967,  //	Brilliant Gem of Vitality
		47968,  //	Brilliant Gem of Chaos
		47969,  //	Clasp of the Elements
		47970,  //	Magical Water Collection Jar
		47971,  //	Brell's Rest Water Sample
		47972,  //	Cooling Chamber Water Sample
		47973,  //	Underquarry Water Sample
		47974,  //	Foundation Water Sample
		47975,  //	Pellucid Grotto Water Sample
		47976,  //	Arthicrex Water Sample
		47977,  //	Fungal Forest Water Sample
		47978,  //	Pure Underfoot Spirits
		47979,  //	Jelly of the Cliknar Queen
		47980,  //	Sacred Cask of Brell
		47981,  //	Brell's Mystical Brew
		47982,  //	Fallen Branch
		47983,  //	Intact Greken Spine
		47984,  //	Rune Carving Tools
		47985,  //	Smooth Stone
		47986,  //	First Sacred Rune Pattern
		47987,  //	Second Sacred Rune Pattern
		47988,  //	Third Sacred Rune Pattern
		47989,  //	Fourth Sacred Rune Pattern
		47990,  //	Fifth Sacred Rune Pattern
		47991,  //	Sixth Sacred Rune Pattern
		47992,  //	First Sacred Prayer Rune
		47993,  //	Second Sacred Prayer Rune
		47994,  //	Third Sacred Prayer Rune
		47995,  //	Fourth Sacred Prayer Rune
		47996,  //	Fifth Sacred Prayer Rune
		47997,  //	Sixth Sacred Prayer Rune
		47998,  //	Wood Panel
		47999,  //	Ornate Trim
		48001,  //	Unfinished Note
		48002,  //	Book of Bindings
		48003,  //	Unfinished Note
		48004,  //	Unfinished Note
		48005,  //	Unfinished Note
		48006,  //	Illegible Message
		48007,  //	Silver Oxide
		48008,  //	Sealed Message
		48009,  //	Champion Crown Plans
		48010,  //	Champion's Mancatcher Crown
		48011,  //	Drix's Satchel
		48012,  //	Drix's Satchel
		48013,  //	Champion's Mancatcher
		48014,  //	Geot's Pack
		48015,  //	Xyzith's Belongings
		48016,  //	Tattered Bracer
		48017,  //	Rusted Shield
		48018,  //	Broken Mancatcher Crown
		48019,  //	Symbol of the Legion
		48020,  //	Seal of Rixiz
		48021,  //	Head of a Traitor
		48022,  //	Band of the Chosen
		48023,  //	Band of the Chosen
		48024,  //	Band of the Chosen
		48025,  //	Klok Scaleroot's Note
		48026,  //	Kyg's Approval
		48027,  //	Broken Medal
		48028,  //	Mirrored Gem
		48029,  //	Black Stone
		48030,  //	Xyzith's Spirit
		48031,  //	Unfinished Mancatcher
		48033,  //	Black Stone
		48034,  //	Brittle Bone
		48035,  //	Poisoned Soul
		48036,  //	Fiery Orb
		48037,  //	Rogue Marauder's Head
		48038,  //	Pure Essence
		48039,  //	Glosk's Sack
		48040,  //	Gem of Reflection
		48041,  //	Ixpacan's Tome
		48042,  //	Ixpacan's Tome
		48043,  //	Demi Lich Skullcap
		48044,  //	Child of Charasis Remains
		48045,  //	Claw of the Cub
		48046,  //	Claw of the Scout
		48047,  //	Claw of the Apprentice
		48048,  //	Claw of the Young Patriarch
		48049,  //	Claw of the Mature Patriarch
		48050,  //	Claw of the Spiritual Elder
		48051,  //	Welgaz's Note
		48052,  //	Battle Plans
		48053,  //	Battle Plans
		48054,  //	Battle Plans
		48055,  //	Battle Plans
		48056,  //	Battle Plans
		48057,  //	Battle Plans
		48058,  //	Battle Plans
		48059,  //	Battle Plans
		48060,  //	Spiritual Tome
		48061,  //	Spiritual Tome
		48062,  //	Hidden Plans
		48063,  //	Hidden Plans
		48064,  //	Lativ's Plans
		48065,  //	Symbol of Lativ
		48066,  //	Symbol of Lativ
		48067,  //	Symbol of Lativ
		48068,  //	Lativ's Remains
		48069,  //	Broken Seal
		48070,  //	Broken Seal
		48071,  //	Broken Seal
		48072,  //	Broken Seal
		48073,  //	Vuzx's Container
		48074,  //	Seal of Choatl
		48075,  //	Sacred Figurine
		48076,  //	Lativ's Remains
		48077,  //	Torch of Truth
		48079,  //	Crypt Key
		48080,  //	Broken Gypsy Lute
		48081,  //	Fibblereaper's Core
		48082,  //	Griffin's Egg
		48083,  //	Bottle of Shared Adventure I
		48084,  //	Bottle of Shared Adventure II
		48085,  //	Bottle of Shared Adventure III
		48086,  //	Brell's Hammer Ornamentation
		48087,  //	Gift: Haiti Relief Plush Bear
		48088,  //	Tread of the Mountain Walker
		48090,  //	Unstable Ooze
		48091,  //	Empathetic Ooze
		48092,  //	Spitting Ooze
		48093,  //	Rune Case
		48094,  //	Case of Runes
		48095,  //	Collapsed Chronal Transmutation Device
		48096,  //	Chronal Transmutation Device
		48097,  //	Chronal Transmutation And You
		48098,  //	Cela's necklace
		48099,  //	Piece of Aeriht's Book
		48100,  //	Seal of Cedrick
		48101,  //	Seal of Mystik
		48102,  //	Medal of Identity
		48104,  //	Section of Parchment
		48105,  //	Section of Parchment
		48106,  //	Section of Parchment
		48107,  //	Section of Parchment
		48109,  //	1/2 of Kaiaren's Diary
		48110,  //	1/2 of Kaiaren's Diary
		48111,  //	Symbol of Moon
		48112,  //	Symbol of Sun
		48113,  //	Symbol of Order
		48114,  //	Symbol of Tranquility
		48115,  //	Symbol of Focus
		48117,  //	Sunless Nectar
		48118,  //	Aged Piece of Meat
		48119,  //	Sugar and Spice Nectar Pie
		48120,  //	Aged Meat and Cheese Sandwich
		48122,  //	Muramite Sand
		48123,  //	Celestial Thread
		48124,  //	Polished Symbol of Sun
		48125,  //	Polished Symbol of Moon
		48126,  //	Polished Symbol of Order
		48127,  //	Initiate's Sash of the Celestial Order
		48128,  //	Letter to Stremstin
		48129,  //	Polished Symbol of Focus
		48130,  //	Polished Symbol of Tranquility
		48132,  //	Danl's Missing Book
		48136,  //	Innoruuk's Dark Blessing
		48137,  //	Kaiaren's Mind
		48138,  //	Kaiaren's Body
		48139,  //	Letter from Tirranun
		48140,  //	Ionized Ore
		48141,  //	Drake Scales
		48142,  //	Magma Carapace
		48143,  //	Poison Sac
		48144,  //	Drake Fang
		48146,  //	Ancient Note of Runes
		48147,  //	Nightbane, Sword of the Valiant
		48148,  //	Symbol of Mistmoore
		48149,  //	The Green Fog
		48156,  //	Fistwrap Pattern
		48157,  //	Sap Pattern
		48158,  //	Blessed Water of Terris Thule
		48159,  //	Blessed Water of Fennin Ro
		48160,  //	Blessed Water of Saryrn
		48161,  //	Blessed Water of Druzzil Ro
		48162,  //	Vanadium Ore
		48163,  //	Immaculate Animal Pelt
		48164,  //	Thalium Barbs
		48165,  //	Thalium Studs
		48166,  //	Thalium Fill
		48167,  //	Fulginate Barbs
		48168,  //	Fulginate Studs
		48169,  //	Fulginate Fill
		48170,  //	Rhenium Barbs
		48171,  //	Rhenium Studs
		48172,  //	Rhenium Fill
		48173,  //	Cobalt Barbs
		48174,  //	Cobalt Studs
		48175,  //	Cobalt Fill
		48176,  //	Titanium Barbs
		48177,  //	Titanium Studs
		48178,  //	Titanium Fill
		48179,  //	Tantalum Barbs
		48180,  //	Tantalum Studs
		48181,  //	Tantalum Fill
		48182,  //	Vanadium Barbs
		48183,  //	Vanadium Studs
		48184,  //	Vanadium Fill
		48185,  //	Sullied Hilt Wrap
		48186,  //	Rough Hilt Wrap
		48187,  //	Fine Hilt Wrap
		48188,  //	Superb Hilt Wrap
		48189,  //	Flawless Hilt Wrap
		48190,  //	Exquisite Hilt Wrap
		48191,  //	Immaculate Hilt Wrap
		48192,  //	Simple Short Haft
		48193,  //	Simple Haft
		48194,  //	Simple Long Haft
		48195,  //	Simple Baton
		48196,  //	Simple Staff
		48197,  //	Simple Handle
		48198,  //	Ornate Short Haft
		48199,  //	Ornate Haft
		48200,  //	Ornate Long Haft
		48201,  //	Ornate Baton
		48202,  //	Ornate Staff
		48203,  //	Ornate Handle
		48204,  //	Intricate Short Haft
		48205,  //	Intricate Haft
		48206,  //	Intricate Long Haft
		48207,  //	Intricate Baton
		48208,  //	Intricate Staff
		48209,  //	Intricate Handle
		48210,  //	Elaborate Short Haft
		48211,  //	Elaborate Haft
		48212,  //	Elaborate Long Haft
		48213,  //	Elaborate Baton
		48214,  //	Elaborate Staff
		48215,  //	Elaborate Handle
		48216,  //	Inured Short Haft
		48217,  //	Inured Haft
		48218,  //	Inured Long Haft
		48219,  //	Inured Baton
		48220,  //	Inured Staff
		48221,  //	Inured Handle
		48222,  //	Elegant Short Haft
		48223,  //	Elegant Haft
		48224,  //	Elegant Long Haft
		48225,  //	Elegant Baton
		48226,  //	Elegant Staff
		48227,  //	Elegant Handle
		48228,  //	Stalwart Short Haft
		48229,  //	Stalwart Haft
		48230,  //	Stalwart Long Haft
		48231,  //	Stalwart Baton
		48232,  //	Stalwart Staff
		48233,  //	Stalwart Handle
		48234,  //	Simple Small Hilt Casting Mold
		48235,  //	Simple Hilt Casting Mold
		48236,  //	Simple Large Hilt Casting Mold
		48237,  //	Simple Blade Shaped Mold
		48238,  //	Simple Heavy Blade Shaped Mold
		48239,  //	Simple Axe Head Shaped Mold
		48240,  //	Simple Small Blade Shaped Mold
		48241,  //	Simple Spear Head Shaped Mold
		48242,  //	Simple Mace Head Shaped Mold
		48243,  //	Simple Hammer Head Shaped Mold
		48244,  //	Simple Metal Wrap Shaped Mold
		48245,  //	Simple Ulak Blade Shaped Mold
		48246,  //	Simple Large Blade Shaped Mold
		48247,  //	Simple Large Axe Head Shaped Mold
		48248,  //	Ornate Small Hilt Casting Mold
		48249,  //	Ornate Hilt Casting Mold
		48250,  //	Ornate Large Hilt Casting Mold
		48251,  //	Ornate Blade Shaped Mold
		48252,  //	Ornate Heavy Blade Shaped Mold
		48253,  //	Ornate Axe Head Shaped Mold
		48254,  //	Ornate Small Blade Shaped Mold
		48255,  //	Ornate Spear Head Shaped Mold
		48256,  //	Ornate Mace Head Shaped Mold
		48257,  //	Ornate Hammer Head Shaped Mold
		48258,  //	Ornate Metal Wrap Shaped Mold
		48259,  //	Ornate Ulak Blade Shaped Mold
		48260,  //	Ornate Large Blade Shaped Mold
		48261,  //	Ornate Large Axe Head Shaped Mold
		48262,  //	Intricate Small Hilt Casting Mold
		48263,  //	Intricate Hilt Casting Mold
		48264,  //	Intricate Large Hilt Casting Mold
		48265,  //	Intricate Blade Shaped Mold
		48266,  //	Intricate Heavy Blade Shaped Mold
		48267,  //	Intricate Axe Head Shaped Mold
		48268,  //	Intricate Small Blade Shaped Mold
		48269,  //	Intricate Spear Head Shaped Mold
		48270,  //	Intricate Mace Head Shaped Mold
		48271,  //	Intricate Hammer Head Shaped Mold
		48272,  //	Intricate Metal Wrap Shaped Mold
		48273,  //	Intricate Ulak Blade Shaped Mold
		48274,  //	Intricate Large Blade Shaped Mold
		48275,  //	Intricate Large Axe Head Shaped Mold
		48276,  //	Elaborate Small Hilt Casting Mold
		48277,  //	Elaborate Hilt Casting Mold
		48278,  //	Elaborate Large Hilt Casting Mold
		48279,  //	Elaborate Blade Shaped Mold
		48280,  //	Elaborate Heavy Blade Shaped Mold
		48281,  //	Elaborate Axe Head Shaped Mold
		48282,  //	Elaborate Small Blade Shaped Mold
		48283,  //	Elaborate Spear Head Shaped Mold
		48284,  //	Elaborate Mace Head Shaped Mold
		48285,  //	Elaborate Hammer Head Shaped Mold
		48286,  //	Elaborate Metal Wrap Shaped Mold
		48287,  //	Elaborate Ulak Blade Shaped Mold
		48288,  //	Elaborate Large Blade Shaped Mold
		48289,  //	Elaborate Large Axe Head Shaped Mold
		48290,  //	Inured Small Hilt Casting Mold
		48291,  //	Inured Hilt Casting Mold
		48292,  //	Inured Large Hilt Casting Mold
		48293,  //	Inured Blade Shaped Mold
		48294,  //	Inured Heavy Blade Shaped Mold
		48295,  //	Inured Axe Head Shaped Mold
		48296,  //	Inured Small Blade Shaped Mold
		48297,  //	Inured Spear Head Shaped Mold
		48298,  //	Inured Mace Head Shaped Mold
		48299,  //	Inured Hammer Head Shaped Mold
		48300,  //	Inured Metal Wrap Shaped Mold
		48301,  //	Inured Ulak Blade Shaped Mold
		48302,  //	Inured Large Blade Shaped Mold
		48303,  //	Inured Large Axe Head Shaped Mold
		48304,  //	Elegant Small Hilt Casting Mold
		48305,  //	Elegant Hilt Casting Mold
		48306,  //	Elegant Large Hilt Casting Mold
		48307,  //	Elegant Blade Shaped Mold
		48308,  //	Elegant Heavy Blade Shaped Mold
		48309,  //	Elegant Axe Head Shaped Mold
		48310,  //	Elegant Small Blade Shaped Mold
		48311,  //	Elegant Spear Head Shaped Mold
		48312,  //	Elegant Mace Head Shaped Mold
		48313,  //	Elegant Hammer Head Shaped Mold
		48314,  //	Elegant Metal Wrap Shaped Mold
		48315,  //	Elegant Ulak Blade Shaped Mold
		48316,  //	Elegant Large Blade Shaped Mold
		48317,  //	Elegant Large Axe Head Shaped Mold
		48318,  //	Stalwart Small Hilt Casting Mold
		48319,  //	Stalwart Hilt Casting Mold
		48320,  //	Stalwart Large Hilt Casting Mold
		48321,  //	Stalwart Blade Shaped Mold
		48322,  //	Stalwart Heavy Blade Shaped Mold
		48323,  //	Stalwart Axe Head Shaped Mold
		48324,  //	Stalwart Small Blade Shaped Mold
		48325,  //	Stalwart Spear Head Shaped Mold
		48326,  //	Stalwart Mace Head Shaped Mold
		48327,  //	Stalwart Hammer Head Shaped Mold
		48328,  //	Stalwart Metal Wrap Shaped Mold
		48329,  //	Stalwart Ulak Blade Shaped Mold
		48330,  //	Stalwart Large Blade Shaped Mold
		48331,  //	Stalwart Large Axe Head Shaped Mold
		48332,  //	Simple Leather Whip
		48333,  //	Simple Leather Fist Wraps
		48334,  //	Simple Sap
		48335,  //	Ornate Leather Whip
		48336,  //	Ornate Leather Fist Wraps
		48337,  //	Ornate Sap
		48338,  //	Intricate Leather Whip
		48339,  //	Intricate Leather Fist Wraps
		48340,  //	Intricate Sap
		48341,  //	Elaborate Leather Whip
		48342,  //	Elaborate Leather Fist Wraps
		48343,  //	Elaborate Sap
		48344,  //	Inured Leather Whip
		48345,  //	Inured Leather Fist Wraps
		48346,  //	Inured Sap
		48347,  //	Elegant Leather Whip
		48348,  //	Elegant Leather Fist Wraps
		48349,  //	Elegant Sap
		48350,  //	Stalwart Leather Whip
		48351,  //	Stalwart Leather Fist Wraps
		48352,  //	Stalwart Sap
		48353,  //	Simple Long Sword
		48354,  //	Simple Knight's Sword
		48355,  //	Simple Hand Axe
		48356,  //	Simple Dagger
		48357,  //	Simple Short Spear
		48358,  //	Simple Mace
		48359,  //	Simple Hammer
		48360,  //	Simple Reinforced Baton
		48361,  //	Simple Ulak
		48362,  //	Simple Great Sword
		48363,  //	Simple Great Axe
		48364,  //	Simple Long Spear
		48365,  //	Simple Reinforced Staff
		48366,  //	Simple Long Hammer
		48367,  //	Ornate Long Sword
		48368,  //	Ornate Knight's Sword
		48369,  //	Ornate Hand Axe
		48370,  //	Ornate Dagger
		48371,  //	Ornate Short Spear
		48372,  //	Ornate Mace
		48373,  //	Ornate Hammer
		48374,  //	Ornate Reinforced Baton
		48375,  //	Ornate Ulak
		48376,  //	Ornate Great Sword
		48377,  //	Ornate Great Axe
		48378,  //	Ornate Long Spear
		48379,  //	Ornate Reinforced Staff
		48380,  //	Ornate Long Hammer
		48381,  //	Intricate Long Sword
		48382,  //	Intricate Knight's Sword
		48383,  //	Intricate Hand Axe
		48384,  //	Intricate Dagger
		48385,  //	Intricate Short Spear
		48386,  //	Intricate Mace
		48387,  //	Intricate Hammer
		48388,  //	Intricate Reinforced Baton
		48389,  //	Intricate Ulak
		48390,  //	Intricate Great Sword
		48391,  //	Intricate Great Axe
		48392,  //	Intricate Long Spear
		48393,  //	Intricate Reinforced Staff
		48394,  //	Intricate Long Hammer
		48395,  //	Elaborate Long Sword
		48396,  //	Elaborate Knight's Sword
		48397,  //	Elaborate Hand Axe
		48398,  //	Elaborate Dagger
		48399,  //	Elaborate Short Spear
		48400,  //	Elaborate Mace
		48401,  //	Elaborate Hammer
		48402,  //	Elaborate Reinforced Baton
		48403,  //	Elaborate Ulak
		48404,  //	Elaborate Great Sword
		48405,  //	Elaborate Great Axe
		48406,  //	Elaborate Long Spear
		48407,  //	Elaborate Reinforced Staff
		48408,  //	Elaborate Long Hammer
		48409,  //	Inured Long Sword
		48410,  //	Inured Knight's Sword
		48411,  //	Inured Hand Axe
		48412,  //	Inured Dagger
		48413,  //	Inured Short Spear
		48414,  //	Inured Mace
		48415,  //	Inured Hammer
		48416,  //	Inured Reinforced Baton
		48417,  //	Inured Ulak
		48418,  //	Inured Great Sword
		48419,  //	Inured Great Axe
		48420,  //	Inured Long Spear
		48421,  //	Inured Reinforced Staff
		48422,  //	Inured Long Hammer
		48423,  //	Elegant Long Sword
		48424,  //	Elegant Knight's Sword
		48425,  //	Elegant Hand Axe
		48426,  //	Elegant Dagger
		48427,  //	Elegant Short Spear
		48428,  //	Elegant Mace
		48429,  //	Elegant Hammer
		48430,  //	Elegant Reinforced Baton
		48431,  //	Elegant Ulak
		48432,  //	Elegant Great Sword
		48433,  //	Elegant Great Axe
		48434,  //	Elegant Long Spear
		48435,  //	Elegant Reinforced Staff
		48436,  //	Elegant Long Hammer
		48437,  //	Stalwart Long Sword
		48438,  //	Stalwart Knight's Sword
		48439,  //	Stalwart Hand Axe
		48440,  //	Stalwart Dagger
		48441,  //	Stalwart Short Spear
		48442,  //	Stalwart Mace
		48443,  //	Stalwart Hammer
		48444,  //	Stalwart Reinforced Baton
		48445,  //	Stalwart Ulak
		48446,  //	Stalwart Great Sword
		48447,  //	Stalwart Great Axe
		48448,  //	Stalwart Long Spear
		48449,  //	Stalwart Reinforced Staff
		48450,  //	Stalwart Long Hammer
		48451,  //	Blessed Nimble Symbol of the Skeptic
		48452,  //	Blessed Penetrating Symbol of the Skeptic
		48453,  //	Blessed Vicious Symbol of the Skeptic
		48454,  //	Blessed Nimble Symbol of Infestation
		48455,  //	Blessed Penetrating Symbol of Infestation
		48456,  //	Blessed Vicious Symbol of Infestation
		48457,  //	Blessed Nimble Symbol of Below
		48458,  //	Blessed Penetrating Symbol of Below
		48459,  //	Blessed Vicious Symbol of Below
		48460,  //	Blessed Nimble Symbol of the Farceur
		48461,  //	Blessed Penetrating Symbol of the Farceur
		48462,  //	Blessed Vicious Symbol of the Farceur
		48463,  //	Blessed Nimble Symbol of Terror
		48464,  //	Blessed Penetrating Symbol of Terror
		48465,  //	Blessed Vicious Symbol of Terror
		48466,  //	Blessed Nimble Symbol of Thaumaturgy
		48467,  //	Blessed Penetrating Symbol of Thaumaturgy
		48468,  //	Blessed Vicious Symbol of Thaumaturgy
		48469,  //	Blessed Nimble Symbol of Devotion
		48470,  //	Blessed Penetrating Symbol of Devotion
		48471,  //	Blessed Vicious Symbol of Devotion
		48472,  //	Blessed Nimble Symbol of the Fire Lord
		48473,  //	Blessed Penetrating Symbol of the Fire Lord
		48474,  //	Blessed Vicious Symbol of the Fire Lord
		48475,  //	Blessed Nimble Symbol of Hate
		48476,  //	Blessed Penetrating Symbol of Hate
		48477,  //	Blessed Vicious Symbol of Hate
		48478,  //	Blessed Nimble Symbol of Storms
		48479,  //	Blessed Penetrating Symbol of Storms
		48480,  //	Blessed Vicious Symbol of Storms
		48481,  //	Blessed Nimble Symbol of Integrity
		48482,  //	Blessed Penetrating Symbol of Integrity
		48483,  //	Blessed Vicious Symbol of Integrity
		48484,  //	Blessed Nimble Symbol of the Depths
		48485,  //	Blessed Penetrating Symbol of the Depths
		48486,  //	Blessed Vicious Symbol of the Depths
		48487,  //	Blessed Nimble Symbol of Stillness
		48488,  //	Blessed Penetrating Symbol of Stillness
		48489,  //	Blessed Vicious Symbol of Stillness
		48490,  //	Blessed Nimble Symbol of the Warmonger
		48491,  //	Blessed Penetrating Symbol of the Warmonger
		48492,  //	Blessed Vicious Symbol of the Warmonger
		48493,  //	Blessed Nimble Symbol of Compassion
		48494,  //	Blessed Penetrating Symbol of Compassion
		48495,  //	Blessed Vicious Symbol of Compassion
		48496,  //	Blessed Nimble Symbol of Malignancy
		48497,  //	Blessed Penetrating Symbol of Malignancy
		48498,  //	Blessed Vicious Symbol of Malignancy
		48499,  //	Blessed Nimble Symbol of Wildfire
		48500,  //	Blessed Penetrating Symbol of Wildfire
		48501,  //	Blessed Vicious Symbol of Wildfire
		48502,  //	Blessed Nimble Symbol of Decay
		48503,  //	Blessed Penetrating Symbol of Decay
		48504,  //	Blessed Vicious Symbol of Decay
		48505,  //	Blessed Nimble Symbol of Judgement
		48506,  //	Blessed Penetrating Symbol of Judgement
		48507,  //	Blessed Vicious Symbol of Judgement
		48508,  //	Blessed Nimble Symbol of Growth
		48509,  //	Blessed Penetrating Symbol of Growth
		48510,  //	Blessed Vicious Symbol of Growth
		48511,  //	Blessed Nimble Symbol of the Brood
		48512,  //	Blessed Penetrating Symbol of the Brood
		48513,  //	Blessed Vicious Symbol of the Brood
		48514,  //	Revered Nimble Symbol of the Skeptic
		48515,  //	Revered Penetrating Symbol of the Skeptic
		48516,  //	Revered Vicious Symbol of the Skeptic
		48517,  //	Revered Nimble Symbol of Infestation
		48518,  //	Revered Penetrating Symbol of Infestation
		48519,  //	Revered Vicious Symbol of Infestation
		48520,  //	Revered Nimble Symbol of Below
		48521,  //	Revered Penetrating Symbol of Below
		48522,  //	Revered Vicious Symbol of Below
		48523,  //	Revered Nimble Symbol of the Farceur
		48524,  //	Revered Penetrating Symbol of the Farceur
		48525,  //	Revered Vicious Symbol of the Farceur
		48526,  //	Revered Nimble Symbol of Terror
		48527,  //	Revered Penetrating Symbol of Terror
		48528,  //	Revered Vicious Symbol of Terror
		48529,  //	Revered Nimble Symbol of Thaumaturgy
		48530,  //	Revered Penetrating Symbol of Thaumaturgy
		48531,  //	Revered Vicious Symbol of Thaumaturgy
		48532,  //	Revered Nimble Symbol of Devotion
		48533,  //	Revered Penetrating Symbol of Devotion
		48534,  //	Revered Vicious Symbol of Devotion
		48535,  //	Revered Nimble Symbol of the Fire Lord
		48536,  //	Revered Penetrating Symbol of the Fire Lord
		48537,  //	Revered Vicious Symbol of the Fire Lord
		48538,  //	Revered Nimble Symbol of Hate
		48539,  //	Revered Penetrating Symbol of Hate
		48540,  //	Revered Vicious Symbol of Hate
		48541,  //	Revered Nimble Symbol of Storms
		48542,  //	Revered Penetrating Symbol of Storms
		48543,  //	Revered Vicious Symbol of Storms
		48544,  //	Revered Nimble Symbol of Integrity
		48545,  //	Revered Penetrating Symbol of Integrity
		48546,  //	Revered Vicious Symbol of Integrity
		48547,  //	Revered Nimble Symbol of the Depths
		48548,  //	Revered Penetrating Symbol of the Depths
		48549,  //	Revered Vicious Symbol of the Depths
		48550,  //	Revered Nimble Symbol of Stillness
		48551,  //	Revered Penetrating Symbol of Stillness
		48552,  //	Revered Vicious Symbol of Stillness
		48553,  //	Revered Nimble Symbol of the Warmonger
		48554,  //	Revered Penetrating Symbol of the Warmonger
		48555,  //	Revered Vicious Symbol of the Warmonger
		48556,  //	Revered Nimble Symbol of Compassion
		48557,  //	Revered Penetrating Symbol of Compassion
		48558,  //	Revered Vicious Symbol of Compassion
		48559,  //	Revered Nimble Symbol of Malignancy
		48560,  //	Revered Penetrating Symbol of Malignancy
		48561,  //	Revered Vicious Symbol of Malignancy
		48562,  //	Revered Nimble Symbol of Wildfire
		48563,  //	Revered Penetrating Symbol of Wildfire
		48564,  //	Revered Vicious Symbol of Wildfire
		48565,  //	Revered Nimble Symbol of Decay
		48566,  //	Revered Penetrating Symbol of Decay
		48567,  //	Revered Vicious Symbol of Decay
		48568,  //	Revered Nimble Symbol of Judgement
		48569,  //	Revered Penetrating Symbol of Judgement
		48570,  //	Revered Vicious Symbol of Judgement
		48571,  //	Revered Nimble Symbol of Growth
		48572,  //	Revered Penetrating Symbol of Growth
		48573,  //	Revered Vicious Symbol of Growth
		48574,  //	Revered Nimble Symbol of the Brood
		48575,  //	Revered Penetrating Symbol of the Brood
		48576,  //	Revered Vicious Symbol of the Brood
		48577,  //	Sacred Nimble Symbol of the Skeptic
		48578,  //	Sacred Penetrating Symbol of the Skeptic
		48579,  //	Sacred Vicious Symbol of the Skeptic
		48580,  //	Sacred Nimble Symbol of Infestation
		48581,  //	Sacred Penetrating Symbol of Infestation
		48582,  //	Sacred Vicious Symbol of Infestation
		48583,  //	Sacred Nimble Symbol of Below
		48584,  //	Sacred Penetrating Symbol of Below
		48585,  //	Sacred Vicious Symbol of Below
		48586,  //	Sacred Nimble Symbol of the Farceur
		48587,  //	Sacred Penetrating Symbol of the Farceur
		48588,  //	Sacred Vicious Symbol of the Farceur
		48589,  //	Sacred Nimble Symbol of Terror
		48590,  //	Sacred Penetrating Symbol of Terror
		48591,  //	Sacred Vicious Symbol of Terror
		48592,  //	Sacred Nimble Symbol of Thaumaturgy
		48593,  //	Sacred Penetrating Symbol of Thaumaturgy
		48594,  //	Sacred Vicious Symbol of Thaumaturgy
		48595,  //	Sacred Nimble Symbol of Devotion
		48596,  //	Sacred Penetrating Symbol of Devotion
		48597,  //	Sacred Vicious Symbol of Devotion
		48598,  //	Sacred Nimble Symbol of the Fire Lord
		48599,  //	Sacred Penetrating Symbol of the Fire Lord
		48600,  //	Sacred Vicious Symbol of the Fire Lord
		48601,  //	Sacred Nimble Symbol of Hate
		48602,  //	Sacred Penetrating Symbol of Hate
		48603,  //	Sacred Vicious Symbol of Hate
		48604,  //	Sacred Nimble Symbol of Storms
		48605,  //	Sacred Penetrating Symbol of Storms
		48606,  //	Sacred Vicious Symbol of Storms
		48607,  //	Sacred Nimble Symbol of Integrity
		48608,  //	Sacred Penetrating Symbol of Integrity
		48609,  //	Sacred Vicious Symbol of Integrity
		48610,  //	Sacred Nimble Symbol of the Depths
		48611,  //	Sacred Penetrating Symbol of the Depths
		48612,  //	Sacred Vicious Symbol of the Depths
		48613,  //	Sacred Nimble Symbol of Stillness
		48614,  //	Sacred Penetrating Symbol of Stillness
		48615,  //	Sacred Vicious Symbol of Stillness
		48616,  //	Sacred Nimble Symbol of the Warmonger
		48617,  //	Sacred Penetrating Symbol of the Warmonger
		48618,  //	Sacred Vicious Symbol of the Warmonger
		48619,  //	Sacred Nimble Symbol of Compassion
		48620,  //	Sacred Penetrating Symbol of Compassion
		48621,  //	Sacred Vicious Symbol of Compassion
		48622,  //	Sacred Nimble Symbol of Malignancy
		48623,  //	Sacred Penetrating Symbol of Malignancy
		48624,  //	Sacred Vicious Symbol of Malignancy
		48625,  //	Sacred Nimble Symbol of Wildfire
		48626,  //	Sacred Penetrating Symbol of Wildfire
		48627,  //	Sacred Vicious Symbol of Wildfire
		48628,  //	Sacred Nimble Symbol of Decay
		48629,  //	Sacred Penetrating Symbol of Decay
		48630,  //	Sacred Vicious Symbol of Decay
		48631,  //	Sacred Nimble Symbol of Judgement
		48632,  //	Sacred Penetrating Symbol of Judgement
		48633,  //	Sacred Vicious Symbol of Judgement
		48634,  //	Sacred Nimble Symbol of Growth
		48635,  //	Sacred Penetrating Symbol of Growth
		48636,  //	Sacred Vicious Symbol of Growth
		48637,  //	Sacred Nimble Symbol of the Brood
		48638,  //	Sacred Penetrating Symbol of the Brood
		48639,  //	Sacred Vicious Symbol of the Brood
		48640,  //	Eminent Nimble Symbol of the Skeptic
		48641,  //	Eminent Penetrating Symbol of the Skeptic
		48642,  //	Eminent Vicious Symbol of the Skeptic
		48643,  //	Eminent Nimble Symbol of Infestation
		48644,  //	Eminent Penetrating Symbol of Infestation
		48645,  //	Eminent Vicious Symbol of Infestation
		48646,  //	Eminent Nimble Symbol of Below
		48647,  //	Eminent Penetrating Symbol of Below
		48648,  //	Eminent Vicious Symbol of Below
		48649,  //	Eminent Nimble Symbol of the Farceur
		48650,  //	Eminent Penetrating Symbol of the Farceur
		48651,  //	Eminent Vicious Symbol of the Farceur
		48652,  //	Eminent Nimble Symbol of Terror
		48653,  //	Eminent Penetrating Symbol of Terror
		48654,  //	Eminent Vicious Symbol of Terror
		48655,  //	Eminent Nimble Symbol of Thaumaturgy
		48656,  //	Eminent Penetrating Symbol of Thaumaturgy
		48657,  //	Eminent Vicious Symbol of Thaumaturgy
		48658,  //	Eminent Nimble Symbol of Devotion
		48659,  //	Eminent Penetrating Symbol of Devotion
		48660,  //	Eminent Vicious Symbol of Devotion
		48661,  //	Eminent Nimble Symbol of the Fire Lord
		48662,  //	Eminent Penetrating Symbol of the Fire Lord
		48663,  //	Eminent Vicious Symbol of the Fire Lord
		48664,  //	Eminent Nimble Symbol of Hate
		48665,  //	Eminent Penetrating Symbol of Hate
		48666,  //	Eminent Vicious Symbol of Hate
		48667,  //	Eminent Nimble Symbol of Storms
		48668,  //	Eminent Penetrating Symbol of Storms
		48669,  //	Eminent Vicious Symbol of Storms
		48670,  //	Eminent Nimble Symbol of Integrity
		48671,  //	Eminent Penetrating Symbol of Integrity
		48672,  //	Eminent Vicious Symbol of Integrity
		48673,  //	Eminent Nimble Symbol of the Depths
		48674,  //	Eminent Penetrating Symbol of the Depths
		48675,  //	Eminent Vicious Symbol of the Depths
		48676,  //	Eminent Nimble Symbol of Stillness
		48677,  //	Eminent Penetrating Symbol of Stillness
		48678,  //	Eminent Vicious Symbol of Stillness
		48679,  //	Eminent Nimble Symbol of the Warmonger
		48680,  //	Eminent Penetrating Symbol of the Warmonger
		48681,  //	Eminent Vicious Symbol of the Warmonger
		48682,  //	Eminent Nimble Symbol of Compassion
		48683,  //	Eminent Penetrating Symbol of Compassion
		48684,  //	Eminent Vicious Symbol of Compassion
		48685,  //	Eminent Nimble Symbol of Malignancy
		48686,  //	Eminent Penetrating Symbol of Malignancy
		48687,  //	Eminent Vicious Symbol of Malignancy
		48688,  //	Eminent Nimble Symbol of Wildfire
		48689,  //	Eminent Penetrating Symbol of Wildfire
		48690,  //	Eminent Vicious Symbol of Wildfire
		48691,  //	Eminent Nimble Symbol of Decay
		48692,  //	Eminent Penetrating Symbol of Decay
		48693,  //	Eminent Vicious Symbol of Decay
		48694,  //	Eminent Nimble Symbol of Judgement
		48695,  //	Eminent Penetrating Symbol of Judgement
		48696,  //	Eminent Vicious Symbol of Judgement
		48697,  //	Eminent Nimble Symbol of Growth
		48698,  //	Eminent Penetrating Symbol of Growth
		48699,  //	Eminent Vicious Symbol of Growth
		48700,  //	Eminent Nimble Symbol of the Brood
		48701,  //	Eminent Penetrating Symbol of the Brood
		48702,  //	Eminent Vicious Symbol of the Brood
		48703,  //	Exalted Nimble Symbol of the Skeptic
		48704,  //	Exalted Penetrating Symbol of the Skeptic
		48705,  //	Exalted Vicious Symbol of the Skeptic
		48706,  //	Exalted Nimble Symbol of Infestation
		48707,  //	Exalted Penetrating Symbol of Infestation
		48708,  //	Exalted Vicious Symbol of Infestation
		48709,  //	Exalted Nimble Symbol of Below
		48710,  //	Exalted Penetrating Symbol of Below
		48711,  //	Exalted Vicious Symbol of Below
		48712,  //	Exalted Nimble Symbol of the Farceur
		48713,  //	Exalted Penetrating Symbol of the Farceur
		48714,  //	Exalted Vicious Symbol of the Farceur
		48715,  //	Exalted Nimble Symbol of Terror
		48716,  //	Exalted Penetrating Symbol of Terror
		48717,  //	Exalted Vicious Symbol of Terror
		48718,  //	Exalted Nimble Symbol of Thaumaturgy
		48719,  //	Exalted Penetrating Symbol of Thaumaturgy
		48720,  //	Exalted Vicious Symbol of Thaumaturgy
		48721,  //	Exalted Nimble Symbol of Devotion
		48722,  //	Exalted Penetrating Symbol of Devotion
		48723,  //	Exalted Vicious Symbol of Devotion
		48724,  //	Exalted Nimble Symbol of the Fire Lord
		48725,  //	Exalted Penetrating Symbol of the Fire Lord
		48726,  //	Exalted Vicious Symbol of the Fire Lord
		48727,  //	Exalted Nimble Symbol of Hate
		48728,  //	Exalted Penetrating Symbol of Hate
		48729,  //	Exalted Vicious Symbol of Hate
		48730,  //	Exalted Nimble Symbol of Storms
		48731,  //	Exalted Penetrating Symbol of Storms
		48732,  //	Exalted Vicious Symbol of Storms
		48733,  //	Exalted Nimble Symbol of Integrity
		48734,  //	Exalted Penetrating Symbol of Integrity
		48735,  //	Exalted Vicious Symbol of Integrity
		48736,  //	Exalted Nimble Symbol of the Depths
		48737,  //	Exalted Penetrating Symbol of the Depths
		48738,  //	Exalted Vicious Symbol of the Depths
		48739,  //	Exalted Nimble Symbol of Stillness
		48740,  //	Exalted Penetrating Symbol of Stillness
		48741,  //	Exalted Vicious Symbol of Stillness
		48742,  //	Exalted Nimble Symbol of the Warmonger
		48743,  //	Exalted Penetrating Symbol of the Warmonger
		48744,  //	Exalted Vicious Symbol of the Warmonger
		48745,  //	Exalted Nimble Symbol of Compassion
		48746,  //	Exalted Penetrating Symbol of Compassion
		48747,  //	Exalted Vicious Symbol of Compassion
		48748,  //	Exalted Nimble Symbol of Malignancy
		48749,  //	Exalted Penetrating Symbol of Malignancy
		48750,  //	Exalted Vicious Symbol of Malignancy
		48751,  //	Exalted Nimble Symbol of Wildfire
		48752,  //	Exalted Penetrating Symbol of Wildfire
		48753,  //	Exalted Vicious Symbol of Wildfire
		48754,  //	Exalted Nimble Symbol of Decay
		48755,  //	Exalted Penetrating Symbol of Decay
		48756,  //	Exalted Vicious Symbol of Decay
		48757,  //	Exalted Nimble Symbol of Judgement
		48758,  //	Exalted Penetrating Symbol of Judgement
		48759,  //	Exalted Vicious Symbol of Judgement
		48760,  //	Exalted Nimble Symbol of Growth
		48761,  //	Exalted Penetrating Symbol of Growth
		48762,  //	Exalted Vicious Symbol of Growth
		48763,  //	Exalted Nimble Symbol of the Brood
		48764,  //	Exalted Penetrating Symbol of the Brood
		48765,  //	Exalted Vicious Symbol of the Brood
		48766,  //	Sublime Nimble Symbol of the Skeptic
		48767,  //	Sublime Penetrating Symbol of the Skeptic
		48768,  //	Sublime Vicious Symbol of the Skeptic
		48769,  //	Sublime Nimble Symbol of Infestation
		48770,  //	Sublime Penetrating Symbol of Infestation
		48771,  //	Sublime Vicious Symbol of Infestation
		48772,  //	Sublime Nimble Symbol of Below
		48773,  //	Sublime Penetrating Symbol of Below
		48774,  //	Sublime Vicious Symbol of Below
		48775,  //	Sublime Nimble Symbol of the Farceur
		48776,  //	Sublime Penetrating Symbol of the Farceur
		48777,  //	Sublime Vicious Symbol of the Farceur
		48778,  //	Sublime Nimble Symbol of Terror
		48779,  //	Sublime Penetrating Symbol of Terror
		48780,  //	Sublime Vicious Symbol of Terror
		48781,  //	Sublime Nimble Symbol of Thaumaturgy
		48782,  //	Sublime Penetrating Symbol of Thaumaturgy
		48783,  //	Sublime Vicious Symbol of Thaumaturgy
		48784,  //	Sublime Nimble Symbol of Devotion
		48785,  //	Sublime Penetrating Symbol of Devotion
		48786,  //	Sublime Vicious Symbol of Devotion
		48787,  //	Sublime Nimble Symbol of the Fire Lord
		48788,  //	Sublime Penetrating Symbol of the Fire Lord
		48789,  //	Sublime Vicious Symbol of the Fire Lord
		48790,  //	Sublime Nimble Symbol of Hate
		48791,  //	Sublime Penetrating Symbol of Hate
		48792,  //	Sublime Vicious Symbol of Hate
		48793,  //	Sublime Nimble Symbol of Storms
		48794,  //	Sublime Penetrating Symbol of Storms
		48795,  //	Sublime Vicious Symbol of Storms
		48796,  //	Sublime Nimble Symbol of Integrity
		48797,  //	Sublime Penetrating Symbol of Integrity
		48798,  //	Sublime Vicious Symbol of Integrity
		48799,  //	Sublime Nimble Symbol of the Depths
		48800,  //	Sublime Penetrating Symbol of the Depths
		48801,  //	Sublime Vicious Symbol of the Depths
		48802,  //	Sublime Nimble Symbol of Stillness
		48803,  //	Sublime Penetrating Symbol of Stillness
		48804,  //	Sublime Vicious Symbol of Stillness
		48805,  //	Sublime Nimble Symbol of the Warmonger
		48806,  //	Sublime Penetrating Symbol of the Warmonger
		48807,  //	Sublime Vicious Symbol of the Warmonger
		48808,  //	Sublime Nimble Symbol of Compassion
		48809,  //	Sublime Penetrating Symbol of Compassion
		48810,  //	Sublime Vicious Symbol of Compassion
		48811,  //	Sublime Nimble Symbol of Malignancy
		48812,  //	Sublime Penetrating Symbol of Malignancy
		48813,  //	Sublime Vicious Symbol of Malignancy
		48814,  //	Sublime Nimble Symbol of Wildfire
		48815,  //	Sublime Penetrating Symbol of Wildfire
		48816,  //	Sublime Vicious Symbol of Wildfire
		48817,  //	Sublime Nimble Symbol of Decay
		48818,  //	Sublime Penetrating Symbol of Decay
		48819,  //	Sublime Vicious Symbol of Decay
		48820,  //	Sublime Nimble Symbol of Judgement
		48821,  //	Sublime Penetrating Symbol of Judgement
		48822,  //	Sublime Vicious Symbol of Judgement
		48823,  //	Sublime Nimble Symbol of Growth
		48824,  //	Sublime Penetrating Symbol of Growth
		48825,  //	Sublime Vicious Symbol of Growth
		48826,  //	Sublime Nimble Symbol of the Brood
		48827,  //	Sublime Penetrating Symbol of the Brood
		48828,  //	Sublime Vicious Symbol of the Brood
		48829,  //	Venerable Nimble Symbol of the Skeptic
		48830,  //	Venerable Penetrating Symbol of the Skeptic
		48831,  //	Venerable Vicious Symbol of the Skeptic
		48832,  //	Venerable Nimble Symbol of Infestation
		48833,  //	Venerable Penetrating Symbol of Infestation
		48834,  //	Venerable Vicious Symbol of Infestation
		48835,  //	Venerable Nimble Symbol of Below
		48836,  //	Venerable Penetrating Symbol of Below
		48837,  //	Venerable Vicious Symbol of Below
		48838,  //	Venerable Nimble Symbol of the Farceur
		48839,  //	Venerable Penetrating Symbol of the Farceur
		48840,  //	Venerable Vicious Symbol of the Farceur
		48841,  //	Venerable Nimble Symbol of Terror
		48842,  //	Venerable Penetrating Symbol of Terror
		48843,  //	Venerable Vicious Symbol of Terror
		48844,  //	Venerable Nimble Symbol of Thaumaturgy
		48845,  //	Venerable Penetrating Symbol of Thaumaturgy
		48846,  //	Venerable Vicious Symbol of Thaumaturgy
		48847,  //	Venerable Nimble Symbol of Devotion
		48848,  //	Venerable Penetrating Symbol of Devotion
		48849,  //	Venerable Vicious Symbol of Devotion
		48850,  //	Venerable Nimble Symbol of the Fire Lord
		48851,  //	Venerable Penetrating Symbol of the Fire Lord
		48852,  //	Venerable Vicious Symbol of the Fire Lord
		48853,  //	Venerable Nimble Symbol of Hate
		48854,  //	Venerable Penetrating Symbol of Hate
		48855,  //	Venerable Vicious Symbol of Hate
		48856,  //	Venerable Nimble Symbol of Storms
		48857,  //	Venerable Penetrating Symbol of Storms
		48858,  //	Venerable Vicious Symbol of Storms
		48859,  //	Venerable Nimble Symbol of Integrity
		48860,  //	Venerable Penetrating Symbol of Integrity
		48861,  //	Venerable Vicious Symbol of Integrity
		48862,  //	Venerable Nimble Symbol of the Depths
		48863,  //	Venerable Penetrating Symbol of the Depths
		48864,  //	Venerable Vicious Symbol of the Depths
		48865,  //	Venerable Nimble Symbol of Stillness
		48866,  //	Venerable Penetrating Symbol of Stillness
		48867,  //	Venerable Vicious Symbol of Stillness
		48868,  //	Venerable Nimble Symbol of the Warmonger
		48869,  //	Venerable Penetrating Symbol of the Warmonger
		48870,  //	Venerable Vicious Symbol of the Warmonger
		48871,  //	Venerable Nimble Symbol of Compassion
		48872,  //	Venerable Penetrating Symbol of Compassion
		48873,  //	Venerable Vicious Symbol of Compassion
		48874,  //	Venerable Nimble Symbol of Malignancy
		48875,  //	Venerable Penetrating Symbol of Malignancy
		48876,  //	Venerable Vicious Symbol of Malignancy
		48877,  //	Venerable Nimble Symbol of Wildfire
		48878,  //	Venerable Penetrating Symbol of Wildfire
		48879,  //	Venerable Vicious Symbol of Wildfire
		48880,  //	Venerable Nimble Symbol of Decay
		48881,  //	Venerable Penetrating Symbol of Decay
		48882,  //	Venerable Vicious Symbol of Decay
		48883,  //	Venerable Nimble Symbol of Judgement
		48884,  //	Venerable Penetrating Symbol of Judgement
		48885,  //	Venerable Vicious Symbol of Judgement
		48886,  //	Venerable Nimble Symbol of Growth
		48887,  //	Venerable Penetrating Symbol of Growth
		48888,  //	Venerable Vicious Symbol of Growth
		48889,  //	Venerable Nimble Symbol of the Brood
		48890,  //	Venerable Penetrating Symbol of the Brood
		32969,  //cracked leather belt
		62243,  //Exquisite Silk Turban
	}
	return
}
