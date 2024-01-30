package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/delivery/router"
)

func main() {
	config.Initialize()

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	fmt.Printf("Listening on http://localhost:%s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router.AppRouter())
}
