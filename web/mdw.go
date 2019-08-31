package web

import (
	"IRIS_WEB/utility/cache"
	"IRIS_WEB/utility/helper"
	"fmt"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"time"
)

// 统一异常处理
func NewRecoverMdw() iris.Handler {
	return func(ctx iris.Context) {
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

				logrus.Errorf("recover => %s", logMessage)

				ctx.StatusCode(500)
				ctx.StopExecution()
			}
		}()

		ctx.Next()
	}
}

// 请求日志记录
func NewAccessLogMdw() iris.Handler {
	return func(ctx iris.Context) {
		// 只有记录在案的ip才会打印请求日志
		realIp := ctx.RemoteAddr()
		if v, _ := cache.Get("debug_" + realIp); v != "1" {
			ctx.Next()
			return
		}

		begin := time.Now()

		reqBody := helper.RequestBody(ctx)
		// 如果请求内容不是json，则不记录
		if strings.Index(reqBody, "{") != 0 {
			reqBody = "non json body..."
		}

		ctx.Record()

		defer func() {
			respBody := string(ctx.Recorder().Body())
			// 响应内容必须是json格式，含有code码的字符串，否则忽略响应内容
			if strings.Index(respBody, "{") != 0 || strings.Index(respBody, "\"code\"") == -1 {
				respBody = "non json body..."
			}

			duration := time.Now().Sub(begin).Nanoseconds() / 1000000

			logrus.WithFields(logrus.Fields{
				"ip":       realIp,
				"method":   ctx.Method(),
				"path":     ctx.Path(),
				"header":   helper.RequestHeader(ctx),
				"queries":  helper.RequestQueries(ctx),
				"body":     reqBody,
				"duration": duration,
			}).Info(respBody)
		}()

		ctx.Next()
	}
}
