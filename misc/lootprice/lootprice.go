package lootprice

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/tools/eqdb"
)

var focus = "lootprice"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	type ItemList struct {
		Name        string  `db:"name"`
		NpcId       int64   `db:"npcid"`
		LootTableId int64   `db:"loottableid"`
		LootDropId  int64   `db:"lootdropid"`
		ItemId      int64   `db:"itemid"`
		Price       float64 `db:"price"`
		Chance      float64 `db:"chance"`
	}

	zone := "wakening"
	fmt.Println("Adjusting pricing for", zone)
	rows, err := db.Queryx(`SELECT se.npcID npcid, nt.name name, i.id itemid, lte.loottable_id loottableid, lde.lootdrop_id lootdropid, lde.chance chance, i.name, i.price price FROM spawn2 s2
			INNER JOIN spawngroup sg ON sg.id = s2.spawngroupID
			INNER JOIN spawnentry se ON se.spawngroupID = sg.id
			INNER JOIN npc_types nt ON nt.id = se.npcID
			INNER JOIN loottable_entries lte ON lte.loottable_id = nt.loottable_id
			INNER JOIN lootdrop_entries lde ON lde.lootdrop_id = nt.loottable_id
			INNER JOIN items i on i.id = lde.item_id
			INNER JOIN merchantlist ml ON ml.item = i.id
			WHERE s2.zone = ? AND lde.chance > 0
			ORDER BY price desc LIMIT 1
		`, zone)
	if err != nil {
		fmt.Println("Error initializing", err.Error())
		return
	}
	//&{Diamond 119111 7741 7741 10037 200000 3.75}
	//iterate results
	for rows.Next() {
		itemList := &ItemList{}
		err = rows.StructScan(itemList)
		if err != nil {
			return
		}

		fmt.Println(itemList)

		newPrice := itemList.Price * (itemList.Chance / 100)
		fmt.Println("Setting price from", itemList.Price, "to", newPrice)

		//Set pricing
		_, err = db.Exec("UPDATE loottable SET mincash = mincash + ?, maxcash = maxcash + ? WHERE id = ?", itemList.Price, itemList.Price, itemList.LootTableId)
		if err != nil {
			return
		}

		//Remove chance of drop
		_, err = db.Exec("UPDATE lootdrop_entries SET chance = 0, disabled_chance = ? WHERE lootdrop_id = ? and item_id = ?", itemList.Chance, itemList.LootDropId, itemList.ItemId)
		if err != nil {
			return
		}
		return
	}
	return
	//&{Diamond 119017 377 10037 200000}
}
