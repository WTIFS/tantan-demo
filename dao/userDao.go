package dao

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type User struct {
	//tableName struct{} `sql:"userasdf, alias:u"`
	Id *int64
	Name string
	Mobile string
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Mobile)
}

type Relation struct {
	user_id int64
	other_user_id int64
	relation int16
	is_matched bool
}

func (s Relation) String() string {
	return fmt.Sprintf("Relation<%d %s %s %s>", s.user_id, s.other_user_id, s.relation, s.is_matched)
}

func ExampleDB_Model() {
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

	//user1 := &User{
	//	Id: nil,
	//	Name: "test1",
	//	Mobile: "13838383388",
	//}
	//err := db.Insert(user1)
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
	var users []User
	var dbModel = db.Model(&users)
	err := dbModel.Select()
	if err != nil {
		panic(err)
	}

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
	//fmt.Println(user)
	// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Relation<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*User)(nil), (*Relation)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}