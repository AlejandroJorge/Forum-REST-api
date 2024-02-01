package tests

import (
	"net/http"

	"github.com/AlejandroJorge/forum-rest-api/delivery/router"
)

var testRouter http.Handler

func MockRouter() http.Handler {
	if testRouter == nil {
		newRouter := router.AppRouter(MockSQLiteDatabase())
		testRouter = newRouter
	}
	return testRouter
}
