package era

import (
	"github.com/turret-io/go-menu/menu"
	"github.com/xackery/eqcleanup/era/aug"
	"github.com/xackery/eqcleanup/era/eraitem"
	"github.com/xackery/eqcleanup/era/fabled"
	"github.com/xackery/eqcleanup/era/halloween"
	"github.com/xackery/eqcleanup/era/ldon"
	"github.com/xackery/eqcleanup/era/named"
	"github.com/xackery/eqcleanup/era/nexus"
	"github.com/xackery/eqcleanup/era/rodent"
	"github.com/xackery/eqcleanup/era/shadowrest"
	"github.com/xackery/eqcleanup/era/soulbinder"
	"github.com/xackery/eqcleanup/era/tribute"
)

func Menu(args ...string) (err error) {
	//var err error
	eraMenu := []menu.CommandOption{
		menu.CommandOption{"aug", "Remove all augments", aug.Clean},
		menu.CommandOption{"fabled", "Remove all fabled NPCs", fabled.Clean},
		menu.CommandOption{"halloween", "Remove halloween event data", halloween.Clean},
		menu.CommandOption{"item", "Remove all out of era items from Luclin onward", eraitem.Clean},
		menu.CommandOption{"ldon", "Remove LDoN NPCs", ldon.Clean},
		menu.CommandOption{"named", "Reduce the spawn rate of named mobs", named.Clean},
		menu.CommandOption{"nexus", "Remove nexus NPCs", nexus.Clean},
		menu.CommandOption{"rodent", "Remove rodents and exterminator quests", rodent.Clean},
		menu.CommandOption{"shadowrest", "Remove shadowrest NPCs from East Commons", shadowrest.Clean},
		menu.CommandOption{"soulbinder", "Remove soulbinder NPCs", soulbinder.Clean},
		menu.CommandOption{"tribute", "Remove tribute NPCs", tribute.Clean},
		menu.CommandOption{"quit", "Return to main menu", nil},
	}
	menuOptions := menu.NewMenuOptions("'menu' for help [era]> ", 0)

	menu := menu.NewMenu(eraMenu, menuOptions)
	menu.Start()
	return
}
