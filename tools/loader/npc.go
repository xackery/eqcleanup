package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/goeq/npc"
	taknpc "github.com/xackery/goeq/takp/npc"
	"os"
)

func generateNPCTypes(db *sqlx.DB) (err error) {
	f, err := os.Create("../../zone/" + zonename + "/npctypes.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	rows, err := db.Queryx(fmt.Sprintf("SELECT * FROM npc_types where id <=  %d and id >= %d", zoneid*1000+999, zoneid*1000))
	if err != nil {
		fmt.Println("Error exec select:", err.Error())
		os.Exit(1)
	}

	outStr := `
package ` + zonename + `
import (
	"github.com/xackery/goeq/npc"
	"database/sql"
	)

var npctypes []npc.NpcTypes = []npc.NpcTypes{
`

	npcCount := 0
	for rows.Next() {
		npcCount++
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

		//Store loot tables for later, when loot gets generated
		if pnpc.Loottable_id != 0 {
			isUnique := true
			for _, lootid := range lootids {
				if lootid == pnpc.Loottable_id {
					isUnique = false
					break
				}
			}
			if isUnique {
				lootids = append(lootids, pnpc.Loottable_id)
			}
		}

		m := mapFields(&pnpc)

		outStr += "npc.NpcTypes{"
		for k, v := range m {
			switch v.(type) {
			case string:
				outStr = fmt.Sprintf("%s%s: \"%s\", ", outStr, k, v)
			case sql.NullString:
				outStr = fmt.Sprintf("%s%s: sql.NullString{String: \"%s\", Valid: true}, ", outStr, k, v.(sql.NullString).String)
			case int:

				outStr = fmt.Sprintf("%s%s: %d, ", outStr, k, v)
			case sql.NullInt64:
				outStr = fmt.Sprintf("%s%s: sql.NullInt64{Int64: %d, Valid: true}, ", outStr, k, v.(sql.NullInt64).Int64)
			case float64:
				outStr = fmt.Sprintf("%s%s: %f, ", outStr, k, v)
			default:
				fmt.Println("Unsupported type:", k, v)
				return
			}
		}
		outStr = outStr[0 : len(outStr)-2]
		outStr += "},\n"
	}
	outStr += "}"

	if _, err = f.WriteString(outStr); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Generated", npcCount, "npcs")
	return
}
