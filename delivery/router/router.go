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

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	initializeUserRoutes(apiRouter)
	initializeProfileRoutes(apiRouter)
	//initializePostRoutes(apiRouter)
	//initializeCommentRoutes(apiRouter)
}

func initializeUserRoutes(router *mux.Router) {
	repository := repository.NewSQLiteUserRepository(config.SQLiteDatabase())
	service := service.NewUserService(repository)
	controller := controller.NewUserController(service)

	router.HandleFunc("/users",
		controller.Create).Methods("POST")

	router.HandleFunc("/users/login",
		controller.CheckCredentials).Methods("POST")

	router.HandleFunc("/users/{id:[0-9]+}",
		middleware.Auth(controller.GetByID)).Methods("GET")

	router.HandleFunc("/users/{id:[0-9]+}/email",
		middleware.Auth(controller.UpdateEmail)).Methods("PUT")

	router.HandleFunc("/users/{id:[0-9]+}/password",
		middleware.Auth(controller.UpdatePassword)).Methods("PUT")

	router.HandleFunc("/users/{id:[0-9]+}",
		middleware.Auth(controller.Delete)).Methods("DELETE")
}

func initializeProfileRoutes(router *mux.Router) {
	repository := repository.NewSQLiteProfileRepository(config.SQLiteDatabase())
	service := service.NewProfileService(repository)
	controller := controller.NewProfileController(service)

	router.HandleFunc("/profiles",
		controller.Create).Methods("POST")

	router.HandleFunc("/profiles/{id:[0-9]+}",
		controller.GetByUserID).Methods("GET")

	router.HandleFunc("/profiles/{tagname:[a-zA-Z][a-zA-Z0-9]*}",
		controller.GetByTagName).Methods("GET")

	router.HandleFunc("/profiles/{id:[0-9]+}/followers",
		controller.GetFollowersByID).Methods("GET")

	router.HandleFunc("/profiles/{tagname:[a-zA-Z][a-zA-Z0-9]*}/followers",
		controller.GetFollowersByTagName).Methods("GET")

	router.HandleFunc("/profiles/{id:[0-9]+}/follows",
		controller.GetFollowsByID).Methods("GET")

	router.HandleFunc("/profiles/{tagname:[a-zA-Z][a-zA-Z0-9]*}/follows",
		controller.GetFollowsByTagName).Methods("GET")

	router.HandleFunc("/profiles/{id:[0-9]+}/tagname",
		controller.UpdateTagName).Methods("PUT")

	router.HandleFunc("/profiles/{id:[0-9]+}/displayname",
		controller.UpdateDisplayName).Methods("PUT")

	router.HandleFunc("/profiles/{id:[0-9]+}/picturepath",
		controller.UpdatePicturePath).Methods("PUT")

	router.HandleFunc("/profiles/{id:[0-9]+}/backgroundpath",
		controller.UpdateBackgroundPath).Methods("PUT")

	router.HandleFunc("/profiles/follows",
		controller.AddFollow).Methods("POST")

	router.HandleFunc("/profiles/follows",
		controller.DeleteFollow).Methods("DELETE")

	router.HandleFunc("/profiles/{id:[0-9]+}",
		controller.Delete).Methods("DELETE")
}

func initializePostRoutes(router *mux.Router) {
	repository := repository.NewSQLitePostRepository(config.SQLiteDatabase())
	service := service.NewPostService(repository)
	controller := controller.NewPostController(service)

	router.HandleFunc("/posts",
		controller.Create).Methods("POST")

	router.HandleFunc("/posts/today",
		controller.GetPopularToday).Methods("GET")

	router.HandleFunc("/posts/week",
		controller.GetPopularLastWeek).Methods("GET")

	router.HandleFunc("/posts/month",
		controller.GetPopularLastMonth).Methods("GET")

	router.HandleFunc("/posts/alltime",
		controller.GetPopularAllTime).Methods("GET")

	router.HandleFunc("/posts/{id:[0-9]+}",
		controller.GetByID).Methods("GET")

	router.HandleFunc("/users/{id:[0-9]+}/posts",
		controller.GetByUser).Methods("GET")

	router.HandleFunc("/posts/{id:[0-9]+}/title",
		controller.UpdateTitle).Methods("PUT")

	router.HandleFunc("/posts/{id:[0-9]+}/description",
		controller.UpdateDescription).Methods("PUT")

	router.HandleFunc("/posts/{id:[0-9]+}/content",
		controller.UpdateContent).Methods("PUT")

	router.HandleFunc("/posts/likes",
		controller.AddLike).Methods("POST")

	router.HandleFunc("/posts/likes",
		controller.DeleteLike).Methods("DELETE")

	router.HandleFunc("/posts/{id:[0-9]+}",
		controller.Delete).Methods("DELETE")
}

/*
func initializeCommentRoutes(router *mux.Router) {
	repository := repository.NewSQLiteCommentRepository(config.SQLiteDatabase())
	service := service.NewCommentService(repository)
	controller := controller.NewCommentController(service)

	router.HandleFunc("/comments",
		controller.Create).Methods("POST")

	router.HandleFunc("/comments/{id:[0-9]+}",
		controller.GetByID).Methods("GET")

	router.HandleFunc("/users/{id:[0-9]+}/comments",
		controller.GetByUser).Methods("GET")

	router.HandleFunc("/posts/{id:[0-9]+/comments}",
		controller.GetByPost).Methods("GET")

	router.HandleFunc("/comments/{id:[0-9]+}",
		controller.Update).Methods("PUT")

	router.HandleFunc("/comments/likes",
		controller.CreateLike).Methods("POST")

	router.HandleFunc("/comments/likes",
		controller.DeleteLike).Methods("DELETE")

	router.HandleFunc("/comments/{id:[0-9]+}",
		controller.Delete).Methods("DELETE")
}
*/
