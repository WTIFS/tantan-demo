package service

import (
	"github.com/WTIFS/tantan-demo/model"
	"github.com/WTIFS/tantan-demo/dao"
	"github.com/WTIFS/tantan-demo/constants"
)

//insert or update relationship
func UpsertRelationship(relationship1 *model.Relationship) (*model.Relationship, error) {
	err1 := dao.UpsertRelationship(relationship1)
	relationship2, err2 := dao.GetRelationByUserId(relationship1.FromUserId, relationship1.ToUserId)
	if (err1 != nil) {
		return relationship1, err1
	}
	if (err2 == nil) { //relationship2 found
		if (relationship1.State == constants.RELATIONSHIP_STATE_LIKED && relationship2.State == constants.RELATIONSHIP_STATE_LIKED) {
			relationship1.State = constants.RELATIONSHIP_STATE_MATCHED
		}
	}
	return relationship1, nil
}

//list relationships of a user
func ListRelationsByUserId(userId int64) ([]model.Relationship, error) {
	return dao.ListRelationsByUserId(userId)
}
