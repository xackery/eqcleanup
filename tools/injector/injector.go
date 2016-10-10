package injector

import (
	"fmt"
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

func PrepareInsertString(entry interface{}, table string) (insertString string, err error) {
	q := "INSERT INTO " + table + " ("
	m := MapFields(entry)
	fields := ""
	for k := range m {
		if table == "lootdrop" && k == "id" {
			continue
		}
		q += fmt.Sprintf("%s, ", k)
		fields += fmt.Sprintf(":%s, ", k)
	}
	if len(fields) < 4 {
		err = fmt.Errorf("Size of %s too small", table)
		return
	}
	fields = fields[0 : len(fields)-2]
	q = q[0:len(q)-2] + ") VALUES (" + fields + ")"

	insertString = q
	return
}
