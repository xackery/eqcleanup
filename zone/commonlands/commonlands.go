package commonlands

import (
	//"database/sql"
	"fmt"
	//"github.com/xackery/eqcleanup/tools/eqconfig"
	//"github.com/xackery/eqcleanup/tools/eqdb"
	"strings"
)

var focus = "commonlands"
var zonename = "commonlands"

func Clean(args ...string) (err error) {
	//db, err := eqdb.Load()
	if err != nil {
		return
	}
	//config, err := eqconfig.Load()
	if err != nil {
		return
	}
	option := ""

	fmt.Println("Running this operation will update zone connections from using the revamped commonlands to using east and west commonlands.")
	fmt.Printf("Continue? [Y/n]")
	for {
		_, err = fmt.Scanln(&option)
		if err != nil {
			return
		}
		if strings.ToLower(option) == "n" {
			fmt.Println("Returning to menu")
			return
		}
		if strings.ToLower(option) == "y" {
			break
		}
	}
	fmt.Println("Updating...")
	return
}
