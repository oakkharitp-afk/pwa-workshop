package advanced

import "fmt"

var addFunc = func(x, y int) int {
	return x + y
}
var multiplyFunc = func(x, y int) int {
	return x * y
}

var subtractFunc = func(x, y int) int {
	return x - y
}

type Myfunc func(x, y int) int

func mathFunc(x, y int, f Myfunc) int {
	n := f(x, y)
	fmt.Printf("x=%d, y=%d, result=%d\n", x, y, n)
	return n
}

func MathFuncCaller() {
	x, y := 6, 2
	mathFunc(x, y, addFunc)
	mathFunc(x, y, multiplyFunc)
	mathFunc(x, y, subtractFunc)
}
