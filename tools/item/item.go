package item

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

//Flushes all possible ways to get an item, that I can think of.. except via quests and GM commands
func RemoveAllInstancesOfItems(db *sqlx.DB, ids []int64) (totalRemoved int64, err error) {

	fmt.Println("Removing", len(ids), "item ids")
	idsString := ""
	for _, id := range ids {
		idsString = fmt.Sprintf("%s%d, ", idsString, id)
	}
	idsString = idsString[0 : len(idsString)-2]

	var affect int64
	var delList map[string]string = map[string]string{
		"inventory":                 "itemid",
		"lootdrop_entries":          "item_id",
		"merchantlist":              "item",
		"merchantlist_temp":         "itemid",
		"guild_bank":                "itemid",
		"ground_spawns":             "item",
		"sharedbank":                "itemid",
		"forage":                    "itemid",
		"merc_inventory":            "item_id",
		"tradeskill_recipe_entries": "item_id",
		"buyer":                     "itemid",
		"character_bandolier":       "item_id",
		"character_corpse_items":    "item_id",
		"character_pet_inventory":   "item_id",
		"character_potionbelt":      "item_id",
		"fishing":                   "itemid",
		"item_tick":                 "it_itemid",
		"keyring":                   "item_id",
		"object_contents":           "itemid",
		"pets_equipmentset_entries": "item_id",
		"starting_items":            "itemid",
		"tasks":                     "rewardid",
		"titles":                    "item_id",
		"trader":                    "item_id",
		"veteran_reward_templates":  "item_id",
	}
	var result sql.Result
	for table, field := range delList {

		result, err = db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s IN (%s)", table, field, idsString))
		if err != nil {
			fmt.Println("Err removing from", table, ":", err.Error())
			return
		}

		affect, _ = result.RowsAffected()
		fmt.Println("Removed", affect, table, "entries")
		totalRemoved += affect
	}

	result, err = db.Exec(fmt.Sprintf("DELETE FROM tradeskill_recipe WHERE tradeskill_recipe.id IN (SELECT recipe_id FROM tradeskill_recipe_entries WHERE tradeskill_recipe_entries.item_id IN (%s))", idsString))
	if err != nil {
		fmt.Println("Err removing from tradeskill_recipe", ":", err.Error())
		return
	}

	affect, _ = result.RowsAffected()
	fmt.Println("Removed", affect, "tradeskill_recipe entries")
	totalRemoved += affect

	fmt.Printf("Done.\n")

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
