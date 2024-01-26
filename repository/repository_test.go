package repository

import (
	"os"
	"path"
	"testing"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/util"
)

func TestMain(t *testing.M) {
	fixWorkingDir()
	config.Initialize()
	t.Run()
}

func fixWorkingDir() {
	currentDir, err := os.Getwd()
	util.PanicIfError(err)
	workingDir := path.Dir(currentDir)
	config.SetWorkingDir(workingDir)
}
