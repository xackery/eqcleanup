package core

import (
	"fmt"

	"github.com/dixonwille/wmenu"
	"github.com/xackery/eqcleanup/client"
	"github.com/xackery/eqcleanup/internal/core/wipe"
)

// Menu represents the core menu options
func Menu(opt wmenu.Opt) (err error) {
	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	ok = true
	menu := wmenu.NewMenu("Choose a core option")
	menu.Option("Character Wipe - Erase all character data", c, false, wipe.Menu)
	menu.Option("Return to main menu", nil, false, func(opt wmenu.Opt) error {
		ok = false
		return nil
	})
	for ok {
		err = menu.Run()
	}
	return
}
