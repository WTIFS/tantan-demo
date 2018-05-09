package service

import (
	"github.com/WTIFS/tantan-demo/model"
	"github.com/WTIFS/tantan-demo/dao"
)

func ListRelationsByUserId(userId int64) ([]model.Relation, error) {
	return dao.ListRelationsByUserId(userId)
}