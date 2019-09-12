package http

import (
	"IRIS_WEB/errors"
	"IRIS_WEB/models"
	"IRIS_WEB/models/dto"
	"IRIS_WEB/services"
	"github.com/kataras/iris"
)

func ActionUsers(ctx iris.Context) {
	var err error
	var params dto.UserDTO
	var users []*models.UserModel

	// 绑定参数
	if err = params.Bind(ctx); err != nil {
		ctx.JSON(err)
		return
	}

	// 根据ID获取用户
	if users, err = services.FetchUsersById(params.UserId); err != nil {
		ctx.JSON(err)
		return
	}

	if len(users) == 0 {
		ctx.JSON(errors.NoDataError())
		return
	}

	ctx.JSON(iris.Map{"code": 1000, "data": users})
}
