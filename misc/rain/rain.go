//Disable rain and snow
package rain

import (
	"fmt"
	"github.com/xackery/eqcleanup/tools/eqdb"
)

var focus = "rain"

func Clean(args ...string) (err error) {
	db, err := eqdb.Load()
	if err != nil {
		return
	}
	result, err := db.Exec("UPDATE zone SET rain_chance1 = 0, rain_chance2 = 0, rain_chance3 = 0, rain_chance4 = 0, rain_duration1 = 0, rain_duration2 = 0, rain_duration3 = 0, rain_duration4 = 0, snow_chance1 = 0, snow_chance2 = 0, snow_chance3 = 0, snow_chance4 = 0, snow_duration1 = 0, snow_duration2 = 0, snow_duration3 = 0, snow_duration4 = 0")
	if err != nil {
		fmt.Println("Error disabling rain and snow: ", err.Error())
		return
	}
	affect, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected for rain & snow", err.Error())
		return
	}

	fmt.Println("Updated", affect, " DB entries related to rain and snow successfully.")

	return
}
