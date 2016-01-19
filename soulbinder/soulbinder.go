//This script will disable soulbinders by removing their spawnentry and spawngroups with a wildcard '%soulbinder%'
package soulbinder

import (
	"database/sql"
	"eqemuconfig"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"spawngroup"
)

var focus = "soulbinder"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	ids, err := spawngroup.GetSpawnGroupIds(db, focus)
	ids = append(ids, 3199)  //Add Romi to delete list
	ids = append(ids, 54932) //And priestess aelea
	if err != nil {
		err = fmt.Errorf("Error getting ", focus, " Ids: %s", err.Error())
		return
	}
	if len(ids) < 1 {
		fmt.Println("No ", focus, "were found to delete")
		return
	}
	totalChanged, err := spawngroup.RemoveSpawnGroupAndEntryById(db, ids)
	if err != nil {
		err = fmt.Errorf("Error removing", focus, "entries: %s", err.Error())
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in spawnentry and spawngroup successfully.")

	filePaths := []string{
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

	delCount, err := removeQuests(config, filePaths)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Deleted", delCount, focus, "related quest files")
	return
}
