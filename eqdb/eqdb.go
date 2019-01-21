package eqdb

import (
	"database/sql"
	"fmt"

	// Used for MYSQL connections
	_ "github.com/go-sql-driver/mysql"

	"github.com/pkg/errors"
	"github.com/xackery/eqemuconfig"
)

// Database wraps db information
type Database struct {
	Instance *sql.DB
}

// New creates a new database instance
func New(config *eqemuconfig.Config) (db *Database, err error) {
	db = &Database{}
	db.Instance, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Db))
	if err != nil {
		return
	}
	fmt.Printf("Connected to MySQL %s as %s\n", config.Database.Host, config.Database.Username)
	return
}

// Delete runs a Delete query
func (db *Database) Delete(query string, value int64, focus string) (count int64, err error) {
	result, err := db.Instance.Exec(query, value)
	if err != nil {
		err = errors.Wrapf(err, "failed to delete %d from %s", value, focus)
		return
	}

	count, err = result.RowsAffected()
	if err != nil {
		err = errors.Wrapf(err, "failed to get rows affected for %s", focus)
		return
	}
	return
}

// ItemInstanceDelete deletes all possible ways to get an item.
func (db *Database) ItemInstanceDelete(ids []int64) (totalDeleted int64, err error) {

	fmt.Println("Removing", len(ids), "item ids")
	idsString := ""
	for _, id := range ids {
		idsString = fmt.Sprintf("%s%d, ", idsString, id)
	}
	idsString = idsString[0 : len(idsString)-2]

	var affect int64
	var delList = map[string]string{
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

		result, err = db.Instance.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s IN (%s)", table, field, idsString))
		if err != nil {
			fmt.Println("Err removing from", table, ":", err.Error())
			return
		}

		affect, _ = result.RowsAffected()
		fmt.Println("Deleted", affect, table, "entries")
		totalDeleted += affect
	}

	result, err = db.Instance.Exec(fmt.Sprintf("DELETE FROM tradeskill_recipe WHERE tradeskill_recipe.id IN (SELECT recipe_id FROM tradeskill_recipe_entries WHERE tradeskill_recipe_entries.item_id IN (%s))", idsString))
	if err != nil {
		fmt.Println("Err removing from tradeskill_recipe", ":", err.Error())
		return
	}

	affect, _ = result.RowsAffected()
	fmt.Println("Deleted", affect, "tradeskill_recipe entries")
	totalDeleted += affect

	fmt.Println("Done")
	return
}

// SpawnGroupIDsByNameWildcardSearch returns a list of spawns for a npc_type
func (db *Database) SpawnGroupIDsByNameWildcardSearch(wildcard string) (ids []int64, err error) {

	rows, err := db.Instance.Query("SELECT sg.id as id FROM npc_types nt INNER JOIN spawnentry se ON se.npcid = nt.id INNER JOIN spawngroup sg ON sg.id = se.spawngroupid WHERE nt.name LIKE ?", "%"+wildcard+"%")
	if err != nil {
		fmt.Println("Error initial")
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
	return
}

// SpawnGroupIDsByEmptyMerchantSearch searches for spawns of empty merchants
func (db *Database) SpawnGroupIDsByEmptyMerchantSearch() (ids []int64, err error) {

	rows, err := db.Instance.Query("SELECT sg.id as id FROM npc_types nt INNER JOIN spawnentry se ON se.npcid = nt.id INNER JOIN spawngroup sg ON sg.id = se.spawngroupid WHERE nt.merchant_id > 0 AND nt.merchant_id not in (select merchantid from merchantlist)")
	if err != nil {
		fmt.Println("Error initial")
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
	return
}

// SpawnGroupIDsByLastNameWildcardSearch returns all spawns with a specified last name
func (db *Database) SpawnGroupIDsByLastNameWildcardSearch(wildcard string) (ids []int64, err error) {

	rows, err := db.Instance.Query("SELECT sg.id as id FROM npc_types nt INNER JOIN spawnentry se ON se.npcid = nt.id INNER JOIN spawngroup sg ON sg.id = se.spawngroupid WHERE nt.lastname LIKE ?", "%"+wildcard+"%")
	if err != nil {
		fmt.Println("Error initial")
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
	return
}

// SpawnGroupIDsByClassSearch returns a list of spawns by class
func (db *Database) SpawnGroupIDsByClassSearch(class int64) (ids []int64, err error) {

	rows, err := db.Instance.Query("SELECT sg.id as id FROM npc_types nt INNER JOIN spawnentry se ON se.npcid = nt.id INNER JOIN spawngroup sg ON sg.id = se.spawngroupid WHERE nt.class = ?", class)
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
	return
}

// GridByZoneDelete deletes a grid by a zone name
func (db *Database) GridByZoneDelete(zone string) (totalDeleted int64, err error) {

	var result sql.Result
	var affect int64

	result, err = db.Instance.Exec(`DELETE grid_entries FROM grid_entries
		INNER JOIN zone ON zone.zoneidnumber = zoneid 
		WHERE zone.short_name = ?`, zone)
	if err != nil {
		fmt.Println("Err removing from gridentries:", err.Error())
		return
	}
	affect, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting gridentries rows affected for", zone, err.Error())
		return
	}
	if affect < 1 {
		//fmt.Println("No rows affected delete spawnentry", id)
	}
	totalDeleted += affect

	result, err = db.Instance.Exec(`DELETE grid FROM grid 
		INNER JOIN zone ON zone.zoneidnumber = grid.zoneid 
		WHERE zone.short_name = ?`, zone)
	if err != nil {
		fmt.Println("Err removing from grid:", err.Error())
		return
	}
	affect, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting grid rows affected for", zone, err.Error())
		return
	}
	if affect < 1 {
		//fmt.Println("No rows affected delete spawnentry", id)
	}
	totalDeleted += affect
	return
}

// SpawnGroupAndEntryByZoneDelete deletes spawn entries by zone
func (db *Database) SpawnGroupAndEntryByZoneDelete(zone string) (totalDeleted int64, err error) {

	var result sql.Result
	var affect int64

	//Delete from spawnentry
	result, err = db.Instance.Exec(`DELETE spawnentry FROM spawnentry 
		INNER JOIN spawngroup ON spawnentry.spawngroupid = spawngroup.id
		INNER JOIN spawn2 ON spawngroup.id = spawn2.spawngroupid 
		WHERE spawn2.zone = ?`, zone)
	if err != nil {
		fmt.Println("Err removing from spawnentry:", err.Error())
		return
	}
	affect, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting spawnentry rows affected for", zone, err.Error())
		return
	}
	if affect < 1 {
		//fmt.Println("No rows affected delete spawnentry", id)
	}
	totalDeleted += affect

	result, err = db.Instance.Exec(`DELETE spawngroup FROM spawngroup INNER JOIN spawn2 ON spawngroup.id = spawn2.spawngroupid WHERE spawn2.zone = ?`, zone)
	if err != nil {
		fmt.Println("Err removing from spawngroup:", err.Error())
		return
	}

	affect, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected for", zone, err.Error())
		return
	}
	if affect < 1 {
		//fmt.Println("No rows affecteted delete spawngroup", id)
	}
	totalDeleted += affect

	//Delete from spawn2

	result, err = db.Instance.Exec("DELETE FROM spawn2 WHERE spawn2.zone = ?", zone)
	if err != nil {
		//fmt.Println("Err removing from spawngroupid:", err.Error())
		return
	}
	affect, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting spawngroupid rows affected for", zone, err.Error())
		return
	}
	if affect < 1 {
		//fmt.Println("No rows affected delete spawngroupid", id)
	}
	totalDeleted += affect
	return
}

// SpawnGroupAndEntryByIDDelete deletes spawn group data
func (db *Database) SpawnGroupAndEntryByIDDelete(ids []int64) (totalDeleted int64, err error) {
	if len(ids) < 1 {
		return
	}
	idsString := ""
	for _, id := range ids {
		idsString = fmt.Sprintf("%s%d, ", idsString, id)
	}
	idsString = idsString[0 : len(idsString)-2]

	var result sql.Result
	var affect int64
	//Delete from spawngroup
	result, err = db.Instance.Exec(fmt.Sprintf("DELETE FROM spawngroup WHERE id IN (%s)", idsString))
	if err != nil {
		fmt.Println("Err removing from spawngroup:", err.Error())
		return
	}

	affect, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected for", err.Error())
		return
	}
	if affect < 1 {
		//fmt.Println("No rows affecteted delete spawngroup", id)
	}
	totalDeleted += affect

	//Delete from spawnentry
	result, err = db.Instance.Exec(fmt.Sprintf("DELETE FROM spawnentry WHERE spawngroupid IN (%s)", idsString))
	if err != nil {
		fmt.Println("Err removing from spawnentry:", err.Error())
		return
	}
	affect, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting spawnentry:", err.Error())
		return
	}
	if affect < 1 {
		//fmt.Println("No rows affected delete spawnentry", id)
	}
	totalDeleted += affect

	//Delete from spawn2

	result, err = db.Instance.Exec(fmt.Sprintf("DELETE FROM spawn2 WHERE spawngroupid IN (%s)", idsString))
	if err != nil {
		//fmt.Println("Err removing from spawngroupid:", err.Error())
		return
	}
	affect, err = result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting spawngroupid:", err.Error())
		return
	}
	if affect < 1 {
		//fmt.Println("No rows affected delete spawngroupid", id)
	}
	totalDeleted += affect

	return
}
