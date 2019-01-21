package aug

import (
	"fmt"

	"github.com/dixonwille/wmenu"
	"github.com/xackery/eqcleanup/client"
)

var focus = "era"

// Menu represents the wipe menu options
func Menu(opt wmenu.Opt) (err error) {
	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	ok = true
	fmt.Println("This will remove any item flagged as an augmentation from being obtainable in the game")

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

// Clean remove augments
func Clean(c *client.Client) (err error) {

	totalChanged := int64(0)

	rows, err := c.Database.Instance.Query("SELECT id FROM items WHERE augtype > 0")
	if err != nil {
		return
	}

	ids := []int64{}

	//iterate results
	for rows.Next() {
		id := int64(0)
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		ids = append(ids, id)
	}
	fmt.Println("Found", len(ids), " augments to remove")
	moreChanged, err := c.Database.ItemInstanceDelete(ids)
	if err != nil {
		return
	}
	totalChanged += moreChanged

	fmt.Println("Updated", totalChanged, " DB entries related to", focus)
	return
}
