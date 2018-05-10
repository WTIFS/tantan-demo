package main

import (
	"github.com/WTIFS/tantan-demo/dao"
	"fmt"
)

func main() {


	//dao.ListRelationsByUserId(1)

	//u := &model.User{
	//	Name: "test",
	//	Mobile: "13838383388",
	//}
	//service.AddUser(u)


	user, err := dao.GetUserById(1)
	fmt.Println(user)
	if (err != nil) {
		panic(err)
	}

	dao.ListUsers()
	dao.ListRelationsByUserId(1)
	a, b := dao.GetRelationByUserId(2,3)
	fmt.Println(a)
	fmt.Println(b)
}

