package main

import (
	"fmt"
	"github.com/turret-io/go-menu/menu"
	"github.com/xackery/eqcleanup/core"
	"github.com/xackery/eqcleanup/era"
	"github.com/xackery/eqcleanup/misc"
	"github.com/xackery/eqcleanup/tools/eqdb"
	"github.com/xackery/eqcleanup/zone"
)

var isDBConnected bool
var isConfigLoaded bool

func main() {
	var err error
	_, err = eqdb.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//var err error
	mainMenu := []menu.CommandOption{
		menu.CommandOption{"era", "Out of era edits, ex. removing items or npcs", era.Menu},
		menu.CommandOption{"zone", "Zone migration tools, ex. newer to classic zone", zone.Menu},
		menu.CommandOption{"core", "System maintainence options, ex. removing all character data", core.Menu},
		menu.CommandOption{"misc", "misc edits, ex. disabling rainfall", misc.Menu},
	}
	menuOptions := menu.NewMenuOptions("'menu' for help []> ", 0)

	menu := menu.NewMenu(mainMenu, menuOptions)
	menu.Start()
}
