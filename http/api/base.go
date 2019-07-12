package api

import (
	"IRIS_WEB/utility/log"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

const (
	SUCCESS      = 1000
	ERROR        = 1001
	ERROR_PARAM  = 1002
	ERROR_NODATA = 1003
)

func Err(ctx context.Context, errCode int, errMsg string, err error) {
	if err != nil {
		log.Error("[%s][%s] => [%v]", ctx.Path(), errMsg, err)
	}
	ctx.JSON(iris.Map{"code": errCode, "data": "", "msg": errMsg})
}

func Suc(ctx context.Context, data interface{}) {
	ctx.JSON(iris.Map{"code": SUCCESS, "data": data, "msg": ""})
}
