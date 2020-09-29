package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func main() {
	v := vertex{3, 4}
	fmt.Println(v)
	v.x = v.x * 3
	v.y = v.y * 2
	fmt.Println(v.ptr())

}
func (v *vertex) ptr() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)

}
