package model

import (
	"time"
	"fmt"
)

type Relationship struct {
	Id int64 `json:"-"`
	FromUserId int64 `json:"-"`
	ToUserId int64 `json:"user_id"`
	State string `json:"state"`
	AddTime time.Time `sql:"default:now()" json:"-"`
	Type string `sql:"-" json:"type"`
}

//type RelationshipExternal struct {
//	UserId int64 `json:"user_id"`
//	State string `json:"state"`
//	Type string `json:"type"`
//}
//
func (r *Relationship) SetType() {
	r.Type = "relationship"
	fmt.Println(r)
}

//func (s Relationship) String() string {
//	return fmt.Sprintf("Relation<%v %v %v>", s.UserId, s.OtherUserId, s.State)
//}