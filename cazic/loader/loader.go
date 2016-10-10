package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/goeq/npc"
	taknpc "github.com/xackery/goeq/takp/npc"
	"os"
	"reflect"
)

func main() {
	f, err := os.Create("../npctypes.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
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

	outStr := `
package cazic
import (
	"github.com/xackery/goeq/npc"
	"database/sql"
	)

var npcids []npc.NpcTypes = []npc.NpcTypes{
`
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

		m := mapFields(&pnpc)
		outStr += "npc.NpcTypes{"
		for k, v := range m {
			switch v.(type) {
			case string:
				outStr = fmt.Sprintf("%s%s: \"%s\", ", outStr, k, v)
			case sql.NullString:
				outStr = fmt.Sprintf("%s%s: sql.NullString{String: \"%s\"}, ", outStr, k, v.(sql.NullString).String)
			case int:

				outStr = fmt.Sprintf("%s%s: %d, ", outStr, k, v)
			case sql.NullInt64:
				outStr = fmt.Sprintf("%s%s: sql.NullInt64{Int64: %d}, ", outStr, k, v.(sql.NullInt64).Int64)
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
	return
}

type M map[string]interface{} // just an alias

func mapFields(x *npc.NpcTypes) M {
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
