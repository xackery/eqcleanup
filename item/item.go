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

		affect, err = DeleteQuery(db, "DELETE FROM merchantlist_temp ml WHERE ml.item_id = ?", id, "merchant list")
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

		affect, err = DeleteQuery(db, "DELETE FROM sharedbank sb WHERE sb.itemid  = ?)", id, "shared bank")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM forage f WHERE f.itemid  = ?)", id, "forage")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM merc_inventory mi WHERE mi.item_id  = ?)", id, "merc inventory")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM tradeskill_recipe_entries tre WHERE tre.item_id = ?)", id, "tradeskill recipe entries")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM buyer b WHERE b.itemid = ?)", id, "buyer")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM character_bandolier cb WHERE cb.item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM character_corpse_items cci WHERE cci.item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM character_pet_inventory cpi WHERE cpi.item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM character_potionbelt cp WHERE cp.item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM fishing f WHERE f.itemid = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM item_tick it WHERE it.it_itemid = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM keyring k WHERE k.item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM object_contents oc WHERE oc.itemid = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM pet_equipmentset_entries pee WHERE pee.item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM starting_items si WHERE si.itemid = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM tasks t WHERE t.rewardid = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM titles t WHERE t.item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM trader t WHERE t.item_id = ?)", id, "bandolier")
		if err != nil {
			return
		}
		totalRemoved += affect

		affect, err = DeleteQuery(db, "DELETE FROM veteran_reward_templates vrt WHERE vrt.item_id = ?)", id, "bandolier")
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
