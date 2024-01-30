package router

import (
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/config"
	"github.com/AlejandroJorge/forum-rest-api/delivery/controller"
	"github.com/AlejandroJorge/forum-rest-api/delivery/middleware"
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
	router.Use(middleware.Logger)

	initializeUserRoutes(router)
	initializeProfileRoutes(router)
	initializePostRoutes(router)
	initializeCommentRoutes(router)
}

func initializeUserRoutes(router *mux.Router) {
	repository := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	service := service.NewUserService(repository)
	controller := controller.NewUserController(service)

	router.HandleFunc("/api/v1/users",
		controller.Create).Methods("POST")

	router.HandleFunc("/api/v1/users/login",
		controller.Login).Methods("POST")

	router.HandleFunc("/api/v1/users/{id:[0-9]+}",
		middleware.Auth(controller.Get)).Methods("GET")

	router.HandleFunc("/api/v1/users/{id:[0-9]+}",
		middleware.Auth(controller.Update)).Methods("PUT")

	router.HandleFunc("/api/v1/users/{id:[0-9]+}",
		middleware.Auth(controller.Delete)).Methods("DELETE")
}

func initializeProfileRoutes(router *mux.Router) {
	repository := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())
	service := service.NewProfileService(repository)
	controller := controller.NewProfileController(service)

	router.HandleFunc("/api/v1/profiles", controller.Create).Methods("POST")
	router.HandleFunc("/api/v1/profiles/{id:[0-9]+}", controller.GetByID).Methods("GET")
	router.HandleFunc("/api/v1/profiles/{tagname:[a-zA-Z][a-zA-Z0-9]*}", controller.GetByTagName).Methods("GET")
	router.HandleFunc("/api/v1/profiles/{id:[0-9]+}/followers", controller.GetFollowersByID).Methods("GET")
	router.HandleFunc("/api/v1/profiles/{tagname:[a-zA-Z][a-zA-Z0-9]*}/followers", controller.GetFollowersByTagName).Methods("GET")
	router.HandleFunc("/api/v1/profiles/{id:[0-9]+}/follows", controller.GetFollowsByID).Methods("GET")
	router.HandleFunc("/api/v1/profiles/{tagname:[a-zA-Z][a-zA-Z0-9]*}/follows", controller.GetFollowsByTagName).Methods("GET")
	router.HandleFunc("/api/v1/profiles/{id:[0-9]+}", controller.Update).Methods("PUT")
	router.HandleFunc("/api/v1/profiles/follows", controller.CreateFollow).Methods("POST")
	router.HandleFunc("/api/v1/profiles/follows", controller.DeleteFollow).Methods("DELETE")
	router.HandleFunc("/api/v1/profiles/{id:[0-9]+}", controller.Delete).Methods("DELETE")
}

func initializePostRoutes(router *mux.Router) {
	repository := repository.NewSQLitePostRepository(config.SQLiteDatabase())
	service := service.NewPostService(repository)
	controller := controller.NewPostController(service)

	router.HandleFunc("/api/v1/posts", controller.Create).Methods("POST")
	router.HandleFunc("/api/v1/posts", controller.GetPopular).Methods("GET")
	router.HandleFunc("/api/v1/posts/{id:[0-9]+}", controller.GetByID).Methods("GET")
	router.HandleFunc("/api/v1/users/{id:[0-9]+}/posts", controller.GetByUserID).Methods("GET")
	router.HandleFunc("/api/v1/posts/{id:[0-9]+}", controller.Update).Methods("PUT")
	router.HandleFunc("/api/v1/posts/likes", controller.AddLike).Methods("POST")
	router.HandleFunc("/api/v1/posts/likes", controller.DeleteLike).Methods("DELETE")
	router.HandleFunc("/api/v1/posts/{id:[0-9]+}", controller.Delete).Methods("DELETE")
}

func initializeCommentRoutes(router *mux.Router) {
	repository := repository.NewSQLiteCommentRepository(config.SQLiteDatabase())
	service := service.NewCommentService(repository)
	controller := controller.NewCommentController(service)

	router.HandleFunc("/api/v1/comments", controller.Create).Methods("POST")
	router.HandleFunc("/api/v1/comments/{id:[0-9]+}", controller.GetByID).Methods("GET")
	router.HandleFunc("/api/v1/users/{id:[0-9]+}/comments", controller.GetByUser).Methods("GET")
	router.HandleFunc("/api/v1/posts/{id:[0-9]+/comments}", controller.GetByPost).Methods("GET")
	router.HandleFunc("/api/v1/comments/{id:[0-9]+}", controller.Update).Methods("PUT")
	router.HandleFunc("/api/v1/comments/likes", controller.CreateLike).Methods("POST")
	router.HandleFunc("/api/v1/comments/likes", controller.DeleteLike).Methods("DELETE")
	router.HandleFunc("/api/v1/comments/{id:[0-9]+}", controller.Delete).Methods("DELETE")
}
