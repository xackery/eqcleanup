package named

import (
	"database/sql"
	"fmt"

	"github.com/dixonwille/wmenu"
	"github.com/xackery/eqcleanup/client"
)

var focus = "nameds"

// Menu represents menu options
func Menu(opt wmenu.Opt) (err error) {
	c, ok := opt.Value.(*client.Client)
	if !ok {
		err = fmt.Errorf("invalid client value passed")
		return
	}
	ok = true
	fmt.Println("This will remove fabled NPCs")

	menu := wmenu.NewMenu("Continue operation?")
	menu.Option("Return to main menu", nil, false, func(opt wmenu.Opt) error {
		ok = false
		return nil
	})
	menu.Option("I understand the risks - continue", nil, false, func(opt wmenu.Opt) error {
		return Clean(c)
	})

	for ok {
		err = menu.Run()
	}
	return
}

// Spawndata represents spawn information
type Spawndata struct {
	Npcid        int
	Spawngroupid int
	Chance       int
	Mindelay     int
	Despawntimer int
}

// Clean reduces the spawn rate of named mobs
func Clean(c *client.Client) (err error) {

	spawns := []Spawndata{
		{Npcid: 37082, Spawngroupid: 2096, Chance: 5},                            //Arthikus was 50
		{Npcid: 35160, Spawngroupid: 4748, Chance: 5},                            //Ghoul of TakishHiz was 20
		{Npcid: 37121, Spawngroupid: 3271, Chance: 5},                            //A sand giant was 50
		{Npcid: 37121, Spawngroupid: 3376, Chance: 5},                            //A sand giant was 50
		{Npcid: 37157, Spawngroupid: 2095, Chance: 5},                            //cazel was 50
		{Npcid: 22187, Spawngroupid: 788, Chance: 5},                             //a griffon was 50
		{Npcid: 22187, Spawngroupid: 792, Chance: 5},                             //a griffon was 50
		{Npcid: 92159, Spawngroupid: 9825, Chance: 5},                            //Overseer_Miklek was 34%
		{Npcid: 110100, Spawngroupid: 16400, Mindelay: 50000, Despawntimer: 100}, //Stormfeather

	}

	totalRemoved := int64(0)

	for _, spawn := range spawns {
		var result sql.Result
		if spawn.Chance > 0 {
			result, err = c.Database.Instance.Exec("UPDATE spawnentry SET chance = ? WHERE npcid = ? AND spawngroupid = ?", spawn.Chance, spawn.Npcid, spawn.Spawngroupid)
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
		if spawn.Mindelay > 0 && spawn.Despawntimer > 0 {
			result, err = c.Database.Instance.Exec("UPDATE spawngroup SET mindelay = ?, despawn_timer = ? WHERE npcid = ? AND spawngroupid = ?", spawn.Mindelay, spawn.Despawntimer, spawn.Npcid, spawn.Spawngroupid)
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
	}
	fmt.Println("Updated", totalRemoved, " DB entries related to", focus, "in spawnentry and spawngroup successfully.")

	return
}
