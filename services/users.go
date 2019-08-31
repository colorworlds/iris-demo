package services

import (
	"IRIS_WEB/models"
	"IRIS_WEB/models/dao"
)

func FetchUsersById(userId int) (provider []*models.UserDataProvider, err error) {
	// 根据ID查询用户
	var users []*models.UserModel
	if users, err = dao.QueryUsersById(userId); err != nil {
		return
	}

	provider = make([]*models.UserDataProvider, 0, 32)
	for _, user := range users {
		provider = append(provider, &models.UserDataProvider{
			ID:       user.ID,
			UserName: user.UserName,
			AuthKey:  user.AuthKey,
			Email:    user.Email,
			Status:   user.Status,
		})
	}

	return
}
