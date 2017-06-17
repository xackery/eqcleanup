package core

import (
	"github.com/turret-io/go-menu/menu"
	"github.com/xackery/eqcleanup/core/wipe"
)

func Menu(args ...string) (err error) {
	//var err error
	eraMenu := []menu.CommandOption{
		{"wipe", "Wipes all Character Data", wipe.Clean},
		{"quit", "Return to main menu", nil},
	}
	menuOptions := menu.NewMenuOptions("'menu' for help [system]> ", 0)

	menu := menu.NewMenu(eraMenu, menuOptions)
	menu.Start()
	return
}
