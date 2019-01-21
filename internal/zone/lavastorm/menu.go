package lavastorm

import (
	"fmt"

	"github.com/dixonwille/wmenu"
	"github.com/xackery/eqcleanup/client"
)

var focus = "lavastorm"
var zoneID = 123

// Menu represents the lavastorm menu options
func Menu(opt wmenu.Opt) (err error) {
	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	ok = true
	fmt.Println("This tool creates a new zone to represent old lavastorm with zone id", zoneID)
	fmt.Println("It then updates spell files, item effects, zone lines, and quests to teleport to the new version")

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

// Clean runs the wipe option
func Clean(c *client.Client) (err error) {
	fmt.Println("this is not yet implementd")
	return
}
