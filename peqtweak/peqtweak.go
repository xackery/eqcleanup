package peqtweak

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/spawngroup"
	"github.com/xackery/eqemuconfig"
)

var focus = "peqtweaks"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {

	result, err = db.Exec("UPDATE npc_types SET runspeed = 1.325 WHERE runspeed = 0 and name = 'a_skeleton'")
	if err != nil {
		fmt.Println("Err updating skeleton movement:", err.Error())
		return
	}

	affect, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected for", zone, err.Error())
		return
	}
	totalRemoved += affect

	fmt.Println("Updated", totalChanged, " DB entries to repair skeleton movement")

	return
}
