package main

import "fmt"

// global variable and constant

const Pi = 3.14

var privateGlobalVal = "variable value"
var PubliceGlobalVal = "variable value"

// multiple variable declaration
var (
	PublicVal1 = "Go Basic Syntex"
	PublicVal2 = "1.0.0"
)

func function1() {
	var localVariable1 = "local variable"
	fmt.Println(localVariable1)

	// access global variable and constant
	fmt.Println(privateGlobalVal)
	fmt.Println(Pi)
	fmt.Println(PublicVal1)
	fmt.Println(PublicVal2)

	var number1 int = 10
	var number2 int // type inference, default value is 0
	number2 = 20

	_ = number1 + number2
}

func function2() {
	var localVariable2 = "local variable"
	localVariable3 := "local variable"
	fmt.Println(localVariable2)
	fmt.Println(localVariable3)

	// fmt.Println(localVariable1) // error: undefined variableName

	boolVar := 1 > 2
	if boolVar {
		number := 100
		fmt.Println("1 is greater than 2")
		fmt.Println(number)
	}
	// fmt.Println(number) // error: undefined variableName

	if !boolVar {
		// localVariable3 = localVariable3 + " changed"  // เขียนทับ variable ด้านบน
		// localVariable3 := localVariable3 + " changed" // ไม่เขียนทับ variable ด้านบน
		localVariable3 := "new val" // shadowing
		fmt.Println(localVariable3)

	}

}
