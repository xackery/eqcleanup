package zone

import (
	"github.com/turret-io/go-menu/menu"
	"github.com/xackery/eqcleanup/zone/cazicthule"
	"github.com/xackery/eqcleanup/zone/commonlands"
	"github.com/xackery/eqcleanup/zone/droga"
	"github.com/xackery/eqcleanup/zone/nurga"
	"github.com/xackery/eqcleanup/zone/paw"
	"github.com/xackery/eqcleanup/zone/veeshan"
)

func Menu(args ...string) (err error) {
	//var err error
	eraMenu := []menu.CommandOption{
		{"cazicthule", "Migrate Cazic Thule", cazicthule.Clean},
		{"paw", "Migrate Splitpaw", paw.Clean},
		{"nurga", "Migrate Nurga", nurga.Clean},
		{"droga", "Migrate Droga", droga.Clean},
		{"veeshan", "Migrate Veeshan's Peak", veeshan.Clean},
		{"commonlands", "Migrate Commonlands", commonlands.Clean},
		{"quit", "Return to main menu", nil},
	}
	menuOptions := menu.NewMenuOptions("'menu' for help [zone]> ", 0)

	menu := menu.NewMenu(eraMenu, menuOptions)
	menu.Start()
	return
}
