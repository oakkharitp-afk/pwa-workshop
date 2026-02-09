package main

import "fmt"

type User struct {
	// <field name> <type> <struct tag>
	ID   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
}

// receiver function
func (u *User) Say() {
	fmt.Println("hello my name is", u.Name)
}

// receiver function --> this user state change
func (u *User) ChangeName(name string) {
	u.Name = name
	fmt.Println("name in receiver function :", u.Name)
}

// receiver function --> this user state not change
func (u User) ChangeName2(name string) {
	u.Name = name
	fmt.Println("name in receiver function :", u.Name)
}
