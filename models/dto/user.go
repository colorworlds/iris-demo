package dto

import (
	"IRIS_WEB/errs"
	"github.com/kataras/iris"
	"IRIS_WEB/utility/validator"
)

type UserDTO struct {
	UserId    int    `form:"user_id" json:"user_id" validate:"min=1"`
	UserEmail string `form:"user_email" json:"user_email" validate:"email"`
	UserPhone string `form:"user_phone" json:"user_phone" validate:"phone"`
}

func (u *UserDTO) Bind(ctx iris.Context) error {
	if err := ctx.ReadForm(u); err != nil {
		return errs.ParamError("invalid form format")
	}

	if err := validator.Validate.Struct(u); err != nil {
		return errs.ParamError(validator.TransError(err))
	}

	return nil
}
