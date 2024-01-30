package tests

import (
	"os"
	"path"

	"github.com/AlejandroJorge/forum-rest-api/util"
)

func FixWorkingDir() {
	currentDir, err := os.Getwd()
	util.PanicIfError(err)
	workingDir := path.Dir(path.Dir(currentDir))
	util.SetWorkingDir(workingDir)
}
