package service

import (
	"github.com/WTIFS/tantan-demo/model"
	"github.com/WTIFS/tantan-demo/dao"
	"github.com/WTIFS/tantan-demo/constants"
)

//insert or update relationship
func UpsertRelationship(relationship1 *model.Relationship) (*model.Relationship, error) {
	dao.UpsertRelationship(relationship1)
	relationship2, err := dao.GetRelationByUserId(relationship1.ToUserId, relationship1.FromUserId)
	if (err != nil) {
		return relationship1, err
	}
	if (relationship1.State == constants.RELATIONSHIP_STATE_LIKED && relationship1.State == relationship2.State) {
		relationship1.State = constants.RELATIONSHIP_STATE_MATCHED
	}
	return relationship1, nil
}

//list relationships of a user
func ListRelationsByUserId(userId int64) ([]model.Relationship, error) {
	return dao.ListRelationsByUserId(userId)
}
