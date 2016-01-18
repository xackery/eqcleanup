package eqemuconfig

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Config struct {
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

	return
}
