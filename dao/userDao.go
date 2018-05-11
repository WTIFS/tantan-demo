package dao

import (
	"github.com/WTIFS/tantan-demo/model"
)

//add a user
func AddUser(user *model.User) (*model.User, error) {
	db := GetConn()
	defer db.Close()
	_, err := db.Model(user).Returning("*").Insert()
	if (err != nil) {
		panic(err)
	}
	return user, err
}

//list all users
func ListUsers() ([]model.User, error) {
	db := GetConn()
	defer db.Close()

	var err error
	var users []model.User
	var dbModel = db.Model(&users)
	err = dbModel.Select()
	if err != nil {
		panic(err)
	}
	//fmt.Println(users)
	return users, err
}
