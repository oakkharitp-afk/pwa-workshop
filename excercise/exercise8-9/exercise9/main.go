package main

import (
	"errors"
	"fmt"
)

//
//	Error and Error Handling
//
//	แก้ไข func validateLength ให้ return error with message "too short string"
//	เมื่อ input string length สั้นกว่า 8 ตัวอักษร
//
//	และแก้ไข func validateAge ให้ return type ใหม่ "type ErrTooYoung int"
//	ที่มี method "Error() string" ซึ่ง return "[age] is too young"
//	ถ้าหาก input มีค่าน้อยกว่า 18
//
//	1. ใช้ errors.New ในการสร้าง error
//	2. type ErrTooYoung int func (err ErrTooYoung) Error() string { return "" }
//

func main() {
	fmt.Println(validateLength("a n t"))        //	too short string
	fmt.Println(validateLength("a n t ant มด")) // nil

	fmt.Println(validateAge(-1)) // your age is negative!
	fmt.Println(validateAge(17)) // 17 is too young
	fmt.Println(validateAge(20)) // nil
}

// validateLength return false when string length less than 8
// Please change return type to error with message "too short string"
func validateLength(s string) bool {
	if len([]rune(s)) < 8 {
		return false
	}
	return true
}

// TODO: สร้าง type ErrTooYoung int และ func (err ErrTooYoung) Error() string

var ageError = errors.New("your age is negative!")

// in case of too young age please create a new type ErrTooYoung int` with method `Error() string`
// example error message: "17 is too young"
func validateAge(n int) error {
	if n < 0 {
		return ageError
	}
	if n < 18 {
		return fmt.Errorf("too young")
	}
	return nil
}
