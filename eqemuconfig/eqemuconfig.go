package eqemuconfig

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Config struct {
	Shortname string   `xml:"world>shortname"`
	Longame   string   `xml:"world>longname"`
	Database  Database `xml:"database"`
	QuestsDir string   `xml:"directories>quests"`
}

type Database struct {
	Host     string `xml:"host"`
	Port     string `xml:"port"`
	Username string `xml:"username"`
	Password string `xml:"password"`
	Db       string `xml:"db"`
}

func LoadConfig() (config Config, err error) {
	f, err := os.Open("eqemu_config.xml")
	if err != nil {
		err = fmt.Errorf("Error opening config: %s", err.Error())
		return
	}

	dec := xml.NewDecoder(f)
	err = dec.Decode(&config)
	if err != nil {
		err = fmt.Errorf("Error decoding config: %s", err.Error())
		return
	}
	if config.QuestsDir == "" {
		config.QuestsDir = "quests"
	}
	return
}
