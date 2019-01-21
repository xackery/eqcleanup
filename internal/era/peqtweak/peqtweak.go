package peqtweak

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqemuconfig"
	"github.com/dixonwille/wmenu"
	"github.com/xackery/eqcleanup/client"
)

var focus = "peqtweaks"

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

// Clean adjusts skeletons to move
func Clean(c *client.Client) (err error) {

	totalChanged := int64(0)

	result, err := db.Exec("UPDATE npc_types SET runspeed = 1.325 WHERE runspeed = 0 and name = 'a_skeleton'")
	if err != nil {
		fmt.Println("Err updating skeleton movement:", err.Error())
		return
	}

	affect, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected for", focus, err.Error())
		return
	}
	totalChanged += affect

	fmt.Println("Updated", totalChanged, " DB entries to repair skeleton movement")

	return
}
