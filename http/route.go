package http

import (
	"IRIS_WEB/http/api"
	"github.com/kataras/iris"
)

func innerRoute(app *iris.Application) {
	app.Any("/users", jwtHandler.Serve, api.ActionUsers)
}
