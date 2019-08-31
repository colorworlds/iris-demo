package dto

import "gopkg.in/go-playground/validator.v9"

type UserDTO struct {
	UserId    int    `form:"user_id" validate:"required,min=1,max=10000"`
	UserEmail string `form:"user_email" validate:"email" json:"user_email"`
}

func (u UserDTO) Validate() error {
	return validator.New().Struct(u)
}
