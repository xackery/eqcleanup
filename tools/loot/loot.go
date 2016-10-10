package loot

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func RemoveLootByZone(db *sqlx.DB, zone string) (totalRemoved int64, err error) {

	var result sql.Result
	var affect int64

	if result, err = db.Exec(`DELETE lootdrop_entries FROM lootdrop_entries
		INNER JOIN loottable_entries ON loottable_entries.lootdrop_id = lootdrop_entries.lootdrop_id 
		INNER JOIN npc_types ON npc_types.loottable_id = loottable_entries.loottable_id 
		INNER JOIN spawnentry ON spawnentry.npcID = npc_types.id
		INNER JOIN spawn2 ON spawn2.spawngroupid = spawnentry.spawngroupid
		WHERE spawn2.zone = ?`, zone); err != nil {
		err = fmt.Errorf("Error removing lootdrop_entries: %s", err.Error())
		return
	}
	if affect, err = result.RowsAffected(); err != nil {
		err = fmt.Errorf("Error getting rows affected: %s", err.Error())
		return
	}
	fmt.Println("Removed", affect, "lootdrop_entries")
	totalRemoved += affect
	if result, err = db.Exec(`DELETE loottable FROM loottable 
		INNER JOIN loottable_entries ON loottable_entries.loottable_id = loottable.id 
		INNER JOIN npc_types ON npc_types.loottable_id = loottable_entries.loottable_id 
		INNER JOIN spawnentry ON spawnentry.npcID = npc_types.id
		INNER JOIN spawn2 ON spawn2.spawngroupid = spawnentry.spawngroupid
		WHERE spawn2.zone = ?`, zone); err != nil {
		err = fmt.Errorf("Error removing loottable: %s", err.Error())
		return
	}
	if affect, err = result.RowsAffected(); err != nil {
		err = fmt.Errorf("Error getting rows affected: %s", err.Error())
		return
	}
	fmt.Println("Removed", affect, "loottable")
	totalRemoved += affect

	if result, err = db.Exec(`DELETE lootdrop FROM lootdrop 
		INNER JOIN loottable_entries ON loottable_entries.lootdrop_id = lootdrop.id 
		INNER JOIN npc_types ON npc_types.loottable_id = loottable_entries.loottable_id 
		INNER JOIN spawnentry ON spawnentry.npcID = npc_types.id
		INNER JOIN spawn2 ON spawn2.spawngroupid = spawnentry.spawngroupid
		WHERE spawn2.zone = ?`, zone); err != nil {
		err = fmt.Errorf("Error removing lootdrop:", err.Error())
		return
	}
	if affect, err = result.RowsAffected(); err != nil {
		err = fmt.Errorf("Error getting rows affected: %s", err.Error())
		return
	}
	fmt.Println("Removed", affect, "lootdrop")
	totalRemoved += affect

	if result, err = db.Exec(`DELETE loottable_entries FROM loottable_entries 
		INNER JOIN npc_types ON npc_types.loottable_id = loottable_entries.loottable_id 
		INNER JOIN spawnentry ON spawnentry.npcID = npc_types.id
		INNER JOIN spawn2 ON spawn2.spawngroupid = spawnentry.spawngroupid
		WHERE spawn2.zone = ?`, zone); err != nil {
		err = fmt.Errorf("Error removing loottable_entries:", err.Error())
		return
	}
	if affect, err = result.RowsAffected(); err != nil {
		err = fmt.Errorf("Error getting rows affected: %s", err.Error())
		return
	}
	fmt.Println("Removed", affect, "loottable_entries")
	totalRemoved += affect

	return
}
