package http

import (
	"IRIS_WEB/http/api"
	"github.com/dgrijalva/jwt-go"
	jwtMdw "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

var jwtHandler = jwtMdw.New(jwtMdw.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("IRIS_WEB"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func innerRoute(app *iris.Application) {
	app.Get("/users", jwtHandler.Serve, api.ActionUsers)
}
