package logging

import (
	"log"

	"github.com/AlejandroJorge/forum-rest-api/config"
)

func LogSetup() {
	configParams := config.GetParams()
	msg := `
	[CONFIG] Configuration:
	[CONFIG] %s
	[CONFIG] %s
	[CONFIG] %d
	[CONFIG] %s
	`

	log.Printf(msg, configParams.DbFolderName, configParams.DbFileName, configParams.Port, configParams.AuthSecret)
}
