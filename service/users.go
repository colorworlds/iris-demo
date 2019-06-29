package service

import (
	"IRIS_WEB/dao"
	"IRIS_WEB/model"
)

func FetchUsersById(userId int) (provider []*model.UserDataProvider, err error) {
	// 根据ID查询用户
	var users []*model.UserModel
	if users, err = dao.QueryUsersById(userId); err != nil {
		return
	}

	provider = make([]*model.UserDataProvider, 0, 32)
	for _, user := range users {
		provider = append(provider, &model.UserDataProvider{
			ID:       user.ID,
			UserName: user.UserName,
			AuthKey:  user.AuthKey,
			Email:    user.Email,
			Status:   user.Status,
		})
	}

	return
}
