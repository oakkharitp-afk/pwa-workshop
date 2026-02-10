package advanced

import "fmt"

// Shape คือ interface
// ใครก็ตามที่มี method Area() float64
// จะถือว่าเป็น Shape ได้ทันที
type Shape interface {
	Area() float64
}

// Rectangle เป็น struct ปกติ
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle implement Shape โดยอัตโนมัติ
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Println("area =", s.Area())
}
