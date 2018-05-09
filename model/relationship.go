package model

import (
	"fmt"
	"time"
)

type Relation struct {
	UserId int64
	OtherUserId int64
	Relation int16
	IsMatched bool
	AddTime time.Time `sql:"default:now()"`
}

func (s Relation) String() string {
	return fmt.Sprintf("Relation<%v %v %d %v>", s.UserId, s.OtherUserId, s.Relation, s.IsMatched)
}