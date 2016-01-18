//This script will disable soulbinders by removing their spawnentry and spawngroups with a wildcard '%soulbinder%'
package soulbinder

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func Clean(db *sqlx.DB) (err error) {
	ids, err := getSoulbinderIds(db)
	if err != nil {
		err = fmt.Errorf("Error getting soulbinder Ids: %s", err.Error())
		return
	}
	if len(ids) < 1 {
		fmt.Println("No soulbinders were found to delete")
		return
	}
	fmt.Println("Found", len(ids), "soulbinder NPC entries")
	totalChanged, err := removeSoulbinderEntries(db, ids)
	if err != nil {
		err = fmt.Errorf("Error removing soulbinder entries: %s", err.Error())
		return
	}
	fmt.Println("Removed", totalChanged, " rows related to Soulbinder in spawnentry and spawngroup successfully.")
	return
}

//Get an array of soulbinder NPC ids
func getSoulbinderIds(db *sqlx.DB) (ids []int64, err error) {

	rows, err := db.Query("SELECT sg.id as id FROM npc_types nt INNER JOIN spawnentry se ON se.npcid = nt.id INNER JOIN spawngroup sg ON sg.id = se.spawngroupid WHERE nt.name LIKE '%soulbinder%'")
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

//Remove soulbinders by taking out spawngroups and spawn entries
func removeSoulbinderEntries(db *sqlx.DB, ids []int64) (totalChanged int64, err error) {
	for _, id := range ids {
		var result sql.Result
		var affect int64
		//Remove from spawngroup
		result, err = db.Exec("DELETE FROM spawngroup WHERE id = ?", id)
		if err != nil {
			fmt.Println("Err removing from spawngroup:", err.Error())
			return
		}

		affect, err = result.RowsAffected()
		if err != nil {
			fmt.Println("Error getting rows affected for", id, err.Error())
			return
		}
		if affect < 1 {
			fmt.Println("No rows affecteted delete spawngroup", id)
		}
		totalChanged += affect

		//Remove from spawnentry
		result, err = db.Exec("DELETE FROM spawnentry WHERE spawngroupid = ?", id)
		if err != nil {
			fmt.Println("Err removing from spawnentry:", err.Error())
			return
		}
		affect, err = result.RowsAffected()
		if err != nil {
			fmt.Println("Error getting spawnentry rows affected for", id, err.Error())
			return
		}
		if affect < 1 {
			fmt.Println("No rows affected delete spawnentry", id)
		}
		totalChanged += affect
	}
	return
}
