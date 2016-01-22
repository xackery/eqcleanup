package item

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func RemoveAllInstancesOfItems(db *sqlx.DB, ids []int64) (totalRemoved int64, err error) {
	for _, id := range ids {
		var result sql.Result
		var affect int64

		result, err = db.Exec("DELETE FROM inventory i WHERE itemid = ?", id)
		if err != nil {
			fmt.Println("Error removing ", id, "from inventory:", err.Error())
			return
		}

		affect, err = result.RowsAffected()
		if err != nil {
			fmt.Println("Error getting rows affected for", id, err.Error())
			return
		}
		if affect < 1 {
			fmt.Println("No rows affecteted deleting inventory", id)
		}
		totalRemoved += affect
	}

	return
}
