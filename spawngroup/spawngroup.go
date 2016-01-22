package spawngroup

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func GetSpawnGroupIdsByNameWildcard(db *sqlx.DB, wildcard string) (ids []int64, err error) {

	rows, err := db.Query("SELECT sg.id as id FROM npc_types nt INNER JOIN spawnentry se ON se.npcid = nt.id INNER JOIN spawngroup sg ON sg.id = se.spawngroupid WHERE nt.name LIKE '%?%'", wildcard)
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

func GetSpawnGroupIdsByClass(db *sqlx.DB, class int64) (ids []int64, err error) {

	rows, err := db.Query("SELECT sg.id as id FROM npc_types nt INNER JOIN spawnentry se ON se.npcid = nt.id INNER JOIN spawngroup sg ON sg.id = se.spawngroupid WHERE nt.class = ?", class)
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

func RemoveSpawnGroupAndEntryById(db *sqlx.DB, ids []int64) (totalRemoved int64, err error) {

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
		totalRemoved += affect

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
		totalRemoved += affect
	}
	return
}
