package main

import (
	//"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/goeq/npc"
	taknpc "github.com/xackery/goeq/takp/npc"
	"os"
)

func main() {
	db, err := sqlx.Open("mysql", fmt.Sprintf("root@tcp(127.0.0.1:3306)/eqmac?charset=utf8&parseTime=true"))
	if err != nil {
		fmt.Println("Error connecting to DB:", err.Error())
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Queryx("SELECT * FROM npc_types where id <= 48999 AND id >= 48000")
	if err != nil {
		fmt.Println("Error exec select:", err.Error())
		os.Exit(1)
	}

	for rows.Next() {
		tnpc := taknpc.NpcTypes{}
		if err = rows.StructScan(&tnpc); err != nil {
			fmt.Println("Error getting rows", err.Error())
			os.Exit(1)
		}
		//fmt.Println(tnpc)
		pnpc := npc.NpcTypes{}

		if err = pnpc.DecodeFromTakp(&tnpc); err != nil {
			fmt.Println(err.Error())
			return
		}
		return
	}
	return
}
