package rodent

import (
	"fmt"

	"github.com/dixonwille/wmenu"
	"github.com/pkg/errors"
	"github.com/xackery/eqcleanup/client"
)

var focus = "rodent"

// Menu represents menu options
func Menu(opt wmenu.Opt) (err error) {
	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	ok = true
	fmt.Println("This will disable rodents and exterminators")

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

// Clean removes exterminators and rodents
func Clean(c *client.Client) (err error) {

	ids, err := c.Database.SpawnGroupIDsByNameWildcardSearch("a_rodent")
	if err != nil {
		err = fmt.Errorf("Error getting %s Ids: %s", focus, err.Error())
		return
	}
	if len(ids) < 1 {
		fmt.Println("No", focus, "were found to delete")
		return
	}

	//Exterminators
	extids, err := c.Database.SpawnGroupIDsByNameWildcardSearch("exterminator_")
	if err != nil {
		err = fmt.Errorf("Error getting exterminator Ids: %s", err.Error())
		return
	}
	if len(extids) > 1 {
		for _, extid := range extids {
			ids = append(ids, extid)
		}
	}

	totalChanged, err := c.Database.SpawnGroupAndEntryByIDDelete(ids)
	if err != nil {
		err = errors.Wrapf(err, "failed to delete %s", focus)
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in spawnentry and spawngroup successfully.")

	filePaths := []string{
		"felwithea/Exterminator_Valern.pl",
		"freeporteast/Exterminator_Larkey.lua",
		"freeporteast/Exterminator_Qalantir.lua",
		"freporte/Exterminator_Larkey.lua",
		"freportw/Exterminator_Qalantir.lua",
		"kaladimb/Exterminator_Vin.pl",
		"neriakb/Exterminator_Damasi.pl",
		"neriakc/Exterminator_Gilea.pl",
		"qeynos/Exterminator_Rasmon.lua",
		"qeynos2/Exterminator_Wintloag.lua",
		"rivervale/Exterminator_Sutten.lua",
	}

	delCount, err := c.Quest.Delete(filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
