package main

import "fmt"

//
//	แก้ไข func เหล่านี้ให้สามารถ return ค่าได้อย่างถูกต้อง
//	และ print ค่าใน main
//		- greeting("Your name") -> "Hello, Your Name"
//		- greetingWithAge("Your name", 20) -> "Hello, Your Name. You are 20 years old."
//		- greetingWithAgeAndDrink("Your name", 20, beer) -> "Hello, Your Name. You are 20 years old and I love beer."
//
//
//	* ใช้ fmt.Sprintf ช่วยในการ format และ concat string โดยใช้ %s แทน string และ %d แทน uint
//
//

func main() {
	fmt.Println(3333, "\n\n", 323)
}

func greeting(name string) string {
	return ""
}

func greetingWithAge(name string, age uint) string {
	return ""
}

func greetingWithAgeAndDrink(name string, age uint) string {
	return ""
}
