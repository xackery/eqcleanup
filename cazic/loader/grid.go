package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/goeq/grid"
	"os"
)

func generateGrids(db *sqlx.DB) (err error) {

	gridstring := "("
	for _, gridid := range gridids {
		gridstring += fmt.Sprintf("%d, ", gridid)
	}
	gridstring = gridstring[0:len(gridstring)-2] + ")"

	f, err := os.Create("../grid.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	rows, err := db.Queryx(fmt.Sprintf("SELECT * FROM grid WHERE zoneid = %d AND id IN %s", zoneid, gridstring))
	if err != nil {
		fmt.Println("Error exec select:", err.Error())
		os.Exit(1)
	}

	outStr := `
package cazic
import (
	"github.com/xackery/goeq/grid"
	)

var grids []grid.Grid = []grid.Grid{
`

	count := 0
	for rows.Next() {
		count++
		gr := grid.Grid{}
		if err = rows.StructScan(&gr); err != nil {
			fmt.Println("Error getting rows", err.Error())
			os.Exit(1)
		}

		m := mapFields(&gr)

		outStr += "grid.Grid{"
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
	fmt.Println("Generated", count, "grids")

	f, err = os.Create("../gridentries.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	rows, err = db.Queryx(fmt.Sprintf("SELECT * FROM grid_entries WHERE zoneid = %d AND gridid IN %s", zoneid, gridstring))
	if err != nil {
		fmt.Println("Error exec select:", err.Error())
		os.Exit(1)
	}

	outStr = `
package cazic
import (
	"github.com/xackery/goeq/grid"
	)

var gridentries []grid.GridEntries = []grid.GridEntries{
`

	count = 0
	for rows.Next() {
		count++
		gr := grid.GridEntries{}
		if err = rows.StructScan(&gr); err != nil {
			fmt.Println("Error getting rows", err.Error())
			os.Exit(1)
		}

		m := mapFields(&gr)

		outStr += "grid.GridEntries{"
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
	fmt.Println("Generated", count, "gridentries")
	return
}
