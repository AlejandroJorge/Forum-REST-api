package main

import (
	"fmt"
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/delivery/router"
	"github.com/AlejandroJorge/forum-rest-api/logging"
	"github.com/gorilla/handlers"
)

func main() {
	config.InitializeAll()
	logging.LogSetup()

	port := config.GetParams().Port
	router := router.AppRouter(config.SQLiteDatabase())

	http.ListenAndServe(fmt.Sprintf(":%d", port), handlers.CORS()(router))
}
