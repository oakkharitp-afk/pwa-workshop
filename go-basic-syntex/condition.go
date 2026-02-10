package main

import "fmt"

// biggest: ฟังก์ชันเปรียบเทียบตัวเลข 2 ตัว
func biggest(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// max: ฟังก์ชันหาค่าที่มากที่สุดใน Slice ของตัวเลข
func max(numbers []int) int {
	// ป้องกันกรณีส่ง Slice ว่างเข้ามา เพื่อไม่ให้เกิด Error (Index out of range)
	if len(numbers) == 0 {
		return 0
	}

	// กำหนดค่าเริ่มต้นเป็นสมาชิกตัวแรก
	max := numbers[0]

	// การใช้ range เพื่อวนลูป (นิยมที่สุดใน Go)
	// _ (blank identifier) ใช้ละเว้น index เพราะเราต้องการแค่ค่า (n)
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	/*
	   // ตัวอย่างการ loop อีกวิธีหนึ่ง
	   for i := 0; i <= len(numbers)-1; i++ {
	       if numbers[i] > max {
	           max = numbers[i]
	       }
	   }
	*/

	return max
}

// exSwitch: ตัวอย่างการใช้ Switch Case
func exSwitch(day string) {
	switch day {
	case "sunday":
		// ใน Go เมื่อจบ Case จะหยุดทันที (Break อัตโนมัติ)
		// ถ้าต้องการให้ไหลลงไปทำงาน Case ถัดไป ต้องใช้คำสั่ง fallthrough
		fallthrough

	case "saturday":
		// ถ้าเป็น sunday (เพราะมี fallthrough) หรือ saturday จะมาทำตรงนี้
		// rest()

	default:
		// ถ้าไม่ตรงกับเงื่อนไขใดเลย จะมาทำที่นี่
		// work()
	}
}

func checkScore(score int) {
	// ประกาศตัวแปร status แล้วเช็คเงื่อนไขทันที
	// ตัวแปร status จะใช้งานได้เฉพาะภายใน Block if/else นี้เท่านั้น
	if status := "Pass"; score >= 50 {
		fmt.Println(status, "with score:", score)
	} else {
		status = "Fail"
		fmt.Println(status)
	}
	// ตรงนี้จะใช้ตัวแปร status ไม่ได้แล้ว (ช่วยประหยัด Memory และลดความซับซ้อน)
}
