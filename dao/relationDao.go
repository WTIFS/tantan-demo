package dao

import (
	"github.com/WTIFS/tantan-demo/model"
	"fmt"
)

//list relation of one user
func ListRelationsByUserId(userId int64) ([]model.Relation, error) {
	db := GetConn()
	defer db.Close()

	var err error
	var relations []model.Relation
	err = db.Model(&relations).Select()
	if err != nil {
		panic(err)
	}
	fmt.Println(relations)
	return relations, err
}