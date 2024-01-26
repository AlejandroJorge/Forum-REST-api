package config

import (
	"os"

	"github.com/AlejandroJorge/forum-rest-api/util"
)

var workingDir string = ""

func GetWorkingDir() string {
	if workingDir == "" {
		wd, err := os.Getwd()
		util.PanicIfError(err)

		workingDir = wd
	}
	return workingDir
}

func SetWorkingDir(path string) {
	workingDir = path
}
