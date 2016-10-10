package zone

import (
	"github.com/turret-io/go-menu/menu"
	"github.com/xackery/eqcleanup/zone/cazicthule"
	"github.com/xackery/eqcleanup/zone/paw"
)

func Menu(args ...string) (err error) {
	//var err error
	eraMenu := []menu.CommandOption{
		menu.CommandOption{"cazicthule", "Migrate Cazic Thule", cazicthule.Clean},
		menu.CommandOption{"paw", "Migrate Splitpaw", paw.Clean},
		menu.CommandOption{"quit", "Return to main menu", nil},
	}
	menuOptions := menu.NewMenuOptions("'menu' for help [zone]> ", 0)

	menu := menu.NewMenu(eraMenu, menuOptions)
	menu.Start()
	return
}
