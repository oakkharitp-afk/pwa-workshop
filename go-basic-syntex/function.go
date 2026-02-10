package main

import "fmt"

// dosomething: ฟังก์ชันแบบพื้นฐานที่สุด
// - ไม่มีพารามิเตอร์ (Input)
// - ไม่มีค่าที่ส่งกลับ (Return Value)
func dosomething() {
	fmt.Println("Hello World")
}

// dosomething2: ฟังก์ชันที่มีการรับค่า (Parameter)
// - รับ msg ซึ่งมีประเภทข้อมูลเป็น string
// - ใช้ในการทำงานภายในฟังก์ชัน แต่ไม่มีการส่งค่ากลับไปยังผู้เรียก
func dosomething2(msg string) {
	fmt.Println("Hello World", msg)
}

// add: ฟังก์ชันที่มีทั้งการรับค่าและส่งค่ากลับ
// - รับ x และ y เป็น int (Input)
// - มีการระบุประเภทข้อมูลที่ส่งกลับเป็น int (อยู่หลังวงเล็บพารามิเตอร์)
// - ต้องมี Keyword 'return' เพื่อส่งผลลัพธ์ออกไป
func add(x int, y int) int {
	return x + y
}

// divide: รับตัวเลข 2 ตัว และส่งกลับทั้ง "ผลหาร" และ "ข้อความ Error"
func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Error: หารด้วยศูนย์ไม่ได้นะ!"
	}
	return a / b, "" // ส่งกลับผลลัพธ์ และ string ว่าง (แทนว่าไม่มี error)
}

// sum: รับตัวเลขกี่ตัวก็ได้ แล้วเอามาบวกกันทั้งหมด
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

// getRectangleArea: คำนวณพื้นที่และเส้นรอบรูป
func getRectangleArea(w, h float64) (area, perimeter float64) {
	area = w * h
	perimeter = (w + h) * 2
	return // จะส่งค่า area และ perimeter ออกไปโดยอัตโนมัติ
}
