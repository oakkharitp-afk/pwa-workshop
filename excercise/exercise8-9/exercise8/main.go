package main

import "fmt"

//
//	Interface
//
// 	implement "volumer" interface ให้ type cube และ triangularPrism
//	โดยที่ volume ของ cube จะเท่ากับ edge * edge * edge
// 	และ volume ของ triangularPrism จะเท่ากับ 0.5 * base * attitude * height
//

func main() {
	cube := cube{
		edge: 2,
	}

	fmt.Println("cube volume:", VolumeOf(cube)) // 8

	triangularPrism := triangularPrism{
		base:     2,
		attitude: 2,
		height:   2,
	}

	fmt.Println("triangular prism:", VolumeOf(triangularPrism)) // 4
}

type volumer interface {
	Volume() float64
}

type cube struct {
	edge float64
} // edge x edge x edge

type triangularPrism struct {
	base     float64
	attitude float64
	height   float64
} // 0.5 x base x attitude x height

func VolumeOf(v volumer) float64 {
	return v.Volume()
}

//
//	TODO
//

func (c cube) Volume() float64 {
	return 0
}

func (t triangularPrism) Volume() float64 {
	return 0
}
