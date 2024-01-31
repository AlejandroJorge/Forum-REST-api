package endpoints

import (
	"testing"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/tests"
)

func TestMain(t *testing.M) {
	tests.FixWorkingDir()
	config.InitializeAll()
	tests.RunMockSQLiteMigration()
	t.Run()
}
