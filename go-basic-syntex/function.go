package main

import "fmt"

func dosomething() {
	fmt.Println("Hello World")
}

func dosomething2(msg string) {
	fmt.Println("Hello World", msg)
}

func add(x int, y int) int {
	return x + y
}
