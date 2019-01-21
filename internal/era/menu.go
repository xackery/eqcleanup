package era

import (
	"fmt"

	"github.com/dixonwille/wmenu"
	"github.com/xackery/eqcleanup/client"
	"github.com/xackery/eqcleanup/internal/era/aug"
	"github.com/xackery/eqcleanup/internal/era/eraitem"
	"github.com/xackery/eqcleanup/internal/era/fabled"
	"github.com/xackery/eqcleanup/internal/era/halloween"
	"github.com/xackery/eqcleanup/internal/era/ldon"
	"github.com/xackery/eqcleanup/internal/era/named"
	"github.com/xackery/eqcleanup/internal/era/nexus"
	"github.com/xackery/eqcleanup/internal/era/rodent"
	"github.com/xackery/eqcleanup/internal/era/shadowrest"
	"github.com/xackery/eqcleanup/internal/era/soulbinder"
	"github.com/xackery/eqcleanup/internal/era/tribute"
	"github.com/xackery/eqcleanup/internal/zone"
)

// Menu represents the era menu options
func Menu(opt wmenu.Opt) (err error) {
	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	ok = true
	menu := wmenu.NewMenu("Choose an era option")
	menu.Option("Kunark - Do fixes to make server kunark-like", c, false, kunark)
	menu.Option("Velious - Do fixes to make server velious-like", c, false, velious)
	menu.Option("Luclin - Do fixes to make server luclin-like", c, false, luclin)
	menu.Option("Manual - Manually choose era fixes", c, false, ManualMenu)
	menu.Option("Return to main menu", nil, false, func(opt wmenu.Opt) error {
		ok = false
		return nil
	})
	for ok {
		err = menu.Run()
	}
	return
}

// ManualMenu represents the era manual menu options
func ManualMenu(opt wmenu.Opt) (err error) {
	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	ok = true
	menu := wmenu.NewMenu("Choose an era option")
	menu.Option("Augments - Remove all augments from game", c, false, aug.Menu)
	menu.Option("Fabled Spawns - Remove all fabled NPCs", c, false, fabled.Menu)
	menu.Option("Halloween - Remove halloween event data", c, false, halloween.Menu)
	menu.Option("Item - Remove all out of era items from Luclin onward", c, false, eraitem.Menu)
	menu.Option("LDON - Remove LDoN NPCs", c, false, ldon.Menu)
	menu.Option("Named - Reduce the spawn rate of named mobs", c, false, named.Menu)
	menu.Option("Nexus - Remove nexus NPCs", c, false, nexus.Menu)
	menu.Option("Rodent - Remove rodents and exterminator quests", c, false, rodent.Menu)
	menu.Option("Shadowrest - Remove shadowrest NPCs from East Commons", c, false, shadowrest.Menu)
	menu.Option("Soulbinder - Remove soulbinder NPCs", c, false, soulbinder.Menu)
	menu.Option("Tribute - Remove tribute NPCs", c, false, tribute.Menu)
	menu.Option("Zone - Change various zones", c, false, zone.Menu)
	menu.Option("Return to main menu", nil, false, func(opt wmenu.Opt) error {
		ok = false
		return nil
	})
	for ok {
		err = menu.Run()
	}
	return
}

func kunark(opt wmenu.Opt) (err error) {
	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	err = aug.Clean(c)
	if err != nil {
		return
	}
	err = fabled.Clean(c)
	if err != nil {
		return
	}
	err = halloween.Clean(c)
	if err != nil {
		return
	}
	err = eraitem.Clean(c)
	if err != nil {
		return
	}
	err = ldon.Clean(c)
	if err != nil {
		return
	}
	err = named.Clean(c)
	if err != nil {
		return
	}
	err = nexus.Clean(c)
	if err != nil {
		return
	}
	err = rodent.Clean(c)
	if err != nil {
		return
	}
	err = shadowrest.Clean(c)
	if err != nil {
		return
	}
	err = soulbinder.Clean(c)
	if err != nil {
		return
	}
	err = tribute.Clean(c)
	if err != nil {
		return
	}
	return
}

func velious(opt wmenu.Opt) (err error) {

	return
}

func luclin(opt wmenu.Opt) (err error) {

	return
}
