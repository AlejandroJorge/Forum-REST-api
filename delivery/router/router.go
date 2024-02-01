package router

import (
	"database/sql"
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/delivery/controller"
	"github.com/AlejandroJorge/forum-rest-api/delivery/middleware"
	"github.com/AlejandroJorge/forum-rest-api/repository"
	"github.com/AlejandroJorge/forum-rest-api/service"
	"github.com/gorilla/mux"
)

var mainRouter http.Handler

func AppRouter(db *sql.DB) http.Handler {
	if mainRouter == nil {
		newRouter := mux.NewRouter()
		initializeRouter(newRouter, db)
		mainRouter = newRouter
	}

	return mainRouter
}

func initializeRouter(router *mux.Router, db *sql.DB) {
	router.Use(middleware.Logger)

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	initializeUserRoutes(apiRouter, db)
	initializeProfileRoutes(apiRouter, db)
	initializePostRoutes(apiRouter, db)
	initializeCommentRoutes(apiRouter, db)
}

func initializeUserRoutes(router *mux.Router, db *sql.DB) {
	repository := repository.NewSQLiteUserRepository(db)
	service := service.NewUserService(repository)
	controller := controller.NewUserController(service)

	router.HandleFunc("/users",
		controller.Create).Methods("POST")

	router.HandleFunc("/users/login",
		controller.CheckCredentials).Methods("POST")

	router.HandleFunc("/users/{userid:[0-9]+}",
		middleware.Auth(controller.GetByID)).Methods("GET")

	router.HandleFunc("/users/{userid:[0-9]+}/email",
		middleware.Auth(controller.UpdateEmail)).Methods("PUT")

	router.HandleFunc("/users/{userid:[0-9]+}/password",
		middleware.Auth(controller.UpdatePassword)).Methods("PUT")

	router.HandleFunc("/users/{userid:[0-9]+}",
		middleware.Auth(controller.Delete)).Methods("DELETE")
}

func initializeProfileRoutes(router *mux.Router, db *sql.DB) {
	repository := repository.NewSQLiteProfileRepository(db)
	service := service.NewProfileService(repository)
	controller := controller.NewProfileController(service)

	router.HandleFunc("/users/{userid:[0-9]+}/profiles",
		middleware.Auth(controller.Create)).Methods("POST")

	router.HandleFunc("/profiles/{userid:[0-9]+}",
		controller.GetByUserID).Methods("GET")

	router.HandleFunc("/profiles/{tagname:[a-zA-Z][a-zA-Z0-9]*}",
		controller.GetByTagName).Methods("GET")

	router.HandleFunc("/profiles/{userid:[0-9]+}/followers",
		controller.GetFollowersByID).Methods("GET")

	router.HandleFunc("/profiles/{tagname:[a-zA-Z][a-zA-Z0-9]*}/followers",
		controller.GetFollowersByTagName).Methods("GET")

	router.HandleFunc("/profiles/{userid:[0-9]+}/follows",
		controller.GetFollowsByID).Methods("GET")

	router.HandleFunc("/profiles/{tagname:[a-zA-Z][a-zA-Z0-9]*}/follows",
		controller.GetFollowsByTagName).Methods("GET")

	router.HandleFunc("/profiles/{userid:[0-9]+}/tagname",
		middleware.Auth(controller.UpdateTagName)).Methods("PUT")

	router.HandleFunc("/profiles/{userid:[0-9]+}/displayname",
		middleware.Auth(controller.UpdateDisplayName)).Methods("PUT")

	router.HandleFunc("/profiles/{userid:[0-9]+}/picturepath",
		middleware.Auth(controller.UpdatePicturePath)).Methods("PUT")

	router.HandleFunc("/profiles/{userid:[0-9]+}/backgroundpath",
		middleware.Auth(controller.UpdateBackgroundPath)).Methods("PUT")

	router.HandleFunc("/profiles/{userid:[0-9]+}/follows/{followedid:[0-9]+}",
		middleware.Auth(controller.AddFollow)).Methods("POST")

	router.HandleFunc("/profiles/{userid:[0-9]+}/follows/{followedid:[0-9]+}",
		middleware.Auth(controller.DeleteFollow)).Methods("DELETE")

	router.HandleFunc("/profiles/{userid:[0-9]+}",
		middleware.Auth(controller.Delete)).Methods("DELETE")
}

func initializePostRoutes(router *mux.Router, db *sql.DB) {
	repository := repository.NewSQLitePostRepository(db)
	service := service.NewPostService(repository)
	controller := controller.NewPostController(service)

	router.HandleFunc("/users/{userid:[0-9]+}/posts",
		middleware.Auth(controller.Create)).Methods("POST")

	router.HandleFunc("/posts/today",
		controller.GetPopularToday).Methods("GET")

	router.HandleFunc("/posts/week",
		controller.GetPopularLastWeek).Methods("GET")

	router.HandleFunc("/posts/month",
		controller.GetPopularLastMonth).Methods("GET")

	router.HandleFunc("/posts/alltime",
		controller.GetPopularAllTime).Methods("GET")

	router.HandleFunc("/posts/{postid:[0-9]+}",
		controller.GetByID).Methods("GET")

	router.HandleFunc("/users/{userid:[0-9]+}/posts",
		controller.GetByUser).Methods("GET")

	router.HandleFunc("/users/{userid:[0-9]+}/posts/{postid:[0-9]+}/title",
		middleware.Auth(controller.UpdateTitle)).Methods("PUT")

	router.HandleFunc("/users/{userid:[0-9]+}/posts/{postid:[0-9]+}/description",
		middleware.Auth(controller.UpdateDescription)).Methods("PUT")

	router.HandleFunc("/users/{userid:[0-9]+}/posts/{postid:[0-9]+}/content",
		middleware.Auth(controller.UpdateContent)).Methods("PUT")

	router.HandleFunc("/users/{userid:[0-9]+}/posts/{postid:[0-9]+}/likes",
		middleware.Auth(controller.AddLike)).Methods("POST")

	router.HandleFunc("/users/{userid:[0-9]+}/posts/{postid:[0-9]+}/likes",
		middleware.Auth(controller.DeleteLike)).Methods("DELETE")

	router.HandleFunc("/users/{userid:[0-9]+}/posts/{postid:[0-9]+}",
		middleware.Auth(controller.Delete)).Methods("DELETE")
}

func initializeCommentRoutes(router *mux.Router, db *sql.DB) {
	repository := repository.NewSQLiteCommentRepository(db)
	service := service.NewCommentService(repository)
	controller := controller.NewCommentController(service)

	router.HandleFunc("/users/{userid:[0-9]+}/comments",
		middleware.Auth(controller.Create)).Methods("POST")

	router.HandleFunc("/comments/{commentid:[0-9]+}",
		controller.GetByID).Methods("GET")

	router.HandleFunc("/users/{userid:[0-9]+}/comments",
		controller.GetByUser).Methods("GET")

	router.HandleFunc("/posts/{postid:[0-9]+}/comments",
		controller.GetByPost).Methods("GET")

	router.HandleFunc("/users/{userid:[0-9]+}/comments/{commentid:[0-9]+}",
		middleware.Auth(controller.UpdateContent)).Methods("PUT")

	router.HandleFunc("/users/{userid:[0-9]+}/comments/{commentid:[0-9]+}/likes",
		middleware.Auth(controller.AddLike)).Methods("POST")

	router.HandleFunc("/users/{userid:[0-9]+}/comments/{commentid:[0-9]+}/likes",
		middleware.Auth(controller.DeleteLike)).Methods("DELETE")

	router.HandleFunc("/users/{userid:[0-9]+}/comments/{commentid:[0-9]+}",
		middleware.Auth(controller.Delete)).Methods("DELETE")
}
