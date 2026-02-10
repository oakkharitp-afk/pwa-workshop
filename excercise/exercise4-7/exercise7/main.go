package main

import "fmt"

//	จง implement method String() string
//	โดยถ้า book มี Name = "Harry Potter" และ Author = "J. K. Rowling"
//	เมื่อรัน book.String() จะได้ "Harry Potter by J. K. Rowling"
//	และ implement method SetName(name string) เพื่อใช้สำหรับแก้ไข Name ของ book

type Book struct {
	Name   string
	Author string
}

func main() {
	// TODO: set value to struct
	book := Book{}

	fmt.Println(book.String()) // Harry Potter by J. K. Rowling

	book.SetName("Harry Potter EP:2")
	fmt.Println(book.String()) // Harry Potter EP:2 by J. K. Rowling
}

// get book detail
func (b *Book) String() string {
	// TODO: concat string and return
	return ""
}

// set book name
func (b *Book) SetName(name string) {
	// TODO: set new book name
}
