package quest

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/xackery/eqemuconfig"
)

// Quest wraps quest operations
type Quest struct {
	Directory string
}

// New creates a new quest instance, or an error if it is not found
func New(config *eqemuconfig.Config) (q *Quest, err error) {
	q = &Quest{}
	questDir := config.QuestsDir
	if config.QuestsDir == "" {
		questDir = "./quests"
	}
	fi, err := os.Stat(questDir)
	if err != nil {
		return
	}
	if !fi.IsDir() {
		err = fmt.Errorf("quest directory %s is not a directory", questDir)
		return
	}
	return
}

// ZoneDelete removes quests for a specified zone
func (q *Quest) ZoneDelete(zone string) (removeCount int64, err error) {
	path := q.Directory + "/" + zone

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

// Delete removes quests based on filepaths provided
func (q *Quest) Delete(filePaths []string) (delCount int, err error) {

	for _, fileName := range filePaths {
		curFile := q.Directory + "/" + fileName
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
