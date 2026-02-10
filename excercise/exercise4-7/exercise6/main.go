package main

import (
	"fmt"
)

//
// 	แก้ไข func wordCount ให้สามารถนับจำนวนคำต่างๆ ได้อย่างถูกต้อง เช่น
//	ถ้า input = "Apple Banana Apple Banana apple"
//	ต้อง return map[string]int{
//		"Apple": 2,
//		"Banana": 2,
//		"apple": 1,
//	}
//

func main() {
	input := "Apple Banana Apple Banana apple"
	result := wordCount(input)
	fmt.Println(result) // map[Apple:2 Banana:2 apple:1]
}

// TODO: split words and count them
func wordCount(s string) map[string]int {
	return nil
}
