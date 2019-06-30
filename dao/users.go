package dao

import (
	"IRIS_WEB/model"
	"IRIS_WEB/utility/db"
)

// 根据用户ID查询
func QueryUsersById(userId int) (users []*model.UserModel, err error) {
	res := db.GetMysql().Where("id = ?", userId).Find(&users)

	err = res.Error

	return
}

// 根据用户名模糊查询
func QueryUsersByName(userName string) (users []*model.UserModel, err error) {
	res := db.GetMysql().Where("user_name like ?", "%" + userName + "%").Find(&users)

	err = res.Error

	return
}
