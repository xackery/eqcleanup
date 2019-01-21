//This script will disable emissary of shadowrest
package shadowrest

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/dixonwille/wmenu"
	"github.com/xackery/eqcleanup/client"
)

var focus = "nexus"

// Menu represents menu options
func Menu(opt wmenu.Opt) (err error) {
	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	ok = true
	fmt.Println("This will remove fabled NPCs")

	menu := wmenu.NewMenu("Continue operation?")
	menu.Option("Return to main menu", nil, false, func(opt wmenu.Opt) error {
		ok = false
		return nil
	})
	menu.Option("I understand the risks - continue", nil, false, func(opt wmenu.Opt) error {
		return Clean(c)
	})

	for ok {
		err = menu.Run()
	}
	return
}

// Clean removes shadowrest
func Clean(c *client.Client) (err error) {
	//SpawnGroups
	ids := []int64{}

	mids, err := c.Database.SpawnGroupIDsByLastNameWildcardSearch("Emissary of Shadowrest")
	if err != nil {
		err = errors.Wrapf(err, "failed to search for %s", focus)
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	fmt.Println("Got", len(ids), "Npcs")

	fmt.Println("Removing spell NPC spawn points")

	totalChanged, err := c.Database.SpawnGroupAndEntryByIDDelete(ids)
	if err != nil {
		err = errors.Wrapf(err, "failed to remove %s", focus)
		return
	}

	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in all player-accessible item locations.")
	return
}
