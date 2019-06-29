package api

import (
	"IRIS_WEB/model"
	"IRIS_WEB/service"
	"github.com/kataras/iris/context"
)

func ActionUsers(ctx context.Context) {
	var err error
	var usersForm model.UserForm
	var users []*model.UserDataProvider

	if err = ctx.ReadForm(&usersForm); err != nil {
		Err(ctx, ERROR_PARAM, "ActionUsers ReadForm Failed", err)
		return
	}

	//查看是否符合验证
	if err = usersForm.Validate(); err != nil {
		Err(ctx, ERROR_PARAM, "ActionUsers usersForm validate Failed", err)
		return
	}

	// 根据ID获取用户
	if users, err = service.FetchUsersById(usersForm.UserId); err != nil {
		Err(ctx, ERROR, "ActionUsers FetchUsers Failed", err)
		return
	}

	if len(users) == 0 {
		Err(ctx, ERROR_NODATA, "ActionUsers FetchUsers No Data", err)
		return
	}

	Suc(ctx, users)
}
