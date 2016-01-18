package main

import (
	//"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/eqemuconfig"
	"github.com/xackery/eqcleanup/soulbinder"
	"os"
)

func main() {
	//Load config
	config, err := eqemuconfig.LoadConfig()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//Connect to DB
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", config.Database.Username, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Db))
	if err != nil {
		fmt.Println("Error connecting to DB:", err.Error())
		os.Exit(1)
	}

	//Cleanup soulbinder
	err = soulbinder.Clean(db)
	if err != nil {
		fmt.Println("Error removing soulbinders:", err.Error())
		os.Exit(1)
	}

}
