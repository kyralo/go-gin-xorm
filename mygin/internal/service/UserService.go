package service

import (
	"mygin/internal/dao"
	"mygin/internal/model"
)

//ListUsers
func ListUsers() []model.User {
	return dao.ListUsers()
}

func SelectUserById(id int64) model.User{
	return dao.SelectUserById(id)
}
