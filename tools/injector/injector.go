package injector

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
)

func MapFields(x interface{}) (m map[string]interface{}) {
	m = make(map[string]interface{})
	v := reflect.ValueOf(x).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := t.FieldByIndex([]int{i})
		// skip unexported fields
		if f.PkgPath != "" {
			continue
		}
		if f.Tag.Get("db") == "" {
			continue
		}
		m[f.Tag.Get("db")] = v.FieldByIndex([]int{i}).Interface()

	}
	return
}

func PrepareInsertString(entry interface{}, table string) (insertString string) {
	q := "INSERT INTO " + table + " ("
	m := MapFields(entry)
	fields := ""
	for k, _ := range m {
		q += fmt.Sprintf("%s, ", k)
		fields += fmt.Sprintf(":%s, ", k)
	}
	fields = fields[0 : len(fields)-2]
	q = q[0:len(q)-2] + ") VALUES (" + fields + ")"

	insertString = q
	return
}

func Inject(db *sqlx.DB, entries []interface{}, table string) (totalChanged int64, err error) {
	//SPAWNENTRY
	q := PrepareInsertString(entries[0], "spawnentry")

	//fmt.Println(q)
	totalChanged = 0
	for _, sg := range entries {
		if _, err = db.NamedExec(q, &sg); err != nil {
			return
		}
		totalChanged++
	}
	return
}
