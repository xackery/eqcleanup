//This script will disable spells by removing every instance of them
package spells

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"github.com/xackery/eqcleanup/item"
)

var focus = "spells"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	//Find all item IDS

	ids := []int64{}
	rows, err := db.Query("SELECT id FROM items WHERE scrolleffect > 0")
	if err != nil {
		return
	}

	//iterate results
	for rows.Next() {
		id := int64(0)
		err = rows.Scan(&id)
		if err != nil {
			return
		}
		ids = append(ids, id)
	}

	totalChanged, err := item.RemoveAllInstancesOfItems(db, ids)
	if err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in all player-accessible item locations.")
	return
}
