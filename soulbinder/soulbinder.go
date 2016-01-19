//This script will disable soulbinders by removing their spawnentry and spawngroups with a wildcard '%soulbinder%'
package soulbinder

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"os"
)

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
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

	err = removeQuests(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func removeQuests(config *eqemuconfig.Config) (err error) {
	if config.QuestsDir == "" {
		err = fmt.Errorf("Quests directory is not set.")
		return
	}
	soulbinders := []string{
		"abysmal/Soulbinder_Jerlin.pl",
		"cabeast/Soulbinder_Shakar.pl",
		"commonlands/Soulbinder_Jubbl.pl",
		"crescent/Priestess_Aelea.pl",
		"ecommons/Soulbinder_Jubbl.pl",
		"everfrost/Soulbinder_Garnog.pl",
		"firiona/Soulbinder_Tardon.pl",
		"gfaydark/Soulbinder_Oakstout.pl",
		"guildlobby/High_Priest_of_Luclin.pl",
		"guildlobby/High_Priestess_of_Luclin.pl",
		"gukta/Soulbinder_Snog.pl",
		"gukta/Soulbinder_Zlippi.pl",
		"gunthak/Soulbinder_Karyin.pl",
		"iceclad/Soulbinder_Cubnitskin.pl",
		"kaladima/Soulbinder_Torvald.pl",
		"neriaka/Soulbinder_Nola_Z-Ret.pl",
		"neriaka/Soulbinder_Novalu.pl",
		"northkarana/Romi.pl",
		"northro/Soulbinder_Ragni.pl",
		"nro/Soulbinder_Ragni.pl",
		"oggok/Soulbinder_Trurg.pl",
		"overthere/Soulbinder_Kardin.pl",
		"paineel/Soulbinder_Tomas.pl",
		//"plugins/default-actions.pl", //There is a soulbinder flag here, ignoring though
		"plugins/soulbinders.pl",
		"poknowledge/Soulbinder_Jera.pl",
		"potranquility/Soulbinder_Derith.pl",
		"rathemtn/Soulbinder_Zlippi.pl",
		"shadowhaven/Soulbinder_Nansin.pl",
		"sharvahl/Soulbinder_Ghula.pl",
		"southro/Soulbinder_Silandra.pl",
		"sro/Soulbinder_Silandra.pl",
	}
	delCount := 0
	for _, filename := range soulbinders {
		curFile := config.QuestsDir + "/" + filename
		_, err = os.Stat(curFile)

		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			fmt.Printf("Error finding %s: %s", curFile, err.Error())
			continue
		}

		err = os.Remove(curFile)
		if err != nil {
			fmt.Printf("Error deleting %s: %s", curFile, err.Error())
			continue
		}
		delCount++
	}
	fmt.Println("Deleted", delCount, " soulbinder related quest files")
	return
}

//Get an array of soulbinder NPC ids
func getSoulbinderIds(db *sqlx.DB) (ids []int64, err error) {

	rows, err := db.Query("SELECT sg.id as id FROM npc_types nt INNER JOIN spawnentry se ON se.npcid = nt.id INNER JOIN spawngroup sg ON sg.id = se.spawngroupid WHERE nt.name LIKE '%soulbinder%'")
	if err != nil {
		return
	}

	ids = append(ids, 3199)  //Add Romi to delete list
	ids = append(ids, 54932) //And priestess aelea
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
