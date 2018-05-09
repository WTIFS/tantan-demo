package service

import (
	"github.com/WTIFS/tantan-demo/dao"
	"github.com/WTIFS/tantan-demo/model"
)

func AddUser(user *model.User) (*model.User, error) {
	return dao.AddUser(user)
}

func ListUsers() ([]model.User, error) {
	return dao.ListUsers()
}