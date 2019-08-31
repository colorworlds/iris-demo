package controllers

import (
	"github.com/dgrijalva/jwt-go"
	mdwJwt "github.com/iris-contrib/middleware/jwt"
)

// jwt中间件
var jwtHandler = mdwJwt.New(mdwJwt.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("IRIS_WEB"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
