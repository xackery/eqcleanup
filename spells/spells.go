//This script will disable spells by removing every instance of them
package spells

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"github.com/xackery/eqcleanup/item"
	"github.com/xackery/eqcleanup/spawngroup"
)

var focus = "spells"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {

	//SpawnGroups
	ids := []int64{
	/*//Neriakc
	//cleric
	3162, //svlia in Neriakc
	3163, //Isshia
	3164, //trik
	3160, //lyniv
	3159, //myrish
	3165, //sol
	//wizard
	3142, //jusar
	3141, //misar
	5791, //drisi
	5792, //ash
	//SK
	3152,
	//Necro
	3154,
	//Neriakb
	//wiz
	3127, //riv
	3122, //tal
	3130,
	3129,
	3121,
	3124,
	3125,
	3132,
	3131,
	3133,
	3128,
	3119,
	3123,
	//rivervale
	7706, //ilscent
	5674, //bumpy
	//staria
	5686, //fralith
	5676, //torth
	//akanon
	48152,
	48151,
	48150,
	48158,
	7159,
	7160,
	7154,
	48136,

	//gfay
	1523,
	1522,
	//grobb
	5581,
	3480,
	5582,
	5558,
	5591,
	5615,
	5628,*/
	}

	fmt.Println("Getting all spell selling NPCs...")

	mids, err := spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Bard Songs")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "General Spells")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Berserker Tomes")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Cleric Spells")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Druid Spells")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Enchanter Spells")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Magician Spells")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Monk Tomes")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Necromancer Spells")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Paladin Spells")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Ranger Spells")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Rogue Tomes")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Shadowknight Spells")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Shaman Spells")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Warrior Tomes")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	mids, err = spawngroup.GetSpawnGroupIdsByLastNameWildcard(db, "Wizard Spells")
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	for _, id := range mids {
		ids = append(ids, id)
	}

	fmt.Println("Got", len(ids), "Npcs")

	fmt.Println("Removing spell NPC spawn points, there's a lot, may take a while")

	totalChanged, err := spawngroup.RemoveSpawnGroupAndEntryById(db, ids)
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}

	//Find all item IDS
	fmt.Println("Now getting all spell item entries, there's a lot here too, expect a while")
	rows, err := db.Query("SELECT id FROM items WHERE scrolleffect > 0")
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
	fmt.Println("Found", len(ids), " spells, deleting them...")
	moreChanged, err := item.RemoveAllInstancesOfItems(db, ids)
	if err != nil {
		return
	}
	totalChanged += moreChanged
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in all player-accessible item locations.")
	return
}
