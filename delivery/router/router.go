package router

import (
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/delivery/controller"
	"github.com/AlejandroJorge/forum-rest-api/repository"
	"github.com/AlejandroJorge/forum-rest-api/service"
	"github.com/gorilla/mux"
)

var mainRouter http.Handler

func AppRouter() http.Handler {
	if mainRouter == nil {
		newRouter := mux.NewRouter()
		initializeRouter(newRouter)
		mainRouter = newRouter
	}

	return mainRouter
}

func initializeRouter(router *mux.Router) {
	initializeUserRoutes(router)
	initializeProfileRoutes(router)
	initializePostRoutes(router)
	initializeCommentRoutes(router)
}

func initializeUserRoutes(router *mux.Router) {
	repository := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	service := service.NewUserService(repository)
	controller := controller.NewUserController(service)

	router.HandleFunc("/api/v1/user", controller.Create).Methods("POST")
	router.HandleFunc("/api/v1/user/{id}", controller.Get).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", controller.Update).Methods("PUT")
	router.HandleFunc("/api/v1/user/{id}", controller.Delete).Methods("DELETE")
}

func initializeProfileRoutes(router *mux.Router) {

}

func initializePostRoutes(router *mux.Router) {

}

func initializeCommentRoutes(router *mux.Router) {

}
