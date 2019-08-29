package http

import (
	"IRIS_WEB/utility/helper"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtMdw "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"runtime"
	"time"
)

// jwt中间件
var jwtHandler = jwtMdw.New(jwtMdw.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("IRIS_WEB"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// 统一异常处理
func NewRecoverMdw() iris.Handler {
	return func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break

					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				// when stack finishes
				logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName())
				logMessage += fmt.Sprintf("At Request: %d %s %s %s\n", ctx.GetStatusCode(), ctx.Path(), ctx.Method(), ctx.RemoteAddr())
				logMessage += fmt.Sprintf("Trace: %s\n", err)
				logMessage += fmt.Sprintf("\n%s", stacktrace)

				golog.Errorf("recover => %s", logMessage)

				ctx.StatusCode(500)
				ctx.StopExecution()
			}
		}()

		ctx.Next()
	}
}

// 请求日志记录
func NewAccessLogMdw() iris.Handler {
	return func(ctx context.Context) {
		begin := time.Now()

		method := ctx.Method()

		path := ctx.Path()

		header := helper.RequestHeader(ctx)

		queries := helper.RequestQueries(ctx)

		body := helper.RequestBody(ctx)

		ctx.Recorder()

		defer func() {
			respBody := string(ctx.Recorder().Body())

			duration := time.Now().Sub(begin).Nanoseconds() / 1000000

			golog.Infof("Method: %s, Path: %s, Header: %s, Queries: %s, Body: %s, response: %v, Duration: %d ms",
				method,
				path,
				header,
				queries,
				body,
				respBody,
				duration,
			)
		}()

		ctx.Next()
	}
}
