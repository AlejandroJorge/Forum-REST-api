package util

import (
	"os"
)

var workingDir string = ""

func GetWorkingDir() string {
	if workingDir == "" {
		wd, err := os.Getwd()
		PanicIfError(err)

		workingDir = wd
	}
	return workingDir
}

func SetWorkingDir(path string) {
	workingDir = path
}
