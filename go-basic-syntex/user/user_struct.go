package user

import "fmt"

type User struct {
	// <field name> <type> <struct tag>
	ID   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
}

func (u *User) Say() {
	fmt.Println("Hello my name is", u.Name)
}
