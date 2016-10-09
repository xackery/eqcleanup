//Removes augments
package aug

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/tools/item"
	"github.com/xackery/eqemuconfig"
)

var focus = "aug"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {

	totalChanged := int64(0)

	fmt.Println("Looking for augment item ids...")
	rows, err := db.Query("SELECT id FROM items WHERE augtype > 0")
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
	fmt.Println("Found", len(ids), " augments, deleting their entries...")
	moreChanged, err := item.RemoveAllInstancesOfItems(db, ids)
	if err != nil {
		return
	}
	totalChanged += moreChanged

	fmt.Println("Updated", totalChanged, " DB entries related to", focus)

	return
}
