package main

import (
	"fmt"
)

// 1. Global Scope & Constants
const (
	Pi     = 3.14159  // Exported (Public) - ขึ้นต้นด้วยตัวใหญ่
	status = "active" // Unexported (Private) - ใช้ได้เฉพาะใน package นี้
)

// 2. Grouping Declarations: จัดกลุ่มตัวแปรที่มีความเกี่ยวข้องกัน
var (
	AppName    = "GoLearning"
	AppVersion = "2.0.1"
	IsRunning  = false
)

func variableDemo() {
	// 3. Short Variable Declaration (:=)
	// นิยมใช้ที่สุดในฟังก์ชัน เพราะสั้นและ Go จะเดา Type (Inference) ให้เอง
	score := 95.5    // float64
	name := "Gopher" // string
	fmt.Println("3. Short Decl:", name, score)

	// 4. Multiple Initializers
	// ประกาศและกำหนดค่าหลายตัวในบรรทัดเดียว
	x, y, z := 10, 20, 30
	fmt.Println("4. Multiple:", x, y, z)

	// 5. Zero Values (หัวใจสำคัญของ Go)
	// Go จะจองหน่วยความจำและให้ค่าเริ่มต้นเสมอ ไม่มี "undefined"
	var (
		a int     // ได้ 0
		b string  // ได้ "" (ข้อความว่าง)
		c bool    // ได้ false
		d float64 // ได้ 0.0
	)
	fmt.Printf("Default values: %v, %q, %v, %v\n", a, b, c, d)

	// 6. Type Conversion (การแปลงประเภทข้อมูล)
	// Go จะไม่แปลงค่าให้โดยอัตโนมัติ (Strongly Typed)
	integerVal := 42
	floatVal := float64(integerVal) // ต้องระบุประเภทที่ต้องการแปลงให้ชัดเจน
	fmt.Println("6. Conversion:", integerVal, "to", floatVal)

	// 7. Shadowing & Block Scope
	limit := 10
	fmt.Println("Before block:", limit) // 10
	{
		// สร้าง scope ใหม่ด้วยเครื่องหมายปีกกา
		// การใช้ := ในนี้จะเป็นการสร้างตัวแปรใหม่ (Shadow) ที่ชื่อเหมือนกัน
		limit := 99
		fmt.Println("Inside block (shadowing):", limit) // 99
	}

	fmt.Println("After block:", limit) // 10 (ค่าเดิมไม่เปลี่ยน)

	// 8. การใช้ Blank Identifier (_)
	// ใช้เพื่อละทิ้งค่าที่เราไม่ต้องการ (ป้องกัน Error: unused variable)
	res, _ := calculateSomething()
	fmt.Println("Result:", res)
}

func calculateSomething() (int, string) {
	return 100, "Success"
}
