package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dixonwille/wmenu"
	"github.com/xackery/eqcleanup/client"
	"github.com/xackery/eqcleanup/internal/core"
	"github.com/xackery/eqcleanup/internal/era"
)

// Version of eqcleanup
var Version = "1.0"

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() (err error) {
	fmt.Println("Welcome to EQ Cleanup", Version, "https://github.com/xackery/eqcleanup")
	fmt.Println("This tool will MODIFY your database. Please back up before doing any operation.")
	fmt.Println("---")

	c, err := client.New()
	if err != nil {
		return
	}
	Menu(c)
	return
}

// Menu is the main menu of eqclient
func Menu(c *client.Client) (err error) {
	ok := true
	menu := wmenu.NewMenu("Choose a category")
	menu.Option("Era - Out of era cleanup", c, false, era.Menu)
	menu.Option("Core - Administrative overhauls", c, false, core.Menu)
	//menu.Option("Misc - Misc edits, ex. disabling rainfall", nil, false, misc.Menu)
	menu.Option("Exit", nil, false, func(opt wmenu.Opt) error {
		ok = false
		fmt.Println("Have a nice day!")
		return nil
	})

	for ok {
		err = menu.Run()
		if err == nil {
			continue
		}

		if strings.Contains(err.Error(), "Invalid") {
			err = nil
			continue
		}
		return
	}
	return
}
