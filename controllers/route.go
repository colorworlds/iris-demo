package controllers

import (
	"time"

	"github.com/kataras/iris"

	"IRIS_WEB/controllers/http"
	"IRIS_WEB/utility/cache"
)

// 定义500错误处理函数
func err500(ctx iris.Context) {
	ctx.WriteString("CUSTOM 500 ERROR")
}

// 定义404错误处理函数
func err404(ctx iris.Context) {
	ctx.WriteString("CUSTOM 404 ERROR")
}

// 调试请求日志
func debug(ctx iris.Context) {
	cache.Set("debug_" + ctx.URLParamDefault("realIp", ctx.RemoteAddr()), "1", time.Hour)
	ctx.WriteString("You can watch in a hour. Ip: " + ctx.RemoteAddr())
}

// 注入路由
func InnerRoute(app *iris.Application) {
	app.OnErrorCode(iris.StatusInternalServerError, err500)
	app.OnErrorCode(iris.StatusNotFound, err404)
	app.Get("/debug", debug)

	app.Any("/users", http.ActionUsers)
	app.Any("/users/auth", jwtHandler.Serve, http.ActionUsers)
}
