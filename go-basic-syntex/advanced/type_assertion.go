package advanced

import "fmt"

func ta() {
	var v interface{} = "hello"

	// เสี่ยง
	s := v.(string) // ถ้าไม่ใช่ string → panic
	fmt.Println(s)

	// ปลอดภัย
	s, ok := v.(string)
	if !ok {
		fmt.Println("not a string")
		return
	}
	fmt.Println(s)

}

func process(v interface{}) {
	switch t := v.(type) {
	case string:
		fmt.Println("string:", t)
	case int:
		fmt.Println("int:", t)
	case bool:
		fmt.Println("bool:", t)
	default:
		fmt.Println("unknown type")
	}
}
