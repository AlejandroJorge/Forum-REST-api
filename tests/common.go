package tests

import (
	"os"
	"path"
	"testing"

	"github.com/AlejandroJorge/forum-rest-api/util"
)

func FixWorkingDir() {
	currentDir, err := os.Getwd()
	util.PanicIfError(err)
	workingDir := path.Dir(path.Dir(currentDir))
	util.SetWorkingDir(workingDir)
}

func EndTestIfError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Unexpected error, weren't testing for this: %v", err)
	}
}

func AssertEqu(expected interface{}, got interface{}, t *testing.T) {
	if expected != got {
		t.Errorf("Expected '%v', got '%v'", expected, got)
	}
}
