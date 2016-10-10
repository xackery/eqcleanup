package spawngroup

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func GetSpawnGroupIdsByNameWildcard(db *sqlx.DB, wildcard string) (ids []int64, err error) {

	rows, err := db.Query("SELECT sg.id as id FROM npc_types nt INNER JOIN spawnentry se ON se.npcid = nt.id INNER JOIN spawngroup sg ON sg.id = se.spawngroupid WHERE nt.name LIKE ?", "%"+wildcard+"%")
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

func GetSpawnGroupIdsByEmptyMerchant(db *sqlx.DB) (ids []int64, err error) {

	rows, err := db.Query("SELECT sg.id as id FROM npc_types nt INNER JOIN spawnentry se ON se.npcid = nt.id INNER JOIN spawngroup sg ON sg.id = se.spawngroupid WHERE nt.merchant_id > 0 AND nt.merchant_id not in (select merchantid from merchantlist)")
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

func GetSpawnGroupIdsByLastNameWildcard(db *sqlx.DB, wildcard string) (ids []int64, err error) {

	rows, err := db.Query("SELECT sg.id as id FROM npc_types nt INNER JOIN spawnentry se ON se.npcid = nt.id INNER JOIN spawngroup sg ON sg.id = se.spawngroupid WHERE nt.lastname LIKE ?", "%"+wildcard+"%")
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

func RemoveSpawnGroupAndEntryByZone(db *sqlx.DB, zone string) (totalRemoved int64, err error) {

	var result sql.Result
	var affect int64

	//Remove from spawnentry
	result, err = db.Exec(`DELETE spawnentry FROM spawnentry 
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
	totalRemoved += affect

	result, err = db.Exec(`DELETE spawngroup FROM spawngroup INNER JOIN spawn2 ON spawngroup.id = spawn2.spawngroupid WHERE spawn2.zone = ?`, zone)
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
	totalRemoved += affect

	//Remove from spawn2

	result, err = db.Exec("DELETE FROM spawn2 WHERE spawn2.zone = ?", zone)
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
	totalRemoved += affect
	return
}

func RemoveSpawnGroupAndEntryById(db *sqlx.DB, ids []int64) (totalRemoved int64, err error) {
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
	//Remove from spawngroup
	result, err = db.Exec("DELETE FROM spawngroup WHERE id IN ?", idsString)
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
	totalRemoved += affect

	//Remove from spawnentry
	result, err = db.Exec("DELETE FROM spawnentry WHERE spawngroupid IN ? ", idsString)
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
	totalRemoved += affect

	//Remove from spawn2

	result, err = db.Exec("DELETE FROM spawn2 WHERE spawngroupid IN ?", idsString)
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
	totalRemoved += affect

	return
}
