package quest

import (
	"fmt"
	"github.com/xackery/eqemuconfig"
	"io/ioutil"
	"os"
	"strings"
)

func RemoveAllQuestsForZone(config *eqemuconfig.Config, zone string) (removeCount int64, err error) {
	if config.QuestsDir == "" {
		err = fmt.Errorf("Quests directory is not set.")
		return
	}
	path := config.QuestsDir + "/" + zone

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}
	removeCount = int64(len(files))
	for _, file := range files {

		if err = os.Remove(fmt.Sprintf("%s/%s", path, file.Name())); err != nil {
			err = fmt.Errorf("error removing %s: %s", file.Name(), err.Error())
			return
		}
	}

	return
}

func Remove(config *eqemuconfig.Config, filePaths []string) (delCount int, err error) {
	if config.QuestsDir == "" {
		err = fmt.Errorf("Quests directory is not set.")
		return
	}
	for _, fileName := range filePaths {
		curFile := config.QuestsDir + "/" + fileName
		_, err = os.Stat(curFile)

		if err != nil {

			if os.IsNotExist(err) || strings.Contains(err.Error(), "no such file") {
				err = nil
				continue
			}
			fmt.Printf("Error finding %s: %s", curFile, err.Error())
			continue
		}

		err = os.Remove(curFile)
		if err != nil {
			fmt.Printf("Error deleting %s: %s", curFile, err.Error())
			continue
		}
		delCount++
	}
	err = nil
	return
}
