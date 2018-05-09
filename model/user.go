package model

import (
	"fmt"
)

//By default all columns except primary keys are nullable and go-pg marshals Go zero values (empty string, 0, zero time, empty map or slice) as SQL NULL. This behavior can be changed using sql:",notnull" tag.
type User struct {
	// tableName is an optional field that specifies custom table name and alias.
	// By default go-pg generates table name and alias from struct name.
	// tableName struct{} `sql:"users, alias:u"`
	Id int64 // Id is automatically detected as primary key
	Name string
	Mobile string
	Type string
	//Birthday time.Time `sql:"-"` // - is used to ignore field
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Mobile)
}