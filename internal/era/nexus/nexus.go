package nexus

import (
	"fmt"

	"github.com/dixonwille/wmenu"
	"github.com/pkg/errors"
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
	fmt.Println("This will remove Nexus teleporting NPCs and announcements")

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

// Clean removes nexus portal npcs
func Clean(c *client.Client) (err error) {
	ids, err := c.Database.SpawnGroupIDsByNameWildcardSearch("nexus_scion")
	if err != nil {
		err = fmt.Errorf("Error getting %s Ids: %s", focus, err.Error())
		return
	}
	vids, err := c.Database.SpawnGroupIDsByNameWildcardSearch("a_mystic_voice")
	for _, id := range vids {
		ids = append(ids, id)
	}

	totalChanged, err := c.Database.SpawnGroupAndEntryByIDDelete(ids)
	if err != nil {
		err = errors.Wrapf(err, "failed to delete spawngroups for %s", focus)
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in spawnentry and spawngroup successfully.")

	filePaths := []string{
		"dreadlands/A_Mystic_Voice.lua",
		"gfaydark/A_Mystic_Voice.pl",
		"greatdivide/A_Mystic_Voice.pl",
		"northkarana/A_Mystic_Voice.lua",
		"tox/A_Mystic_Voice.pl",
		"toxxulia/A_Mystic_Voice.pl",
	}

	delCount, err := c.Quest.Delete(filePaths)
	if err != nil {
		err = errors.Wrapf(err, "failed to delete quests for %s", focus)
		return
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
