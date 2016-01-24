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
		affect, err = DeleteQuery(db, "DELETE FROM inventory WHERE itemid = ?", id, "inventory")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM lootdrop_entries WHERE item_id = ?", id, "lootdrop entries")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM merchantlist WHERE item_id = ?", id, "merchant list")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM merchantlist_temp WHEREitem_id = ?", id, "merchant list")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM guild_bank WHERE itemid = ?", id, "guild bank")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM ground_spawns WHERE item = ?", id, "ground spawn")
		if err != nil {
			return
		}
		totalRemoved += affect

		//Delete any tradeskill recipes that use this as a reagent
		affect, err = DeleteQuery(db, "DELETE FROM tradeskill_recipe WHERE tradeskill_recipe.id IN (SELECT recipe_id FROM tradeskill_recipe_entries WHERE tradeskill_recipe_entries.item_id = ?)", id, "tradeskill recipes")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM sharedbank WHERE itemid  = ?)", id, "shared bank")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM forage WHERE itemid  = ?)", id, "forage")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM merc_inventory WHERE item_id  = ?)", id, "merc inventory")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM tradeskill_recipe_entries WHERE item_id = ?)", id, "tradeskill recipe entries")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM buyer WHERE itemid = ?)", id, "buyer")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM character_bandolier WHERE item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM character_corpse_items WHERE item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM character_pet_inventory WHERE item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM character_potionbelt WHERE item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM fishing WHERE itemid = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM item_tick WHERE it_itemid = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM keyring WHERE item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM object_contents WHERE itemid = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM pet_equipmentset_entries WHERE item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM starting_items WHERE itemid = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM tasks WHERE rewardid = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM titles WHERE item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM trader WHERE item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM veteran_reward_templates WHERE item_id = ?)", id, "bandolier")
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
