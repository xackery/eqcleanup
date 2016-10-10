package eqdb

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/tools/eqconfig"
)

var db *sqlx.DB

func Load() (d *sqlx.DB, err error) {
	config, err := eqconfig.Load()
	if err != nil {
		return
	}

	if db != nil {
		d = db
		return
	}
	db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Db))
	if err != nil {
		err = fmt.Errorf("error connecting to db: %s", err.Error())
		return
	}
	d = db
	return
}
