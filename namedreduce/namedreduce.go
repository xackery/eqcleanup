package namedreduce

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqemuconfig"
)

var focus = "nameds"

type Spawndata struct {
	Npcid        int
	Spawngroupid int
	Chance       int
}

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {

	spawns := []Spawndata{
		Spawndata{Npcid: 37082, Spawngroupid: 2096, Chance: 5}, //Arthikus was 50
		Spawndata{Npcid: 35160, Spawngroupid: 4748, Chance: 5}, //Ghoul of TakishHiz was 20
		Spawndata{Npcid: 37121, Spawngroupid: 3271, Chance: 5}, //A sand giant was 50
		Spawndata{Npcid: 37121, Spawngroupid: 3376, Chance: 5}, //A sand giant was 50
		Spawndata{Npcid: 37157, Spawngroupid: 2095, Chance: 5}, //cazel was 50
		Spawndata{Npcid: 22187, Spawngroupid: 788, Chance: 5},  //a griffon was 50
		Spawndata{Npcid: 22187, Spawngroupid: 792, Chance: 5},  //a griffon was 50
	}

	totalRemoved := int64(0)

	for _, spawn := range spawns {
		var result sql.Result
		result, err = db.Exec("UPDATE spawnentry SET chance = ? WHERE npcid = ? AND spawngroupid = ?", spawn.Chance, spawn.Npcid, spawn.Spawngroupid)
		if err != nil {
			fmt.Println("Err updating spawngroup:", err.Error())
			return
		}
		var affect int64
		affect, err = result.RowsAffected()
		if err != nil {
			fmt.Println("Error getting rows affected for", focus, err.Error())
			return
		}
		totalRemoved += affect
	}
	fmt.Println("Updated", totalRemoved, " DB entries related to", focus, "in spawnentry and spawngroup successfully.")

	return
}
