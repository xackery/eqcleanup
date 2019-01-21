package zone

import (
	"fmt"

	"github.com/dixonwille/wmenu"
	"github.com/xackery/eqcleanup/client"
	"github.com/xackery/eqcleanup/internal/zone/lavastorm"
)

// Menu represents the core menu options
func Menu(opt wmenu.Opt) (err error) {

	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	ok = true
	menu := wmenu.NewMenu("Choose a zone to modify")
	menu.Option("Lavastorm", c, false, lavastorm.Menu)
	/*menu.Option("Splitpaw", c, false, paw.Menu)
	menu.Option("Nurga", c, false, nurga.Menu)
	menu.Option("Droga", c, false, droga.Menu)
	menu.Option("Veeshan", c, false, veeshan.Menu)
	menu.Option("East Commonlands", c, false, ecommons.Menu)
	menu.Option("West Commonlands", c, false, wcommons.Menu)*/
	menu.Option("Return to main menu", nil, false, func(opt wmenu.Opt) error {
		ok = false
		return nil
	})
	for ok {
		err = menu.Run()
	}
	return
}
