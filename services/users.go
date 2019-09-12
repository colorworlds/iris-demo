package services

import (
	"IRIS_WEB/errors"
	"IRIS_WEB/models"
	"IRIS_WEB/models/dao"
)

func FetchUsersById(userId int) ([]*models.UserModel, error) {
	// 根据ID查询用户
	var err error
	var users []*models.UserModel
	if users, err = dao.QueryUsersById(userId); err != nil {
		return nil, errors.DBError("query users by id", err)
	}

	return users, nil
}

