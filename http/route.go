package http

import (
	"IRIS_WEB/http/api"
	"github.com/kataras/iris"
)

func innerRoute(app *iris.Application) {
	app.Get("/users", api.ActionUsers)
}