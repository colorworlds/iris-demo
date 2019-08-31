package http

import (
	"github.com/kataras/iris"

	"IRIS_WEB/models"
	"IRIS_WEB/models/dto"
	"IRIS_WEB/models/errs"
	"IRIS_WEB/services"
)

func ActionUsers(ctx iris.Context) {
	var err error
	var params dto.UserDTO
	var users []*models.UserDataProvider

	if err = ctx.ReadForm(&params); err != nil {
		ctx.JSON(errs.ParamError(err))
		return
	}

	//查看是否符合验证
	if err = params.Validate(); err != nil {
		ctx.JSON(errs.ParamError(err))
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
