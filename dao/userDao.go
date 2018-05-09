package dao

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/WTIFS/tantan-demo/model"
)

//add a user
func AddUser(user *model.User) (*model.User, error) {
	db := GetConn()
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



func ExampleDB_Model() {
	var err error
	db := pg.Connect(&pg.Options{
		Addr: "localhost:5432",
		User: "postgres",
		Password: "www",
		Database: "putong",
	})
	defer db.Close()

	//err := createSchema(db)
	//if err != nil {
	//	panic(err)
	//}

	//user1 := &model.User{
	//	Name: "test1",
	//	Mobile: "13838383388",
	//}
	//err = db.Insert(user1)
	//if err != nil {
	//	panic(err)
	//}

	/*user2 := &User{
		name: "test2",
		mobile: "13838383389",
	}
	err = db.Insert(user2)
	if err != nil {
		panic(err)
	}

	relation1 := &Relation{
		user_id: 1,
		other_user_id: user1.id,
		relation: 1,
		is_matched: false,
	}
	err = db.Insert(relation1)
	if err != nil {
		panic(err)
	}

	// Select user by primary key.
	user3 := &User{id: user1.id}
	err = db.Select(user3)
	if err != nil {
		panic(err)
	}*/

	// Select all users.
	var users []model.User
	var dbModel = db.Model(&users)
	err = dbModel.Select()
	if err != nil {
		panic(err)
	}

	//无效的query
	//var user2 User
	//a,b := db.Query(pg.Scan(&user2), "SELECT id, name, mobile FROM users WHERE id=?", 1)
	//fmt.Println(a)
	//fmt.Println(b)

	var userName, userMobile string
	err = db.Model((*model.User)(nil)).Column("name", "mobile").Where("id=?", 1).Select(&userName, &userMobile)
	fmt.Println(userName)
	fmt.Println(userMobile)
	// Select relation and associated author in one query.
	//relation := new(Relation)
	//err = db.Model(relation).
	//	Relation("other_user_id").
	//	Where("user.id = ?", 1).
	//	Select()
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println(user1.String())
	fmt.Println(users)
	//fmt.Println(user2)
	//fmt.Println(user)
	// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Relation<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
}
