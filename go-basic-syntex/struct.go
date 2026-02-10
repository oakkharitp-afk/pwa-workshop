package main

import "fmt"

// User: โครงสร้างข้อมูล (Struct) ที่เก็บข้อมูลผู้ใช้
type User struct {
	// <field name> <type> <struct tag>
	// Struct Tag ใช้สำหรับกำหนดชื่อฟิลด์เวลาแปลงเป็น JSON หรือเชื่อมต่อ Database (BSON)
	ID   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name" json:"name"`
}

// 1. Say: Receiver function แบบ Pointer (*User)
// นิยมใช้แบบ Pointer เพื่อประสิทธิภาพ (ไม่ต้อง Copy ข้อมูลใหม่ทั้งก้อน)
func (u *User) Say() {
	fmt.Printf("Hello, my name is %s (ID: %s)\n", u.Name, u.ID)
}

// 2. ChangeName: Receiver function แบบ Pointer
// เมื่อใช้ * (Pointer) การแก้ไขค่า u.Name จะส่งผลถึงตัวแปรต้นฉบับ (State change)
func (u *User) ChangeName(name string) {
	u.Name = name
	fmt.Println("Inside ChangeName (Pointer):", u.Name)
}

// 3. ChangeName2: Receiver function แบบ Value
// เมื่อไม่ใส่ * (Value) Go จะทำการ Copy ข้อมูล User มาใหม่ทั้งก้อน
// การแก้ไขค่าในนี้จะ "ไม่ส่งผล" ถึงตัวแปรต้นฉบับข้างนอก (No state change)
func (u User) ChangeName2(name string) {
	u.Name = name
	fmt.Println("Inside ChangeName2 (Value):", u.Name)
}

func StructEx() {
	// การสร้าง Object จาก Struct
	me := User{ID: "U001", Name: "Somchai"}

	fmt.Println("--- เริ่มต้น ---")
	me.Say()

	fmt.Println("\n--- ทดสอบ ChangeName2 (Value) ---")
	me.ChangeName2("Wichai")
	fmt.Println("ผลลัพธ์ข้างนอก:", me.Name) // จะยังเป็น Somchai อยู่

	fmt.Println("\n--- ทดสอบ ChangeName (Pointer) ---")
	me.ChangeName("Wichai")
	fmt.Println("ผลลัพธ์ข้างนอก:", me.Name) // เปลี่ยนเป็น Wichai แล้ว
}
