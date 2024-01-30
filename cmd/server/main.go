package main

import (
	"fmt"
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/delivery/router"
	"github.com/AlejandroJorge/forum-rest-api/logging"
)

func main() {
	config.InitializeAll()

	logging.LogSetup()
	fmt.Printf("Listening on http://localhost:%d", config.GetParams().Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.GetParams().Port), router.AppRouter())

}
