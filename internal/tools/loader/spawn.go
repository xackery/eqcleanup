package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/goeq/spawn"
	takspawn "github.com/xackery/goeq/takp/spawn"
	"os"
)

func generateSpawngroups(db *sqlx.DB) (err error) {
	f, err := os.Create("../../zone/" + zonename + "/spawngroups.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	rows, err := db.Queryx(fmt.Sprintf("SELECT * FROM spawngroup where id <=  %d and id >= %d", zoneid*1000+999, zoneid*1000))
	if err != nil {
		fmt.Println("Error exec select:", err.Error())
		os.Exit(1)
	}

	outStr := `
package ` + zonename + `
import (
	"github.com/xackery/goeq/spawn"
	"database/sql"
	)

var spawngroups []spawn.SpawnGroup = []spawn.SpawnGroup{
`
	spawngroupCount := 0
	spawngroups := []*spawn.SpawnGroup{}

	for rows.Next() {
		spawngroupCount++
		sg := spawn.SpawnGroup{}
		if err = rows.StructScan(&sg); err != nil {
			fmt.Println("Error getting rows", err.Error())
			os.Exit(1)
		}
		//special exception, CT spawngroup improper.
		if sg.Name.String == "cshomeAsherah_the_Torch_Bearer00083304628" {
			continue
		}

		m := mapFields(&sg)
		spawngroups = append(spawngroups, &sg)

		outStr += "spawn.SpawnGroup{"
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
	fmt.Println("Generated", spawngroupCount, "spawngroups")

	//build a string for in clause
	spawngroupstring := "("
	for _, sg := range spawngroups {
		spawngroupstring = fmt.Sprintf("%s%d, ", spawngroupstring, sg.Id.Int64)
	}
	spawngroupstring = spawngroupstring[0:len(spawngroupstring)-2] + ")"

	f, err = os.Create("../../zone/" + zonename + "/spawnentries.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	//grab spawnentries

	rows, err = db.Queryx(fmt.Sprintf("SELECT * FROM spawnentry where spawngroupid in %s", spawngroupstring))
	if err != nil {
		fmt.Println("Error exec select:", err.Error())
		os.Exit(1)
	}

	outStr = `
package ` + zonename + `
import (
	"github.com/xackery/goeq/spawn"
	//"database/sql"
	)

var spawnentries []spawn.SpawnEntry = []spawn.SpawnEntry{
`
	spawngroupCount = 0

	for rows.Next() {
		spawngroupCount++
		tse := takspawn.SpawnEntry{}
		if err = rows.StructScan(&tse); err != nil {
			fmt.Println("Error getting rows for spawnentry", err.Error())
			os.Exit(1)
		}

		se := spawn.SpawnEntry{}
		if err = se.DecodeFromTakp(&tse); err != nil {
			fmt.Println(err.Error())
			return
		}
		m := mapFields(&se)

		outStr += "spawn.SpawnEntry{"
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
	fmt.Println("Generated", spawngroupCount, "spawnentries")

	f, err = os.Create("../../zone/" + zonename + "/spawn.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	//grab spawnentries

	rows, err = db.Queryx(fmt.Sprintf("SELECT * FROM spawn2 where spawngroupid in %s", spawngroupstring))
	if err != nil {
		fmt.Println("Error exec select:", err.Error())
		os.Exit(1)
	}

	outStr = `
package ` + zonename + `
import (
	"github.com/xackery/goeq/spawn"
	"database/sql"
	)

var spawns []spawn.Spawn2 = []spawn.Spawn2{
`
	spawngroupCount = 0

	for rows.Next() {
		spawngroupCount++
		ts := takspawn.Spawn2{}
		if err = rows.StructScan(&ts); err != nil {
			fmt.Println("Error getting rows for spawn2", err.Error())
			os.Exit(1)
		}

		se := spawn.Spawn2{}
		if err = se.DecodeFromTakp(&ts); err != nil {
			fmt.Println(err.Error())
			return
		}
		m := mapFields(&se)

		if ts.Pathgrid != 0 {
			isUnique := true
			for _, gridid := range gridids {
				if gridid == ts.Pathgrid {
					isUnique = false
					break
				}
			}
			if isUnique {
				gridids = append(gridids, ts.Pathgrid)
			}
		}

		outStr += "spawn.Spawn2{"
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
	fmt.Println("Generated", spawngroupCount, "spawn2")

	return
}
