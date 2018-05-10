package dao

import (
	"github.com/WTIFS/tantan-demo/model"
	"fmt"
)

func GetRelationByUserId(fromUserId int64, toUserId int64) (model.Relationship, error) {
	db := GetConn()
	defer db.Close()

	var err error

	var relationship model.Relationship
	err = db.Model(&relationship).
		Where("from_user_id=?", fromUserId).
		Where("to_user_id=?", toUserId).
		Select()
	return relationship, err
}

//add or update relationship
func UpsertRelationship(relationship *model.Relationship) error {
	db := GetConn()
	defer db.Close()

	_, err := db.Model(relationship).
		OnConflict("(from_user_id, to_user_id) DO UPDATE").
		Set("state=?state").
		Insert()
	return err
}

//list relation of one user
func ListRelationsByUserId(fromUserId int64) ([]model.Relationship, error) {
	db := GetConn()
	defer db.Close()

	var err error
	var relationships []model.Relationship
	err = db.Model(&relationships).Where("from_user_id=?", fromUserId).Select()
	if err != nil {
		panic(err)
	}
	fmt.Println(relationships)
	return relationships, err
}