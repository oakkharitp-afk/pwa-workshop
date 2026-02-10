package main

import "fmt"

//
// 	แก้ไข func ต่อไปนี้ ให้สามารถตัดแบ่ง slice ของ string ดังกล่าว และสามารถ print ค่าได้อย่างถูกต้อง
//
//	* ใช้ slice operator ด้วย syntax slice[low:high] ในการตัดแบ่ง slice ได้
//

func main() {
	abc() // [apple banana coconut]
	efg()
	cde()
}

func abc() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "apple", "banana" and "coconut" needed

	fmt.Print(s) // [apple banana coconut]

}

func efg() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "elderberries", "figs" and "grapes" needed

	fmt.Print(s) // [elderberries figs grapes]
}

func cde() {
	s := []string{"apple", "banana", "coconut", "durian", "elderberries", "figs", "grapes"}

	// TODO: slice the s to the new variable, only "coconut", "durian" and "elderberries" needed

	fmt.Print(s) // [coconut durian elderberries]
}
