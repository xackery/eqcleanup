package tribute

import (
	"fmt"

	"github.com/dixonwille/wmenu"
	"github.com/pkg/errors"
	"github.com/xackery/eqcleanup/client"
)

var focus = "tribute"

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

// Clean removes tribute npcs
func Clean(c *client.Client) (err error) {

	//tribute master
	ids, err := c.Database.SpawnGroupIDsByClassSearch(63)
	if err != nil {
		err = fmt.Errorf("Error getting %s Ids: %s", focus, err.Error())
		return
	}

	//tribute grand master
	mids, err := c.Database.SpawnGroupIDsByClassSearch(64)
	for _, id := range mids {
		ids = append(ids, id)
	}

	if len(ids) < 1 {
		fmt.Println("No ", focus, "were found to delete")
		return
	}

	totalChanged, err := c.Database.SpawnGroupAndEntryByIDDelete(ids)
	if err != nil {
		err = errors.Wrapf(err, "failed to remove spawn groups for %s", focus)
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in spawnentry and spawngroup successfully.")
	return
}
