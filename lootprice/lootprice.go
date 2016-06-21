package lootprice

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqemuconfig"
)

var focus = "lootprice"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	type ItemList struct {
		Name   string `db:"name"`
		ItemId int64  `db:"itemid"`
		Price  int64  `db:"price"`
	}

	zone := "wakening"
	fmt.Println("Adjusting pricing for", zone)
	rows, err := db.Query(`SELECT se.npcID npcid, nt.name, i.id itemid, i.name, i.price price, s2.zone FROM spawn2 s2
			INNER JOIN spawngroup sg ON sg.id = s2.spawngroupID
			INNER JOIN spawnentry se ON se.spawngroupID = sg.id
			INNER JOIN npc_types nt ON nt.id = se.npcID
			INNER JOIN loottable_entries lte ON lte.loottable_id = nt.loottable_id
			INNER JOIN lootdrop_entries lde ON lde.lootdrop_id = nt.loottable_id
			INNER JOIN items i on i.id = lde.item_id
			INNER JOIN merchantlist ml ON ml.item = i.id
			WHERE s2.zone = ?
			ORDER BY price desc LIMIT 1
		`, zone)
	if err != nil {
		fmt.Println("Error initializing", err.Error())
		return
	}

	//iterate results
	for rows.Next() {
		itemList := &ItemList{}
		err = rows.Scan(itemList)
		if err != nil {
			return
		}
		fmt.Println(itemList)
		return
	}
	return
}
