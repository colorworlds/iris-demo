package api

import (
	"IRIS_WEB/errs"
	"IRIS_WEB/model"
	"IRIS_WEB/service"
	"fmt"
	"github.com/kataras/iris/context"
	"gopkg.in/go-playground/validator.v9"
)

func ActionUsers(ctx context.Context) {
	var err error
	var usersForm model.UserForm
	var users []*model.UserDataProvider

	if err = ctx.ReadForm(&usersForm); err != nil {
		ctx.JSON(errs.ParamError(err))
		return
	}

	//查看是否符合验证
	if err = usersForm.Validate(); err != nil {
		fmt.Printf("%#v\n", err)
		if e, ok := err.(validator.ValidationErrors); ok {
			fe, _ := e[0].(validator.FieldError)
			fmt.Printf("%+v\n", fe.ActualTag())
			fmt.Printf("%+v\n", fe.Tag())
			fmt.Printf("%+v\n", fe.Field())
			fmt.Printf("%+v\n", fe.Namespace())
			fmt.Printf("%+v\n", fe.StructField())
			fmt.Printf("%+v\n", fe.Param())
			fmt.Printf("%+v\n", fe.Value())
		}
		ctx.JSON(errs.ParamError(err))
		return
	}

	// 根据ID获取用户
	if users, err = service.FetchUsersById(usersForm.UserId); err != nil {
		ctx.JSON(errs.DBError(err))
		return
	}

	if len(users) == 0 {
		ctx.JSON(errs.NoDataError())
		return
	}

	ctx.JSON(errs.NoError(users))
}
