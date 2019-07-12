package http

import (
	"IRIS_WEB/utility/helper"
	"IRIS_WEB/utility/log"
	"github.com/dgrijalva/jwt-go"
	jwtMdw "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"time"
)

// jwt中间件
var jwtHandler = jwtMdw.New(jwtMdw.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("IRIS_WEB"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// 请求日志记录
func NewAccessLogMdw() iris.Handler {
	return func(ctx context.Context) {
		begin := time.Now()

		method := ctx.Method()

		path := ctx.Path()

		header := helper.RequestHeader(ctx)

		queries := helper.RequestQueries(ctx)

		body := helper.RequestBody(ctx)

		defer func() {
			code := ctx.ResponseWriter().StatusCode()

			duration := time.Now().Sub(begin).Nanoseconds() / 1000000

			log.Info("[ACCESS-LOG] Method: %s, Path: %s, Header: %s, Queries: %s, Body: %s, StatusCode: %d, Duration: %d ms",
				method,
				path,
				header,
				queries,
				body,
				code,
				duration,
			)
		}()

		ctx.Next()
	}
}
