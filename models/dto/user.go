package dto

import (
	. "IRIS_WEB/errors"
	"IRIS_WEB/utility/validator"
	"github.com/kataras/iris"
)

type UserDTO struct {
	UserId    int    `form:"user_id" json:"user_id" validate:"min=1"`
	UserEmail string `form:"user_email" json:"user_email" validate:"email"`
	UserPhone string `form:"user_phone" json:"user_phone" validate:"phone"`
}

func (u *UserDTO) Bind(ctx iris.Context) error {
	if err := ctx.ReadForm(u); err != nil {
		return ParamError("invalid form format")
	}

	if err, errMsg := validator.Check(u); err != nil {
		return ParamError(err, errMsg)
	}

	return nil
}
