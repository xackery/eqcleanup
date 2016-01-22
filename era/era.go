package main

import (
	//	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"os"
)

//skyfire
//50386 Flawed Combatant's Hoe
//50396
//46178
//46177
//gfaydark - crude/newbie

/*
SELECT id, NAME, ac, hp, mana, astr, asta FROM items WHERE id IN (

) ORDER BY hp DESC;
*/
func main() {

	config, err := eqemuconfig.LoadConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	db, err := connectDB(&config)
	if err != nil {
		fmt.Println(err.Error())
	}
	query := `SELECT lde.item_id FROM npc_types nt 
INNER JOIN spawnentry se ON se.npcid = nt.id 
INNER JOIN spawn2 s ON s.spawngroupid = se.spawngroupid
INNER JOIN spawngroup sg ON sg.id = se.spawngroupid
INNER JOIN loottable lt ON lt.id = nt.loottable_id
INNER JOIN loottable_entries le ON le.loottable_id = lt.id
INNER JOIN lootdrop_entries lde ON lde.lootdrop_id = le.lootdrop_id
WHERE s.zone = 'gfaydark'`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
	}
	newStr := ""
	ids := []int64{}
	//iterate results
	for rows.Next() {
		id := int64(0)
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		isNew := true
		for _, curId := range ids {
			if curId == id {
				isNew = false
			}
		}
		if isNew {
			ids = append(ids, id)
			newStr = fmt.Sprintf("%s,%d", newStr, id)
		}
	}
	fmt.Println("Win:", newStr)
}

func connectDB(config *eqemuconfig.Config) (db *sqlx.DB, err error) {
	//Connect to DB
	db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Db))
	if err != nil {
		fmt.Println("Error connecting to DB:", err.Error())
		os.Exit(1)
	}
	return
}
