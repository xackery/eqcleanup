package item

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

//Flushes all possible ways to get an item, that I can think of.. except via quests and GM commands
func RemoveAllInstancesOfItems(db *sqlx.DB, ids []int64) (totalRemoved int64, err error) {
	for _, id := range ids {
		var affect int64
		affect, err = DeleteQuery(db, "DELETE FROM inventory i WHERE i.itemid = ?", id, "inventory")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM lootdrop_entries le WHERE le.item_id = ?", id, "lootdrop entries")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM merchantlist ml WHERE ml.item_id = ?", id, "merchant list")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM guild_bank gb WHERE gb.itemid = ?", id, "guild bank")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM ground_spawns gs WHERE gs.item = ?", id, "ground spawn")
		if err != nil {
			return
		}
		totalRemoved += affect

		//Delete any tradeskill recipes that use this as a reagent
		affect, err = DeleteQuery(db, "DELETE FROM tradeskill_recipe tr WHERE tr.id IN (SELECT recipe_id FROM tradeskill_recipe_entries tre WHERE tre.item_id = ?)", id, "tradeskill recipes")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM tradeskill_recipe_entries tre WHERE tre.item_id = ?)", id, "tradeskill recipe entries")
		if err != nil {
			return
		}
		totalRemoved += affect
	}

	return
}

func DeleteQuery(db *sqlx.DB, query string, itemid int64, focus string) (affect int64, err error) {
	var result sql.Result

	result, err = db.Exec(query, itemid)
	if err != nil {
		fmt.Println("Error removing ", itemid, "from", focus, ":", err.Error())
		return
	}

	affect, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected for", focus, itemid, err.Error())
		return
	}
	return
}
