package http

import (
	"IRIS_WEB/models"
	"IRIS_WEB/models/dto"
	"IRIS_WEB/errs"
	"IRIS_WEB/services"
	"github.com/kataras/iris"
)

func ActionUsers(ctx iris.Context) {
	var err error
	var params dto.UserDTO
	var users []*models.UserDataProvider

	// 绑定参数
	if err = params.Bind(ctx); err != nil {
		ctx.JSON(err)
		return
	}

	// 根据ID获取用户
	if users, err = services.FetchUsersById(params.UserId); err != nil {
		ctx.JSON(errs.DBError(err))
		return
	}

	if len(users) == 0 {
		ctx.JSON(errs.NoDataError())
		return
	}

	ctx.JSON(errs.NoError(users))
}
