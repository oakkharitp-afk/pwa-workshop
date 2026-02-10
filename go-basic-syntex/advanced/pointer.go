package advanced

import "fmt"

func Pointer() {
	x := 10
	p := &x // p เก็บ address ของ x

	fmt.Println(x)  // 10
	fmt.Println(&x) // address ของ x
	fmt.Println(p)  // address เดียวกับ &x
	fmt.Println(*p) // dereference → ค่า 10
}
