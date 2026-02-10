package main

import (
	"fmt"
)

func BasicType() {
	// --- 1. ตัวแปรประเภทข้อความ (Strings) ---
	str1 := "Hello\nWorld" // สตริงปกติ (ตีความอักขระพิเศษเช่น \n)
	// Raw String (แสดงผลตามที่พิมพ์เป๊ะๆ มักใช้กับ JSON หรือ SQL)
	str2 := `Hello
World`
	// การพิมพ์ ` ใน Windows (รหัส Alt): กดปุ่ม Alt ค้างไว้ แล้วพิมพ์ 9 จากนั้นพิมพ์ 6 บนแป้นตัวเลข

	fmt.Println("--- Strings ---")
	fmt.Println(str1)
	fmt.Println(str2)

	// --- 2. ตัวแปรประเภทตัวเลข (Numbers) ---
	numInt := 3          // Default คือ int (ขนาดตาม OS)
	numFloat := 3.       // Default คือ float64
	var u uint = 7       // Unsigned int (เก็บเฉพาะค่าบวก)
	var p float32 = 22.7 // บังคับให้เป็น 32-bit เพื่อประหยัดพื้นที่

	fmt.Printf("\n--- Numbers ---\nType: %T, Value: %v\n", numInt, numInt)
	fmt.Printf("Type: %T, Value: %v\n", numFloat, numFloat)
	fmt.Printf("Type: %T, Value: %v\n", u, u)
	fmt.Printf("Type: %T, Value: %v\n", p, p)

	// --- 3. อาร์เรย์ (Arrays) vs สไลซ์ (Slices) ---
	// Array: ขนาดคงที่ [ขนาด]ประเภท
	var arr [3]int = [3]int{1, 2, 3}

	// Slice: ขนาดปรับเปลี่ยนได้ []ประเภท (ไม่มีตัวเลขในก้ามปู)
	// นิยมใช้มากกว่า Array ในการเขียน Go ทั่วไป
	slice := []int{10, 20, 30}
	slice = append(slice, 40) // เพิ่มข้อมูลได้ (Array ทำไม่ได้)

	fmt.Println("\n--- Array vs Slice ---")
	fmt.Println("Array (Fixed):", arr)
	fmt.Println("Slice (Dynamic):", slice)

	// --- 4. การแปลงประเภทข้อมูล (Type Conversions) ---
	// Go ไม่มีการแปลงค่าอัตโนมัติ (No Implicit Conversion)
	i := 42
	f := float64(i)
	u2 := uint(f)

	fmt.Println("\n--- Type Conversions ---")
	fmt.Printf("Converted: %v (Type: %T)\n", f, f)
	fmt.Printf("Converted back: %v (Type: %T)\n", u2, u2)
}
