package http

import (
	"IRIS_WEB/http/api"
	"github.com/kataras/iris"
)

func innerRoute(app *iris.Application) {
	app.Any("/users", api.ActionUsers)
	app.Any("/users/auth", jwtHandler.Serve, api.ActionUsers)
}
