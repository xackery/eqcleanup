package misc

import (
	"github.com/turret-io/go-menu/menu"
	"github.com/xackery/eqcleanup/misc/merchant"
	"github.com/xackery/eqcleanup/misc/priest"
	"github.com/xackery/eqcleanup/misc/rain"
	"github.com/xackery/eqcleanup/misc/spell"
)

func Menu(args ...string) (err error) {
	//var err error
	eraMenu := []menu.CommandOption{
		menu.CommandOption{"spell", "Removes all spell scrolls from game", spell.Clean},
		menu.CommandOption{"rain", "Disables Rain & Snow on all zones", rain.Clean},
		menu.CommandOption{"priest", "Removes Priest of Discord", priest.Clean},
		menu.CommandOption{"merchant", "Removes any merchants with empty inventory", merchant.Clean},
		menu.CommandOption{"quit", "Return to main menu", nil},
	}
	menuOptions := menu.NewMenuOptions("'menu' for help [system]> ", 0)

	menu := menu.NewMenu(eraMenu, menuOptions)
	menu.Start()
	return
}
