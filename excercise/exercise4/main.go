package main

import (
	"fmt"
)

//	แก้ไข func ต่อไปนี้ ให้สามารถทำงานได้อย่างถูกต้อง โดยที่
//
//		func double จะเพิ่มค่า value ที่อยู่ใน pointer เป็น 2 เท่า
//		func appendGreeting จะเติม "Hi, " เข้าไปข้างหน้า
//
//	* assign ค่ากลับเข้าไปที่ pointer
//

func main() {
	n := 2
	fmt.Println(n) // 2

	double(&n)
	fmt.Println(n) // 4

	name := "Your name"
	fmt.Println(name) // Your name

	appendGreeting(&name)
	fmt.Println(name) // Hi, Your name
}

// n := 2
// double(&n)
// n == 4
func double(n *int) {

}

// name := "Your name"
// appendGreeting(&name)
// name == "Hi, Your name"
func appendGreeting(s *string) {

}
