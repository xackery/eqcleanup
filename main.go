package main

import (
	//"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/aug"
	"github.com/xackery/eqcleanup/cazicthule"
	"github.com/xackery/eqcleanup/characterwipe"
	"github.com/xackery/eqcleanup/defiant"
	"github.com/xackery/eqcleanup/emptymerchant"
	"github.com/xackery/eqcleanup/era"
	"github.com/xackery/eqcleanup/fabled"
	"github.com/xackery/eqcleanup/ldon"
	"github.com/xackery/eqcleanup/lootprice"
	"github.com/xackery/eqcleanup/namedreduce"
	"github.com/xackery/eqcleanup/nexus"
	"github.com/xackery/eqcleanup/peqtweak"
	"github.com/xackery/eqcleanup/priests"
	"github.com/xackery/eqcleanup/rainsnow"
	"github.com/xackery/eqcleanup/rodent"
	"github.com/xackery/eqcleanup/shadowrest"
	"github.com/xackery/eqcleanup/soulbinder"
	"github.com/xackery/eqcleanup/spells"
	"github.com/xackery/eqcleanup/tribute"
	"github.com/xackery/eqcleanup/trickortreat"
	"github.com/xackery/eqemuconfig"
	"os"
	"strings"
)

func main() {

	for {
		showMenu()
	}

}

func connectDB(config *eqemuconfig.Config) (db *sqlx.DB, err error) {
	//Connect to DB
	db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Db))
	if err != nil {
		fmt.Println("Error connecting to DB:", err.Error())
		os.Exit(1)
	}
	return
}

var isDBConnected bool
var isConfigLoaded bool

func showMenu() {
	var err error
	var option string
	fmt.Println("===EQ Cleanup===")
	fmt.Println("Remember: Always back up before using this tool!")
	fmt.Println("Choose an option:")
	config := menuConfig()
	db := menuDB(config)
	defer db.Close()
	if isDBConnected && isConfigLoaded {
		fmt.Println("3) Delete Soulbinders - Remove Soulbinder NPCs from classic zones.")
		fmt.Println("4) Delete Rodents - Remove rats from starting zones. May not get all")
		fmt.Println("5) Delete Trick or Treat Quests - Halloween quests were added in latest PEQ, this strips them out.")
		fmt.Println("6) Delete Tribute Masters - Removes the tribue NPCs")
		fmt.Println("7) Delete All spells, tomes (scribable items) - This is experimental")
		fmt.Println("8) Delete Priests of Discord - Self explanatory")
		fmt.Println("9) Delete Defiant Armor - Removes specifically defiant armor (better to use out of era)")
		fmt.Println("10) Delete LDON Npcs - Delete the Adventuerer merchants and other NPCs for LDON")
		fmt.Println("11) Delete Nexus portal NPCs - Removes the Luclin Nexus representatives across norrath (NK, DL, etc)")
		fmt.Println("12) Disable Rain and Snow - Disables Rain and Snow from ever happening on any zone.")
		fmt.Println("13) Delete Emissary of Shadowrest NPCs - Cleans out the Shadowrest NPCs in East Commonlands")
		fmt.Println("14) Delete Empty Merchant NPCs - If you remove spells, this will remove any empty merchants.")
		fmt.Println("15) Remove Fabled - Removes fabled mobs from spawning")
		fmt.Println("16) PEQ Tweaks - Misc tweaks, like fixing skeleton movement speed 0 problems")
		fmt.Println("17) Named Spawn Rate Reduction - Adjusts spawn rate on some key named monsters")
		fmt.Println("18) Delete Augments - Delete all augments on server")
		fmt.Println("19) Character Wipe - Flushes all character/account data")
		fmt.Println("20) Era Cleanup - Strips out items post-velious")
		fmt.Println("21) Loot Price - Special price adjustment script (IGNORE)")
		fmt.Println("22) Migrate Cazic Thule - Migrates NPCs from New Cazic Thule to Old Cazic Thule")
	} else {
		fmt.Println("-Commands are disabled until database and config is good-")
	}
	fmt.Println("Q) Quit")

	fmt.Scan(&option)
	fmt.Println("You chose option:", option)
	option = strings.ToLower(option)
	if option == "q" || option == "exit" {
		fmt.Println("Exiting")
		os.Exit(0)
	}

	if !isDBConnected || !isConfigLoaded {
		return
	}
	if option == "3" { //Clean up Soulbinders
		err = soulbinder.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing soulbinders:", err.Error())
		}
	} else if option == "4" {
		err = rodent.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing rodents:", err.Error())
		}
	} else if option == "5" {
		err = trickortreat.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing trick or treat:", err.Error())
		}
	} else if option == "6" {
		err = tribute.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing tribute:", err.Error())
		}
	} else if option == "7" {
		err = spells.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing spells:", err.Error())
		}
	} else if option == "8" {
		err = priests.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing priest of discords:", err.Error())
		}
	} else if option == "9" {
		err = defiant.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing defiant:", err.Error())
		}
	} else if option == "10" {
		err = ldon.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing ldon:", err.Error())
		}
	} else if option == "11" {
		err = nexus.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing nexus:", err.Error())
		}
	} else if option == "12" {
		err = rainsnow.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing rain and snow:", err.Error())
		}
	} else if option == "13" {
		err = shadowrest.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing shadowrest:", err.Error())
		}
	} else if option == "14" {
		err = emptymerchant.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing empty merchants:", err.Error())
		}
	} else if option == "15" {
		err = fabled.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing fabled:", err.Error())
		}
	} else if option == "16" {
		err = peqtweak.Clean(db, config)
		if err != nil {
			fmt.Println("Error appling peq tweaks:", err.Error())
		}
	} else if option == "17" {
		err = namedreduce.Clean(db, config)
		if err != nil {
			fmt.Println("Error adjusting nameds:", err.Error())
		}
	} else if option == "18" {
		err = aug.Clean(db, config)
		if err != nil {
			fmt.Println("Error removing aug:", err.Error())
		}
	} else if option == "19" {
		err = characterwipe.Clean(db, config)
		if err != nil {
			fmt.Println("Error wiping characters:", err.Error())
		}
	} else if option == "20" {
		err = era.Clean(db, config)
		if err != nil {
			fmt.Println("Error removed era:", err.Error())
		}
	} else if option == "21" {
		err = lootprice.Clean(db, config)
		if err != nil {
			fmt.Println("Error adjusting loot pricing:", err.Error())
		}
	} else if option == "22" {
		if err = cazicthule.Clean(db, config); err != nil {
			fmt.Println("Error migrating cazic:", err.Error())
		}
	}
	return
}

func menuConfig() (config *eqemuconfig.Config) {
	config, err := eqemuconfig.GetConfig()
	status := "Good"
	if err != nil {
		isConfigLoaded = false
		status = fmt.Sprintf("Bad (%s)", err.Error())
	} else {
		isConfigLoaded = true
		status = fmt.Sprintf("Good (%s)", config.World.Longname)
	}
	fmt.Printf("1) Reload eqemu_config.xml (Status: %s)\n", status)
	return
}

func menuDB(config *eqemuconfig.Config) (db *sqlx.DB) {
	var err error
	db, err = connectDB(config)
	status := "Good"
	if err != nil {
		isDBConnected = false
		status = fmt.Sprintf("Bad (%s)", err.Error())
	} else {
		isDBConnected = true
	}
	fmt.Printf("2) Test DB Connection Settings (Status: %s)\n", status)
	return
}
