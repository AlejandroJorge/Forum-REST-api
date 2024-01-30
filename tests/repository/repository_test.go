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
	config.InitializeAll()
	t.Run()
}

func fixWorkingDir() {
	currentDir, err := os.Getwd()
	util.PanicIfError(err)
	workingDir := path.Dir(path.Dir(currentDir))
	util.SetWorkingDir(workingDir)
}
