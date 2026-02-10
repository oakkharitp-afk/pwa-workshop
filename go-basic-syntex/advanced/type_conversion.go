package advanced

import (
	"fmt"
	"strconv"
)

var a int = 10
var b float64 = float64(a) // ต้องแปลงเอง
// var b float64 = a // compile error

func tc1() {
	i := 42

	s := strconv.Itoa(i)      // int → string
	n, err := strconv.Atoi(s) // string → int
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	fmt.Println(n)
}
