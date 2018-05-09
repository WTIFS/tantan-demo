package main

import (
	"github.com/WTIFS/tantan-demo/dao"
)

func main() {


	//dao.ListRelationsByUserId(1)

	//u := &model.User{
	//	Name: "test",
	//	Mobile: "13838383388",
	//}
	//service.AddUser(u)

	dao.ListUsers()
	dao.ListRelationsByUserId(1)
}

