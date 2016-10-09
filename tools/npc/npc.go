package spawngroup

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func RemoveNPCByZone(db *sqlx.DB, zone string) (totalRemoved int64, err error) {

	var zoneid int64
	var rows *sql.Rows

	if rows, err = db.Exec("SELECT zonenumberid from zone where short_name = ?", zone); err != nil {
		err = fmt.Errorf("Could not find zone %s: %s", zone, err.Error())
		return
	}
	for rows.Next() {
		if err = rows.Scan(&zoneid); err != nil {
			err = fmt.Errorf("Error getting zone id: %s", err.Error())
			return
		}
	}

	if zoneid < 1 {
		err = fmt.Errorf("Error getting zone id")
		return
	}

	var result sql.Result
	var affect int64
	//Remove from spawngroup

	if result, err = db.Exec("DELETE FROM npc_types WHERE id <= ? AND id >= ?", zoneid*1000+999, zoneid*1000); err != nil {
		err = fmt.Errorf("Error deleting npc_types: %s", err.Error())
		return
	}

	if affect, err = result.RowsAffected(); err != nil {
		err = fmt.Errorf("Error gettign rows affected for npc_type: %s", err.Error())
		return
	}
	totalRemoved += affect
	return
}
