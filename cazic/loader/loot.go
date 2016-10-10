package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/goeq/loot"
	takloot "github.com/xackery/goeq/takp/loot"
	"os"
)

func generateLoot(db *sqlx.DB) (err error) {
	f, err := os.Create("../loottable.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	lootstring := "("
	for _, lootid := range lootids {
		lootstring += fmt.Sprintf("%d, ", lootid)
	}
	lootstring = lootstring[0:len(lootstring)-2] + ")"

	rows, err := db.Queryx(fmt.Sprintf("SELECT * FROM loottable where id in %s", lootstring))
	if err != nil {
		fmt.Println("Error exec select:", err.Error())
		os.Exit(1)
	}

	outStr := `
package cazic
import (
	"github.com/xackery/goeq/loot"
	"database/sql"
	)

var loottables []loot.LootTable = []loot.LootTable{
`

	counter := 0

	for rows.Next() {
		counter++
		tl := takloot.LootTable{}
		if err = rows.StructScan(&tl); err != nil {
			fmt.Println("Error getting rows", err.Error())
			os.Exit(1)
		}
		l := loot.LootTable{}

		if err = l.DecodeFromTakp(&tl); err != nil {
			fmt.Println(err.Error())
			return
		}

		m := mapFields(&l)

		outStr += "loot.LootTable{"
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
	fmt.Println("Generated", counter, "loottable")

	f, err = os.Create("../loottableentries.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	lootstring = "("
	for _, lootid := range lootids {
		lootstring += fmt.Sprintf("%d, ", lootid)
	}
	lootstring = lootstring[0:len(lootstring)-2] + ")"

	rows, err = db.Queryx(fmt.Sprintf("SELECT * FROM loottable_entries where loottable_id in %s", lootstring))
	if err != nil {
		fmt.Println("Error exec select:", err.Error())
		os.Exit(1)
	}

	outStr = `
package cazic
import (
	"github.com/xackery/goeq/loot"
	//"database/sql"
	)

var loottableentries []loot.LootTableEntries = []loot.LootTableEntries{
`

	counter = 0

	for rows.Next() {
		counter++
		tl := takloot.LootTableEntries{}
		if err = rows.StructScan(&tl); err != nil {
			fmt.Println("Error getting rows", err.Error())
			os.Exit(1)
		}
		l := loot.LootTableEntries{}

		if err = l.DecodeFromTakp(&tl); err != nil {
			fmt.Println(err.Error())
			return
		}

		m := mapFields(&l)

		outStr += "loot.LootTableEntries{"
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
	fmt.Println("Generated", counter, "loottableentry")
	return
}
