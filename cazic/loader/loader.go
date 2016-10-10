package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/goeq/npc"
	"github.com/xackery/goeq/spawn"
	taknpc "github.com/xackery/goeq/takp/npc"
	takspawn "github.com/xackery/goeq/takp/spawn"
	"os"
	"reflect"
)

func main() {
	db, err := sqlx.Open("mysql", fmt.Sprintf("root@tcp(127.0.0.1:3306)/eqmac?charset=utf8&parseTime=true"))
	if err != nil {
		fmt.Println("Error connecting to DB:", err.Error())
		os.Exit(1)
	}
	defer db.Close()
	generateNPCTypes(db, 48)
	generateSpawngroups(db, 48)
	return
}

func generateNPCTypes(db *sqlx.DB, zoneid int) (err error) {
	f, err := os.Create("../npctypes.go")
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
package cazic
import (
	"github.com/xackery/goeq/npc"
	"database/sql"
	)

var npcids []npc.NpcTypes = []npc.NpcTypes{
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

func generateSpawngroups(db *sqlx.DB, zoneid int) (err error) {
	f, err := os.Create("../spawngroups.go")
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
package cazic
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

	f, err = os.Create("../spawnentries.go")
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
package cazic
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

	f, err = os.Create("../spawn.go")
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
package cazic
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

type M map[string]interface{} // just an alias

func mapFields(x interface{}) M {
	o := make(M)
	v := reflect.ValueOf(x).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := t.FieldByIndex([]int{i})
		// skip unexported fields
		if f.PkgPath != "" {
			continue
		}
		o[f.Name] = v.FieldByIndex([]int{i}).Interface()
	}
	return o
}
